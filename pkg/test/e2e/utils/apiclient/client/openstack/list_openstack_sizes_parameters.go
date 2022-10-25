// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListOpenstackSizesParams creates a new ListOpenstackSizesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOpenstackSizesParams() *ListOpenstackSizesParams {
	return &ListOpenstackSizesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackSizesParamsWithTimeout creates a new ListOpenstackSizesParams object
// with the ability to set a timeout on a request.
func NewListOpenstackSizesParamsWithTimeout(timeout time.Duration) *ListOpenstackSizesParams {
	return &ListOpenstackSizesParams{
		timeout: timeout,
	}
}

// NewListOpenstackSizesParamsWithContext creates a new ListOpenstackSizesParams object
// with the ability to set a context for a request.
func NewListOpenstackSizesParamsWithContext(ctx context.Context) *ListOpenstackSizesParams {
	return &ListOpenstackSizesParams{
		Context: ctx,
	}
}

// NewListOpenstackSizesParamsWithHTTPClient creates a new ListOpenstackSizesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOpenstackSizesParamsWithHTTPClient(client *http.Client) *ListOpenstackSizesParams {
	return &ListOpenstackSizesParams{
		HTTPClient: client,
	}
}

/*
ListOpenstackSizesParams contains all the parameters to send to the API endpoint

	for the list openstack sizes operation.

	Typically these are written to a http.Request.
*/
type ListOpenstackSizesParams struct {

	// ApplicationCredentialID.
	ApplicationCredentialID *string

	// ApplicationCredentialSecret.
	ApplicationCredentialSecret *string

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Domain.
	Domain *string

	// OIDCAuthentication.
	OIDCAuthentication *bool

	// Password.
	Password *string

	// Project.
	Project *string

	// ProjectID.
	ProjectID *string

	// Tenant.
	Tenant *string

	// TenantID.
	TenantID *string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list openstack sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackSizesParams) WithDefaults() *ListOpenstackSizesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list openstack sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenstackSizesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithTimeout(timeout time.Duration) *ListOpenstackSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithContext(ctx context.Context) *ListOpenstackSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithHTTPClient(client *http.Client) *ListOpenstackSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCredentialID adds the applicationCredentialID to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithApplicationCredentialID(applicationCredentialID *string) *ListOpenstackSizesParams {
	o.SetApplicationCredentialID(applicationCredentialID)
	return o
}

// SetApplicationCredentialID adds the applicationCredentialId to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetApplicationCredentialID(applicationCredentialID *string) {
	o.ApplicationCredentialID = applicationCredentialID
}

// WithApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithApplicationCredentialSecret(applicationCredentialSecret *string) *ListOpenstackSizesParams {
	o.SetApplicationCredentialSecret(applicationCredentialSecret)
	return o
}

// SetApplicationCredentialSecret adds the applicationCredentialSecret to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetApplicationCredentialSecret(applicationCredentialSecret *string) {
	o.ApplicationCredentialSecret = applicationCredentialSecret
}

// WithCredential adds the credential to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithCredential(credential *string) *ListOpenstackSizesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithDatacenterName(datacenterName *string) *ListOpenstackSizesParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithDomain(domain *string) *ListOpenstackSizesParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithOIDCAuthentication adds the oIDCAuthentication to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithOIDCAuthentication(oIDCAuthentication *bool) *ListOpenstackSizesParams {
	o.SetOIDCAuthentication(oIDCAuthentication)
	return o
}

// SetOIDCAuthentication adds the oIdCAuthentication to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetOIDCAuthentication(oIDCAuthentication *bool) {
	o.OIDCAuthentication = oIDCAuthentication
}

// WithPassword adds the password to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithPassword(password *string) *ListOpenstackSizesParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetPassword(password *string) {
	o.Password = password
}

// WithProject adds the project to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithProject(project *string) *ListOpenstackSizesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetProject(project *string) {
	o.Project = project
}

// WithProjectID adds the projectID to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithProjectID(projectID *string) *ListOpenstackSizesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithTenant adds the tenant to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithTenant(tenant *string) *ListOpenstackSizesParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithTenantID(tenantID *string) *ListOpenstackSizesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithUsername adds the username to the list openstack sizes params
func (o *ListOpenstackSizesParams) WithUsername(username *string) *ListOpenstackSizesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list openstack sizes params
func (o *ListOpenstackSizesParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCredentialID != nil {

		// header param ApplicationCredentialID
		if err := r.SetHeaderParam("ApplicationCredentialID", *o.ApplicationCredentialID); err != nil {
			return err
		}
	}

	if o.ApplicationCredentialSecret != nil {

		// header param ApplicationCredentialSecret
		if err := r.SetHeaderParam("ApplicationCredentialSecret", *o.ApplicationCredentialSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.Domain != nil {

		// header param Domain
		if err := r.SetHeaderParam("Domain", *o.Domain); err != nil {
			return err
		}
	}

	if o.OIDCAuthentication != nil {

		// header param OIDCAuthentication
		if err := r.SetHeaderParam("OIDCAuthentication", swag.FormatBool(*o.OIDCAuthentication)); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Project != nil {

		// header param Project
		if err := r.SetHeaderParam("Project", *o.Project); err != nil {
			return err
		}
	}

	if o.ProjectID != nil {

		// header param ProjectID
		if err := r.SetHeaderParam("ProjectID", *o.ProjectID); err != nil {
			return err
		}
	}

	if o.Tenant != nil {

		// header param Tenant
		if err := r.SetHeaderParam("Tenant", *o.Tenant); err != nil {
			return err
		}
	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
