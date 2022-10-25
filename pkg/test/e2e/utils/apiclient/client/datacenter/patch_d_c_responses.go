// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// PatchDCReader is a Reader for the PatchDC structure.
type PatchDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchDCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchDCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchDCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDCOK creates a PatchDCOK with default headers values
func NewPatchDCOK() *PatchDCOK {
	return &PatchDCOK{}
}

/*
PatchDCOK describes a response with status code 200, with default header values.

Datacenter
*/
type PatchDCOK struct {
	Payload *models.Datacenter
}

// IsSuccess returns true when this patch d c o k response has a 2xx status code
func (o *PatchDCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch d c o k response has a 3xx status code
func (o *PatchDCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch d c o k response has a 4xx status code
func (o *PatchDCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch d c o k response has a 5xx status code
func (o *PatchDCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch d c o k response a status code equal to that given
func (o *PatchDCOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchDCOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCOK  %+v", 200, o.Payload)
}

func (o *PatchDCOK) String() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCOK  %+v", 200, o.Payload)
}

func (o *PatchDCOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *PatchDCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDCUnauthorized creates a PatchDCUnauthorized with default headers values
func NewPatchDCUnauthorized() *PatchDCUnauthorized {
	return &PatchDCUnauthorized{}
}

/*
PatchDCUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type PatchDCUnauthorized struct {
}

// IsSuccess returns true when this patch d c unauthorized response has a 2xx status code
func (o *PatchDCUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch d c unauthorized response has a 3xx status code
func (o *PatchDCUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch d c unauthorized response has a 4xx status code
func (o *PatchDCUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch d c unauthorized response has a 5xx status code
func (o *PatchDCUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch d c unauthorized response a status code equal to that given
func (o *PatchDCUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchDCUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCUnauthorized ", 401)
}

func (o *PatchDCUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCUnauthorized ", 401)
}

func (o *PatchDCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDCForbidden creates a PatchDCForbidden with default headers values
func NewPatchDCForbidden() *PatchDCForbidden {
	return &PatchDCForbidden{}
}

/*
PatchDCForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type PatchDCForbidden struct {
}

// IsSuccess returns true when this patch d c forbidden response has a 2xx status code
func (o *PatchDCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch d c forbidden response has a 3xx status code
func (o *PatchDCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch d c forbidden response has a 4xx status code
func (o *PatchDCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch d c forbidden response has a 5xx status code
func (o *PatchDCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch d c forbidden response a status code equal to that given
func (o *PatchDCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchDCForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCForbidden ", 403)
}

func (o *PatchDCForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDCForbidden ", 403)
}

func (o *PatchDCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDCDefault creates a PatchDCDefault with default headers values
func NewPatchDCDefault(code int) *PatchDCDefault {
	return &PatchDCDefault{
		_statusCode: code,
	}
}

/*
PatchDCDefault describes a response with status code -1, with default header values.

errorResponse
*/
type PatchDCDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch d c default response
func (o *PatchDCDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this patch d c default response has a 2xx status code
func (o *PatchDCDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch d c default response has a 3xx status code
func (o *PatchDCDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch d c default response has a 4xx status code
func (o *PatchDCDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch d c default response has a 5xx status code
func (o *PatchDCDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch d c default response a status code equal to that given
func (o *PatchDCDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PatchDCDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDC default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDCDefault) String() string {
	return fmt.Sprintf("[PATCH /api/v1/seed/{seed_name}/dc/{dc}][%d] patchDC default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDCDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchDCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
