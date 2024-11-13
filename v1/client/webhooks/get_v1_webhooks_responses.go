// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// GetV1WebhooksReader is a Reader for the GetV1Webhooks structure.
type GetV1WebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/webhooks] GetV1Webhooks", response, response.Code())
	}
}

// NewGetV1WebhooksOK creates a GetV1WebhooksOK with default headers values
func NewGetV1WebhooksOK() *GetV1WebhooksOK {
	return &GetV1WebhooksOK{}
}

/*
GetV1WebhooksOK describes a response with status code 200, with default header values.

QueryResponse with list of Webhooks
*/
type GetV1WebhooksOK struct {
	Payload *GetV1WebhooksOKBody
}

// IsSuccess returns true when this get v1 webhooks o k response has a 2xx status code
func (o *GetV1WebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 webhooks o k response has a 3xx status code
func (o *GetV1WebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 webhooks o k response has a 4xx status code
func (o *GetV1WebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 webhooks o k response has a 5xx status code
func (o *GetV1WebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 webhooks o k response a status code equal to that given
func (o *GetV1WebhooksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 webhooks o k response
func (o *GetV1WebhooksOK) Code() int {
	return 200
}

func (o *GetV1WebhooksOK) Error() string {
	return fmt.Sprintf("[GET /v1/webhooks][%d] getV1WebhooksOK  %+v", 200, o.Payload)
}

func (o *GetV1WebhooksOK) String() string {
	return fmt.Sprintf("[GET /v1/webhooks][%d] getV1WebhooksOK  %+v", 200, o.Payload)
}

func (o *GetV1WebhooksOK) GetPayload() *GetV1WebhooksOKBody {
	return o.Payload
}

func (o *GetV1WebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1WebhooksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1WebhooksOKBody get v1 webhooks o k body
swagger:model GetV1WebhooksOKBody
*/
type GetV1WebhooksOKBody struct {
	models.QueryResponse

	// items
	Items []*models.Webhook `json:"items"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1WebhooksOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1WebhooksOKBodyAO0
	var getV1WebhooksOKBodyAO0 models.QueryResponse
	if err := swag.ReadJSON(raw, &getV1WebhooksOKBodyAO0); err != nil {
		return err
	}
	o.QueryResponse = getV1WebhooksOKBodyAO0

	// GetV1WebhooksOKBodyAO1
	var dataGetV1WebhooksOKBodyAO1 struct {
		Items []*models.Webhook `json:"items"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1WebhooksOKBodyAO1); err != nil {
		return err
	}

	o.Items = dataGetV1WebhooksOKBodyAO1.Items

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1WebhooksOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1WebhooksOKBodyAO0, err := swag.WriteJSON(o.QueryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1WebhooksOKBodyAO0)
	var dataGetV1WebhooksOKBodyAO1 struct {
		Items []*models.Webhook `json:"items"`
	}

	dataGetV1WebhooksOKBodyAO1.Items = o.Items

	jsonDataGetV1WebhooksOKBodyAO1, errGetV1WebhooksOKBodyAO1 := swag.WriteJSON(dataGetV1WebhooksOKBodyAO1)
	if errGetV1WebhooksOKBodyAO1 != nil {
		return nil, errGetV1WebhooksOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1WebhooksOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 webhooks o k body
func (o *GetV1WebhooksOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1WebhooksOKBody) validateItems(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1WebhooksOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1WebhooksOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 webhooks o k body based on the context it is used
func (o *GetV1WebhooksOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1WebhooksOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {

			if swag.IsZero(o.Items[i]) { // not required
				return nil
			}

			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1WebhooksOK" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1WebhooksOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1WebhooksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1WebhooksOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1WebhooksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
