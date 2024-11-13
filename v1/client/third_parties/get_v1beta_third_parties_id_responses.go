// Code generated by go-swagger; DO NOT EDIT.

package third_parties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/get-momo/payments-atlar-sdk/v1/models"
)

// GetV1betaThirdPartiesIDReader is a Reader for the GetV1betaThirdPartiesID structure.
type GetV1betaThirdPartiesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1betaThirdPartiesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1betaThirdPartiesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetV1betaThirdPartiesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1beta/third-parties/{id}] GetV1betaThirdPartiesID", response, response.Code())
	}
}

// NewGetV1betaThirdPartiesIDOK creates a GetV1betaThirdPartiesIDOK with default headers values
func NewGetV1betaThirdPartiesIDOK() *GetV1betaThirdPartiesIDOK {
	return &GetV1betaThirdPartiesIDOK{}
}

/*
GetV1betaThirdPartiesIDOK describes a response with status code 200, with default header values.

The identified third party.
*/
type GetV1betaThirdPartiesIDOK struct {
	Payload *models.ThirdParty
}

// IsSuccess returns true when this get v1beta third parties Id o k response has a 2xx status code
func (o *GetV1betaThirdPartiesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1beta third parties Id o k response has a 3xx status code
func (o *GetV1betaThirdPartiesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1beta third parties Id o k response has a 4xx status code
func (o *GetV1betaThirdPartiesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1beta third parties Id o k response has a 5xx status code
func (o *GetV1betaThirdPartiesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1beta third parties Id o k response a status code equal to that given
func (o *GetV1betaThirdPartiesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1beta third parties Id o k response
func (o *GetV1betaThirdPartiesIDOK) Code() int {
	return 200
}

func (o *GetV1betaThirdPartiesIDOK) Error() string {
	return fmt.Sprintf("[GET /v1beta/third-parties/{id}][%d] getV1betaThirdPartiesIdOK  %+v", 200, o.Payload)
}

func (o *GetV1betaThirdPartiesIDOK) String() string {
	return fmt.Sprintf("[GET /v1beta/third-parties/{id}][%d] getV1betaThirdPartiesIdOK  %+v", 200, o.Payload)
}

func (o *GetV1betaThirdPartiesIDOK) GetPayload() *models.ThirdParty {
	return o.Payload
}

func (o *GetV1betaThirdPartiesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThirdParty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1betaThirdPartiesIDNotFound creates a GetV1betaThirdPartiesIDNotFound with default headers values
func NewGetV1betaThirdPartiesIDNotFound() *GetV1betaThirdPartiesIDNotFound {
	return &GetV1betaThirdPartiesIDNotFound{}
}

/*
GetV1betaThirdPartiesIDNotFound describes a response with status code 404, with default header values.

The identified third party doesn't exist
*/
type GetV1betaThirdPartiesIDNotFound struct {
	Payload string
}

// IsSuccess returns true when this get v1beta third parties Id not found response has a 2xx status code
func (o *GetV1betaThirdPartiesIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1beta third parties Id not found response has a 3xx status code
func (o *GetV1betaThirdPartiesIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1beta third parties Id not found response has a 4xx status code
func (o *GetV1betaThirdPartiesIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1beta third parties Id not found response has a 5xx status code
func (o *GetV1betaThirdPartiesIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1beta third parties Id not found response a status code equal to that given
func (o *GetV1betaThirdPartiesIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1beta third parties Id not found response
func (o *GetV1betaThirdPartiesIDNotFound) Code() int {
	return 404
}

func (o *GetV1betaThirdPartiesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1beta/third-parties/{id}][%d] getV1betaThirdPartiesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1betaThirdPartiesIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1beta/third-parties/{id}][%d] getV1betaThirdPartiesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1betaThirdPartiesIDNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetV1betaThirdPartiesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
