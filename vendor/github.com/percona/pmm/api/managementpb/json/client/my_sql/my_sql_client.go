// Code generated by go-swagger; DO NOT EDIT.

package my_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new my sql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for my sql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddMySQL adds my SQL adds my SQL service and starts several agents it automatically adds a service to inventory which is running on provided node id then adds mysqld exporter and qan mysql perfschema agents with provided pmm agent id and other parameters
*/
func (a *Client) AddMySQL(params *AddMySQLParams) (*AddMySQLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySQLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQL",
		Method:             "POST",
		PathPattern:        "/v0/management/MySQL/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySQLReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddMySQLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMySQLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
