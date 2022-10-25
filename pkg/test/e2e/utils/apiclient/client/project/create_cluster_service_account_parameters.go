// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewCreateClusterServiceAccountParams creates a new CreateClusterServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterServiceAccountParams() *CreateClusterServiceAccountParams {
	return &CreateClusterServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterServiceAccountParamsWithTimeout creates a new CreateClusterServiceAccountParams object
// with the ability to set a timeout on a request.
func NewCreateClusterServiceAccountParamsWithTimeout(timeout time.Duration) *CreateClusterServiceAccountParams {
	return &CreateClusterServiceAccountParams{
		timeout: timeout,
	}
}

// NewCreateClusterServiceAccountParamsWithContext creates a new CreateClusterServiceAccountParams object
// with the ability to set a context for a request.
func NewCreateClusterServiceAccountParamsWithContext(ctx context.Context) *CreateClusterServiceAccountParams {
	return &CreateClusterServiceAccountParams{
		Context: ctx,
	}
}

// NewCreateClusterServiceAccountParamsWithHTTPClient creates a new CreateClusterServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterServiceAccountParamsWithHTTPClient(client *http.Client) *CreateClusterServiceAccountParams {
	return &CreateClusterServiceAccountParams{
		HTTPClient: client,
	}
}

/*
CreateClusterServiceAccountParams contains all the parameters to send to the API endpoint

	for the create cluster service account operation.

	Typically these are written to a http.Request.
*/
type CreateClusterServiceAccountParams struct {

	// Body.
	Body *models.ClusterServiceAccount

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterServiceAccountParams) WithDefaults() *CreateClusterServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithTimeout(timeout time.Duration) *CreateClusterServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithContext(ctx context.Context) *CreateClusterServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithHTTPClient(client *http.Client) *CreateClusterServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithBody(body *models.ClusterServiceAccount) *CreateClusterServiceAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetBody(body *models.ClusterServiceAccount) {
	o.Body = body
}

// WithClusterID adds the clusterID to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithClusterID(clusterID string) *CreateClusterServiceAccountParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the create cluster service account params
func (o *CreateClusterServiceAccountParams) WithProjectID(projectID string) *CreateClusterServiceAccountParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create cluster service account params
func (o *CreateClusterServiceAccountParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
