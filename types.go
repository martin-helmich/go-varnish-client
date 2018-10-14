package varnishclient

import (
	"bufio"
	"io"
)

type Response struct {
	Code int
	Body []byte
}

type Backend struct {
	Name string
}

type BackendsResponse []Backend

type Parameter struct {
	Name      string
	Value     string
	Unit      string
	IsDefault bool
}

type ParametersResponse []Parameter

type client struct {
	authChallenge []byte
	reader        io.Reader
	writer        io.Writer
	scanner       *bufio.Scanner

	authenticationRequired bool
	authenticated          bool
}

type Client interface {
	AuthenticationRequired() bool
	Authenticate([]byte) error
	ListBackends(pattern string) (BackendsResponse, error)

	SetParameter(name, value string) error
	ListParameters() (ParametersResponse, error)

	DiscardVCL(configName string) error
	DefineInlineVCL(configName string, vcl []byte, mode string) error
	AddLabelToVCL(label string, configName string) error
	LoadVCL(configName, filename string, mode string) error
	UseVCL(configName string) error
}
