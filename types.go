package varnishclient

import (
	"context"
)

// Backend is a single item of the list returned by the `ListBackends` method
type Backend struct {
	Name string
}

// BackendsResponse is the respose type of the `ListBackends` method
type BackendsResponse []Backend

// Parameter is a single item of the list returned by the `ListParameters` method
type Parameter struct {
	Name      string
	Value     string
	Unit      string
	IsDefault bool
}

// ParametersResponse is the response type of the `ListParameters` method
type ParametersResponse []Parameter

// Client contains the most common Varnish administration operations (and the
// necessary tools to build your own that are not yet implemented)
type Client struct {
	authChallenge []byte
	roundtrip     Roundtrip

	authenticationRequired bool
	authenticated          bool
}

// Type guard to assert that *Client actually implements ClientInterface
var _ ClientInterface = &Client{}

// ClientInterface describes the common methods offered by the Varnish client
type ClientInterface interface {
	AuthenticationRequired() bool
	Authenticate(ctx context.Context, secret []byte) error
	ListBackends(ctx context.Context, pattern string) (BackendsResponse, error)

	SetParameter(ctx context.Context, name, value string) error
	ListParameters(ctx context.Context) (ParametersResponse, error)

	DiscardVCL(ctx context.Context, configName string) error
	DefineInlineVCL(ctx context.Context, configName string, vcl []byte, mode VCLState) error
	AddLabelToVCL(ctx context.Context, label string, configName string) error
	LoadVCL(ctx context.Context, configName, filename string, mode VCLState) error
	UseVCL(ctx context.Context, configName string) error
	SetVCLState(ctx context.Context, configName string, state VCLState) error
}
