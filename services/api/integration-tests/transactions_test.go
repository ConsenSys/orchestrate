// +build integration

package integrationtests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/api"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
	clientutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/client-utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/sdk/client"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/testutils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/tx"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/utils"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/service/controllers"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/chain-registry/store/models"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/contract-registry/proto"
	"gopkg.in/h2non/gock.v1"
)

const (
	waitForEnvelopeTimeOut = 2 * time.Second
)

// transactionsTestSuite is a test suite for Transaction Scheduler Transactions controller
type transactionsTestSuite struct {
	suite.Suite
	client client.OrchestrateClient
	faucet *api.FaucetResponse
	env    *IntegrationEnvironment
}

func (s *transactionsTestSuite) TestValidation() {
	ctx := s.env.ctx
	chain := testutils.FakeChain()

	s.T().Run("should fail if payload is invalid", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransactionRequest()
		txRequest.ChainName = ""

		_, err := s.client.SendContractTransaction(ctx, txRequest)

		assert.True(t, errors.IsInvalidFormatError(err))
	})

	s.T().Run("should fail with 400 if from account is set and oneTimeKeyEnabled", func(t *testing.T) {
		txRequest := testutils.FakeSendTransactionRequest()
		txRequest.Params.OneTimeKey = true

		_, err := s.client.SendContractTransaction(ctx, txRequest)

		assert.True(t, errors.IsInvalidFormatError(err))
	})

	s.T().Run("should fail if idempotency key is identical but different params", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransactionRequest()
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: utils.RandomString(16),
		})

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)
		gock.New(chainRegistryURL).
			URL(fmt.Sprintf("%s?chain_uuid=%s&account=%s", chainRegistryURL, chain.UUID, txRequest.Params.From)).
			Reply(404).Done()
		_, err := s.client.SendContractTransaction(rctx, txRequest)
		assert.NoError(t, err)

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)
		txRequest.Params.MethodSignature = "differentMethodSignature()"
		_, err = s.client.SendContractTransaction(rctx, txRequest)
		assert.True(t, errors.IsConstraintViolatedError(err))
	})

	s.T().Run("should fail with 422 if chains cannot be fetched", func(t *testing.T) {
		defer gock.Off()
		gock.New(chainRegistryURL).Get("/chains").Reply(404)
		txRequest := testutils.FakeSendTransactionRequest()

		_, err := s.client.SendContractTransaction(ctx, txRequest)

		assert.True(t, errors.IsInvalidParameterError(err))
	})

	s.T().Run("should fail with 422 if account does not exist", func(t *testing.T) {
		defer gock.Off()

		// Create a txRequest with an inexisting account
		txRequest := testutils.FakeSendTransactionRequest()
		txRequest.Params.From = "0x905B88EFf8Bda1543d4d6f4aA05afef143D27E18"

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)

		_, err := s.client.SendContractTransaction(ctx, txRequest)

		assert.True(t, errors.IsInvalidParameterError(err))
	})

	s.T().Run("should fail with 422 if chainUUID does not exist", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransactionRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(404)
		gock.New(chainRegistryURL).
			URL(fmt.Sprintf("%s?chain_uuid=%s&account=%s", chainRegistryURL, chain.UUID, txRequest.Params.From)).
			Reply(404)

		_, err := s.client.SendContractTransaction(ctx, txRequest)

		assert.True(t, errors.IsInvalidParameterError(err))
	})
}

