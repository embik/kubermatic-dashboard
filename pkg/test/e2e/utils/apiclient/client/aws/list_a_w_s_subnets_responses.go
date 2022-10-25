// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListAWSSubnetsReader is a Reader for the ListAWSSubnets structure.
type ListAWSSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSSubnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSSubnetsOK creates a ListAWSSubnetsOK with default headers values
func NewListAWSSubnetsOK() *ListAWSSubnetsOK {
	return &ListAWSSubnetsOK{}
}

/*
ListAWSSubnetsOK describes a response with status code 200, with default header values.

AWSSubnetList
*/
type ListAWSSubnetsOK struct {
	Payload models.AWSSubnetList
}

// IsSuccess returns true when this list a w s subnets o k response has a 2xx status code
func (o *ListAWSSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list a w s subnets o k response has a 3xx status code
func (o *ListAWSSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list a w s subnets o k response has a 4xx status code
func (o *ListAWSSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list a w s subnets o k response has a 5xx status code
func (o *ListAWSSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list a w s subnets o k response a status code equal to that given
func (o *ListAWSSubnetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAWSSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/subnets][%d] listAWSSubnetsOK  %+v", 200, o.Payload)
}

func (o *ListAWSSubnetsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/subnets][%d] listAWSSubnetsOK  %+v", 200, o.Payload)
}

func (o *ListAWSSubnetsOK) GetPayload() models.AWSSubnetList {
	return o.Payload
}

func (o *ListAWSSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSSubnetsDefault creates a ListAWSSubnetsDefault with default headers values
func NewListAWSSubnetsDefault(code int) *ListAWSSubnetsDefault {
	return &ListAWSSubnetsDefault{
		_statusCode: code,
	}
}

/*
ListAWSSubnetsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListAWSSubnetsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list a w s subnets default response
func (o *ListAWSSubnetsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list a w s subnets default response has a 2xx status code
func (o *ListAWSSubnetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list a w s subnets default response has a 3xx status code
func (o *ListAWSSubnetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list a w s subnets default response has a 4xx status code
func (o *ListAWSSubnetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list a w s subnets default response has a 5xx status code
func (o *ListAWSSubnetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list a w s subnets default response a status code equal to that given
func (o *ListAWSSubnetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAWSSubnetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/subnets][%d] listAWSSubnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSubnetsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/providers/aws/{dc}/subnets][%d] listAWSSubnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSSubnetsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAWSSubnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}