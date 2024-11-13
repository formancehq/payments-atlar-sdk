// Code generated by go-swagger; DO NOT EDIT.

package counterparties

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

// GetV1CounterpartiesReader is a Reader for the GetV1Counterparties structure.
type GetV1CounterpartiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CounterpartiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CounterpartiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1CounterpartiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/counterparties] GetV1Counterparties", response, response.Code())
	}
}

// NewGetV1CounterpartiesOK creates a GetV1CounterpartiesOK with default headers values
func NewGetV1CounterpartiesOK() *GetV1CounterpartiesOK {
	return &GetV1CounterpartiesOK{}
}

/*
GetV1CounterpartiesOK describes a response with status code 200, with default header values.

QueryResponse with list of Counterparties
*/
type GetV1CounterpartiesOK struct {
	Payload *GetV1CounterpartiesOKBody
}

// IsSuccess returns true when this get v1 counterparties o k response has a 2xx status code
func (o *GetV1CounterpartiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 counterparties o k response has a 3xx status code
func (o *GetV1CounterpartiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 counterparties o k response has a 4xx status code
func (o *GetV1CounterpartiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 counterparties o k response has a 5xx status code
func (o *GetV1CounterpartiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 counterparties o k response a status code equal to that given
func (o *GetV1CounterpartiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 counterparties o k response
func (o *GetV1CounterpartiesOK) Code() int {
	return 200
}

func (o *GetV1CounterpartiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/counterparties][%d] getV1CounterpartiesOK  %+v", 200, o.Payload)
}

func (o *GetV1CounterpartiesOK) String() string {
	return fmt.Sprintf("[GET /v1/counterparties][%d] getV1CounterpartiesOK  %+v", 200, o.Payload)
}

func (o *GetV1CounterpartiesOK) GetPayload() *GetV1CounterpartiesOKBody {
	return o.Payload
}

func (o *GetV1CounterpartiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1CounterpartiesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1CounterpartiesBadRequest creates a GetV1CounterpartiesBadRequest with default headers values
func NewGetV1CounterpartiesBadRequest() *GetV1CounterpartiesBadRequest {
	return &GetV1CounterpartiesBadRequest{}
}

/*
GetV1CounterpartiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetV1CounterpartiesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 counterparties bad request response has a 2xx status code
func (o *GetV1CounterpartiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 counterparties bad request response has a 3xx status code
func (o *GetV1CounterpartiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 counterparties bad request response has a 4xx status code
func (o *GetV1CounterpartiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 counterparties bad request response has a 5xx status code
func (o *GetV1CounterpartiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 counterparties bad request response a status code equal to that given
func (o *GetV1CounterpartiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 counterparties bad request response
func (o *GetV1CounterpartiesBadRequest) Code() int {
	return 400
}

func (o *GetV1CounterpartiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/counterparties][%d] getV1CounterpartiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1CounterpartiesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/counterparties][%d] getV1CounterpartiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1CounterpartiesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1CounterpartiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1CounterpartiesOKBody get v1 counterparties o k body
swagger:model GetV1CounterpartiesOKBody
*/
type GetV1CounterpartiesOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Counterparty `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1CounterpartiesOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1CounterpartiesOKBodyAO0
	var getV1CounterpartiesOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1CounterpartiesOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1CounterpartiesOKBodyAO0

	// GetV1CounterpartiesOKBodyAO1
	var dataGetV1CounterpartiesOKBodyAO1 struct {
		Items []*models.Counterparty `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1CounterpartiesOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1CounterpartiesOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1CounterpartiesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1CounterpartiesOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1CounterpartiesOKBodyAO0)
	var dataGetV1CounterpartiesOKBodyAO1 struct {
		Items []*models.Counterparty `json:"items"`
	}

	dataGetV1CounterpartiesOKBodyAO1.Items = o.Items

	jsonDataGetV1CounterpartiesOKBodyAO1, errGetV1CounterpartiesOKBodyAO1 := swag.WriteJSON(dataGetV1CounterpartiesOKBodyAO1)
	if errGetV1CounterpartiesOKBodyAO1 != nil {
		return nil, errGetV1CounterpartiesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1CounterpartiesOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 counterparties o k body
func (o *GetV1CounterpartiesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1CounterpartiesOKBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1CounterpartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1CounterpartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 counterparties o k body based on the context it is used
func (o *GetV1CounterpartiesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1CounterpartiesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1CounterpartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1CounterpartiesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1CounterpartiesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1CounterpartiesOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1CounterpartiesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
