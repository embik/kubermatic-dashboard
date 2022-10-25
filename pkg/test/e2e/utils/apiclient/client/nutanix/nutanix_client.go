// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nutanix API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nutanix API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListNutanixCategories(params *ListNutanixCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoriesOK, error)

	ListNutanixCategoriesNoCredentials(params *ListNutanixCategoriesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoriesNoCredentialsOK, error)

	ListNutanixCategoryValues(params *ListNutanixCategoryValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoryValuesOK, error)

	ListNutanixCategoryValuesNoCredentials(params *ListNutanixCategoryValuesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoryValuesNoCredentialsOK, error)

	ListNutanixClusters(params *ListNutanixClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixClustersOK, error)

	ListNutanixProjects(params *ListNutanixProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixProjectsOK, error)

	ListNutanixSubnets(params *ListNutanixSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixSubnetsOK, error)

	ListNutanixSubnetsNoCredentials(params *ListNutanixSubnetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixSubnetsNoCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListNutanixCategories List category keys from Nutanix
*/
func (a *Client) ListNutanixCategories(params *ListNutanixCategoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixCategoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixCategories",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/nutanix/{dc}/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixCategoriesReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixCategoriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixCategoriesNoCredentials Lists available Nutanix categories
*/
func (a *Client) ListNutanixCategoriesNoCredentials(params *ListNutanixCategoriesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoriesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixCategoriesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixCategoriesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixCategoriesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixCategoriesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixCategoriesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixCategoryValues List available category values for a specific category from Nutanix
*/
func (a *Client) ListNutanixCategoryValues(params *ListNutanixCategoryValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoryValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixCategoryValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixCategoryValues",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/nutanix/{dc}/categories/{category}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixCategoryValuesReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixCategoryValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixCategoryValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixCategoryValuesNoCredentials Lists available Nutanix category values for a specific category
*/
func (a *Client) ListNutanixCategoryValuesNoCredentials(params *ListNutanixCategoryValuesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixCategoryValuesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixCategoryValuesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixCategoryValuesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/categories/{category}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixCategoryValuesNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixCategoryValuesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixCategoryValuesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixClusters List clusters from Nutanix
*/
func (a *Client) ListNutanixClusters(params *ListNutanixClustersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixClusters",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/nutanix/{dc}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixClustersReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixProjects List projects from Nutanix
*/
func (a *Client) ListNutanixProjects(params *ListNutanixProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixProjects",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/nutanix/{dc}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixProjectsReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixProjectsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixSubnets List subnets from Nutanix
*/
func (a *Client) ListNutanixSubnets(params *ListNutanixSubnetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixSubnetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixSubnets",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/nutanix/{dc}/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixSubnetsReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixSubnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixSubnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNutanixSubnetsNoCredentials Lists available Nutanix Subnets
*/
func (a *Client) ListNutanixSubnetsNoCredentials(params *ListNutanixSubnetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNutanixSubnetsNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNutanixSubnetsNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNutanixSubnetsNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNutanixSubnetsNoCredentialsReader{formats: a.formats},
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
	success, ok := result.(*ListNutanixSubnetsNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNutanixSubnetsNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
