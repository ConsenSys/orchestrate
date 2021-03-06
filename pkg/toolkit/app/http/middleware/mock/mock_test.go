// +build unit

package mock

import (
	"testing"

	mockhandler "github.com/consensys/orchestrate/pkg/toolkit/app/http/handler/mock"
	"github.com/golang/mock/gomock"
)

func TestNewMockMiddleware(t *testing.T) {
	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	mockHandler1 := mockhandler.NewMockHandler(ctrlr)
	mid := NewMockMiddleware(mockHandler1)

	mockHandler2 := mockhandler.NewMockHandler(ctrlr)

	h := mid(mockHandler2)
	mockHandler1.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())
	mockHandler2.EXPECT().ServeHTTP(gomock.Any(), gomock.Any())

	h.ServeHTTP(nil, nil)
}
