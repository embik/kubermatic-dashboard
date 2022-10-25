// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetAddonReader is a Reader for the GetAddon structure.
type GetAddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAddonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAddonOK creates a GetAddonOK with default headers values
func NewGetAddonOK() *GetAddonOK {
	return &GetAddonOK{}
}

/*
GetAddonOK describes a response with status code 200, with default header values.

Addon
*/
type GetAddonOK struct {
	Payload *models.Addon
}

// IsSuccess returns true when this get addon o k response has a 2xx status code
func (o *GetAddonOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get addon o k response has a 3xx status code
func (o *GetAddonOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get addon o k response has a 4xx status code
func (o *GetAddonOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get addon o k response has a 5xx status code
func (o *GetAddonOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get addon o k response a status code equal to that given
func (o *GetAddonOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAddonOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonOK  %+v", 200, o.Payload)
}

func (o *GetAddonOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonOK  %+v", 200, o.Payload)
}

func (o *GetAddonOK) GetPayload() *models.Addon {
	return o.Payload
}

func (o *GetAddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Addon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddonUnauthorized creates a GetAddonUnauthorized with default headers values
func NewGetAddonUnauthorized() *GetAddonUnauthorized {
	return &GetAddonUnauthorized{}
}

/*
GetAddonUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetAddonUnauthorized struct {
}

// IsSuccess returns true when this get addon unauthorized response has a 2xx status code
func (o *GetAddonUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get addon unauthorized response has a 3xx status code
func (o *GetAddonUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get addon unauthorized response has a 4xx status code
func (o *GetAddonUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get addon unauthorized response has a 5xx status code
func (o *GetAddonUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get addon unauthorized response a status code equal to that given
func (o *GetAddonUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAddonUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonUnauthorized ", 401)
}

func (o *GetAddonUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonUnauthorized ", 401)
}

func (o *GetAddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAddonForbidden creates a GetAddonForbidden with default headers values
func NewGetAddonForbidden() *GetAddonForbidden {
	return &GetAddonForbidden{}
}

/*
GetAddonForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetAddonForbidden struct {
}

// IsSuccess returns true when this get addon forbidden response has a 2xx status code
func (o *GetAddonForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get addon forbidden response has a 3xx status code
func (o *GetAddonForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get addon forbidden response has a 4xx status code
func (o *GetAddonForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get addon forbidden response has a 5xx status code
func (o *GetAddonForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get addon forbidden response a status code equal to that given
func (o *GetAddonForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAddonForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonForbidden ", 403)
}

func (o *GetAddonForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddonForbidden ", 403)
}

func (o *GetAddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAddonDefault creates a GetAddonDefault with default headers values
func NewGetAddonDefault(code int) *GetAddonDefault {
	return &GetAddonDefault{
		_statusCode: code,
	}
}

/*
GetAddonDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetAddonDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get addon default response
func (o *GetAddonDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get addon default response has a 2xx status code
func (o *GetAddonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get addon default response has a 3xx status code
func (o *GetAddonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get addon default response has a 4xx status code
func (o *GetAddonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get addon default response has a 5xx status code
func (o *GetAddonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get addon default response a status code equal to that given
func (o *GetAddonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAddonDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddon default  %+v", o._statusCode, o.Payload)
}

func (o *GetAddonDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}][%d] getAddon default  %+v", o._statusCode, o.Payload)
}

func (o *GetAddonDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAddonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
