// Code generated by go-swagger; DO NOT EDIT.

package pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPipelinesParams creates a new PipelinesParams object
// with the default values initialized.
func NewPipelinesParams() *PipelinesParams {

	return &PipelinesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPipelinesParamsWithTimeout creates a new PipelinesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPipelinesParamsWithTimeout(timeout time.Duration) *PipelinesParams {

	return &PipelinesParams{

		timeout: timeout,
	}
}

// NewPipelinesParamsWithContext creates a new PipelinesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPipelinesParamsWithContext(ctx context.Context) *PipelinesParams {

	return &PipelinesParams{

		Context: ctx,
	}
}

// NewPipelinesParamsWithHTTPClient creates a new PipelinesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPipelinesParamsWithHTTPClient(client *http.Client) *PipelinesParams {

	return &PipelinesParams{
		HTTPClient: client,
	}
}

/*PipelinesParams contains all the parameters to send to the API endpoint
for the pipelines operation typically these are written to a http.Request
*/
type PipelinesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pipelines params
func (o *PipelinesParams) WithTimeout(timeout time.Duration) *PipelinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipelines params
func (o *PipelinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipelines params
func (o *PipelinesParams) WithContext(ctx context.Context) *PipelinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipelines params
func (o *PipelinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipelines params
func (o *PipelinesParams) WithHTTPClient(client *http.Client) *PipelinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipelines params
func (o *PipelinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PipelinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
