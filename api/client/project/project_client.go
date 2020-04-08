// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateProject creates a project if the given ID already exists a conflict is returned
*/
func (a *Client) CreateProject(params *CreateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProject",
		Method:             "PUT",
		PathPattern:        "/v1/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateProjectCreated), nil

}

/*
DeleteProject deletes a project and returns the deleted entity
*/
func (a *Client) DeleteProject(params *DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProject",
		Method:             "DELETE",
		PathPattern:        "/v1/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProjectOK), nil

}

/*
FindProject gets project by id
*/
func (a *Client) FindProject(params *FindProjectParams, authInfo runtime.ClientAuthInfoWriter) (*FindProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findProject",
		Method:             "GET",
		PathPattern:        "/v1/project/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindProjectOK), nil

}

/*
ListProjects gets all projects
*/
func (a *Client) ListProjects(params *ListProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProjects",
		Method:             "GET",
		PathPattern:        "/v1/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListProjectsOK), nil

}

/*
UpdateProject updates a project optimistic lock error can occur
*/
func (a *Client) UpdateProject(params *UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProject",
		Method:             "POST",
		PathPattern:        "/v1/project",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProjectOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
