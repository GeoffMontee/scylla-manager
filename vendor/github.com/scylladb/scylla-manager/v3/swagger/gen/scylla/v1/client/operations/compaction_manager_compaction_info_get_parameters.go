// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCompactionManagerCompactionInfoGetParams creates a new CompactionManagerCompactionInfoGetParams object
// with the default values initialized.
func NewCompactionManagerCompactionInfoGetParams() *CompactionManagerCompactionInfoGetParams {

	return &CompactionManagerCompactionInfoGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompactionManagerCompactionInfoGetParamsWithTimeout creates a new CompactionManagerCompactionInfoGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompactionManagerCompactionInfoGetParamsWithTimeout(timeout time.Duration) *CompactionManagerCompactionInfoGetParams {

	return &CompactionManagerCompactionInfoGetParams{

		timeout: timeout,
	}
}

// NewCompactionManagerCompactionInfoGetParamsWithContext creates a new CompactionManagerCompactionInfoGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompactionManagerCompactionInfoGetParamsWithContext(ctx context.Context) *CompactionManagerCompactionInfoGetParams {

	return &CompactionManagerCompactionInfoGetParams{

		Context: ctx,
	}
}

// NewCompactionManagerCompactionInfoGetParamsWithHTTPClient creates a new CompactionManagerCompactionInfoGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompactionManagerCompactionInfoGetParamsWithHTTPClient(client *http.Client) *CompactionManagerCompactionInfoGetParams {

	return &CompactionManagerCompactionInfoGetParams{
		HTTPClient: client,
	}
}

/*
CompactionManagerCompactionInfoGetParams contains all the parameters to send to the API endpoint
for the compaction manager compaction info get operation typically these are written to a http.Request
*/
type CompactionManagerCompactionInfoGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) WithTimeout(timeout time.Duration) *CompactionManagerCompactionInfoGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) WithContext(ctx context.Context) *CompactionManagerCompactionInfoGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) WithHTTPClient(client *http.Client) *CompactionManagerCompactionInfoGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the compaction manager compaction info get params
func (o *CompactionManagerCompactionInfoGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CompactionManagerCompactionInfoGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
