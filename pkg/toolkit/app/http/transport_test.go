// +build unit

package http

import (
	"context"
	"net/http"
	"testing"
	"time"

	authutils "github.com/consensys/orchestrate/pkg/toolkit/app/auth/utils"
	"github.com/consensys/orchestrate/pkg/toolkit/app/http/transport"
	"github.com/consensys/orchestrate/pkg/toolkit/app/multitenancy"
	"github.com/stretchr/testify/assert"
)

type MockTransport struct {
	roundTrips int
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.roundTrips++
	return nil, nil
}

func TestAuthHeadersTransport(t *testing.T) {
	mockTransport := &MockTransport{}
	tt := transport.NewContextAuthHeadersTransport()(mockTransport)

	// Test without setting authorization in context
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	_, _ = tt.RoundTrip(req)
	assert.Equal(t, 1, mockTransport.roundTrips, "Mock transport should have been called")
	authorization := req.Header.Get("Authorization")
	assert.Equal(t, "", authorization, "Authorization header shuld be empty")

	authToken := "test-auth"
	// Test setting authorization in context
	userInfo := &multitenancy.UserInfo{
		AuthMode: multitenancy.AuthMethodJWT,
		AuthValue: authToken,
	}
	ctxOne := multitenancy.WithUserInfo(context.Background(), userInfo)
	req, _ = http.NewRequestWithContext(ctxOne, http.MethodGet, "", nil)
	
	_, _ = tt.RoundTrip(req)
	assert.Equal(t, 2, mockTransport.roundTrips, "Mock transport should have been called")
	authorization = req.Header.Get("Authorization")
	assert.Equal(t, authToken, authorization, "Authorization header should be empty")
}

func TestAuthAPIKeyHeadersTransport(t *testing.T) {
	mockTransport := &MockTransport{}
	tt := transport.NewXAPIKeyHeadersTransport("test-auth")(mockTransport)

	// Test setting X-API-Key in context
	ctxTwo := authutils.WithAPIKey(context.Background(), "test-auth")
	req, _ := http.NewRequestWithContext(ctxTwo, http.MethodGet, "", nil)
	_, _ = tt.RoundTrip(req)
	assert.Equal(t, 1, mockTransport.roundTrips, "Mock transport should have been called")
	authorization := authutils.GetAPIKeyHeaderValue(req)
	assert.Equal(t, "test-auth", authorization, "Authorization header should be empty")
}

type Mock409Transport struct {
	roundTrips int
}

func (t *Mock409Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.roundTrips++
	if t.roundTrips > 1 {
		return &http.Response{StatusCode: http.StatusOK}, nil
	}

	header := make(http.Header)
	header.Set("Retry-After", "1")
	return &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header:     header,
	}, nil
}

func TestRetry429Transport(t *testing.T) {
	mockT := &Mock409Transport{}
	transport := transport.NewRetry429Transport()(mockT)

	now := time.Now()
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	_, _ = transport.RoundTrip(req)
	elapsed := time.Since(now)

	assert.Equal(t, 2, mockT.roundTrips, "RoundTrip should have retried")
	assert.Greater(t, int64(elapsed), int64(time.Second), "Time should have elapsed during retry")
}
