package varnishclient

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

type roundtrip struct {
	reader io.Reader
	writer io.Writer
}

func (rt *roundtrip) writeBytes(in []byte) (int, error) {
	glog.V(8).Infof("writing %d bytes to server: %x (%s)", len(in), in, strconv.Quote(string(in)))
	return rt.writer.Write(in)
}

func (rt *roundtrip) writeBytesCtx(ctx context.Context, in []byte) (int, error) {
	ci := make(chan int)
	ce := make(chan error)

	go func() {
		n, err := rt.writeBytes(in)
		if err != nil {
			ce <- err
		} else {
			ci <- n
		}
	}()

	select {
	case n := <-ci:
		return n, nil
	case err := <-ce:
		return 0, err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func (rt *roundtrip) writeString(ctx context.Context, in string) (int, error) {
	return rt.writeBytesCtx(ctx, []byte(in))
}

func (rt *roundtrip) readResponse() (*Response, error) {
	header := make([]byte, 13)

	n, err := io.ReadFull(rt.reader, header)
	if err != nil {
		return nil, err
	}

	glog.V(8).Infof("read %d bytes of header", n)
	glog.V(8).Infof("header: %s", strconv.Quote(string(header)))

	code, err := strconv.Atoi(string(header[0:3]))
	if err != nil {
		return nil, err
	}

	blen, err := strconv.Atoi(strings.TrimSpace(string(header[3:11])))
	if err != nil {
		return nil, fmt.Errorf("invalid length: %s", err.Error())
	}

	glog.V(8).Infof("received message from Varnish server: response code %d, body length %d", code, blen)

	body := make([]byte, blen+1)
	m, err := io.ReadFull(rt.reader, body)

	if m != blen+1 {
		return nil, fmt.Errorf("incomplete body: only %d bytes read, %d expected", m, blen)
	}

	glog.V(8).Infof("%d bytes read", m)
	glog.V(8).Infof("message body: %s", strconv.Quote(string(body)))

	response := Response{}
	response.Code = code
	response.Body = body

	return &response, nil
}

func (rt *roundtrip) readResponseCtx(ctx context.Context) (*Response, error) {
	rc := make(chan *Response)
	re := make(chan error)

	go func() {
		res, err := rt.readResponse()
		if err != nil {
			re <- err
		} else {
			rc <- res
		}
	}()

	select {
	case res := <-rc:
		return res, nil
	case err := <-re:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ReadHello reads the initial message from the server to the client that is
// sent upon connecting.
func (rt *roundtrip) ReadHello(ctx context.Context) (*Response, error) {
	return rt.readResponseCtx(ctx)
}

// SendRequest sends a generic request to the Varnish admin server and returns
// the response.
func (rt *roundtrip) Execute(ctx context.Context, req *Request) (*Response, error) {
	cmd := req.Method

	glog.V(8).Infof("writing to server: %s", cmd)
	if len(req.Arguments) > 0 {
		_, err := rt.writeString(ctx, req.Method+" ")
		if err != nil {
			return nil, err
		}

		for i := range req.Arguments {
			_, err := rt.writeString(ctx, req.Arguments[i])
			if err != nil {
				return nil, err
			}

			if i < len(req.Arguments)-1 {
				_, err := rt.writeBytes([]byte{' '})
				if err != nil {
					return nil, err
				}
			}
		}

		_, err = rt.writeBytes([]byte{0x0a})
		if err != nil {
			return nil, err
		}
	} else {
		_, err := rt.writeString(ctx, req.Method+"\n")
		if err != nil {
			return nil, err
		}
	}

	glog.V(8).Infof("request written; waiting for response")

	return rt.readResponseCtx(ctx)
}