func (s *transactionsTestSuite) TestSuccess() {
	ctx := s.env.ctx
	chain := testutils.FakeChain()

	s.T().Run("should send a contract transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransactionRequest()

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)

		txRequest.Params.From = ""
		txRequest.Params.OneTimeKey = true
		IdempotencyKey := utils.RandomString(16)
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: IdempotencyKey,
		})
		txResponse, err := s.client.SendContractTransaction(rctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.NotEmpty(t, txResponse.UUID)
		assert.NotEmpty(t, txResponse.IdempotencyKey)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		job := txResponseGET.Jobs[0]

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.NotEmpty(t, job.UUID)
		assert.Equal(t, job.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusStarted, job.Status)
		assert.Equal(t, txRequest.Params.From, job.Transaction.From)
		assert.Equal(t, txRequest.Params.To, job.Transaction.To)
		assert.Equal(t, utils.EthereumTransaction, job.Type)

		evlp, err := s.env.consumer.WaitForEnvelope(job.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.Equal(t, job.ScheduleUUID, evlp.GetID())
		assert.Equal(t, job.UUID, evlp.GetJobUUID())
		assert.True(t, evlp.IsOneTimeKeySignature())
		assert.Equal(t, tx.JobTypeMap[utils.EthereumTransaction].String(), evlp.GetJobTypeString())
		assert.Equal(t, evlp.GetChainIDString(), chain.ChainID)
		assert.Equal(t, evlp.PartitionKey(), "")
	})

	s.T().Run("should send a transaction with an additional faucet job", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransferTransactionRequest()
		chainWithFaucet := testutils.FakeChain()
		chainWithFaucet.UUID = s.faucet.ChainRule

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chainWithFaucet})
		gock.New(chainRegistryURL).Get("/chains/" + chainWithFaucet.UUID).Times(2).Reply(200).JSON(chainWithFaucet)
		gock.New(chainWithFaucet.URLs[0]).Post("").Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x1000000000\"}")
		gock.New(chainWithFaucet.URLs[0]).Post("").Reply(200).BodyString("{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x0\"}")

		IdempotencyKey := utils.RandomString(16)
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: IdempotencyKey,
		})
		txResponse, err := s.client.SendTransferTransaction(rctx, txRequest)
		require.NoError(t, err)
		assert.NotEmpty(t, txResponse.UUID)
		assert.NotEmpty(t, txResponse.IdempotencyKey)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		require.NoError(t, err)
		require.Len(t, txResponseGET.Jobs, 2)

		faucetJob := txResponseGET.Jobs[1]
		txJob := txResponseGET.Jobs[0]
		assert.Equal(t, faucetJob.ChainUUID, chainWithFaucet.UUID)
		assert.Equal(t, utils.StatusStarted, faucetJob.Status)
		assert.Equal(t, utils.EthereumTransaction, faucetJob.Type)
		assert.Equal(t, faucetJob.Transaction.To, txJob.Transaction.From)
		assert.Equal(t, faucetJob.Transaction.Value, s.faucet.Amount)

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.NotEmpty(t, txJob.UUID)
		assert.Equal(t, txJob.ChainUUID, chainWithFaucet.UUID)
		assert.Equal(t, utils.StatusStarted, txJob.Status)
		assert.Equal(t, txRequest.Params.From, txJob.Transaction.From)
		assert.Equal(t, txRequest.Params.To, txJob.Transaction.To)
		assert.Equal(t, utils.EthereumTransaction, txJob.Type)

		fctEvlp, err := s.env.consumer.WaitForEnvelope(faucetJob.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.Equal(t, faucetJob.ScheduleUUID, fctEvlp.GetID())
		assert.Equal(t, faucetJob.UUID, fctEvlp.GetJobUUID())
		assert.Equal(t, tx.JobTypeMap[utils.EthereumTransaction].String(), fctEvlp.GetJobTypeString())
		assert.Equal(t, fctEvlp.GetChainIDString(), chain.ChainID)

		jobEvlp, err := s.env.consumer.WaitForEnvelope(txJob.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.Equal(t, txJob.ScheduleUUID, jobEvlp.GetID())
		assert.Equal(t, txJob.UUID, jobEvlp.GetJobUUID())
	})

	s.T().Run("should send a tessera transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTesseraRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Times(2).Reply(200).JSON(chain)

		IdempotencyKey := utils.RandomString(16)
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: IdempotencyKey,
		})
		txResponse, err := s.client.SendContractTransaction(rctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.NotEmpty(t, txResponse.UUID)
		assert.NotEmpty(t, txResponse.IdempotencyKey)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.Len(t, txResponseGET.Jobs, 2)

		privTxJob := txResponseGET.Jobs[0]
		assert.NotEmpty(t, privTxJob.UUID)
		assert.Equal(t, privTxJob.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusStarted, privTxJob.Status)
		assert.Equal(t, txRequest.Params.From, privTxJob.Transaction.From)
		assert.Equal(t, txRequest.Params.To, privTxJob.Transaction.To)
		assert.Equal(t, utils.TesseraPrivateTransaction, privTxJob.Type)

		markingTxJob := txResponseGET.Jobs[1]
		assert.NotEmpty(t, markingTxJob.UUID)
		assert.Equal(t, markingTxJob.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusCreated, markingTxJob.Status)
		assert.Equal(t, utils.TesseraMarkingTransaction, markingTxJob.Type)

		privTxEvlp, err := s.env.consumer.WaitForEnvelope(privTxJob.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, privTxJob.ScheduleUUID, privTxEvlp.GetID())
		assert.Equal(t, privTxJob.UUID, privTxEvlp.GetJobUUID())
		assert.False(t, privTxEvlp.IsOneTimeKeySignature())
		assert.Equal(t, tx.JobTypeMap[utils.TesseraPrivateTransaction].String(), privTxEvlp.GetJobTypeString())
	})

	s.T().Run("should send an orion transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendOrionRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Times(2).Reply(200).JSON(chain)

		txResponse, err := s.client.SendContractTransaction(ctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.NotEmpty(t, txResponse.UUID)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.Len(t, txResponseGET.Jobs, 2)

		privTxJob := txResponseGET.Jobs[0]
		assert.NotEmpty(t, privTxJob.UUID)
		assert.Equal(t, privTxJob.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusStarted, privTxJob.Status)
		assert.Equal(t, txRequest.Params.From, privTxJob.Transaction.From)
		assert.Equal(t, txRequest.Params.To, privTxJob.Transaction.To)
		assert.Equal(t, utils.OrionEEATransaction, privTxJob.Type)

		markingTxJob := txResponseGET.Jobs[1]
		assert.NotEmpty(t, markingTxJob.UUID)
		assert.Equal(t, markingTxJob.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusCreated, markingTxJob.Status)
		assert.Equal(t, utils.OrionMarkingTransaction, markingTxJob.Type)

		privTxEvlp, err := s.env.consumer.WaitForEnvelope(privTxJob.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, privTxJob.ScheduleUUID, privTxEvlp.GetID())
		assert.Equal(t, privTxJob.UUID, privTxEvlp.GetJobUUID())
		assert.Equal(t, tx.JobTypeMap[utils.OrionEEATransaction].String(), privTxEvlp.GetJobTypeString())
	})

	s.T().Run("should send a deploy contract transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeDeployContractRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Times(2).Reply(200).JSON(chain)

		txRequest.Params.Args = testutils.ParseIArray(123) // FakeContract arguments

		s.env.contractRegistryResponseFaker.GetContract = func() (*proto.GetContractResponse, error) {
			return &proto.GetContractResponse{
				Contract: testutils.FakeContract(),
			}, nil
		}
		txResponse, err := s.client.SendDeployTransaction(ctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.NotEmpty(t, txResponse.UUID)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		job := txResponseGET.Jobs[0]

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.NotEmpty(t, job.UUID)
		assert.Equal(t, job.ChainUUID, chain.UUID)
		assert.Equal(t, utils.StatusStarted, job.Status)
		assert.Equal(t, txRequest.Params.From, job.Transaction.From)
		assert.Equal(t, utils.EthereumTransaction, job.Type)

		evlp, err := s.env.consumer.WaitForEnvelope(job.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.Equal(t, job.ScheduleUUID, evlp.GetID())
		assert.Equal(t, job.UUID, evlp.GetJobUUID())
		assert.Equal(t, tx.JobTypeMap[utils.EthereumTransaction].String(), evlp.GetJobTypeString())
	})

	s.T().Run("should send a raw transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendRawTransactionRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)
		IdempotencyKey := utils.RandomString(16)
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: IdempotencyKey,
		})
		txResponse, err := s.client.SendRawTransaction(rctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.NotEmpty(t, txResponse.UUID)
		assert.NotEmpty(t, txResponse.IdempotencyKey)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		job := txResponseGET.Jobs[0]

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.NotEmpty(t, job.UUID)
		assert.Equal(t, utils.StatusStarted, job.Status)
		assert.Equal(t, txRequest.Params.Raw, job.Transaction.Raw)
		assert.Equal(t, utils.EthereumRawTransaction, job.Type)

		evlp, err := s.env.consumer.WaitForEnvelope(job.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		assert.Equal(t, job.ScheduleUUID, evlp.GetID())
		assert.Equal(t, job.UUID, evlp.GetJobUUID())
		assert.Equal(t, tx.JobTypeMap[utils.EthereumRawTransaction].String(), evlp.GetJobTypeString())
	})

	s.T().Run("should send a transfer transaction successfully", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransferTransactionRequest()
		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)

		txResponse, err := s.client.SendTransferTransaction(ctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		
		assert.NotEmpty(t, txResponse.UUID)

		txResponseGET, err := s.client.GetTxRequest(ctx, txResponse.UUID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		job := txResponseGET.Jobs[0]

		assert.NotEmpty(t, txResponseGET.UUID)
		assert.NotEmpty(t, job.UUID)
		assert.Equal(t, utils.StatusStarted, job.Status)
		assert.Equal(t, txRequest.Params.Value, job.Transaction.Value)
		assert.Equal(t, txRequest.Params.To, job.Transaction.To)
		assert.Equal(t, txRequest.Params.From, job.Transaction.From)
		assert.Equal(t, utils.EthereumTransaction, job.Type)

		evlp, err := s.env.consumer.WaitForEnvelope(job.ScheduleUUID, s.env.kafkaTopicConfig.Sender, waitForEnvelopeTimeOut)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		assert.Equal(t, job.ScheduleUUID, evlp.GetID())
		assert.Equal(t, job.UUID, evlp.GetJobUUID())
		assert.Equal(t, tx.JobTypeMap[utils.EthereumTransaction].String(), evlp.GetJobTypeString())
	})

	s.T().Run("should succeed if payloads and idempotency key are the same and return same schedule", func(t *testing.T) {
		defer gock.Off()
		txRequest := testutils.FakeSendTransactionRequest()
		idempotencyKey := utils.RandomString(16)
		rctx := context.WithValue(ctx, clientutils.RequestHeaderKey, map[string]string{
			controllers.IdempotencyKeyHeader: idempotencyKey,
		})

		// Kill Kafka on first call so data is added in DB and status is CREATED but does not get updated to STARTED
		err := s.env.client.Stop(rctx, kafkaContainerID)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)

		_, err = s.client.SendContractTransaction(rctx, txRequest)
		assert.Error(t, err)

		err = s.env.client.StartServiceAndWait(rctx, kafkaContainerID, 10*time.Second)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}

		gock.New(chainRegistryURL).Get("/chains").Reply(200).JSON([]*models.Chain{chain})
		gock.New(chainRegistryURL).Get("/chains/" + chain.UUID).Reply(200).JSON(chain)

		txResponse, err := s.client.SendContractTransaction(rctx, txRequest)
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
		job := txResponse.Jobs[0]
		assert.Equal(t, idempotencyKey, txResponse.IdempotencyKey)
		assert.Equal(t, utils.StatusStarted, job.Status)
	})
}
