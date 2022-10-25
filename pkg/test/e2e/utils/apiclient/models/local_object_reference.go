// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LocalObjectReference LocalObjectReference contains enough information to let you locate the
// referenced object inside the same namespace.
// +structType=atomic
//
// swagger:model LocalObjectReference
type LocalObjectReference struct {

	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// TODO: Add other useful fields. apiVersion, kind, uid?
	// +optional
	Name string `json:"name,omitempty"`
}

// Validate validates this local object reference
func (m *LocalObjectReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this local object reference based on context it is used
func (m *LocalObjectReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocalObjectReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalObjectReference) UnmarshalBinary(b []byte) error {
	var res LocalObjectReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
