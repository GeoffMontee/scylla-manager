// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/scylla/v2/models"
)

// FindConfigAlternatorEnforceAuthorizationReader is a Reader for the FindConfigAlternatorEnforceAuthorization structure.
type FindConfigAlternatorEnforceAuthorizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigAlternatorEnforceAuthorizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigAlternatorEnforceAuthorizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigAlternatorEnforceAuthorizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigAlternatorEnforceAuthorizationOK creates a FindConfigAlternatorEnforceAuthorizationOK with default headers values
func NewFindConfigAlternatorEnforceAuthorizationOK() *FindConfigAlternatorEnforceAuthorizationOK {
	return &FindConfigAlternatorEnforceAuthorizationOK{}
}

/*
FindConfigAlternatorEnforceAuthorizationOK handles this case with default header values.

Config value
*/
type FindConfigAlternatorEnforceAuthorizationOK struct {
	Payload bool
}

func (o *FindConfigAlternatorEnforceAuthorizationOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigAlternatorEnforceAuthorizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigAlternatorEnforceAuthorizationDefault creates a FindConfigAlternatorEnforceAuthorizationDefault with default headers values
func NewFindConfigAlternatorEnforceAuthorizationDefault(code int) *FindConfigAlternatorEnforceAuthorizationDefault {
	return &FindConfigAlternatorEnforceAuthorizationDefault{
		_statusCode: code,
	}
}

/*
FindConfigAlternatorEnforceAuthorizationDefault handles this case with default header values.

unexpected error
*/
type FindConfigAlternatorEnforceAuthorizationDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config alternator enforce authorization default response
func (o *FindConfigAlternatorEnforceAuthorizationDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigAlternatorEnforceAuthorizationDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigAlternatorEnforceAuthorizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigAlternatorEnforceAuthorizationDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
