// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewListTemplatesParams creates a new ListTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTemplatesParams() *ListTemplatesParams {
	return &ListTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTemplatesParamsWithTimeout creates a new ListTemplatesParams object
// with the ability to set a timeout on a request.
func NewListTemplatesParamsWithTimeout(timeout time.Duration) *ListTemplatesParams {
	return &ListTemplatesParams{
		timeout: timeout,
	}
}

// NewListTemplatesParamsWithContext creates a new ListTemplatesParams object
// with the ability to set a context for a request.
func NewListTemplatesParamsWithContext(ctx context.Context) *ListTemplatesParams {
	return &ListTemplatesParams{
		Context: ctx,
	}
}

// NewListTemplatesParamsWithHTTPClient creates a new ListTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTemplatesParamsWithHTTPClient(client *http.Client) *ListTemplatesParams {
	return &ListTemplatesParams{
		HTTPClient: client,
	}
}

/* ListTemplatesParams contains all the parameters to send to the API endpoint
   for the list templates operation.

   Typically these are written to a http.Request.
*/
type ListTemplatesParams struct {
	// Body.
	Body ListTemplatesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTemplatesParams) WithDefaults() *ListTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list templates params
func (o *ListTemplatesParams) WithTimeout(timeout time.Duration) *ListTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list templates params
func (o *ListTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list templates params
func (o *ListTemplatesParams) WithContext(ctx context.Context) *ListTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list templates params
func (o *ListTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list templates params
func (o *ListTemplatesParams) WithHTTPClient(client *http.Client) *ListTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list templates params
func (o *ListTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list templates params
func (o *ListTemplatesParams) WithBody(body ListTemplatesBody) *ListTemplatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list templates params
func (o *ListTemplatesParams) SetBody(body ListTemplatesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
