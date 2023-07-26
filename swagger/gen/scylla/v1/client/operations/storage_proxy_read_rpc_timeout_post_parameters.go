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

// NewStorageProxyReadRPCTimeoutPostParams creates a new StorageProxyReadRPCTimeoutPostParams object
// with the default values initialized.
func NewStorageProxyReadRPCTimeoutPostParams() *StorageProxyReadRPCTimeoutPostParams {
	var ()
	return &StorageProxyReadRPCTimeoutPostParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyReadRPCTimeoutPostParamsWithTimeout creates a new StorageProxyReadRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyReadRPCTimeoutPostParamsWithTimeout(timeout time.Duration) *StorageProxyReadRPCTimeoutPostParams {
	var ()
	return &StorageProxyReadRPCTimeoutPostParams{

		requestTimeout: timeout,
	}
}

// NewStorageProxyReadRPCTimeoutPostParamsWithContext creates a new StorageProxyReadRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyReadRPCTimeoutPostParamsWithContext(ctx context.Context) *StorageProxyReadRPCTimeoutPostParams {
	var ()
	return &StorageProxyReadRPCTimeoutPostParams{

		Context: ctx,
	}
}

// NewStorageProxyReadRPCTimeoutPostParamsWithHTTPClient creates a new StorageProxyReadRPCTimeoutPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyReadRPCTimeoutPostParamsWithHTTPClient(client *http.Client) *StorageProxyReadRPCTimeoutPostParams {
	var ()
	return &StorageProxyReadRPCTimeoutPostParams{
		HTTPClient: client,
	}
}

/*
StorageProxyReadRPCTimeoutPostParams contains all the parameters to send to the API endpoint
for the storage proxy read Rpc timeout post operation typically these are written to a http.Request
*/
type StorageProxyReadRPCTimeoutPostParams struct {

	/*Timeout
	  The timeout in second

	*/
	Timeout string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) WithRequestTimeout(timeout time.Duration) *StorageProxyReadRPCTimeoutPostParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) WithContext(ctx context.Context) *StorageProxyReadRPCTimeoutPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) WithHTTPClient(client *http.Client) *StorageProxyReadRPCTimeoutPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) WithTimeout(timeout string) *StorageProxyReadRPCTimeoutPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy read Rpc timeout post params
func (o *StorageProxyReadRPCTimeoutPostParams) SetTimeout(timeout string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyReadRPCTimeoutPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := qrTimeout
	if qTimeout != "" {
		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
