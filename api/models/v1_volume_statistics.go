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

// V1VolumeStatistics v1 volume statistics
//
// swagger:model v1.VolumeStatistics
type V1VolumeStatistics struct {

	// compression ratio
	// Required: true
	CompressionRatio *float64 `json:"CompressionRatio"`

	// logical used storage
	// Required: true
	LogicalUsedStorage *int64 `json:"LogicalUsedStorage"`

	// physical used storage
	// Required: true
	PhysicalUsedStorage *int64 `json:"PhysicalUsedStorage"`
}

// Validate validates this v1 volume statistics
func (m *V1VolumeStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompressionRatio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalUsedStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalUsedStorage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VolumeStatistics) validateCompressionRatio(formats strfmt.Registry) error {

	if err := validate.Required("CompressionRatio", "body", m.CompressionRatio); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeStatistics) validateLogicalUsedStorage(formats strfmt.Registry) error {

	if err := validate.Required("LogicalUsedStorage", "body", m.LogicalUsedStorage); err != nil {
		return err
	}

	return nil
}

func (m *V1VolumeStatistics) validatePhysicalUsedStorage(formats strfmt.Registry) error {

	if err := validate.Required("PhysicalUsedStorage", "body", m.PhysicalUsedStorage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 volume statistics based on context it is used
func (m *V1VolumeStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VolumeStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeStatistics) UnmarshalBinary(b []byte) error {
	var res V1VolumeStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
