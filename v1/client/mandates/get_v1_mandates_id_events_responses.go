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

	"github.com/formancehq/payments-atlar-sdk/v1/models"
)

// GetV1MandatesIDEventsReader is a Reader for the GetV1MandatesIDEvents structure.
type GetV1MandatesIDEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MandatesIDEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MandatesIDEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/mandates/{id}/events] GetV1MandatesIDEvents", response, response.Code())
	}
}

// NewGetV1MandatesIDEventsOK creates a GetV1MandatesIDEventsOK with default headers values
func NewGetV1MandatesIDEventsOK() *GetV1MandatesIDEventsOK {
	return &GetV1MandatesIDEventsOK{}
}

/*
GetV1MandatesIDEventsOK describes a response with status code 200, with default header values.

QueryResponse with a list of Mandate events
*/
type GetV1MandatesIDEventsOK struct {
	Payload *GetV1MandatesIDEventsOKBody
}

// IsSuccess returns true when this get v1 mandates Id events o k response has a 2xx status code
func (o *GetV1MandatesIDEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 mandates Id events o k response has a 3xx status code
func (o *GetV1MandatesIDEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 mandates Id events o k response has a 4xx status code
func (o *GetV1MandatesIDEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 mandates Id events o k response has a 5xx status code
func (o *GetV1MandatesIDEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 mandates Id events o k response a status code equal to that given
func (o *GetV1MandatesIDEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 mandates Id events o k response
func (o *GetV1MandatesIDEventsOK) Code() int {
	return 200
}

func (o *GetV1MandatesIDEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/mandates/{id}/events][%d] getV1MandatesIdEventsOK  %+v", 200, o.Payload)
}

func (o *GetV1MandatesIDEventsOK) String() string {
	return fmt.Sprintf("[GET /v1/mandates/{id}/events][%d] getV1MandatesIdEventsOK  %+v", 200, o.Payload)
}

func (o *GetV1MandatesIDEventsOK) GetPayload() *GetV1MandatesIDEventsOKBody {
	return o.Payload
}

func (o *GetV1MandatesIDEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1MandatesIDEventsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1MandatesIDEventsOKBody get v1 mandates ID events o k body
swagger:model GetV1MandatesIDEventsOKBody
*/
type GetV1MandatesIDEventsOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Event `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1MandatesIDEventsOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1MandatesIDEventsOKBodyAO0
	var getV1MandatesIDEventsOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1MandatesIDEventsOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1MandatesIDEventsOKBodyAO0

	// GetV1MandatesIDEventsOKBodyAO1
	var dataGetV1MandatesIDEventsOKBodyAO1 struct {
		Items []*models.Event `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1MandatesIDEventsOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1MandatesIDEventsOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1MandatesIDEventsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1MandatesIDEventsOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1MandatesIDEventsOKBodyAO0)
	var dataGetV1MandatesIDEventsOKBodyAO1 struct {
		Items []*models.Event `json:"items"`
	}

	dataGetV1MandatesIDEventsOKBodyAO1.Items = o.Items

	jsonDataGetV1MandatesIDEventsOKBodyAO1, errGetV1MandatesIDEventsOKBodyAO1 := swag.WriteJSON(dataGetV1MandatesIDEventsOKBodyAO1)
	if errGetV1MandatesIDEventsOKBodyAO1 != nil {
		return nil, errGetV1MandatesIDEventsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1MandatesIDEventsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 mandates ID events o k body
func (o *GetV1MandatesIDEventsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1MandatesIDEventsOKBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1MandatesIdEventsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1MandatesIdEventsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 mandates ID events o k body based on the context it is used
func (o *GetV1MandatesIDEventsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1MandatesIDEventsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1MandatesIdEventsOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1MandatesIdEventsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1MandatesIDEventsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1MandatesIDEventsOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1MandatesIDEventsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
