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

// ListProjectKubeVirtInstancetypesReader is a Reader for the ListProjectKubeVirtInstancetypes structure.
type ListProjectKubeVirtInstancetypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectKubeVirtInstancetypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectKubeVirtInstancetypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectKubeVirtInstancetypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectKubeVirtInstancetypesOK creates a ListProjectKubeVirtInstancetypesOK with default headers values
func NewListProjectKubeVirtInstancetypesOK() *ListProjectKubeVirtInstancetypesOK {
	return &ListProjectKubeVirtInstancetypesOK{}
}

/*
ListProjectKubeVirtInstancetypesOK describes a response with status code 200, with default header values.

VirtualMachineInstancetypeList
*/
type ListProjectKubeVirtInstancetypesOK struct {
	Payload *models.VirtualMachineInstancetypeList
}

// IsSuccess returns true when this list project kube virt instancetypes o k response has a 2xx status code
func (o *ListProjectKubeVirtInstancetypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project kube virt instancetypes o k response has a 3xx status code
func (o *ListProjectKubeVirtInstancetypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project kube virt instancetypes o k response has a 4xx status code
func (o *ListProjectKubeVirtInstancetypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project kube virt instancetypes o k response has a 5xx status code
func (o *ListProjectKubeVirtInstancetypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project kube virt instancetypes o k response a status code equal to that given
func (o *ListProjectKubeVirtInstancetypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectKubeVirtInstancetypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/instancetypes][%d] listProjectKubeVirtInstancetypesOK  %+v", 200, o.Payload)
}

func (o *ListProjectKubeVirtInstancetypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/instancetypes][%d] listProjectKubeVirtInstancetypesOK  %+v", 200, o.Payload)
}

func (o *ListProjectKubeVirtInstancetypesOK) GetPayload() *models.VirtualMachineInstancetypeList {
	return o.Payload
}

func (o *ListProjectKubeVirtInstancetypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineInstancetypeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectKubeVirtInstancetypesDefault creates a ListProjectKubeVirtInstancetypesDefault with default headers values
func NewListProjectKubeVirtInstancetypesDefault(code int) *ListProjectKubeVirtInstancetypesDefault {
	return &ListProjectKubeVirtInstancetypesDefault{
		_statusCode: code,
	}
}

/*
ListProjectKubeVirtInstancetypesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectKubeVirtInstancetypesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project kube virt instancetypes default response
func (o *ListProjectKubeVirtInstancetypesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project kube virt instancetypes default response has a 2xx status code
func (o *ListProjectKubeVirtInstancetypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project kube virt instancetypes default response has a 3xx status code
func (o *ListProjectKubeVirtInstancetypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project kube virt instancetypes default response has a 4xx status code
func (o *ListProjectKubeVirtInstancetypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project kube virt instancetypes default response has a 5xx status code
func (o *ListProjectKubeVirtInstancetypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project kube virt instancetypes default response a status code equal to that given
func (o *ListProjectKubeVirtInstancetypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectKubeVirtInstancetypesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/instancetypes][%d] listProjectKubeVirtInstancetypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectKubeVirtInstancetypesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/instancetypes][%d] listProjectKubeVirtInstancetypes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectKubeVirtInstancetypesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectKubeVirtInstancetypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}