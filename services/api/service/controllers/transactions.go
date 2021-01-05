package controllers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/api"

	"github.com/gorilla/mux"
	jsonutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/encoding/json"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/httputil"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/multitenancy"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/business/use-cases"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/api/service/formatters"
)

var _ entities.ETHTransactionParams

const (
	IdempotencyKeyHeader = "X-Idempotency-Key"
)

type TransactionsController struct {
	ucs usecases.TransactionUseCases
}

func NewTransactionsController(ucs usecases.TransactionUseCases) *TransactionsController {
	return &TransactionsController{
		ucs: ucs,
	}
}

// Add routes to router
func (c *TransactionsController) Append(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/transactions/send").
		Handler(http.HandlerFunc(c.send))
	router.Methods(http.MethodPost).Path("/transactions/send-raw").
		Handler(http.HandlerFunc(c.sendRaw))
	router.Methods(http.MethodPost).Path("/transactions/transfer").
		Handler(http.HandlerFunc(c.transfer))
	router.Methods(http.MethodPost).Path("/transactions/deploy-contract").
		Handler(http.HandlerFunc(c.deployContract))
	router.Methods(http.MethodGet).Path("/transactions/{uuid}").
		Handler(http.HandlerFunc(c.getOne))
	router.Methods(http.MethodGet).Path("/transactions").
		Handler(http.HandlerFunc(c.search))
}

// @Summary Creates and sends a new contract transaction
// @Description Creates and executes a new smart contract transaction request
// @Description The transaction can be private (Tessera, Orion).
// @Description The transaction can be a One Time Key transaction in 0 gas private networks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param request body api.SendTransactionRequest{params=api.TransactionParams{gasPricePolicy=api.GasPriceParams{retryPolicy=api.RetryParams}}} true "Contract transaction request"
// @Success 202 {object} api.TransactionResponse "Created contract transaction request"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 409 {object} httputil.ErrorResponse "Already existing transaction"
// @Failure 422 {object} httputil.ErrorResponse "Unprocessable parameters were sent"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions/send [post]
func (c *TransactionsController) send(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	txRequest := &api.SendTransactionRequest{}
	if err := jsonutils.UnmarshalBody(request.Body, txRequest); err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err := txRequest.Params.Validate(); err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	txReq := formatters.FormatSendTxRequest(txRequest, request.Header.Get(IdempotencyKeyHeader))
	txResponse, err := c.ucs.SendContractTransaction().Execute(ctx, txReq, multitenancy.TenantIDFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(rw).Encode(formatters.FormatTxResponse(txResponse))
}

// @Summary Creates and sends a new contract deployment
// @Description Creates and executes a new contract deployment request
// @Description The transaction can be private (Tessera, Orion).
// @Description The transaction can be a One Time Key transaction in 0 gas private networks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param request body api.DeployContractRequest{params=api.DeployContractParams{gasPricePolicy=api.GasPriceParams{retryPolicy=api.RetryParams}}} true "Deployment transaction request"
// @Success 202 {object} api.TransactionResponse "Created deployment transaction request"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 409 {object} httputil.ErrorResponse "Already existing transaction"
// @Failure 422 {object} httputil.ErrorResponse "Unprocessable parameters were sent"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions/deploy-contract [post]
func (c *TransactionsController) deployContract(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	txRequest := &api.DeployContractRequest{}
	if err := jsonutils.UnmarshalBody(request.Body, txRequest); err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err := txRequest.Params.Validate(); err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	txReq := formatters.FormatDeployContractRequest(txRequest, request.Header.Get(IdempotencyKeyHeader))
	txResponse, err := c.ucs.SendDeployTransaction().Execute(ctx, txReq, multitenancy.TenantIDFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(rw).Encode(formatters.FormatTxResponse(txResponse))
}

// @Summary Creates and sends a raw transaction
// @Description Creates and executes a new raw transaction request
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param request body api.RawTransactionRequest{params=api.RawTransactionParams{retryPolicy=api.IntervalRetryParams}} true "Raw transaction request"
// @Success 202 {object} api.TransactionResponse "Created raw transaction request"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 409 {object} httputil.ErrorResponse "Already existing transaction"
// @Failure 422 {object} httputil.ErrorResponse "Unprocessable parameters were sent"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions/send-raw [post]
func (c *TransactionsController) sendRaw(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	txRequest := &api.RawTransactionRequest{}
	err := jsonutils.UnmarshalBody(request.Body, txRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	txReq := formatters.FormatSendRawRequest(txRequest, request.Header.Get(IdempotencyKeyHeader))
	txResponse, err := c.ucs.SendTransaction().Execute(ctx, txReq, "", multitenancy.TenantIDFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(rw).Encode(formatters.FormatTxResponse(txResponse))
}

// @Summary Creates and sends a transfer transaction
// @Description Creates and executes a new transfer request
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param request body api.TransferRequest{params=api.TransferParams{gasPricePolicy=api.GasPriceParams{retryPolicy=api.RetryParams}}} true "Transfer transaction request"
// @Success 202 {object} api.TransactionResponse "Created transfer transaction request"
// @Failure 400 {object} httputil.ErrorResponse "Invalid request"
// @Failure 409 {object} httputil.ErrorResponse "Already existing transaction"
// @Failure 422 {object} httputil.ErrorResponse "Unprocessable parameters were sent"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions/transfer [post]
func (c *TransactionsController) transfer(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	txRequest := &api.TransferRequest{}
	err := jsonutils.UnmarshalBody(request.Body, txRequest)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if err = txRequest.Params.Validate(); err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	txReq := formatters.FormatTransferRequest(txRequest, request.Header.Get(IdempotencyKeyHeader))
	txResponse, err := c.ucs.SendTransaction().Execute(ctx, txReq, "", multitenancy.TenantIDFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(rw).Encode(formatters.FormatTxResponse(txResponse))
}

// @Summary Fetch a transaction request by uuid
// @Description Fetch a single transaction request by uuid
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param uuid path string true "UUID of the transaction request"
// @Success 200 {object} api.TransactionResponse "Transaction request found"
// @Failure 404 {object} httputil.ErrorResponse "Transaction request not found"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions/{uuid} [get]
func (c *TransactionsController) getOne(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	uuid := mux.Vars(request)["uuid"]

	txRequest, err := c.ucs.GetTransaction().Execute(ctx, uuid, multitenancy.AllowedTenantsFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(formatters.FormatTxResponse(txRequest))
}

// @Summary Search transaction requests by provided filters
// @Description Get a list of filtered transaction requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security JWTAuth
// @Param idempotency_keys query []string false "List of idempotency keys" collectionFormat(csv)
// @Success 200 {array} api.TransactionResponse "List of transaction requests found"
// @Failure 400 {object} httputil.ErrorResponse "Invalid filter in the request"
// @Failure 500 {object} httputil.ErrorResponse "Internal server error"
// @Router /transactions [get]
func (c *TransactionsController) search(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	ctx := request.Context()

	filters, err := formatters.FormatTransactionsFilterRequest(request)
	if err != nil {
		httputil.WriteError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	txRequests, err := c.ucs.SearchTransactions().Execute(ctx, filters, multitenancy.AllowedTenantsFromContext(ctx))
	if err != nil {
		httputil.WriteHTTPErrorResponse(rw, err)
		return
	}

	var responses []*api.TransactionResponse
	for _, txRequest := range txRequests {
		responses = append(responses, formatters.FormatTxResponse(txRequest))
	}

	_ = json.NewEncoder(rw).Encode(responses)
}
