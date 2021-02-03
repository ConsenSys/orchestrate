package sender

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/ethclient"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/log"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/sdk/client"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/tx-sender/tx-sender/use-cases"
	utils2 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/tx-sender/tx-sender/utils"
)

const sendTesseraPrivateTxComponent = "use-cases.send-tessera-private-tx"

type sendTesseraPrivateTxUseCase struct {
	ec               ethclient.QuorumTransactionSender
	chainRegistryURL string
	jobClient        client.JobClient
	crafter          usecases.CraftTransactionUseCase
	logger           *log.Logger
}

func NewSendTesseraPrivateTxUseCase(ec ethclient.QuorumTransactionSender, crafter usecases.CraftTransactionUseCase,
	jobClient client.JobClient, chainRegistryURL string) usecases.SendTesseraPrivateTxUseCase {
	return &sendTesseraPrivateTxUseCase{
		ec:               ec,
		chainRegistryURL: chainRegistryURL,
		jobClient:        jobClient,
		crafter:          crafter,
		logger:           log.NewLogger().SetComponent(sendTesseraPrivateTxComponent),
	}
}

func (uc *sendTesseraPrivateTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	ctx = log.With(log.WithFields(ctx, log.Field("job", job.UUID)), uc.logger)
	logger := uc.logger.WithContext(ctx)
	logger.Debug("processing tessera private transaction job")

	job.Transaction.Nonce = "0"
	err := uc.crafter.Execute(ctx, job)
	if err != nil {
		return errors.FromError(err).ExtendComponent(sendTesseraMarkingTxComponent)
	}

	job.Transaction.EnclaveKey, err = uc.sendTx(ctx, job)
	if err != nil {
		return errors.FromError(err).ExtendComponent(sendTesseraPrivateTxComponent)
	}

	err = utils2.UpdateJobStatus(ctx, uc.jobClient, job.UUID, utils.StatusStored, "", job.Transaction)
	if err != nil {
		return errors.FromError(err).ExtendComponent(sendTesseraPrivateTxComponent)
	}

	logger.Info("tessera private job was sent successfully")
	return nil
}

func (uc *sendTesseraPrivateTxUseCase) sendTx(ctx context.Context, job *entities.Job) (string, error) {
	logger := uc.logger.WithContext(ctx)
	proxyTessera := utils.GetProxyTesseraURL(uc.chainRegistryURL, job.ChainUUID)
	data, err := hexutil.Decode(job.Transaction.Data)
	if err != nil {
		errMsg := "cannot decode transaction data"
		logger.WithError(err).Errorf(errMsg)
		return "", err
	}

	enclaveKey, err := uc.ec.StoreRaw(ctx, proxyTessera, data, job.Transaction.PrivateFrom)
	if err != nil {
		errMsg := "cannot send tessera private transaction"
		logger.WithError(err).Errorf(errMsg)
		return "", err
	}

	return enclaveKey, nil
}