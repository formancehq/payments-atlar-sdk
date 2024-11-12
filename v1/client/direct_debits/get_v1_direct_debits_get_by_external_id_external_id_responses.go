// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// GetV1DirectDebitsGetByExternalIDExternalIDReader is a Reader for the GetV1DirectDebitsGetByExternalIDExternalID structure.
type GetV1DirectDebitsGetByExternalIDExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DirectDebitsGetByExternalIDExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DirectDebitsGetByExternalIDExternalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetV1DirectDebitsGetByExternalIDExternalIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/direct-debits:getByExternalId/{externalId}] GetV1DirectDebitsGetByExternalIDExternalID", response, response.Code())
	}
}

// NewGetV1DirectDebitsGetByExternalIDExternalIDOK creates a GetV1DirectDebitsGetByExternalIDExternalIDOK with default headers values
func NewGetV1DirectDebitsGetByExternalIDExternalIDOK() *GetV1DirectDebitsGetByExternalIDExternalIDOK {
	return &GetV1DirectDebitsGetByExternalIDExternalIDOK{}
}

/*
GetV1DirectDebitsGetByExternalIDExternalIDOK describes a response with status code 200, with default header values.

The identified Direct Debit.
*/
type GetV1DirectDebitsGetByExternalIDExternalIDOK struct {
	Payload *models.DirectDebit
}

// IsSuccess returns true when this get v1 direct debits get by external Id external Id o k response has a 2xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 direct debits get by external Id external Id o k response has a 3xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 direct debits get by external Id external Id o k response has a 4xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 direct debits get by external Id external Id o k response has a 5xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 direct debits get by external Id external Id o k response a status code equal to that given
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 direct debits get by external Id external Id o k response
func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) Code() int {
	return 200
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/direct-debits:getByExternalId/{externalId}][%d] getV1DirectDebitsGetByExternalIdExternalIdOK  %+v", 200, o.Payload)
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) String() string {
	return fmt.Sprintf("[GET /v1/direct-debits:getByExternalId/{externalId}][%d] getV1DirectDebitsGetByExternalIdExternalIdOK  %+v", 200, o.Payload)
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) GetPayload() *models.DirectDebit {
	return o.Payload
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectDebit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1DirectDebitsGetByExternalIDExternalIDNotFound creates a GetV1DirectDebitsGetByExternalIDExternalIDNotFound with default headers values
func NewGetV1DirectDebitsGetByExternalIDExternalIDNotFound() *GetV1DirectDebitsGetByExternalIDExternalIDNotFound {
	return &GetV1DirectDebitsGetByExternalIDExternalIDNotFound{}
}

/*
GetV1DirectDebitsGetByExternalIDExternalIDNotFound describes a response with status code 404, with default header values.

The identified Direct Debit doesn't exist.
*/
type GetV1DirectDebitsGetByExternalIDExternalIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this get v1 direct debits get by external Id external Id not found response has a 2xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 direct debits get by external Id external Id not found response has a 3xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 direct debits get by external Id external Id not found response has a 4xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 direct debits get by external Id external Id not found response has a 5xx status code
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 direct debits get by external Id external Id not found response a status code equal to that given
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 direct debits get by external Id external Id not found response
func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) Code() int {
	return 404
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/direct-debits:getByExternalId/{externalId}][%d] getV1DirectDebitsGetByExternalIdExternalIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/direct-debits:getByExternalId/{externalId}][%d] getV1DirectDebitsGetByExternalIdExternalIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetV1DirectDebitsGetByExternalIDExternalIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
