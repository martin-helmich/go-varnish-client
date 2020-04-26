package varnishclient

import "context"

// Response contains the data that was received from Varnish in response to a request
type Response struct {
	Code int
	Body []byte
}

// Request contains the data sent to Varnish as a request
type Request struct {
	Method    string
	Arguments []string
}

// Roundtrip defines the interface for sending requests to Varnish and receiving
// responses
type Roundtrip interface {
	ReadHello(ctx context.Context) (*Response, error)
	Execute(ctx context.Context, req *Request) (*Response, error)
}
