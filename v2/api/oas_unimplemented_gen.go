// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// FinancialDataV2AccountsIDBalancesGet implements GET /financial-data/v2/accounts/{id}/balances operation.
//
// Returns a list of account balances. For further details on Pagination, see general section above.
//
// GET /financial-data/v2/accounts/{id}/balances
func (UnimplementedHandler) FinancialDataV2AccountsIDBalancesGet(ctx context.Context, params FinancialDataV2AccountsIDBalancesGetParams) (r FinancialDataV2AccountsIDBalancesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FinancialDataV2EntitiesGet implements GET /financial-data/v2/entities operation.
//
// Returns a list of entities. For further details on Pagination, see general section above.
//
// GET /financial-data/v2/entities
func (UnimplementedHandler) FinancialDataV2EntitiesGet(ctx context.Context, params FinancialDataV2EntitiesGetParams) (r FinancialDataV2EntitiesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FinancialDataV2EntitiesIDDelete implements DELETE /financial-data/v2/entities/{id} operation.
//
// Deletes the specified entity. It is only possible to delete an entity if it has no children or has
// no associated accounts.
//
// DELETE /financial-data/v2/entities/{id}
func (UnimplementedHandler) FinancialDataV2EntitiesIDDelete(ctx context.Context, params FinancialDataV2EntitiesIDDeleteParams) (r FinancialDataV2EntitiesIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FinancialDataV2EntitiesIDGet implements GET /financial-data/v2/entities/{id} operation.
//
// Retrieve a specific entity.
//
// GET /financial-data/v2/entities/{id}
func (UnimplementedHandler) FinancialDataV2EntitiesIDGet(ctx context.Context, params FinancialDataV2EntitiesIDGetParams) (r FinancialDataV2EntitiesIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FinancialDataV2EntitiesPost implements POST /financial-data/v2/entities operation.
//
// Create entity.
//
// POST /financial-data/v2/entities
func (UnimplementedHandler) FinancialDataV2EntitiesPost(ctx context.Context, req *CreateEntityRequest) (r FinancialDataV2EntitiesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FinancialDataV2PendingTransactionsIDEventsGet implements GET /financial-data/v2/pending-transactions/{id}/events operation.
//
// Retrieve events for a specific pending transaction. Returns a list of events. For further details
// on Pagination, see the section above.
//
// GET /financial-data/v2/pending-transactions/{id}/events
func (UnimplementedHandler) FinancialDataV2PendingTransactionsIDEventsGet(ctx context.Context, params FinancialDataV2PendingTransactionsIDEventsGetParams) (r FinancialDataV2PendingTransactionsIDEventsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

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
func (UnimplementedHandler) IamV2betaOAuth2TokenPost(ctx context.Context, req *IamV2betaOAuth2TokenPostReq) (r IamV2betaOAuth2TokenPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2CounterpartiesIDDelete implements DELETE /payments/v2/counterparties/{id} operation.
//
// Deletes the specified counterparty.
//
// DELETE /payments/v2/counterparties/{id}
func (UnimplementedHandler) PaymentsV2CounterpartiesIDDelete(ctx context.Context, params PaymentsV2CounterpartiesIDDeleteParams) (r PaymentsV2CounterpartiesIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2CreditTransferBatchesIDApprovePost implements POST /payments/v2/credit-transfer-batches/{id}:approve operation.
//
// Approve the batch of credit transfers.
// Approval of a batch is only applicable when the batch `treatment` is `BATCH`.
//
// POST /payments/v2/credit-transfer-batches/{id}:approve
func (UnimplementedHandler) PaymentsV2CreditTransferBatchesIDApprovePost(ctx context.Context, req *ApproveBatchRequest, params PaymentsV2CreditTransferBatchesIDApprovePostParams) (r PaymentsV2CreditTransferBatchesIDApprovePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2CreditTransferBatchesIDRejectPost implements POST /payments/v2/credit-transfer-batches/{id}:reject operation.
//
// Reject the batch of credit transfers.
// Rejection of a batch is only applicable when the batch `treatment` is `BATCH`.
//
// POST /payments/v2/credit-transfer-batches/{id}:reject
func (UnimplementedHandler) PaymentsV2CreditTransferBatchesIDRejectPost(ctx context.Context, req *RejectPaymentRequest, params PaymentsV2CreditTransferBatchesIDRejectPostParams) (r PaymentsV2CreditTransferBatchesIDRejectPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2CreditTransfersIDEventsGet implements GET /payments/v2/credit-transfers/{id}/events operation.
//
// Retrieve events for a specific credit transfer. Returns a list of events. For further details on
// Pagination, see the section above.
//
// GET /payments/v2/credit-transfers/{id}/events
func (UnimplementedHandler) PaymentsV2CreditTransfersIDEventsGet(ctx context.Context, params PaymentsV2CreditTransfersIDEventsGetParams) (r PaymentsV2CreditTransfersIDEventsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2DirectDebitsIDEventsGet implements GET /payments/v2/direct-debits/{id}/events operation.
//
// Retrieve events for a specific direct debit. Returns a list of events. For further details on
// Pagination, see the section above.
//
// GET /payments/v2/direct-debits/{id}/events
func (UnimplementedHandler) PaymentsV2DirectDebitsIDEventsGet(ctx context.Context, params PaymentsV2DirectDebitsIDEventsGetParams) (r PaymentsV2DirectDebitsIDEventsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2ExternalAccountsIDDelete implements DELETE /payments/v2/external-accounts/{id} operation.
//
// Deletes the specified external account.
//
// DELETE /payments/v2/external-accounts/{id}
func (UnimplementedHandler) PaymentsV2ExternalAccountsIDDelete(ctx context.Context, params PaymentsV2ExternalAccountsIDDeleteParams) (r PaymentsV2ExternalAccountsIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2MandatesIDEventsGet implements GET /payments/v2/mandates/{id}/events operation.
//
// Retrieve events for a specific mandate. Returns a list of events. For further details on
// Pagination, see the section above.
//
// GET /payments/v2/mandates/{id}/events
func (UnimplementedHandler) PaymentsV2MandatesIDEventsGet(ctx context.Context, params PaymentsV2MandatesIDEventsGetParams) (r PaymentsV2MandatesIDEventsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2betaCreditTransferBatchesGet implements GET /payments/v2beta/credit-transfer-batches operation.
//
// Returns a list of credit transfer batches.
//
// GET /payments/v2beta/credit-transfer-batches
func (UnimplementedHandler) PaymentsV2betaCreditTransferBatchesGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesGetParams) (r PaymentsV2betaCreditTransferBatchesGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2betaCreditTransferBatchesIDGet implements GET /payments/v2beta/credit-transfer-batches/{id} operation.
//
// Retrieve a specific credit transfer batch.
//
// GET /payments/v2beta/credit-transfer-batches/{id}
func (UnimplementedHandler) PaymentsV2betaCreditTransferBatchesIDGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesIDGetParams) (r PaymentsV2betaCreditTransferBatchesIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2betaCreditTransferBatchesIDResultsGet implements GET /payments/v2beta/credit-transfer-batches/{id}/results operation.
//
// List processing results for individual credit transfer requests that were part of a credit
// transfer batch.
// The result items represent the result(s) of processing the batch request input but doesn't
// <b>not</b> represent the end-to-end processing of the payment.
//
// GET /payments/v2beta/credit-transfer-batches/{id}/results
func (UnimplementedHandler) PaymentsV2betaCreditTransferBatchesIDResultsGet(ctx context.Context, params PaymentsV2betaCreditTransferBatchesIDResultsGetParams) (r PaymentsV2betaCreditTransferBatchesIDResultsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PaymentsV2betaCreditTransferBatchesPost implements POST /payments/v2beta/credit-transfer-batches operation.
//
// Create a batch of credit transfer payments.
// See https://docs.atlar.com/docs/batch-payments for more information.
//
// POST /payments/v2beta/credit-transfer-batches
func (UnimplementedHandler) PaymentsV2betaCreditTransferBatchesPost(ctx context.Context, req *PaymentsV2betaCreditTransferBatchesPostReqMultipartFormData) (r PaymentsV2betaCreditTransferBatchesPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *UnexpectedErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *UnexpectedErrorStatusCode) {
	r = new(UnexpectedErrorStatusCode)
	return r
}
