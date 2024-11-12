// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateRejectRequest create reject request
//
// swagger:model CreateRejectRequest
type CreateRejectRequest struct {

	// An optional rejection reason that will propagate to the reject event on the payment.
	// Example: This looks like a duplicate
	// Max Length: 512
	// Min Length: 0
	Reason *string `json:"reason,omitempty"`
}

// Validate validates this create reject request
func (m *CreateRejectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRejectRequest) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := validate.MinLength("reason", "body", *m.Reason, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("reason", "body", *m.Reason, 512); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create reject request based on context it is used
func (m *CreateRejectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateRejectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRejectRequest) UnmarshalBinary(b []byte) error {
	var res CreateRejectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
