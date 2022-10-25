// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubeOneOpenStackCloudSpec KubeOneOpenStackCloudSpec specifies access data to an OpenStack cloud.
//
// swagger:model KubeOneOpenStackCloudSpec
type KubeOneOpenStackCloudSpec struct {

	// auth URL
	AuthURL string `json:"authURL,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// Project, formally known as tenant.
	Project string `json:"project,omitempty"`

	// ProjectID, formally known as tenantID.
	ProjectID string `json:"projectID,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this kube one open stack cloud spec
func (m *KubeOneOpenStackCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kube one open stack cloud spec based on context it is used
func (m *KubeOneOpenStackCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeOneOpenStackCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeOneOpenStackCloudSpec) UnmarshalBinary(b []byte) error {
	var res KubeOneOpenStackCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
