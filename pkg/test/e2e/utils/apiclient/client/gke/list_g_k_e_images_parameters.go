// Code generated by go-swagger; DO NOT EDIT.

package gke

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

// NewListGKEImagesParams creates a new ListGKEImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGKEImagesParams() *ListGKEImagesParams {
	return &ListGKEImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGKEImagesParamsWithTimeout creates a new ListGKEImagesParams object
// with the ability to set a timeout on a request.
func NewListGKEImagesParamsWithTimeout(timeout time.Duration) *ListGKEImagesParams {
	return &ListGKEImagesParams{
		timeout: timeout,
	}
}

// NewListGKEImagesParamsWithContext creates a new ListGKEImagesParams object
// with the ability to set a context for a request.
func NewListGKEImagesParamsWithContext(ctx context.Context) *ListGKEImagesParams {
	return &ListGKEImagesParams{
		Context: ctx,
	}
}

// NewListGKEImagesParamsWithHTTPClient creates a new ListGKEImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGKEImagesParamsWithHTTPClient(client *http.Client) *ListGKEImagesParams {
	return &ListGKEImagesParams{
		HTTPClient: client,
	}
}

/*
ListGKEImagesParams contains all the parameters to send to the API endpoint

	for the list g k e images operation.

	Typically these are written to a http.Request.
*/
type ListGKEImagesParams struct {

	/* Credential.

	   The credential name used in the preset for the GCP provider
	*/
	Credential *string

	/* ServiceAccount.

	   The plain GCP service account
	*/
	ServiceAccount *string

	/* Zone.

	   The zone name
	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list g k e images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGKEImagesParams) WithDefaults() *ListGKEImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list g k e images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGKEImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list g k e images params
func (o *ListGKEImagesParams) WithTimeout(timeout time.Duration) *ListGKEImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g k e images params
func (o *ListGKEImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g k e images params
func (o *ListGKEImagesParams) WithContext(ctx context.Context) *ListGKEImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g k e images params
func (o *ListGKEImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g k e images params
func (o *ListGKEImagesParams) WithHTTPClient(client *http.Client) *ListGKEImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g k e images params
func (o *ListGKEImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list g k e images params
func (o *ListGKEImagesParams) WithCredential(credential *string) *ListGKEImagesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list g k e images params
func (o *ListGKEImagesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithServiceAccount adds the serviceAccount to the list g k e images params
func (o *ListGKEImagesParams) WithServiceAccount(serviceAccount *string) *ListGKEImagesParams {
	o.SetServiceAccount(serviceAccount)
	return o
}

// SetServiceAccount adds the serviceAccount to the list g k e images params
func (o *ListGKEImagesParams) SetServiceAccount(serviceAccount *string) {
	o.ServiceAccount = serviceAccount
}

// WithZone adds the zone to the list g k e images params
func (o *ListGKEImagesParams) WithZone(zone *string) *ListGKEImagesParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list g k e images params
func (o *ListGKEImagesParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *ListGKEImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.ServiceAccount != nil {

		// header param ServiceAccount
		if err := r.SetHeaderParam("ServiceAccount", *o.ServiceAccount); err != nil {
			return err
		}
	}

	if o.Zone != nil {

		// header param Zone
		if err := r.SetHeaderParam("Zone", *o.Zone); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
