package usecases

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockstore "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/mock"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/models"
)

func TestGetChain_ByUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	chainAgent := mockstore.NewMockChainAgent(mockCtrl)

	getChainUC := NewGetChain(chainAgent)
	chainUUID := uuid.Must(uuid.NewV4()).String()

	expectedChain := &models.Chain{
		UUID: chainUUID,
		Name: "testChain",
	}
	chainAgent.EXPECT().GetChainByUUID(gomock.Any(), gomock.Eq(chainUUID)).Return(expectedChain, nil).Times(1)

	actualChain, err := getChainUC.Execute(context.Background(), chainUUID, "")
	assert.Nil(t, err)
	assert.Equal(t, expectedChain, actualChain)
}

func TestGetChain_ByUUIDAndTenantID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	chainAgent := mockstore.NewMockChainAgent(mockCtrl)

	getChainUC := NewGetChain(chainAgent)
	chainUUID := uuid.Must(uuid.NewV4()).String()
	tenantID := "tenantID_3"

	expectedChain := &models.Chain{
		UUID:     chainUUID,
		TenantID: tenantID,
		Name:     "testChain",
	}
	chainAgent.EXPECT().GetChainByUUIDAndTenant(gomock.Any(), gomock.Eq(chainUUID), tenantID).Return(expectedChain, nil).Times(1)

	actualChain, err := getChainUC.Execute(context.Background(), chainUUID, tenantID)
	assert.Nil(t, err)
	assert.Equal(t, expectedChain, actualChain)
}
