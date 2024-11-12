// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// FinancialDataV2AccountsIDBalancesGet implements GET /financial-data/v2/accounts/{id}/balances operation.
	//
	// Returns a list of account balances. For further details on Pagination, see general section above.
	//
	// GET /financial-data/v2/accounts/{id}/balances
	FinancialDataV2AccountsIDBalancesGet(ctx context.Context, params FinancialDataV2AccountsIDBalancesGetParams) (FinancialDataV2AccountsIDBalancesGetRes, error)
	// FinancialDataV2EntitiesGet implements GET /financial-data/v2/entities operation.
	//
	// Returns a list of entities. For further details on Pagination, see general section above.
	//
	// GET /financial-data/v2/entities
	FinancialDataV2EntitiesGet(ctx context.Context, params FinancialDataV2EntitiesGetParams) (FinancialDataV2EntitiesGetRes, error)
	// FinancialDataV2EntitiesIDDelete implements DELETE /financial-data/v2/entities/{id} operation.
	//
	// Deletes the specified entity. It is only possible to delete an entity if it has no children or has
	// no associated accounts.
	//
	// DELETE /financial-data/v2/entities/{id}
	FinancialDataV2EntitiesIDDelete(ctx context.Context, params FinancialDataV2EntitiesIDDeleteParams) (FinancialDataV2EntitiesIDDeleteRes, error)
	// FinancialDataV2EntitiesIDGet implements GET /financial-data/v2/entities/{id} operation.
	//
	// Retrieve a specific entity.
	//
	// GET /financial-data/v2/entities/{id}
	FinancialDataV2EntitiesIDGet(ctx context.Context, params FinancialDataV2EntitiesIDGetParams) (FinancialDataV2EntitiesIDGetRes, error)
	// FinancialDataV2EntitiesPost implements POST /financial-data/v2/entities operation.
	//
	// Create entity.
	//
	// POST /financial-data/v2/entities
	FinancialDataV2EntitiesPost(ctx context.Context, req *CreateEntityRequest) (FinancialDataV2EntitiesPostRes, error)
	// FinancialDataV2PendingTransactionsIDEventsGet implements GET /financial-data/v2/pending-transactions/{id}/events operation.
	//
	// Retrieve events for a specific pending transaction. Returns a list of events. For further details
	// on Pagination, see the section above.
	//
	// GET /financial-data/v2/pending-transactions/{id}/events
	FinancialDataV2PendingTransactionsIDEventsGet(ctx context.Context, params FinancialDataV2PendingTransactionsIDEventsGetParams) (FinancialDataV2PendingTransactionsIDEventsGetRes, error)
	// IamV2betaOAuth2TokenPost implements POST /iam/v2beta/oauth2/token operation.
	//
	// This endpoint is an OAuth 2.0 <a href="https://datatracker.ietf.org/doc/html/rfc6749#section-3.
	// 2">token endpoint</a> and can be used by a client to obtain an access token.
	// Use <a href="https://app.atlar.com/users/all">programmatic access user</a> credentials as OAuth 2.
	// 0 client credentials.
	// Use the <i>access key</i> as <a href="https://datatracker.ietf.org/doc/html/rfc6749#section-2.
	// 2">client identifier (`client_id`)</a>.
	// Use the <i>secret</i> as <a href="https://datatracker.ietf.org/doc/html/rfc6749#section-2.3.
	// 1">client password (`client_secret`)</a>.
	// This endpoint supports <a href="https://datatracker.ietf.org/doc/html/rfc6749#section-2.3.
	// 1">authentication using HTTP Basic auth</a>.
	// This endpoint only supports the OAuth 2.0 <a href="https://datatracker.ietf.
	// org/doc/html/rfc6749#section-4.4">client credentials grant type</a>.
	//
	// POST /iam/v2beta/oauth2/token
	IamV2betaOAuth2TokenPost(ctx context.Context, req *IamV2betaOAuth2TokenPostReq) (IamV2betaOAuth2TokenPostRes, error)
	// PaymentsV2CounterpartiesIDDelete implements DELETE /payments/v2/counterparties/{id} operation.
	//
	// Deletes the specified counterparty.
	//
	// DELETE /payments/v2/counterparties/{id}
	PaymentsV2CounterpartiesIDDelete(ctx context.Context, params PaymentsV2CounterpartiesIDDeleteParams) (PaymentsV2CounterpartiesIDDeleteRes, error)
	// PaymentsV2CreditTransferBatchesIDApprovePost implements POST /payments/v2/credit-transfer-batches/{id}:approve operation.
	//
	// Approve the batch of credit transfers.
	// Approval of a batch is only applicable when the batch `treatment` is `BATCH`.
	//
	// POST /payments/v2/credit-transfer-batches/{id}:approve
	PaymentsV2CreditTransferBatchesIDApprovePost(ctx context.Context, req *ApproveBatchRequest, params PaymentsV2CreditTransferBatchesIDApprovePostParams) (PaymentsV2CreditTransferBatchesIDApprovePostRes, error)
	// PaymentsV2CreditTransferBatchesIDRejectPost implements POST /payments/v2/credit-transfer-batches/{id}:reject operation.
	//
	// Reject the batch of credit transfers.
	// Rejection of a batch is only applicable when the batch `treatment` is `BATCH`.
	//
	// POST /payments/v2/credit-transfer-batches/{id}:reject
	PaymentsV2CreditTransferBatchesIDRejectPost(ctx context.Context, req *RejectPaymentRequest, params PaymentsV2CreditTransferBatchesIDRejectPostParams) (PaymentsV2CreditTransferBatchesIDRejectPostRes, error)
	// PaymentsV2CreditTransfersIDEventsGet implements GET /payments/v2/credit-transfers/{id}/events operation.
	//
	// Retrieve events for a specific credit transfer. Returns a list of events. For further details on
	// Pagination, see the section above.
	//
	// GET /payments/v2/credit-transfers/{id}/events
	PaymentsV2CreditTransfersIDEventsGet(ctx context.Context, params PaymentsV2CreditTransfersIDEventsGetParams) (PaymentsV2CreditTransfersIDEventsGetRes, error)
	// PaymentsV2DirectDebitsIDEventsGet implements GET /payments/v2/direct-debits/{id}/events operation.
	//
	// Retrieve events for a specific direct debit. Returns a list of events. For further details on
	// Pagination, see the section above.
	//
	// GET /payments/v2/direct-debits/{id}/events
	PaymentsV2DirectDebitsIDEventsGet(ctx context.Context, params PaymentsV2DirectDebitsIDEventsGetParams) (PaymentsV2DirectDebitsIDEventsGetRes, error)
	// PaymentsV2ExternalAccountsIDDelete implements DELETE /payments/v2/external-accounts/{id} operation.
	//
	// Deletes the specified external account.
	//
	// DELETE /payments/v2/external-accounts/{id}
	PaymentsV2ExternalAccountsIDDelete(ctx context.Context, params PaymentsV2ExternalAccountsIDDeleteParams) (PaymentsV2ExternalAccountsIDDeleteRes, error)
	// PaymentsV2MandatesIDEventsGet implements GET /payments/v2/mandates/{id}/events operation.
	//
	// Retrieve events for a specific mandate. Returns a list of events. For further details on
	// Pagination, see the section above.
	//
	// GET /payments/v2/mandates/{id}/events
	PaymentsV2MandatesIDEventsGet(ctx context.Context, params PaymentsV2MandatesIDEventsGetParams) (PaymentsV2MandatesIDEventsGetRes, error)
	// PaymentsV2betaCreditTransferBatchesGet implements GET /payments/v2beta/credit-transfer-batches operation.
	//
	// Returns a list of credit transfer batches.
	//
	// GET /payments/v2beta/credit-transfer-batches
	PaymentsV2betaCreditTransferBatchesGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesGetParams) (PaymentsV2betaCreditTransferBatchesGetRes, error)
	// PaymentsV2betaCreditTransferBatchesIDGet implements GET /payments/v2beta/credit-transfer-batches/{id} operation.
	//
	// Retrieve a specific credit transfer batch.
	//
	// GET /payments/v2beta/credit-transfer-batches/{id}
	PaymentsV2betaCreditTransferBatchesIDGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesIDGetParams) (PaymentsV2betaCreditTransferBatchesIDGetRes, error)
	// PaymentsV2betaCreditTransferBatchesIDResultsGet implements GET /payments/v2beta/credit-transfer-batches/{id}/results operation.
	//
	// List processing results for individual credit transfer requests that were part of a credit
	// transfer batch.
	// The result items represent the result(s) of processing the batch request input but doesn't
	// <b>not</b> represent the end-to-end processing of the payment.
	//
	// GET /payments/v2beta/credit-transfer-batches/{id}/results
	PaymentsV2betaCreditTransferBatchesIDResultsGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesIDResultsGetParams) (PaymentsV2betaCreditTransferBatchesIDResultsGetRes, error)
	// PaymentsV2betaCreditTransferBatchesPost implements POST /payments/v2beta/credit-transfer-batches operation.
	//
	// Create a batch of credit transfer payments.
	// See https://docs.atlar.com/docs/batch-payments for more information.
	//
	// POST /payments/v2beta/credit-transfer-batches
	PaymentsV2betaCreditTransferBatchesPost(ctx context.Context, req *PaymentsV2betaCreditTransferBatchesPostReqMultipartFormData) (PaymentsV2betaCreditTransferBatchesPostRes, error)
	// NewError creates *UnexpectedErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *UnexpectedErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
