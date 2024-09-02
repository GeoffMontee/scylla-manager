// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/scylladb/scylla-manager/v3/swagger/gen/agent/models"
)

// OperationsDeletepathsReader is a Reader for the OperationsDeletepaths structure.
type OperationsDeletepathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsDeletepathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsDeletepathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsDeletepathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsDeletepathsOK creates a OperationsDeletepathsOK with default headers values
func NewOperationsDeletepathsOK() *OperationsDeletepathsOK {
	return &OperationsDeletepathsOK{}
}

/*
OperationsDeletepathsOK handles this case with default header values.

Deleted files count
*/
type OperationsDeletepathsOK struct {
	Payload *OperationsDeletepathsOKBody
	JobID   int64
}

func (o *OperationsDeletepathsOK) GetPayload() *OperationsDeletepathsOKBody {
	return o.Payload
}

func (o *OperationsDeletepathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(OperationsDeletepathsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsDeletepathsDefault creates a OperationsDeletepathsDefault with default headers values
func NewOperationsDeletepathsDefault(code int) *OperationsDeletepathsDefault {
	return &OperationsDeletepathsDefault{
		_statusCode: code,
	}
}

/*
OperationsDeletepathsDefault handles this case with default header values.

Server error
*/
type OperationsDeletepathsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations deletepaths default response
func (o *OperationsDeletepathsDefault) Code() int {
	return o._statusCode
}

func (o *OperationsDeletepathsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsDeletepathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsDeletepathsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}

/*
OperationsDeletepathsOKBody operations deletepaths o k body
swagger:model OperationsDeletepathsOKBody
*/
type OperationsDeletepathsOKBody struct {

	// deletes
	Deletes int64 `json:"deletes,omitempty"`
}

// Validate validates this operations deletepaths o k body
func (o *OperationsDeletepathsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OperationsDeletepathsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OperationsDeletepathsOKBody) UnmarshalBinary(b []byte) error {
	var res OperationsDeletepathsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
