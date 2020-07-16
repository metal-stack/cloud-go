// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1S3UpdateRequest v1 s3 update request
// swagger:model v1.S3UpdateRequest
type V1S3UpdateRequest struct {

	// add keys
	// Required: true
	AddKeys []*V1S3Key `json:"add_keys"`

	// id
	// Required: true
	ID *string `json:"id"`

	// partition
	// Required: true
	Partition *string `json:"partition"`

	// project
	// Required: true
	Project *string `json:"project"`

	// remove access keys
	// Required: true
	RemoveAccessKeys []string `json:"remove_access_keys"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Validate validates this v1 s3 update request
func (m *V1S3UpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoveAccessKeys(formats); err != nil {
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

func (m *V1S3UpdateRequest) validateAddKeys(formats strfmt.Registry) error {

	if err := validate.Required("add_keys", "body", m.AddKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.AddKeys); i++ {
		if swag.IsZero(m.AddKeys[i]) { // not required
			continue
		}

		if m.AddKeys[i] != nil {
			if err := m.AddKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("add_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1S3UpdateRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UpdateRequest) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UpdateRequest) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UpdateRequest) validateRemoveAccessKeys(formats strfmt.Registry) error {

	if err := validate.Required("remove_access_keys", "body", m.RemoveAccessKeys); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UpdateRequest) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1S3UpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3UpdateRequest) UnmarshalBinary(b []byte) error {
	var res V1S3UpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}