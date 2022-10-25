// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListOpenstackSubnetsNoCredentialsReader is a Reader for the ListOpenstackSubnetsNoCredentials structure.
type ListOpenstackSubnetsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackSubnetsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackSubnetsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackSubnetsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackSubnetsNoCredentialsOK creates a ListOpenstackSubnetsNoCredentialsOK with default headers values
func NewListOpenstackSubnetsNoCredentialsOK() *ListOpenstackSubnetsNoCredentialsOK {
	return &ListOpenstackSubnetsNoCredentialsOK{}
}

/*
ListOpenstackSubnetsNoCredentialsOK describes a response with status code 200, with default header values.

OpenstackSubnet
*/
type ListOpenstackSubnetsNoCredentialsOK struct {
	Payload []*models.OpenstackSubnet
}

// IsSuccess returns true when this list openstack subnets no credentials o k response has a 2xx status code
func (o *ListOpenstackSubnetsNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list openstack subnets no credentials o k response has a 3xx status code
func (o *ListOpenstackSubnetsNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list openstack subnets no credentials o k response has a 4xx status code
func (o *ListOpenstackSubnetsNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list openstack subnets no credentials o k response has a 5xx status code
func (o *ListOpenstackSubnetsNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list openstack subnets no credentials o k response a status code equal to that given
func (o *ListOpenstackSubnetsNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListOpenstackSubnetsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets][%d] listOpenstackSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackSubnetsNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets][%d] listOpenstackSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackSubnetsNoCredentialsOK) GetPayload() []*models.OpenstackSubnet {
	return o.Payload
}

func (o *ListOpenstackSubnetsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackSubnetsNoCredentialsDefault creates a ListOpenstackSubnetsNoCredentialsDefault with default headers values
func NewListOpenstackSubnetsNoCredentialsDefault(code int) *ListOpenstackSubnetsNoCredentialsDefault {
	return &ListOpenstackSubnetsNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListOpenstackSubnetsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListOpenstackSubnetsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack subnets no credentials default response
func (o *ListOpenstackSubnetsNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list openstack subnets no credentials default response has a 2xx status code
func (o *ListOpenstackSubnetsNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list openstack subnets no credentials default response has a 3xx status code
func (o *ListOpenstackSubnetsNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list openstack subnets no credentials default response has a 4xx status code
func (o *ListOpenstackSubnetsNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list openstack subnets no credentials default response has a 5xx status code
func (o *ListOpenstackSubnetsNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list openstack subnets no credentials default response a status code equal to that given
func (o *ListOpenstackSubnetsNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListOpenstackSubnetsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets][%d] listOpenstackSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackSubnetsNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/subnets][%d] listOpenstackSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackSubnetsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackSubnetsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
