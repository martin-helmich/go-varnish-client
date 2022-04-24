package varnishclient

import (
	"context"
	"time"
)

// Backend is a single item of the list returned by the `ListBackends` method
type Backend struct {
	Name string
}

// BackendsResponse is the respose type of the `ListBackends` method
type BackendsResponse []Backend

type VCLConfigStatus int

const (
	VCLActive VCLConfigStatus = iota
	VCLAvailable
	VCLDiscarded
	VCLUnknown
)

func (v VCLConfigStatus) String() string {
	return [...]string{"active", "available", "discarded", "unknown"}[v]
}

type VCLConfig struct {
	Name           string
	ActiveBackends int
	Status         VCLConfigStatus
}

type VCLConfigsResponse []VCLConfig
type BanListResponse []Ban

const (
	BanActive BanStatus = iota
	BanGone
	BanComplete
)

type BanStatus int

func (v BanStatus) String() string {
	return [...]string{"active", "gone", "complete"}[v]
}

type Ban struct {
	Time    time.Time
	Objects int64
	Status  BanStatus
	Spec    string
}

func (b Ban) Equals(ban Ban) bool {
	if b.Spec != ban.Spec || b.Objects != ban.Objects || b.Status != ban.Status || b.Time != ban.Time {
		return false
	}
	return true
}

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
	Ban(ctx context.Context, args ...string) error
	ListBans(ctx context.Context) (BanListResponse, error)
	SetParameter(ctx context.Context, name, value string) error
	ListParameters(ctx context.Context) (ParametersResponse, error)

	DiscardVCL(ctx context.Context, configName string) error
	DefineInlineVCL(ctx context.Context, configName string, vcl []byte, mode VCLState) error
	AddLabelToVCL(ctx context.Context, label string, configName string) error
	ListVCL(ctx context.Context) (VCLConfigsResponse, error)
	LoadVCL(ctx context.Context, configName, filename string, mode VCLState) error
	UseVCL(ctx context.Context, configName string) error
	SetVCLState(ctx context.Context, configName string, state VCLState) error
}
