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

// ModelsV1MachineProvisioningEvent models v1 machine provisioning event
//
// swagger:model models.V1MachineProvisioningEvent
type ModelsV1MachineProvisioningEvent struct {

	// event
	// Required: true
	Event *string `json:"event"`

	// message
	Message string `json:"message,omitempty"`

	// time
	Time string `json:"time,omitempty"`
}

// Validate validates this models v1 machine provisioning event
func (m *ModelsV1MachineProvisioningEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsV1MachineProvisioningEvent) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models v1 machine provisioning event based on context it is used
func (m *ModelsV1MachineProvisioningEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsV1MachineProvisioningEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsV1MachineProvisioningEvent) UnmarshalBinary(b []byte) error {
	var res ModelsV1MachineProvisioningEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
