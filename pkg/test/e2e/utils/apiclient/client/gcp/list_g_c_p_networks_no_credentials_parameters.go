// Code generated by go-swagger; DO NOT EDIT.

package gcp

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
)

// NewListGCPNetworksNoCredentialsParams creates a new ListGCPNetworksNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGCPNetworksNoCredentialsParams() *ListGCPNetworksNoCredentialsParams {
	return &ListGCPNetworksNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGCPNetworksNoCredentialsParamsWithTimeout creates a new ListGCPNetworksNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListGCPNetworksNoCredentialsParamsWithTimeout(timeout time.Duration) *ListGCPNetworksNoCredentialsParams {
	return &ListGCPNetworksNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListGCPNetworksNoCredentialsParamsWithContext creates a new ListGCPNetworksNoCredentialsParams object
// with the ability to set a context for a request.
func NewListGCPNetworksNoCredentialsParamsWithContext(ctx context.Context) *ListGCPNetworksNoCredentialsParams {
	return &ListGCPNetworksNoCredentialsParams{
		Context: ctx,
	}
}

// NewListGCPNetworksNoCredentialsParamsWithHTTPClient creates a new ListGCPNetworksNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGCPNetworksNoCredentialsParamsWithHTTPClient(client *http.Client) *ListGCPNetworksNoCredentialsParams {
	return &ListGCPNetworksNoCredentialsParams{
		HTTPClient: client,
	}
}

/*
ListGCPNetworksNoCredentialsParams contains all the parameters to send to the API endpoint

	for the list g c p networks no credentials operation.

	Typically these are written to a http.Request.
*/
type ListGCPNetworksNoCredentialsParams struct {

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list g c p networks no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGCPNetworksNoCredentialsParams) WithDefaults() *ListGCPNetworksNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list g c p networks no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGCPNetworksNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithTimeout(timeout time.Duration) *ListGCPNetworksNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithContext(ctx context.Context) *ListGCPNetworksNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithHTTPClient(client *http.Client) *ListGCPNetworksNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithClusterID(clusterID string) *ListGCPNetworksNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithDC(dc string) *ListGCPNetworksNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) WithProjectID(projectID string) *ListGCPNetworksNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list g c p networks no credentials params
func (o *ListGCPNetworksNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListGCPNetworksNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
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
