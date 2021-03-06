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

// V1S3PartitionResponse v1 s3 partition response
//
// swagger:model v1.S3PartitionResponse
type V1S3PartitionResponse struct {

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ready
	// Required: true
	Ready *bool `json:"ready"`
}

// Validate validates this v1 s3 partition response
func (m *V1S3PartitionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReady(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3PartitionResponse) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *V1S3PartitionResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *V1S3PartitionResponse) validateReady(formats strfmt.Registry) error {

	if err := validate.Required("ready", "body", m.Ready); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 s3 partition response based on context it is used
func (m *V1S3PartitionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1S3PartitionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3PartitionResponse) UnmarshalBinary(b []byte) error {
	var res V1S3PartitionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
