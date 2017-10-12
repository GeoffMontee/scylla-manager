// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scylladb/mermaid/scylla/internal/models"
)

// DescribeAnyRingReader is a Reader for the DescribeAnyRing structure.
type DescribeAnyRingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAnyRingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAnyRingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAnyRingOK creates a DescribeAnyRingOK with default headers values
func NewDescribeAnyRingOK() *DescribeAnyRingOK {
	return &DescribeAnyRingOK{}
}

/*DescribeAnyRingOK handles this case with default header values.

DescribeAnyRingOK describe any ring o k
*/
type DescribeAnyRingOK struct {
	Payload models.DescribeAnyRingOKBody
}

func (o *DescribeAnyRingOK) Error() string {
	return fmt.Sprintf("[GET /storage_service/describe_ring/][%d] describeAnyRingOK  %+v", 200, o.Payload)
}

func (o *DescribeAnyRingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
