// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OperatingSystemSpec OperatingSystemSpec represents the collection of os specific settings. Only one must be set at a time.
//
// swagger:model OperatingSystemSpec
type OperatingSystemSpec struct {

	// amzn2
	Amzn2 *AmazonLinuxSpec `json:"amzn2,omitempty"`

	// centos
	Centos *CentOSSpec `json:"centos,omitempty"`

	// flatcar
	Flatcar *FlatcarSpec `json:"flatcar,omitempty"`

	// rhel
	Rhel *RHELSpec `json:"rhel,omitempty"`

	// rockylinux
	Rockylinux *RockyLinuxSpec `json:"rockylinux,omitempty"`

	// sles
	Sles *SLESSpec `json:"sles,omitempty"`

	// ubuntu
	Ubuntu *UbuntuSpec `json:"ubuntu,omitempty"`
}

// Validate validates this operating system spec
func (m *OperatingSystemSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmzn2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCentos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlatcar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRhel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRockylinux(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUbuntu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatingSystemSpec) validateAmzn2(formats strfmt.Registry) error {
	if swag.IsZero(m.Amzn2) { // not required
		return nil
	}

	if m.Amzn2 != nil {
		if err := m.Amzn2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amzn2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amzn2")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateCentos(formats strfmt.Registry) error {
	if swag.IsZero(m.Centos) { // not required
		return nil
	}

	if m.Centos != nil {
		if err := m.Centos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("centos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("centos")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateFlatcar(formats strfmt.Registry) error {
	if swag.IsZero(m.Flatcar) { // not required
		return nil
	}

	if m.Flatcar != nil {
		if err := m.Flatcar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flatcar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flatcar")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateRhel(formats strfmt.Registry) error {
	if swag.IsZero(m.Rhel) { // not required
		return nil
	}

	if m.Rhel != nil {
		if err := m.Rhel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rhel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rhel")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateRockylinux(formats strfmt.Registry) error {
	if swag.IsZero(m.Rockylinux) { // not required
		return nil
	}

	if m.Rockylinux != nil {
		if err := m.Rockylinux.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rockylinux")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rockylinux")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateSles(formats strfmt.Registry) error {
	if swag.IsZero(m.Sles) { // not required
		return nil
	}

	if m.Sles != nil {
		if err := m.Sles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sles")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) validateUbuntu(formats strfmt.Registry) error {
	if swag.IsZero(m.Ubuntu) { // not required
		return nil
	}

	if m.Ubuntu != nil {
		if err := m.Ubuntu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ubuntu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ubuntu")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this operating system spec based on the context it is used
func (m *OperatingSystemSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmzn2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCentos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlatcar(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRhel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRockylinux(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUbuntu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatingSystemSpec) contextValidateAmzn2(ctx context.Context, formats strfmt.Registry) error {

	if m.Amzn2 != nil {
		if err := m.Amzn2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amzn2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amzn2")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateCentos(ctx context.Context, formats strfmt.Registry) error {

	if m.Centos != nil {
		if err := m.Centos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("centos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("centos")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateFlatcar(ctx context.Context, formats strfmt.Registry) error {

	if m.Flatcar != nil {
		if err := m.Flatcar.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flatcar")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flatcar")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateRhel(ctx context.Context, formats strfmt.Registry) error {

	if m.Rhel != nil {
		if err := m.Rhel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rhel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rhel")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateRockylinux(ctx context.Context, formats strfmt.Registry) error {

	if m.Rockylinux != nil {
		if err := m.Rockylinux.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rockylinux")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rockylinux")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateSles(ctx context.Context, formats strfmt.Registry) error {

	if m.Sles != nil {
		if err := m.Sles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sles")
			}
			return err
		}
	}

	return nil
}

func (m *OperatingSystemSpec) contextValidateUbuntu(ctx context.Context, formats strfmt.Registry) error {

	if m.Ubuntu != nil {
		if err := m.Ubuntu.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ubuntu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ubuntu")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperatingSystemSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatingSystemSpec) UnmarshalBinary(b []byte) error {
	var res OperatingSystemSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
