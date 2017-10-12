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

// GetCollectdItemsReader is a Reader for the GetCollectdItems structure.
type GetCollectdItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectdItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCollectdItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCollectdItemsOK creates a GetCollectdItemsOK with default headers values
func NewGetCollectdItemsOK() *GetCollectdItemsOK {
	return &GetCollectdItemsOK{}
}

/*GetCollectdItemsOK handles this case with default header values.

GetCollectdItemsOK get collectd items o k
*/
type GetCollectdItemsOK struct {
	Payload models.GetCollectdItemsOKBody
}

func (o *GetCollectdItemsOK) Error() string {
	return fmt.Sprintf("[GET /collectd/][%d] getCollectdItemsOK  %+v", 200, o.Payload)
}

func (o *GetCollectdItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
