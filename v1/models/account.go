// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
//
// swagger:model Account
type Account struct {

	// alias
	// Example: My Business Account
	Alias string `json:"alias,omitempty"`

	// balance
	Balance *Balance `json:"balance,omitempty"`

	// DEPRECATED
	Bank *BankSlim `json:"bank,omitempty"`

	// BIC of the Bank & Branch of the account
	// Example: DEUTDEFFXXX
	Bic string `json:"bic,omitempty"`

	// created
	// Example: 2022-05-04T18:31:12.889104898Z
	Created string `json:"created,omitempty"`

	// 3-letter ISO currency code
	// Example: EUR
	Currency string `json:"currency,omitempty"`

	// ID of the associated entity, e.g. the company that owns the account.
	EntityID string `json:"entityId,omitempty"`

	// external metadata
	ExternalMetadata ExternalMetadata `json:"externalMetadata,omitempty"`

	// fictive
	// Example: true
	Fictive bool `json:"fictive,omitempty"`

	// id
	// Example: 31d593d7-fff9-4783-aa1d-92acb7b21a19
	// Required: true
	ID *string `json:"id"`

	// identifiers
	Identifiers []*AccountIdentifier `json:"identifiers"`

	// 2-letter ISO country code
	// Example: DE
	Market string `json:"market,omitempty"`

	// name
	// Example: Company Account
	Name string `json:"name,omitempty"`

	// organization Id
	// Example: 605e26fc-4fce-495a-a92f-2c3592d7287e
	// Required: true
	OrganizationID *string `json:"organizationId"`

	// owner
	Owner *PartyIdentification `json:"owner,omitempty"`

	// ID of the associated third-party, e.g. the bank in which the account is held.
	ThirdPartyID string `json:"thirdPartyId,omitempty"`

	// updated
	// Example: 2022-05-04T18:31:12.889104898Z
	Updated string `json:"updated,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateBank(formats strfmt.Registry) error {
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

func (m *Account) validateExternalMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalMetadata) { // not required
		return nil
	}

	if m.ExternalMetadata != nil {
		if err := m.ExternalMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalMetadata")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {
		if swag.IsZero(m.Identifiers[i]) { // not required
			continue
		}

		if m.Identifiers[i] != nil {
			if err := m.Identifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Account) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account based on the context it is used
func (m *Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBank(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.Balance != nil {

		if swag.IsZero(m.Balance) { // not required
			return nil
		}

		if err := m.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidateBank(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Account) contextValidateExternalMetadata(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalMetadata) { // not required
		return nil
	}

	if err := m.ExternalMetadata.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("externalMetadata")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("externalMetadata")
		}
		return err
	}

	return nil
}

func (m *Account) contextValidateIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Identifiers); i++ {

		if m.Identifiers[i] != nil {

			if swag.IsZero(m.Identifiers[i]) { // not required
				return nil
			}

			if err := m.Identifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Account) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
