// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMwareCloudDirectorCatalog VMwareCloudDirectorCatalog represents a VMware Cloud Director catalog.
//
// swagger:model VMwareCloudDirectorCatalog
type VMwareCloudDirectorCatalog struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v mware cloud director catalog
func (m *VMwareCloudDirectorCatalog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v mware cloud director catalog based on context it is used
func (m *VMwareCloudDirectorCatalog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMwareCloudDirectorCatalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareCloudDirectorCatalog) UnmarshalBinary(b []byte) error {
	var res VMwareCloudDirectorCatalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
