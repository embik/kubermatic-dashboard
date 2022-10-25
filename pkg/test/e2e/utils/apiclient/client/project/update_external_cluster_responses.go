// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// UpdateExternalClusterReader is a Reader for the UpdateExternalCluster structure.
type UpdateExternalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateExternalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateExternalClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateExternalClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateExternalClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateExternalClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateExternalClusterOK creates a UpdateExternalClusterOK with default headers values
func NewUpdateExternalClusterOK() *UpdateExternalClusterOK {
	return &UpdateExternalClusterOK{}
}

/*
UpdateExternalClusterOK describes a response with status code 200, with default header values.

ExternalCluster
*/
type UpdateExternalClusterOK struct {
	Payload *models.ExternalCluster
}

// IsSuccess returns true when this update external cluster o k response has a 2xx status code
func (o *UpdateExternalClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update external cluster o k response has a 3xx status code
func (o *UpdateExternalClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update external cluster o k response has a 4xx status code
func (o *UpdateExternalClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update external cluster o k response has a 5xx status code
func (o *UpdateExternalClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update external cluster o k response a status code equal to that given
func (o *UpdateExternalClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateExternalClusterOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateExternalClusterOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterOK  %+v", 200, o.Payload)
}

func (o *UpdateExternalClusterOK) GetPayload() *models.ExternalCluster {
	return o.Payload
}

func (o *UpdateExternalClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalClusterUnauthorized creates a UpdateExternalClusterUnauthorized with default headers values
func NewUpdateExternalClusterUnauthorized() *UpdateExternalClusterUnauthorized {
	return &UpdateExternalClusterUnauthorized{}
}

/*
UpdateExternalClusterUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateExternalClusterUnauthorized struct {
}

// IsSuccess returns true when this update external cluster unauthorized response has a 2xx status code
func (o *UpdateExternalClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update external cluster unauthorized response has a 3xx status code
func (o *UpdateExternalClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update external cluster unauthorized response has a 4xx status code
func (o *UpdateExternalClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update external cluster unauthorized response has a 5xx status code
func (o *UpdateExternalClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update external cluster unauthorized response a status code equal to that given
func (o *UpdateExternalClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateExternalClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterUnauthorized ", 401)
}

func (o *UpdateExternalClusterUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterUnauthorized ", 401)
}

func (o *UpdateExternalClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateExternalClusterForbidden creates a UpdateExternalClusterForbidden with default headers values
func NewUpdateExternalClusterForbidden() *UpdateExternalClusterForbidden {
	return &UpdateExternalClusterForbidden{}
}

/*
UpdateExternalClusterForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateExternalClusterForbidden struct {
}

// IsSuccess returns true when this update external cluster forbidden response has a 2xx status code
func (o *UpdateExternalClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update external cluster forbidden response has a 3xx status code
func (o *UpdateExternalClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update external cluster forbidden response has a 4xx status code
func (o *UpdateExternalClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update external cluster forbidden response has a 5xx status code
func (o *UpdateExternalClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update external cluster forbidden response a status code equal to that given
func (o *UpdateExternalClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateExternalClusterForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterForbidden ", 403)
}

func (o *UpdateExternalClusterForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalClusterForbidden ", 403)
}

func (o *UpdateExternalClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateExternalClusterDefault creates a UpdateExternalClusterDefault with default headers values
func NewUpdateExternalClusterDefault(code int) *UpdateExternalClusterDefault {
	return &UpdateExternalClusterDefault{
		_statusCode: code,
	}
}

/*
UpdateExternalClusterDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateExternalClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update external cluster default response
func (o *UpdateExternalClusterDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update external cluster default response has a 2xx status code
func (o *UpdateExternalClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update external cluster default response has a 3xx status code
func (o *UpdateExternalClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update external cluster default response has a 4xx status code
func (o *UpdateExternalClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update external cluster default response has a 5xx status code
func (o *UpdateExternalClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update external cluster default response a status code equal to that given
func (o *UpdateExternalClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateExternalClusterDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateExternalClusterDefault) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}][%d] updateExternalCluster default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateExternalClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateExternalClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateExternalClusterBody update external cluster body
swagger:model UpdateExternalClusterBody
*/
type UpdateExternalClusterBody struct {

	// Kubeconfig Base64 encoded kubeconfig
	Kubeconfig string `json:"kubeconfig,omitempty"`

	// Name is human readable name for the external cluster
	Name string `json:"name,omitempty"`
}

// Validate validates this update external cluster body
func (o *UpdateExternalClusterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update external cluster body based on context it is used
func (o *UpdateExternalClusterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateExternalClusterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateExternalClusterBody) UnmarshalBinary(b []byte) error {
	var res UpdateExternalClusterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
