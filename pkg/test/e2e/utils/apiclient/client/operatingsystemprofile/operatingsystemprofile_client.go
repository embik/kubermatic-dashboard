// Code generated by go-swagger; DO NOT EDIT.

package operatingsystemprofile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operatingsystemprofile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operatingsystemprofile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListOperatingSystemProfiles(params *ListOperatingSystemProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOperatingSystemProfilesOK, error)

	ListOperatingSystemProfilesForCluster(params *ListOperatingSystemProfilesForClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOperatingSystemProfilesForClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListOperatingSystemProfiles lists operating system profiles
*/
func (a *Client) ListOperatingSystemProfiles(params *ListOperatingSystemProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOperatingSystemProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOperatingSystemProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOperatingSystemProfiles",
		Method:             "GET",
		PathPattern:        "/api/v2/seeds/{seed_name}/operatingsystemprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOperatingSystemProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOperatingSystemProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOperatingSystemProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListOperatingSystemProfilesForCluster Lists all available Operating System Profiles for a cluster
*/
func (a *Client) ListOperatingSystemProfilesForCluster(params *ListOperatingSystemProfilesForClusterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOperatingSystemProfilesForClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOperatingSystemProfilesForClusterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOperatingSystemProfilesForCluster",
		Method:             "GET",
		PathPattern:        "/projects/{project_id}/clusters/{cluster_id}/operatingsystemprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOperatingSystemProfilesForClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOperatingSystemProfilesForClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOperatingSystemProfilesForClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
