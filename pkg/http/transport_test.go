// +build unit

package http
// 
// import (
// 	"context"
// 	"net/http"
// 	"testing"
// 	"time"
// 
// 	"github.com/stretchr/testify/assert"
// 	authutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/auth/utils"
// )
// 
// type MockTransport struct {
// 	roundTrips int
// }
// 
// func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	t.roundTrips++
// 	return nil, nil
// }
// 
// func TestAuthHeadersTransport(t *testing.T) {
// 	mockTransport := &MockTransport{}
// 	transport := NewAuthHeadersTransport(mockTransport)
// 
// 	// Test without setting authorization in context
// 	req, _ := http.NewRequest(http.MethodGet, "", nil)
// 	_, _ = transport.RoundTrip(req)
// 	assert.Equal(t, 1, mockTransport.roundTrips, "Mock transport should have been called")
// 	authorization := req.Header.Get("Authorization")
// 	assert.Equal(t, "", authorization, "Authorization header shuld be empty")
// 
// 	// Test setting authorization in context
// 	req, _ = http.NewRequestWithContext(
// 		authutils.WithAuthorization(context.Background(), "test-auth"),
// 		http.MethodGet, "", nil,
// 	)
// 	_, _ = transport.RoundTrip(req)
// 	assert.Equal(t, 2, mockTransport.roundTrips, "Mock transport should have been called")
// 	authorization = req.Header.Get("Authorization")
// 	assert.Equal(t, "test-auth", authorization, "Authorization header shuld be empty")
// 
// 	// Test setting X-API-Key in context
// 	req, _ = http.NewRequestWithContext(
// 		authutils.WithAPIKey(context.Background(), "test-auth"),
// 		http.MethodGet, "", nil,
// 	)
// 	_, _ = transport.RoundTrip(req)
// 	assert.Equal(t, 3, mockTransport.roundTrips, "Mock transport should have been called")
// 	authorization = req.Header.Get(authutils.APIKeyHeader)
// 	assert.Equal(t, "test-auth", authorization, "Authorization header shuld be empty")
// }
// 
// type Mock409Transport struct {
// 	roundTrips int
// }
// 
// func (t *Mock409Transport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	t.roundTrips++
// 	if t.roundTrips > 1 {
// 		return &http.Response{StatusCode: http.StatusOK}, nil
// 	}
// 
// 	header := make(http.Header)
// 	header.Set("Retry-After", "1")
// 	return &http.Response{
// 		StatusCode: http.StatusTooManyRequests,
// 		Header:     header,
// 	}, nil
// }
// 
// func TestRetry429Transport(t *testing.T) {
// 	mockT := &Mock409Transport{}
// 	transport := NewRetry429Transport(mockT)
// 
// 	now := time.Now()
// 	req, _ := http.NewRequest(http.MethodGet, "", nil)
// 	_, _ = transport.RoundTrip(req)
// 	elapsed := time.Since(now)
// 
// 	assert.Equal(t, 2, mockT.roundTrips, "RoundTrip should have retried")
// 	assert.Greater(t, int64(elapsed), int64(time.Second), "Time should have elapsed during retry")
// }
