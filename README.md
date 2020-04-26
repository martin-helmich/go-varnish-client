# Go client for the Varnish administration port

This repository contains a client library that can be used to connect to the Varnish administration port (typically used by the [varnishadm](https://varnish-cache.org/docs/trunk/reference/varnishadm.html) tool).

**NOTE**: Experimental. Use at own peril!

## Installation

Install this library via `go get`:

```
$ go get github.com/martin-helmich/go-varnish-client
```

## Usage

### Establish a connection

First, connect to the administration port using the `varnishclient.DialTCP` method:

```go
client, err := varnishclient.DialTCP(context.Background(), "127.0.0.1:6082")
if err != nil {
    panic(err)
}
```

### Authenticate

You can then use the `client.AuthenticationRequired()` method to check if the Varnish server requires authentication.
Then, use the `client.Authenticate()` method to perform the authentication:

```go
secret, _ := ioutil.ReadFile("/etc/varnish/secret")

if client.AuthenticationRequired() {
    err := client.Authenticate(context.Background(), secret)
    if err != nil {
        panic(err)
    }
}
```

### Define and update a new configuration

```go
ctx := context.Background()

err := client.DefineInlineVCL(ctx, "my-new-config", vclCode, "warm")
if err != nil {
    panic(err)
}

err = client.UseVCL(ctx, "my-new-config")
if err != nil {
    panic(err)
}
```

### Define timeouts/cancellations on operations

All operations accept a `context.Context` parameter that can be used for timeouts and/or cancellations:

```go
ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
defer cancel()

err := client.DefineInlineVCL(ctx, "my-new-config", vclCode, "warm")
if err != nil {
    panic(err)
}
```