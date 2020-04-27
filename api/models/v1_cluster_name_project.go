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

// V1ClusterNameProject v1 cluster name project
// swagger:model v1.ClusterNameProject
type V1ClusterNameProject struct {

	// cluster name
	// Required: true
	ClusterName *string `json:"ClusterName"`

	// project
	// Required: true
	Project *string `json:"Project"`
}

// Validate validates this v1 cluster name project
func (m *V1ClusterNameProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ClusterNameProject) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("ClusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *V1ClusterNameProject) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("Project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ClusterNameProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ClusterNameProject) UnmarshalBinary(b []byte) error {
	var res V1ClusterNameProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}