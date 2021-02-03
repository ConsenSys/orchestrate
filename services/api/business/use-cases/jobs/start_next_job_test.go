// +build unit

package jobs

import (
	"context"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
	mocks2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases/mocks"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/mocks"
	testutils2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/store/models/testutils"
)

func TestStartNextJob_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	mockJobDA := mocks.NewMockJobAgent(ctrl)
	mockTxDA := mocks.NewMockTransactionAgent(ctrl)
	mockStartJobUC := mocks2.NewMockStartJobUseCase(ctrl)

	mockDB.EXPECT().Job().Return(mockJobDA).AnyTimes()
	mockDB.EXPECT().Transaction().Return(mockTxDA).AnyTimes()

	usecase := NewStartNextJobUseCase(mockDB, mockStartJobUC)

	ctx := context.Background()
	tenants := []string{"tenantID"}

	t.Run("should execute use case for orion marking transaction successfully", func(t *testing.T) {
		jobModel := testutils2.FakeJobModel(0)
		nextJobModel := testutils2.FakeJobModel(0)
		txHash := ethcommon.HexToHash("0x123")

		jobModel.NextJobUUID = nextJobModel.UUID
		jobModel.Transaction.Hash = txHash.String()
		jobModel.Logs = append(jobModel.Logs, &models.Log{
			ID:        1,
			Status:    utils.StatusStored,
			CreatedAt: time.Now(),
		})
		jobModel.Type = utils.OrionEEATransaction
		nextJobModel.Type = utils.OrionMarkingTransaction

		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), jobModel.UUID, tenants).
			Return(jobModel, nil)
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), nextJobModel.UUID, tenants).
			Return(nextJobModel, nil)
		nextJobModel.Transaction.Data = txHash.String()
		mockTxDA.EXPECT().Update(gomock.Any(), nextJobModel.Transaction).Return(nil)

		mockStartJobUC.EXPECT().Execute(gomock.Any(), nextJobModel.UUID, tenants)
		err := usecase.Execute(ctx, jobModel.UUID, tenants)

		assert.NoError(t, err)
	})

	t.Run("should execute use case for tessera marking transaction successfully", func(t *testing.T) {
		jobModel := testutils2.FakeJobModel(0)
		nextJobModel := testutils2.FakeJobModel(0)
		enclaveKey := ethcommon.HexToHash("0x123").String()

		jobModel.NextJobUUID = nextJobModel.UUID
		jobModel.Transaction.EnclaveKey = enclaveKey
		jobModel.Transaction.Gas = "0x1"
		jobModel.Logs = append(jobModel.Logs, &models.Log{
			ID:        1,
			Status:    utils.StatusStored,
			CreatedAt: time.Now(),
		})
		jobModel.Type = utils.TesseraPrivateTransaction
		nextJobModel.Type = utils.TesseraMarkingTransaction

		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), jobModel.UUID, tenants).
			Return(jobModel, nil)
		mockJobDA.EXPECT().FindOneByUUID(gomock.Any(), nextJobModel.UUID, tenants).
			Return(nextJobModel, nil)
		nextJobModel.Transaction.Data = enclaveKey
		nextJobModel.Transaction.Gas = "0x1"
		mockTxDA.EXPECT().Update(gomock.Any(), nextJobModel.Transaction).Return(nil)

		mockStartJobUC.EXPECT().Execute(gomock.Any(), nextJobModel.UUID, tenants)
		err := usecase.Execute(ctx, jobModel.UUID, tenants)

		assert.NoError(t, err)
	})
}