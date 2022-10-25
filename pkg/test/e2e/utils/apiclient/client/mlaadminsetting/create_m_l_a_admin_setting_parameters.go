// Code generated by go-swagger; DO NOT EDIT.

package mlaadminsetting

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

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewCreateMLAAdminSettingParams creates a new CreateMLAAdminSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMLAAdminSettingParams() *CreateMLAAdminSettingParams {
	return &CreateMLAAdminSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMLAAdminSettingParamsWithTimeout creates a new CreateMLAAdminSettingParams object
// with the ability to set a timeout on a request.
func NewCreateMLAAdminSettingParamsWithTimeout(timeout time.Duration) *CreateMLAAdminSettingParams {
	return &CreateMLAAdminSettingParams{
		timeout: timeout,
	}
}

// NewCreateMLAAdminSettingParamsWithContext creates a new CreateMLAAdminSettingParams object
// with the ability to set a context for a request.
func NewCreateMLAAdminSettingParamsWithContext(ctx context.Context) *CreateMLAAdminSettingParams {
	return &CreateMLAAdminSettingParams{
		Context: ctx,
	}
}

// NewCreateMLAAdminSettingParamsWithHTTPClient creates a new CreateMLAAdminSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMLAAdminSettingParamsWithHTTPClient(client *http.Client) *CreateMLAAdminSettingParams {
	return &CreateMLAAdminSettingParams{
		HTTPClient: client,
	}
}

/*
CreateMLAAdminSettingParams contains all the parameters to send to the API endpoint

	for the create m l a admin setting operation.

	Typically these are written to a http.Request.
*/
type CreateMLAAdminSettingParams struct {

	// Body.
	Body *models.MLAAdminSetting

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create m l a admin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMLAAdminSettingParams) WithDefaults() *CreateMLAAdminSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create m l a admin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMLAAdminSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithTimeout(timeout time.Duration) *CreateMLAAdminSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithContext(ctx context.Context) *CreateMLAAdminSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithHTTPClient(client *http.Client) *CreateMLAAdminSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithBody(body *models.MLAAdminSetting) *CreateMLAAdminSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetBody(body *models.MLAAdminSetting) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithClusterID(clusterID string) *CreateMLAAdminSettingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) WithProjectID(projectID string) *CreateMLAAdminSettingParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create m l a admin setting params
func (o *CreateMLAAdminSettingParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMLAAdminSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
