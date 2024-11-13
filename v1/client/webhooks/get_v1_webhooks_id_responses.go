// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// GetV1WebhooksIDReader is a Reader for the GetV1WebhooksID structure.
type GetV1WebhooksIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WebhooksIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WebhooksIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetV1WebhooksIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/webhooks/{id}] GetV1WebhooksID", response, response.Code())
	}
}

// NewGetV1WebhooksIDOK creates a GetV1WebhooksIDOK with default headers values
func NewGetV1WebhooksIDOK() *GetV1WebhooksIDOK {
	return &GetV1WebhooksIDOK{}
}

/*
GetV1WebhooksIDOK describes a response with status code 200, with default header values.

The identified webhook
*/
type GetV1WebhooksIDOK struct {
	Payload *models.Webhook
}

// IsSuccess returns true when this get v1 webhooks Id o k response has a 2xx status code
func (o *GetV1WebhooksIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 webhooks Id o k response has a 3xx status code
func (o *GetV1WebhooksIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 webhooks Id o k response has a 4xx status code
func (o *GetV1WebhooksIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 webhooks Id o k response has a 5xx status code
func (o *GetV1WebhooksIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 webhooks Id o k response a status code equal to that given
func (o *GetV1WebhooksIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 webhooks Id o k response
func (o *GetV1WebhooksIDOK) Code() int {
	return 200
}

func (o *GetV1WebhooksIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/webhooks/{id}][%d] getV1WebhooksIdOK  %+v", 200, o.Payload)
}

func (o *GetV1WebhooksIDOK) String() string {
	return fmt.Sprintf("[GET /v1/webhooks/{id}][%d] getV1WebhooksIdOK  %+v", 200, o.Payload)
}

func (o *GetV1WebhooksIDOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *GetV1WebhooksIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1WebhooksIDNotFound creates a GetV1WebhooksIDNotFound with default headers values
func NewGetV1WebhooksIDNotFound() *GetV1WebhooksIDNotFound {
	return &GetV1WebhooksIDNotFound{}
}

/*
GetV1WebhooksIDNotFound describes a response with status code 404, with default header values.

The identified webhook doesn't exist
*/
type GetV1WebhooksIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this get v1 webhooks Id not found response has a 2xx status code
func (o *GetV1WebhooksIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 webhooks Id not found response has a 3xx status code
func (o *GetV1WebhooksIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 webhooks Id not found response has a 4xx status code
func (o *GetV1WebhooksIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 webhooks Id not found response has a 5xx status code
func (o *GetV1WebhooksIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 webhooks Id not found response a status code equal to that given
func (o *GetV1WebhooksIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 webhooks Id not found response
func (o *GetV1WebhooksIDNotFound) Code() int {
	return 404
}

func (o *GetV1WebhooksIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/webhooks/{id}][%d] getV1WebhooksIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1WebhooksIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/webhooks/{id}][%d] getV1WebhooksIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1WebhooksIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetV1WebhooksIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
