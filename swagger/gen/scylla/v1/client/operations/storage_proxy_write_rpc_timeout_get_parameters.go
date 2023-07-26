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

// NewStorageProxyWriteRPCTimeoutGetParams creates a new StorageProxyWriteRPCTimeoutGetParams object
// with the default values initialized.
func NewStorageProxyWriteRPCTimeoutGetParams() *StorageProxyWriteRPCTimeoutGetParams {

	return &StorageProxyWriteRPCTimeoutGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyWriteRPCTimeoutGetParamsWithTimeout creates a new StorageProxyWriteRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyWriteRPCTimeoutGetParamsWithTimeout(timeout time.Duration) *StorageProxyWriteRPCTimeoutGetParams {

	return &StorageProxyWriteRPCTimeoutGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyWriteRPCTimeoutGetParamsWithContext creates a new StorageProxyWriteRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyWriteRPCTimeoutGetParamsWithContext(ctx context.Context) *StorageProxyWriteRPCTimeoutGetParams {

	return &StorageProxyWriteRPCTimeoutGetParams{

		Context: ctx,
	}
}

// NewStorageProxyWriteRPCTimeoutGetParamsWithHTTPClient creates a new StorageProxyWriteRPCTimeoutGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyWriteRPCTimeoutGetParamsWithHTTPClient(client *http.Client) *StorageProxyWriteRPCTimeoutGetParams {

	return &StorageProxyWriteRPCTimeoutGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyWriteRPCTimeoutGetParams contains all the parameters to send to the API endpoint
for the storage proxy write Rpc timeout get operation typically these are written to a http.Request
*/
type StorageProxyWriteRPCTimeoutGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) WithTimeout(timeout time.Duration) *StorageProxyWriteRPCTimeoutGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) WithContext(ctx context.Context) *StorageProxyWriteRPCTimeoutGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) WithHTTPClient(client *http.Client) *StorageProxyWriteRPCTimeoutGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy write Rpc timeout get params
func (o *StorageProxyWriteRPCTimeoutGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyWriteRPCTimeoutGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
