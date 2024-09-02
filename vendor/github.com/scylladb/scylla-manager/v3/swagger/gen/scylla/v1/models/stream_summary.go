// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StreamSummary stream_summary
//
// # Stream summary info
//
// swagger:model stream_summary
type StreamSummary struct {

	// The ID
	CfID string `json:"cf_id,omitempty"`

	// Number of files to transfer. Can be 0 if nothing to transfer for some streaming request.
	Files int32 `json:"files,omitempty"`

	// total size
	TotalSize interface{} `json:"total_size,omitempty"`
}

// Validate validates this stream summary
func (m *StreamSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StreamSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamSummary) UnmarshalBinary(b []byte) error {
	var res StreamSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
