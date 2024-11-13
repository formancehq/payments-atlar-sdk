// Code generated by go-swagger; DO NOT EDIT.

package third_parties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new third parties API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for third parties API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1betaThirdParties(params *GetV1betaThirdPartiesParams, opts ...ClientOption) (*GetV1betaThirdPartiesOK, error)

	GetV1betaThirdPartiesID(params *GetV1betaThirdPartiesIDParams, opts ...ClientOption) (*GetV1betaThirdPartiesIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1betaThirdParties queries third parties

Returns a list of ThirdParties. For further details on Pagination, see general section above.
*/
func (a *Client) GetV1betaThirdParties(params *GetV1betaThirdPartiesParams, opts ...ClientOption) (*GetV1betaThirdPartiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1betaThirdPartiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1betaThirdParties",
		Method:             "GET",
		PathPattern:        "/v1beta/third-parties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1betaThirdPartiesReader{formats: a.formats},
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
	success, ok := result.(*GetV1betaThirdPartiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1betaThirdParties: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1betaThirdPartiesID reads third party

Returns the identified ThirdParty.
*/
func (a *Client) GetV1betaThirdPartiesID(params *GetV1betaThirdPartiesIDParams, opts ...ClientOption) (*GetV1betaThirdPartiesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1betaThirdPartiesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1betaThirdPartiesID",
		Method:             "GET",
		PathPattern:        "/v1beta/third-parties/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1betaThirdPartiesIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1betaThirdPartiesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1betaThirdPartiesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
