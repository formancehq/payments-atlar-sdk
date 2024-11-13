// Code generated by go-swagger; DO NOT EDIT.

package external_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteV1ExternalAccountsIDParams creates a new DeleteV1ExternalAccountsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ExternalAccountsIDParams() *DeleteV1ExternalAccountsIDParams {
	return &DeleteV1ExternalAccountsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ExternalAccountsIDParamsWithTimeout creates a new DeleteV1ExternalAccountsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ExternalAccountsIDParamsWithTimeout(timeout time.Duration) *DeleteV1ExternalAccountsIDParams {
	return &DeleteV1ExternalAccountsIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ExternalAccountsIDParamsWithContext creates a new DeleteV1ExternalAccountsIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ExternalAccountsIDParamsWithContext(ctx context.Context) *DeleteV1ExternalAccountsIDParams {
	return &DeleteV1ExternalAccountsIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ExternalAccountsIDParamsWithHTTPClient creates a new DeleteV1ExternalAccountsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ExternalAccountsIDParamsWithHTTPClient(client *http.Client) *DeleteV1ExternalAccountsIDParams {
	return &DeleteV1ExternalAccountsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ExternalAccountsIDParams contains all the parameters to send to the API endpoint

	for the delete v1 external accounts ID operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ExternalAccountsIDParams struct {

	/* ID.

	   External account ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 external accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ExternalAccountsIDParams) WithDefaults() *DeleteV1ExternalAccountsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 external accounts ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ExternalAccountsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) WithTimeout(timeout time.Duration) *DeleteV1ExternalAccountsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) WithContext(ctx context.Context) *DeleteV1ExternalAccountsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) WithHTTPClient(client *http.Client) *DeleteV1ExternalAccountsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) WithID(id string) *DeleteV1ExternalAccountsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 external accounts ID params
func (o *DeleteV1ExternalAccountsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ExternalAccountsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
