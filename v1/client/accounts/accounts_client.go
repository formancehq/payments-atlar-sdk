// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new accounts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accounts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Accounts(params *GetV1AccountsParams, opts ...ClientOption) (*GetV1AccountsOK, error)

	GetV1AccountsID(params *GetV1AccountsIDParams, opts ...ClientOption) (*GetV1AccountsIDOK, error)

	PutV1AccountsID(params *PutV1AccountsIDParams, opts ...ClientOption) (*PutV1AccountsIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1Accounts queries accounts

Returns a list of accounts. For further details on Pagination, see section above.
*/
func (a *Client) GetV1Accounts(params *GetV1AccountsParams, opts ...ClientOption) (*GetV1AccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Accounts",
		Method:             "GET",
		PathPattern:        "/v1/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AccountsReader{formats: a.formats},
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
	success, ok := result.(*GetV1AccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Accounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AccountsID reads account

Returns the identified account
*/
func (a *Client) GetV1AccountsID(params *GetV1AccountsIDParams, opts ...ClientOption) (*GetV1AccountsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AccountsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1AccountsID",
		Method:             "GET",
		PathPattern:        "/v1/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AccountsIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1AccountsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1AccountsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutV1AccountsID modifies account

Returns the identified account
*/
func (a *Client) PutV1AccountsID(params *PutV1AccountsIDParams, opts ...ClientOption) (*PutV1AccountsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1AccountsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutV1AccountsID",
		Method:             "PUT",
		PathPattern:        "/v1/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1AccountsIDReader{formats: a.formats},
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
	success, ok := result.(*PutV1AccountsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutV1AccountsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
