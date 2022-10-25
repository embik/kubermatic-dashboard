// Code generated by go-swagger; DO NOT EDIT.

package etcdrestore

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

// NewGetEtcdRestoreParams creates a new GetEtcdRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEtcdRestoreParams() *GetEtcdRestoreParams {
	return &GetEtcdRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEtcdRestoreParamsWithTimeout creates a new GetEtcdRestoreParams object
// with the ability to set a timeout on a request.
func NewGetEtcdRestoreParamsWithTimeout(timeout time.Duration) *GetEtcdRestoreParams {
	return &GetEtcdRestoreParams{
		timeout: timeout,
	}
}

// NewGetEtcdRestoreParamsWithContext creates a new GetEtcdRestoreParams object
// with the ability to set a context for a request.
func NewGetEtcdRestoreParamsWithContext(ctx context.Context) *GetEtcdRestoreParams {
	return &GetEtcdRestoreParams{
		Context: ctx,
	}
}

// NewGetEtcdRestoreParamsWithHTTPClient creates a new GetEtcdRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEtcdRestoreParamsWithHTTPClient(client *http.Client) *GetEtcdRestoreParams {
	return &GetEtcdRestoreParams{
		HTTPClient: client,
	}
}

/*
GetEtcdRestoreParams contains all the parameters to send to the API endpoint

	for the get etcd restore operation.

	Typically these are written to a http.Request.
*/
type GetEtcdRestoreParams struct {

	// ClusterID.
	ClusterID string

	// ErName.
	EtcdRestoreName string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get etcd restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEtcdRestoreParams) WithDefaults() *GetEtcdRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get etcd restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEtcdRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get etcd restore params
func (o *GetEtcdRestoreParams) WithTimeout(timeout time.Duration) *GetEtcdRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get etcd restore params
func (o *GetEtcdRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get etcd restore params
func (o *GetEtcdRestoreParams) WithContext(ctx context.Context) *GetEtcdRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get etcd restore params
func (o *GetEtcdRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get etcd restore params
func (o *GetEtcdRestoreParams) WithHTTPClient(client *http.Client) *GetEtcdRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get etcd restore params
func (o *GetEtcdRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get etcd restore params
func (o *GetEtcdRestoreParams) WithClusterID(clusterID string) *GetEtcdRestoreParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get etcd restore params
func (o *GetEtcdRestoreParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithEtcdRestoreName adds the erName to the get etcd restore params
func (o *GetEtcdRestoreParams) WithEtcdRestoreName(erName string) *GetEtcdRestoreParams {
	o.SetEtcdRestoreName(erName)
	return o
}

// SetEtcdRestoreName adds the erName to the get etcd restore params
func (o *GetEtcdRestoreParams) SetEtcdRestoreName(erName string) {
	o.EtcdRestoreName = erName
}

// WithProjectID adds the projectID to the get etcd restore params
func (o *GetEtcdRestoreParams) WithProjectID(projectID string) *GetEtcdRestoreParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get etcd restore params
func (o *GetEtcdRestoreParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEtcdRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param er_name
	if err := r.SetPathParam("er_name", o.EtcdRestoreName); err != nil {
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
