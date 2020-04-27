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

// V1S3CredentialsResponse v1 s3 credentials response
// swagger:model v1.S3CredentialsResponse
type V1S3CredentialsResponse struct {

	// access key
	// Required: true
	AccessKey *string `json:"access_key"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// name
	// Required: true
	Name *string `json:"name"`

	// partition
	// Required: true
	Partition *string `json:"partition"`

	// secret key
	// Required: true
	SecretKey *string `json:"secret_key"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Validate validates this v1 s3 credentials response
func (m *V1S3CredentialsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
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

func (m *V1S3CredentialsResponse) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("access_key", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validatePartition(formats strfmt.Registry) error {

	if err := validate.Required("partition", "body", m.Partition); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secret_key", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

func (m *V1S3CredentialsResponse) validateTenant(formats strfmt.Registry) error {

	if err := validate.Required("tenant", "body", m.Tenant); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1S3CredentialsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3CredentialsResponse) UnmarshalBinary(b []byte) error {
	var res V1S3CredentialsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}