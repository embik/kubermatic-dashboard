// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListKubeVirtInstancetypesNoCredentialsReader is a Reader for the ListKubeVirtInstancetypesNoCredentials structure.
type ListKubeVirtInstancetypesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubeVirtInstancetypesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubeVirtInstancetypesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubeVirtInstancetypesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubeVirtInstancetypesNoCredentialsOK creates a ListKubeVirtInstancetypesNoCredentialsOK with default headers values
func NewListKubeVirtInstancetypesNoCredentialsOK() *ListKubeVirtInstancetypesNoCredentialsOK {
	return &ListKubeVirtInstancetypesNoCredentialsOK{}
}

/*
ListKubeVirtInstancetypesNoCredentialsOK describes a response with status code 200, with default header values.

VirtualMachineInstancetypeList
*/
type ListKubeVirtInstancetypesNoCredentialsOK struct {
	Payload *models.VirtualMachineInstancetypeList
}

// IsSuccess returns true when this list kube virt instancetypes no credentials o k response has a 2xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list kube virt instancetypes no credentials o k response has a 3xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list kube virt instancetypes no credentials o k response has a 4xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list kube virt instancetypes no credentials o k response has a 5xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list kube virt instancetypes no credentials o k response a status code equal to that given
func (o *ListKubeVirtInstancetypesNoCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListKubeVirtInstancetypesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/instancetypes][%d] listKubeVirtInstancetypesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtInstancetypesNoCredentialsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/instancetypes][%d] listKubeVirtInstancetypesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtInstancetypesNoCredentialsOK) GetPayload() *models.VirtualMachineInstancetypeList {
	return o.Payload
}

func (o *ListKubeVirtInstancetypesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineInstancetypeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubeVirtInstancetypesNoCredentialsDefault creates a ListKubeVirtInstancetypesNoCredentialsDefault with default headers values
func NewListKubeVirtInstancetypesNoCredentialsDefault(code int) *ListKubeVirtInstancetypesNoCredentialsDefault {
	return &ListKubeVirtInstancetypesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*
ListKubeVirtInstancetypesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListKubeVirtInstancetypesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list kube virt instancetypes no credentials default response
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list kube virt instancetypes no credentials default response has a 2xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list kube virt instancetypes no credentials default response has a 3xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list kube virt instancetypes no credentials default response has a 4xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list kube virt instancetypes no credentials default response has a 5xx status code
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list kube virt instancetypes no credentials default response a status code equal to that given
func (o *ListKubeVirtInstancetypesNoCredentialsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListKubeVirtInstancetypesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/instancetypes][%d] listKubeVirtInstancetypesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtInstancetypesNoCredentialsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/kubevirt/instancetypes][%d] listKubeVirtInstancetypesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtInstancetypesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListKubeVirtInstancetypesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
