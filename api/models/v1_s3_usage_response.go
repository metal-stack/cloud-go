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

// V1S3UsageResponse v1 s3 usage response
// swagger:model v1.S3UsageResponse
type V1S3UsageResponse struct {

	// just the usage data of the individual s3 buckets summed up
	// Required: true
	Accumulatedusage *V1S3UsageAccumuluated `json:"accumulatedusage"`

	// the start time in the accounting window to look at
	// Required: true
	// Format: date-time
	From *strfmt.DateTime `json:"from"`

	// the end time in the accounting window to look at (defaults to current system time)
	// Format: date-time
	To strfmt.DateTime `json:"to,omitempty"`

	// the usage data of the individual s3 buckets
	// Required: true
	Usage []*V1S3Usage `json:"usage"`
}

// Validate validates this v1 s3 usage response
func (m *V1S3UsageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccumulatedusage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1S3UsageResponse) validateAccumulatedusage(formats strfmt.Registry) error {

	if err := validate.Required("accumulatedusage", "body", m.Accumulatedusage); err != nil {
		return err
	}

	if m.Accumulatedusage != nil {
		if err := m.Accumulatedusage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accumulatedusage")
			}
			return err
		}
	}

	return nil
}

func (m *V1S3UsageResponse) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	if err := validate.FormatOf("from", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UsageResponse) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.FormatOf("to", "body", "date-time", m.To.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1S3UsageResponse) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	for i := 0; i < len(m.Usage); i++ {
		if swag.IsZero(m.Usage[i]) { // not required
			continue
		}

		if m.Usage[i] != nil {
			if err := m.Usage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1S3UsageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1S3UsageResponse) UnmarshalBinary(b []byte) error {
	var res V1S3UsageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}