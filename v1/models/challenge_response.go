// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChallengeResponse challenge response
//
// swagger:model ChallengeResponse
type ChallengeResponse struct {

	// challenge
	Challenge string `json:"challenge,omitempty"`

	// session
	Session string `json:"session,omitempty"`
}

// Validate validates this challenge response
func (m *ChallengeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this challenge response based on context it is used
func (m *ChallengeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChallengeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChallengeResponse) UnmarshalBinary(b []byte) error {
	var res ChallengeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}