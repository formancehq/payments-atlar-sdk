// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

	"github.com/get-momo/payments-atlar-sdk/v1/models"
)

// NewPostV1WebhooksParams creates a new PostV1WebhooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1WebhooksParams() *PostV1WebhooksParams {
	return &PostV1WebhooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1WebhooksParamsWithTimeout creates a new PostV1WebhooksParams object
// with the ability to set a timeout on a request.
func NewPostV1WebhooksParamsWithTimeout(timeout time.Duration) *PostV1WebhooksParams {
	return &PostV1WebhooksParams{
		timeout: timeout,
	}
}

// NewPostV1WebhooksParamsWithContext creates a new PostV1WebhooksParams object
// with the ability to set a context for a request.
func NewPostV1WebhooksParamsWithContext(ctx context.Context) *PostV1WebhooksParams {
	return &PostV1WebhooksParams{
		Context: ctx,
	}
}

// NewPostV1WebhooksParamsWithHTTPClient creates a new PostV1WebhooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1WebhooksParamsWithHTTPClient(client *http.Client) *PostV1WebhooksParams {
	return &PostV1WebhooksParams{
		HTTPClient: client,
	}
}

/*
PostV1WebhooksParams contains all the parameters to send to the API endpoint

	for the post v1 webhooks operation.

	Typically these are written to a http.Request.
*/
type PostV1WebhooksParams struct {

	/* Webhook.

	   CreateWebhookRequest
	*/
	Webhook *models.CreateWebhookRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1WebhooksParams) WithDefaults() *PostV1WebhooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 webhooks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1WebhooksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 webhooks params
func (o *PostV1WebhooksParams) WithTimeout(timeout time.Duration) *PostV1WebhooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 webhooks params
func (o *PostV1WebhooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 webhooks params
func (o *PostV1WebhooksParams) WithContext(ctx context.Context) *PostV1WebhooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 webhooks params
func (o *PostV1WebhooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 webhooks params
func (o *PostV1WebhooksParams) WithHTTPClient(client *http.Client) *PostV1WebhooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 webhooks params
func (o *PostV1WebhooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhook adds the webhook to the post v1 webhooks params
func (o *PostV1WebhooksParams) WithWebhook(webhook *models.CreateWebhookRequest) *PostV1WebhooksParams {
	o.SetWebhook(webhook)
	return o
}

// SetWebhook adds the webhook to the post v1 webhooks params
func (o *PostV1WebhooksParams) SetWebhook(webhook *models.CreateWebhookRequest) {
	o.Webhook = webhook
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1WebhooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Webhook != nil {
		if err := r.SetBodyParam(o.Webhook); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
