// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// FindClustersReader is a Reader for the FindClusters structure.
type FindClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindClustersOK creates a FindClustersOK with default headers values
func NewFindClustersOK() *FindClustersOK {
	return &FindClustersOK{}
}

/* FindClustersOK describes a response with status code 200, with default header values.

OK
*/
type FindClustersOK struct {
	Payload []*models.V1ClusterResponse
}

func (o *FindClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/find][%d] findClustersOK  %+v", 200, o.Payload)
}
func (o *FindClustersOK) GetPayload() []*models.V1ClusterResponse {
	return o.Payload
}

func (o *FindClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindClustersDefault creates a FindClustersDefault with default headers values
func NewFindClustersDefault(code int) *FindClustersDefault {
	return &FindClustersDefault{
		_statusCode: code,
	}
}

/* FindClustersDefault describes a response with status code -1, with default header values.

Error
*/
type FindClustersDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find clusters default response
func (o *FindClustersDefault) Code() int {
	return o._statusCode
}

func (o *FindClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/find][%d] findClusters default  %+v", o._statusCode, o.Payload)
}
func (o *FindClustersDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
