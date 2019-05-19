// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRule retrieves a rule

Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.
*/
func (a *Client) GetRule(params *GetRuleParams) (*GetRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRule",
		Method:             "GET",
		PathPattern:        "/rules/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRuleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRuleOK), nil

}

/*
GetVersion gets service version

This endpoint returns the service version typically notated using semantic versioning.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) GetVersion(params *GetVersionParams) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionOK), nil

}

/*
GetWellKnownJSONWebKeys lists cryptographic keys

This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.
*/
func (a *Client) GetWellKnownJSONWebKeys(params *GetWellKnownJSONWebKeysParams) (*GetWellKnownJSONWebKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWellKnownJSONWebKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWellKnownJSONWebKeys",
		Method:             "GET",
		PathPattern:        "/.well-known/jwks.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWellKnownJSONWebKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWellKnownJSONWebKeysOK), nil

}

/*
IsInstanceAlive checks alive status

This endpoint returns a 200 status code when the HTTP server is up running.
This status does currently not include checks whether the database connection is working.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceAlive(params *IsInstanceAliveParams) (*IsInstanceAliveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceAliveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isInstanceAlive",
		Method:             "GET",
		PathPattern:        "/health/alive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceAliveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsInstanceAliveOK), nil

}

/*
IsInstanceReady checks readiness status

This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.
the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceReady(params *IsInstanceReadyParams) (*IsInstanceReadyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceReadyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isInstanceReady",
		Method:             "GET",
		PathPattern:        "/health/ready",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceReadyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsInstanceReadyOK), nil

}

/*
ListRules lists all rules

This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full
view of what rules you have currently in place.
*/
func (a *Client) ListRules(params *ListRulesParams) (*ListRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRules",
		Method:             "GET",
		PathPattern:        "/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRulesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRulesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}