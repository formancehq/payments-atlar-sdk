// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/get-momo/payments-atlar-sdk/v1/models"
)

// GetV1MandatesReader is a Reader for the GetV1Mandates structure.
type GetV1MandatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MandatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MandatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1MandatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/mandates] GetV1Mandates", response, response.Code())
	}
}

// NewGetV1MandatesOK creates a GetV1MandatesOK with default headers values
func NewGetV1MandatesOK() *GetV1MandatesOK {
	return &GetV1MandatesOK{}
}

/*
GetV1MandatesOK describes a response with status code 200, with default header values.

QueryResponse with list of Mandates
*/
type GetV1MandatesOK struct {
	Payload *GetV1MandatesOKBody
}

// IsSuccess returns true when this get v1 mandates o k response has a 2xx status code
func (o *GetV1MandatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 mandates o k response has a 3xx status code
func (o *GetV1MandatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 mandates o k response has a 4xx status code
func (o *GetV1MandatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 mandates o k response has a 5xx status code
func (o *GetV1MandatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 mandates o k response a status code equal to that given
func (o *GetV1MandatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 mandates o k response
func (o *GetV1MandatesOK) Code() int {
	return 200
}

func (o *GetV1MandatesOK) Error() string {
	return fmt.Sprintf("[GET /v1/mandates][%d] getV1MandatesOK  %+v", 200, o.Payload)
}

func (o *GetV1MandatesOK) String() string {
	return fmt.Sprintf("[GET /v1/mandates][%d] getV1MandatesOK  %+v", 200, o.Payload)
}

func (o *GetV1MandatesOK) GetPayload() *GetV1MandatesOKBody {
	return o.Payload
}

func (o *GetV1MandatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1MandatesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1MandatesBadRequest creates a GetV1MandatesBadRequest with default headers values
func NewGetV1MandatesBadRequest() *GetV1MandatesBadRequest {
	return &GetV1MandatesBadRequest{}
}

/*
GetV1MandatesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetV1MandatesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 mandates bad request response has a 2xx status code
func (o *GetV1MandatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 mandates bad request response has a 3xx status code
func (o *GetV1MandatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 mandates bad request response has a 4xx status code
func (o *GetV1MandatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 mandates bad request response has a 5xx status code
func (o *GetV1MandatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 mandates bad request response a status code equal to that given
func (o *GetV1MandatesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 mandates bad request response
func (o *GetV1MandatesBadRequest) Code() int {
	return 400
}

func (o *GetV1MandatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/mandates][%d] getV1MandatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1MandatesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/mandates][%d] getV1MandatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1MandatesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1MandatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1MandatesOKBody get v1 mandates o k body
swagger:model GetV1MandatesOKBody
*/
type GetV1MandatesOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Mandate `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1MandatesOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1MandatesOKBodyAO0
	var getV1MandatesOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1MandatesOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1MandatesOKBodyAO0

	// GetV1MandatesOKBodyAO1
	var dataGetV1MandatesOKBodyAO1 struct {
		Items []*models.Mandate `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1MandatesOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1MandatesOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1MandatesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1MandatesOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1MandatesOKBodyAO0)
	var dataGetV1MandatesOKBodyAO1 struct {
		Items []*models.Mandate `json:"items"`
	}

	dataGetV1MandatesOKBodyAO1.Items = o.Items

	jsonDataGetV1MandatesOKBodyAO1, errGetV1MandatesOKBodyAO1 := swag.WriteJSON(dataGetV1MandatesOKBodyAO1)
	if errGetV1MandatesOKBodyAO1 != nil {
		return nil, errGetV1MandatesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1MandatesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 mandates o k body
func (o *GetV1MandatesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1MandatesOKBody) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1MandatesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1MandatesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 mandates o k body based on the context it is used
func (o *GetV1MandatesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.QueryResponse
	if err := o.QueryResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1MandatesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1MandatesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1MandatesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1MandatesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1MandatesOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1MandatesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
