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
	"github.com/go-openapi/swag"
)

// NewCacheServiceRowCacheKeysToSavePostParams creates a new CacheServiceRowCacheKeysToSavePostParams object
// with the default values initialized.
func NewCacheServiceRowCacheKeysToSavePostParams() *CacheServiceRowCacheKeysToSavePostParams {
	var ()
	return &CacheServiceRowCacheKeysToSavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceRowCacheKeysToSavePostParamsWithTimeout creates a new CacheServiceRowCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceRowCacheKeysToSavePostParamsWithTimeout(timeout time.Duration) *CacheServiceRowCacheKeysToSavePostParams {
	var ()
	return &CacheServiceRowCacheKeysToSavePostParams{

		timeout: timeout,
	}
}

// NewCacheServiceRowCacheKeysToSavePostParamsWithContext creates a new CacheServiceRowCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceRowCacheKeysToSavePostParamsWithContext(ctx context.Context) *CacheServiceRowCacheKeysToSavePostParams {
	var ()
	return &CacheServiceRowCacheKeysToSavePostParams{

		Context: ctx,
	}
}

// NewCacheServiceRowCacheKeysToSavePostParamsWithHTTPClient creates a new CacheServiceRowCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceRowCacheKeysToSavePostParamsWithHTTPClient(client *http.Client) *CacheServiceRowCacheKeysToSavePostParams {
	var ()
	return &CacheServiceRowCacheKeysToSavePostParams{
		HTTPClient: client,
	}
}

/*
CacheServiceRowCacheKeysToSavePostParams contains all the parameters to send to the API endpoint
for the cache service row cache keys to save post operation typically these are written to a http.Request
*/
type CacheServiceRowCacheKeysToSavePostParams struct {

	/*Rckts
	  row cache keys to save

	*/
	Rckts int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) WithTimeout(timeout time.Duration) *CacheServiceRowCacheKeysToSavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) WithContext(ctx context.Context) *CacheServiceRowCacheKeysToSavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) WithHTTPClient(client *http.Client) *CacheServiceRowCacheKeysToSavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRckts adds the rckts to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) WithRckts(rckts int32) *CacheServiceRowCacheKeysToSavePostParams {
	o.SetRckts(rckts)
	return o
}

// SetRckts adds the rckts to the cache service row cache keys to save post params
func (o *CacheServiceRowCacheKeysToSavePostParams) SetRckts(rckts int32) {
	o.Rckts = rckts
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceRowCacheKeysToSavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param rckts
	qrRckts := o.Rckts
	qRckts := swag.FormatInt32(qrRckts)
	if qRckts != "" {
		if err := r.SetQueryParam("rckts", qRckts); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
