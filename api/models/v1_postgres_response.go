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

// V1PostgresResponse v1 postgres response
//
// swagger:model v1.PostgresResponse
type V1PostgresResponse struct {

	// description
	// Required: true
	Description *string `json:"Description"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// partition ID
	// Required: true
	PartitionID *string `json:"PartitionID"`

	// project ID
	// Required: true
	ProjectID *string `json:"ProjectID"`

	// tenant
	// Required: true
	Tenant *string `json:"Tenant"`

	// access list
	AccessList *V1AccessList `json:"accessList,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// creation timestamp
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creationTimestamp,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// maintenance
	Maintenance *V1MaintenanceWindow `json:"maintenance,omitempty"`

	// number of instances
	NumberOfInstances int32 `json:"numberOfInstances,omitempty"`

	// size
	Size *V1Size `json:"size,omitempty"`

	// status
	// Required: true
	Status *V1PostgresStatus `json:"status"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 postgres response
func (m *V1PostgresResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresResponse) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validatePartitionID(formats strfmt.Registry) error {

	if err := validate.Required("PartitionID", "body", m.PartitionID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("ProjectID", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("Tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validateAccessList(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessList) { // not required
		return nil
	}

	if m.AccessList != nil {
		if err := m.AccessList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PostgresResponse) validateMaintenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Maintenance) { // not required
		return nil
	}

	if m.Maintenance != nil {
		if err := m.Maintenance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 postgres response based on the context it is used
func (m *V1PostgresResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaintenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1PostgresResponse) contextValidateAccessList(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessList != nil {
		if err := m.AccessList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessList")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) contextValidateMaintenance(ctx context.Context, formats strfmt.Registry) error {

	if m.Maintenance != nil {
		if err := m.Maintenance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

func (m *V1PostgresResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1PostgresResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PostgresResponse) UnmarshalBinary(b []byte) error {
	var res V1PostgresResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}