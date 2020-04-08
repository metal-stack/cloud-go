// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1MasterdataLookupResponse v1 masterdata lookup response
// swagger:model v1.MasterdataLookupResponse
type V1MasterdataLookupResponse struct {

	// project
	// Required: true
	Project *V1Project `json:"Project"`

	// tenant
	// Required: true
	Tenant *V1Tenant `json:"Tenant"`
}

// Validate validates this v1 masterdata lookup response
func (m *V1MasterdataLookupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1MasterdataLookupResponse) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("Project", "body", m.Project); err != nil {
		return err
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Project")
			}
			return err
		}
	}

	return nil
}

func (m *V1MasterdataLookupResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("Tenant", "body", m.Tenant); err != nil {
		return err
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Tenant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1MasterdataLookupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1MasterdataLookupResponse) UnmarshalBinary(b []byte) error {
	var res V1MasterdataLookupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
