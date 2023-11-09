// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAzureAvailabilityZonesReader is a Reader for the ListAzureAvailabilityZones structure.
type ListAzureAvailabilityZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureAvailabilityZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureAvailabilityZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureAvailabilityZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureAvailabilityZonesOK creates a ListAzureAvailabilityZonesOK with default headers values
func NewListAzureAvailabilityZonesOK() *ListAzureAvailabilityZonesOK {
	return &ListAzureAvailabilityZonesOK{}
}

/*
ListAzureAvailabilityZonesOK describes a response with status code 200, with default header values.

AzureAvailabilityZonesList
*/
type ListAzureAvailabilityZonesOK struct {
	Payload *models.AzureAvailabilityZonesList
}

// IsSuccess returns true when this list azure availability zones o k response has a 2xx status code
func (o *ListAzureAvailabilityZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list azure availability zones o k response has a 3xx status code
func (o *ListAzureAvailabilityZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list azure availability zones o k response has a 4xx status code
func (o *ListAzureAvailabilityZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list azure availability zones o k response has a 5xx status code
func (o *ListAzureAvailabilityZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list azure availability zones o k response a status code equal to that given
func (o *ListAzureAvailabilityZonesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAzureAvailabilityZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZonesOK  %+v", 200, o.Payload)
}

func (o *ListAzureAvailabilityZonesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZonesOK  %+v", 200, o.Payload)
}

func (o *ListAzureAvailabilityZonesOK) GetPayload() *models.AzureAvailabilityZonesList {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureAvailabilityZonesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAvailabilityZonesDefault creates a ListAzureAvailabilityZonesDefault with default headers values
func NewListAzureAvailabilityZonesDefault(code int) *ListAzureAvailabilityZonesDefault {
	return &ListAzureAvailabilityZonesDefault{
		_statusCode: code,
	}
}

/*
ListAzureAvailabilityZonesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAzureAvailabilityZonesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure availability zones default response
func (o *ListAzureAvailabilityZonesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list azure availability zones default response has a 2xx status code
func (o *ListAzureAvailabilityZonesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list azure availability zones default response has a 3xx status code
func (o *ListAzureAvailabilityZonesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list azure availability zones default response has a 4xx status code
func (o *ListAzureAvailabilityZonesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list azure availability zones default response has a 5xx status code
func (o *ListAzureAvailabilityZonesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list azure availability zones default response a status code equal to that given
func (o *ListAzureAvailabilityZonesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAzureAvailabilityZonesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureAvailabilityZonesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureAvailabilityZonesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}