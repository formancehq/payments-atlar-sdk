// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegulatoryReportingDetail RegulatoryReporting is an optional field that can be set when additional information is needed due to regulatory and statutory requirements.
//
// swagger:model RegulatoryReportingDetail
type RegulatoryReportingDetail struct {

	// Code specifies the nature, purpose and reason for the transaction.
	Code string `json:"code,omitempty"`

	// Indicator indentifies wether the regulatory reporting information applies to the debit/credit side or both.
	// Example: CREDITOR
	// Enum: [DEBTOR CREDITOR BOTH]
	Indicator string `json:"indicator,omitempty"`

	// Market is the ISO 3166-1 alpha-2 country code of the entity requiring the details.
	Market string `json:"market,omitempty"`
}

// Validate validates this regulatory reporting detail
func (m *RegulatoryReportingDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndicator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var regulatoryReportingDetailTypeIndicatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEBTOR","CREDITOR","BOTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		regulatoryReportingDetailTypeIndicatorPropEnum = append(regulatoryReportingDetailTypeIndicatorPropEnum, v)
	}
}

const (

	// RegulatoryReportingDetailIndicatorDEBTOR captures enum value "DEBTOR"
	RegulatoryReportingDetailIndicatorDEBTOR string = "DEBTOR"

	// RegulatoryReportingDetailIndicatorCREDITOR captures enum value "CREDITOR"
	RegulatoryReportingDetailIndicatorCREDITOR string = "CREDITOR"

	// RegulatoryReportingDetailIndicatorBOTH captures enum value "BOTH"
	RegulatoryReportingDetailIndicatorBOTH string = "BOTH"
)

// prop value enum
func (m *RegulatoryReportingDetail) validateIndicatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, regulatoryReportingDetailTypeIndicatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RegulatoryReportingDetail) validateIndicator(formats strfmt.Registry) error {
	if swag.IsZero(m.Indicator) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndicatorEnum("indicator", "body", m.Indicator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this regulatory reporting detail based on context it is used
func (m *RegulatoryReportingDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegulatoryReportingDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegulatoryReportingDetail) UnmarshalBinary(b []byte) error {
	var res RegulatoryReportingDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
