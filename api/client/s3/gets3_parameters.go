// Code generated by go-swagger; DO NOT EDIT.

package s3

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

	models "github.com/metal-stack/cloud-go/api/models"
)

// NewGets3Params creates a new Gets3Params object
// with the default values initialized.
func NewGets3Params() *Gets3Params {
	var ()
	return &Gets3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGets3ParamsWithTimeout creates a new Gets3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGets3ParamsWithTimeout(timeout time.Duration) *Gets3Params {
	var ()
	return &Gets3Params{

		timeout: timeout,
	}
}

// NewGets3ParamsWithContext creates a new Gets3Params object
// with the default values initialized, and the ability to set a context for a request
func NewGets3ParamsWithContext(ctx context.Context) *Gets3Params {
	var ()
	return &Gets3Params{

		Context: ctx,
	}
}

// NewGets3ParamsWithHTTPClient creates a new Gets3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGets3ParamsWithHTTPClient(client *http.Client) *Gets3Params {
	var ()
	return &Gets3Params{
		HTTPClient: client,
	}
}

/*Gets3Params contains all the parameters to send to the API endpoint
for the gets3 operation typically these are written to a http.Request
*/
type Gets3Params struct {

	/*Body*/
	Body *models.V1S3GetRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the gets3 params
func (o *Gets3Params) WithTimeout(timeout time.Duration) *Gets3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gets3 params
func (o *Gets3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gets3 params
func (o *Gets3Params) WithContext(ctx context.Context) *Gets3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gets3 params
func (o *Gets3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gets3 params
func (o *Gets3Params) WithHTTPClient(client *http.Client) *Gets3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gets3 params
func (o *Gets3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the gets3 params
func (o *Gets3Params) WithBody(body *models.V1S3GetRequest) *Gets3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the gets3 params
func (o *Gets3Params) SetBody(body *models.V1S3GetRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Gets3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
