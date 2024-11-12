// Code generated by go-swagger; DO NOT EDIT.

package counterparties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// PostV1CounterpartiesReader is a Reader for the PostV1Counterparties structure.
type PostV1CounterpartiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1CounterpartiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1CounterpartiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1CounterpartiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/counterparties] PostV1Counterparties", response, response.Code())
	}
}

// NewPostV1CounterpartiesCreated creates a PostV1CounterpartiesCreated with default headers values
func NewPostV1CounterpartiesCreated() *PostV1CounterpartiesCreated {
	return &PostV1CounterpartiesCreated{}
}

/*
PostV1CounterpartiesCreated describes a response with status code 201, with default header values.

The created counterparty
*/
type PostV1CounterpartiesCreated struct {
	Payload *models.Counterparty
}

// IsSuccess returns true when this post v1 counterparties created response has a 2xx status code
func (o *PostV1CounterpartiesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 counterparties created response has a 3xx status code
func (o *PostV1CounterpartiesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 counterparties created response has a 4xx status code
func (o *PostV1CounterpartiesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 counterparties created response has a 5xx status code
func (o *PostV1CounterpartiesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 counterparties created response a status code equal to that given
func (o *PostV1CounterpartiesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 counterparties created response
func (o *PostV1CounterpartiesCreated) Code() int {
	return 201
}

func (o *PostV1CounterpartiesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/counterparties][%d] postV1CounterpartiesCreated  %+v", 201, o.Payload)
}

func (o *PostV1CounterpartiesCreated) String() string {
	return fmt.Sprintf("[POST /v1/counterparties][%d] postV1CounterpartiesCreated  %+v", 201, o.Payload)
}

func (o *PostV1CounterpartiesCreated) GetPayload() *models.Counterparty {
	return o.Payload
}

func (o *PostV1CounterpartiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Counterparty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1CounterpartiesBadRequest creates a PostV1CounterpartiesBadRequest with default headers values
func NewPostV1CounterpartiesBadRequest() *PostV1CounterpartiesBadRequest {
	return &PostV1CounterpartiesBadRequest{}
}

/*
PostV1CounterpartiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostV1CounterpartiesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this post v1 counterparties bad request response has a 2xx status code
func (o *PostV1CounterpartiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 counterparties bad request response has a 3xx status code
func (o *PostV1CounterpartiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 counterparties bad request response has a 4xx status code
func (o *PostV1CounterpartiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 counterparties bad request response has a 5xx status code
func (o *PostV1CounterpartiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 counterparties bad request response a status code equal to that given
func (o *PostV1CounterpartiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 counterparties bad request response
func (o *PostV1CounterpartiesBadRequest) Code() int {
	return 400
}

func (o *PostV1CounterpartiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/counterparties][%d] postV1CounterpartiesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1CounterpartiesBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/counterparties][%d] postV1CounterpartiesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1CounterpartiesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostV1CounterpartiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
