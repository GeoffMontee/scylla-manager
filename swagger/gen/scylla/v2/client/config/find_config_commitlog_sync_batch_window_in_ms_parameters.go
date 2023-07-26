// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigCommitlogSyncBatchWindowInMsParams creates a new FindConfigCommitlogSyncBatchWindowInMsParams object
// with the default values initialized.
func NewFindConfigCommitlogSyncBatchWindowInMsParams() *FindConfigCommitlogSyncBatchWindowInMsParams {

	return &FindConfigCommitlogSyncBatchWindowInMsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogSyncBatchWindowInMsParamsWithTimeout creates a new FindConfigCommitlogSyncBatchWindowInMsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogSyncBatchWindowInMsParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogSyncBatchWindowInMsParams {

	return &FindConfigCommitlogSyncBatchWindowInMsParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogSyncBatchWindowInMsParamsWithContext creates a new FindConfigCommitlogSyncBatchWindowInMsParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogSyncBatchWindowInMsParamsWithContext(ctx context.Context) *FindConfigCommitlogSyncBatchWindowInMsParams {

	return &FindConfigCommitlogSyncBatchWindowInMsParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogSyncBatchWindowInMsParamsWithHTTPClient creates a new FindConfigCommitlogSyncBatchWindowInMsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogSyncBatchWindowInMsParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogSyncBatchWindowInMsParams {

	return &FindConfigCommitlogSyncBatchWindowInMsParams{
		HTTPClient: client,
	}
}

/*
FindConfigCommitlogSyncBatchWindowInMsParams contains all the parameters to send to the API endpoint
for the find config commitlog sync batch window in ms operation typically these are written to a http.Request
*/
type FindConfigCommitlogSyncBatchWindowInMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogSyncBatchWindowInMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) WithContext(ctx context.Context) *FindConfigCommitlogSyncBatchWindowInMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogSyncBatchWindowInMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog sync batch window in ms params
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogSyncBatchWindowInMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
