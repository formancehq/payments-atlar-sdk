// Code generated by go-swagger; DO NOT EDIT.

package external_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/get-momo/payments-atlar-sdk/v1/models"
)

// DeleteV1ExternalAccountsIDReader is a Reader for the DeleteV1ExternalAccountsID structure.
type DeleteV1ExternalAccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ExternalAccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1ExternalAccountsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteV1ExternalAccountsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/external-accounts/{id}] DeleteV1ExternalAccountsID", response, response.Code())
	}
}

// NewDeleteV1ExternalAccountsIDNoContent creates a DeleteV1ExternalAccountsIDNoContent with default headers values
func NewDeleteV1ExternalAccountsIDNoContent() *DeleteV1ExternalAccountsIDNoContent {
	return &DeleteV1ExternalAccountsIDNoContent{}
}

/*
DeleteV1ExternalAccountsIDNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteV1ExternalAccountsIDNoContent struct {
}

// IsSuccess returns true when this delete v1 external accounts Id no content response has a 2xx status code
func (o *DeleteV1ExternalAccountsIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 external accounts Id no content response has a 3xx status code
func (o *DeleteV1ExternalAccountsIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 external accounts Id no content response has a 4xx status code
func (o *DeleteV1ExternalAccountsIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 external accounts Id no content response has a 5xx status code
func (o *DeleteV1ExternalAccountsIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 external accounts Id no content response a status code equal to that given
func (o *DeleteV1ExternalAccountsIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete v1 external accounts Id no content response
func (o *DeleteV1ExternalAccountsIDNoContent) Code() int {
	return 204
}

func (o *DeleteV1ExternalAccountsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/external-accounts/{id}][%d] deleteV1ExternalAccountsIdNoContent ", 204)
}

func (o *DeleteV1ExternalAccountsIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/external-accounts/{id}][%d] deleteV1ExternalAccountsIdNoContent ", 204)
}

func (o *DeleteV1ExternalAccountsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1ExternalAccountsIDNotFound creates a DeleteV1ExternalAccountsIDNotFound with default headers values
func NewDeleteV1ExternalAccountsIDNotFound() *DeleteV1ExternalAccountsIDNotFound {
	return &DeleteV1ExternalAccountsIDNotFound{}
}

/*
DeleteV1ExternalAccountsIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found
*/
type DeleteV1ExternalAccountsIDNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete v1 external accounts Id not found response has a 2xx status code
func (o *DeleteV1ExternalAccountsIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 external accounts Id not found response has a 3xx status code
func (o *DeleteV1ExternalAccountsIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 external accounts Id not found response has a 4xx status code
func (o *DeleteV1ExternalAccountsIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 external accounts Id not found response has a 5xx status code
func (o *DeleteV1ExternalAccountsIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 external accounts Id not found response a status code equal to that given
func (o *DeleteV1ExternalAccountsIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete v1 external accounts Id not found response
func (o *DeleteV1ExternalAccountsIDNotFound) Code() int {
	return 404
}

func (o *DeleteV1ExternalAccountsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/external-accounts/{id}][%d] deleteV1ExternalAccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV1ExternalAccountsIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/external-accounts/{id}][%d] deleteV1ExternalAccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV1ExternalAccountsIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteV1ExternalAccountsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
