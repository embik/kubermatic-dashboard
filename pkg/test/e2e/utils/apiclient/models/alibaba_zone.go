// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlibabaZone AlibabaZone represents a object of Alibaba zone.
//
// swagger:model AlibabaZone
type AlibabaZone struct {

	// ID
	ID string `json:"id,omitempty"`
}

// Validate validates this alibaba zone
func (m *AlibabaZone) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alibaba zone based on context it is used
func (m *AlibabaZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlibabaZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlibabaZone) UnmarshalBinary(b []byte) error {
	var res AlibabaZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
