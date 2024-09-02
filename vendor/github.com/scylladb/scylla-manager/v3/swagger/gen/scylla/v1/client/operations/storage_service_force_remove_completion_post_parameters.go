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

// NewStorageServiceForceRemoveCompletionPostParams creates a new StorageServiceForceRemoveCompletionPostParams object
// with the default values initialized.
func NewStorageServiceForceRemoveCompletionPostParams() *StorageServiceForceRemoveCompletionPostParams {

	return &StorageServiceForceRemoveCompletionPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceForceRemoveCompletionPostParamsWithTimeout creates a new StorageServiceForceRemoveCompletionPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceForceRemoveCompletionPostParamsWithTimeout(timeout time.Duration) *StorageServiceForceRemoveCompletionPostParams {

	return &StorageServiceForceRemoveCompletionPostParams{

		timeout: timeout,
	}
}

// NewStorageServiceForceRemoveCompletionPostParamsWithContext creates a new StorageServiceForceRemoveCompletionPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceForceRemoveCompletionPostParamsWithContext(ctx context.Context) *StorageServiceForceRemoveCompletionPostParams {

	return &StorageServiceForceRemoveCompletionPostParams{

		Context: ctx,
	}
}

// NewStorageServiceForceRemoveCompletionPostParamsWithHTTPClient creates a new StorageServiceForceRemoveCompletionPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceForceRemoveCompletionPostParamsWithHTTPClient(client *http.Client) *StorageServiceForceRemoveCompletionPostParams {

	return &StorageServiceForceRemoveCompletionPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceForceRemoveCompletionPostParams contains all the parameters to send to the API endpoint
for the storage service force remove completion post operation typically these are written to a http.Request
*/
type StorageServiceForceRemoveCompletionPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) WithTimeout(timeout time.Duration) *StorageServiceForceRemoveCompletionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) WithContext(ctx context.Context) *StorageServiceForceRemoveCompletionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) WithHTTPClient(client *http.Client) *StorageServiceForceRemoveCompletionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service force remove completion post params
func (o *StorageServiceForceRemoveCompletionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceForceRemoveCompletionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
