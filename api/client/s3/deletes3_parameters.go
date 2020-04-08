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

// NewDeletes3Params creates a new Deletes3Params object
// with the default values initialized.
func NewDeletes3Params() *Deletes3Params {
	var ()
	return &Deletes3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletes3ParamsWithTimeout creates a new Deletes3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletes3ParamsWithTimeout(timeout time.Duration) *Deletes3Params {
	var ()
	return &Deletes3Params{

		timeout: timeout,
	}
}

// NewDeletes3ParamsWithContext creates a new Deletes3Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeletes3ParamsWithContext(ctx context.Context) *Deletes3Params {
	var ()
	return &Deletes3Params{

		Context: ctx,
	}
}

// NewDeletes3ParamsWithHTTPClient creates a new Deletes3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletes3ParamsWithHTTPClient(client *http.Client) *Deletes3Params {
	var ()
	return &Deletes3Params{
		HTTPClient: client,
	}
}

/*Deletes3Params contains all the parameters to send to the API endpoint
for the deletes3 operation typically these are written to a http.Request
*/
type Deletes3Params struct {

	/*Body*/
	Body *models.V1S3DeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the deletes3 params
func (o *Deletes3Params) WithTimeout(timeout time.Duration) *Deletes3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deletes3 params
func (o *Deletes3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deletes3 params
func (o *Deletes3Params) WithContext(ctx context.Context) *Deletes3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deletes3 params
func (o *Deletes3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deletes3 params
func (o *Deletes3Params) WithHTTPClient(client *http.Client) *Deletes3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deletes3 params
func (o *Deletes3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the deletes3 params
func (o *Deletes3Params) WithBody(body *models.V1S3DeleteRequest) *Deletes3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the deletes3 params
func (o *Deletes3Params) SetBody(body *models.V1S3DeleteRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *Deletes3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
