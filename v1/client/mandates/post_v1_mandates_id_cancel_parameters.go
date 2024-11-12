// Code generated by go-swagger; DO NOT EDIT.

package mandates

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

// NewPostV1MandatesIDCancelParams creates a new PostV1MandatesIDCancelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1MandatesIDCancelParams() *PostV1MandatesIDCancelParams {
	return &PostV1MandatesIDCancelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1MandatesIDCancelParamsWithTimeout creates a new PostV1MandatesIDCancelParams object
// with the ability to set a timeout on a request.
func NewPostV1MandatesIDCancelParamsWithTimeout(timeout time.Duration) *PostV1MandatesIDCancelParams {
	return &PostV1MandatesIDCancelParams{
		timeout: timeout,
	}
}

// NewPostV1MandatesIDCancelParamsWithContext creates a new PostV1MandatesIDCancelParams object
// with the ability to set a context for a request.
func NewPostV1MandatesIDCancelParamsWithContext(ctx context.Context) *PostV1MandatesIDCancelParams {
	return &PostV1MandatesIDCancelParams{
		Context: ctx,
	}
}

// NewPostV1MandatesIDCancelParamsWithHTTPClient creates a new PostV1MandatesIDCancelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1MandatesIDCancelParamsWithHTTPClient(client *http.Client) *PostV1MandatesIDCancelParams {
	return &PostV1MandatesIDCancelParams{
		HTTPClient: client,
	}
}

/*
PostV1MandatesIDCancelParams contains all the parameters to send to the API endpoint

	for the post v1 mandates ID cancel operation.

	Typically these are written to a http.Request.
*/
type PostV1MandatesIDCancelParams struct {

	/* IfMatch.

	   The ETag of the Mandate
	*/
	IfMatch *string

	/* ID.

	   Mandate ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 mandates ID cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MandatesIDCancelParams) WithDefaults() *PostV1MandatesIDCancelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 mandates ID cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MandatesIDCancelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) WithTimeout(timeout time.Duration) *PostV1MandatesIDCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) WithContext(ctx context.Context) *PostV1MandatesIDCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) WithHTTPClient(client *http.Client) *PostV1MandatesIDCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) WithIfMatch(ifMatch *string) *PostV1MandatesIDCancelParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithID adds the id to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) WithID(id string) *PostV1MandatesIDCancelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post v1 mandates ID cancel params
func (o *PostV1MandatesIDCancelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1MandatesIDCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
