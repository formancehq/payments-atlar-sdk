// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Transactions(params *GetV1TransactionsParams, opts ...ClientOption) (*GetV1TransactionsOK, error)

	GetV1TransactionsID(params *GetV1TransactionsIDParams, opts ...ClientOption) (*GetV1TransactionsIDOK, error)

	GetV1TransactionsIDEvents(params *GetV1TransactionsIDEventsParams, opts ...ClientOption) (*GetV1TransactionsIDEventsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetV1Transactions queries transactions

	Returns a list of transactions. For further details on Pagination, see section above.

Sort order is always descending on date
*/
func (a *Client) GetV1Transactions(params *GetV1TransactionsParams, opts ...ClientOption) (*GetV1TransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TransactionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Transactions",
		Method:             "GET",
		PathPattern:        "/v1/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TransactionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Transactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1TransactionsID reads transaction

Returns the identified transaction
*/
func (a *Client) GetV1TransactionsID(params *GetV1TransactionsIDParams, opts ...ClientOption) (*GetV1TransactionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TransactionsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1TransactionsID",
		Method:             "GET",
		PathPattern:        "/v1/transactions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TransactionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TransactionsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1TransactionsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1TransactionsIDEvents queries transaction events

Returns a list of Transaction events. For further details on Pagination, see the section above.
*/
func (a *Client) GetV1TransactionsIDEvents(params *GetV1TransactionsIDEventsParams, opts ...ClientOption) (*GetV1TransactionsIDEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TransactionsIDEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1TransactionsIDEvents",
		Method:             "GET",
		PathPattern:        "/v1/transactions/{id}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1TransactionsIDEventsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TransactionsIDEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1TransactionsIDEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
