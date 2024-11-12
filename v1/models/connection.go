// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Connection connection
//
// swagger:model Connection
type Connection struct {

	// affiliation
	Affiliation *SlimAffiliation `json:"affiliation,omitempty"`

	// alias
	Alias string `json:"alias,omitempty"`

	// bank
	Bank *BankSlim `json:"bank,omitempty"`

	// connection type
	ConnectionType *SlimConnectionType `json:"connectionType,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// inactive
	Inactive bool `json:"inactive,omitempty"`

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updated
	Updated string `json:"updated,omitempty"`

	// verified
	Verified string `json:"verified,omitempty"`
}

// Validate validates this connection
func (m *Connection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffiliation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Connection) validateAffiliation(formats strfmt.Registry) error {
	if swag.IsZero(m.Affiliation) { // not required
		return nil
	}

	if m.Affiliation != nil {
		if err := m.Affiliation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affiliation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affiliation")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) validateBank(formats strfmt.Registry) error {
	if swag.IsZero(m.Bank) { // not required
		return nil
	}

	if m.Bank != nil {
		if err := m.Bank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bank")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) validateConnectionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionType) { // not required
		return nil
	}

	if m.ConnectionType != nil {
		if err := m.ConnectionType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this connection based on the context it is used
func (m *Connection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAffiliation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBank(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Connection) contextValidateAffiliation(ctx context.Context, formats strfmt.Registry) error {

	if m.Affiliation != nil {

		if swag.IsZero(m.Affiliation) { // not required
			return nil
		}

		if err := m.Affiliation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affiliation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affiliation")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) contextValidateBank(ctx context.Context, formats strfmt.Registry) error {

	if m.Bank != nil {

		if swag.IsZero(m.Bank) { // not required
			return nil
		}

		if err := m.Bank.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bank")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) contextValidateConnectionType(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionType != nil {

		if swag.IsZero(m.ConnectionType) { // not required
			return nil
		}

		if err := m.ConnectionType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Connection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Connection) UnmarshalBinary(b []byte) error {
	var res Connection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
