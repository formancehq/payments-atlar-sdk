// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// PutV1AccountsIDReader is a Reader for the PutV1AccountsID structure.
type PutV1AccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1AccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1AccountsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1AccountsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutV1AccountsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/accounts/{id}] PutV1AccountsID", response, response.Code())
	}
}

// NewPutV1AccountsIDOK creates a PutV1AccountsIDOK with default headers values
func NewPutV1AccountsIDOK() *PutV1AccountsIDOK {
	return &PutV1AccountsIDOK{}
}

/*
PutV1AccountsIDOK describes a response with status code 200, with default header values.

The identified Account.
*/
type PutV1AccountsIDOK struct {
	Payload *models.Account
}

// IsSuccess returns true when this put v1 accounts Id o k response has a 2xx status code
func (o *PutV1AccountsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 accounts Id o k response has a 3xx status code
func (o *PutV1AccountsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 accounts Id o k response has a 4xx status code
func (o *PutV1AccountsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 accounts Id o k response has a 5xx status code
func (o *PutV1AccountsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 accounts Id o k response a status code equal to that given
func (o *PutV1AccountsIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 accounts Id o k response
func (o *PutV1AccountsIDOK) Code() int {
	return 200
}

func (o *PutV1AccountsIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdOK  %+v", 200, o.Payload)
}

func (o *PutV1AccountsIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdOK  %+v", 200, o.Payload)
}

func (o *PutV1AccountsIDOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *PutV1AccountsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AccountsIDBadRequest creates a PutV1AccountsIDBadRequest with default headers values
func NewPutV1AccountsIDBadRequest() *PutV1AccountsIDBadRequest {
	return &PutV1AccountsIDBadRequest{}
}

/*
PutV1AccountsIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutV1AccountsIDBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this put v1 accounts Id bad request response has a 2xx status code
func (o *PutV1AccountsIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 accounts Id bad request response has a 3xx status code
func (o *PutV1AccountsIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 accounts Id bad request response has a 4xx status code
func (o *PutV1AccountsIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 accounts Id bad request response has a 5xx status code
func (o *PutV1AccountsIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 accounts Id bad request response a status code equal to that given
func (o *PutV1AccountsIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put v1 accounts Id bad request response
func (o *PutV1AccountsIDBadRequest) Code() int {
	return 400
}

func (o *PutV1AccountsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1AccountsIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1AccountsIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutV1AccountsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AccountsIDNotFound creates a PutV1AccountsIDNotFound with default headers values
func NewPutV1AccountsIDNotFound() *PutV1AccountsIDNotFound {
	return &PutV1AccountsIDNotFound{}
}

/*
PutV1AccountsIDNotFound describes a response with status code 404, with default header values.

The identified Account doesn't exist.
*/
type PutV1AccountsIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this put v1 accounts Id not found response has a 2xx status code
func (o *PutV1AccountsIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 accounts Id not found response has a 3xx status code
func (o *PutV1AccountsIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 accounts Id not found response has a 4xx status code
func (o *PutV1AccountsIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 accounts Id not found response has a 5xx status code
func (o *PutV1AccountsIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 accounts Id not found response a status code equal to that given
func (o *PutV1AccountsIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put v1 accounts Id not found response
func (o *PutV1AccountsIDNotFound) Code() int {
	return 404
}

func (o *PutV1AccountsIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *PutV1AccountsIDNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/accounts/{id}][%d] putV1AccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *PutV1AccountsIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *PutV1AccountsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}