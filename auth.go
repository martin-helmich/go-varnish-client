package varnishclient

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Authenticate authenticates the client with Varnish.
// This method only does something then the Varnish admin server actually
// requires authentication (you can test for that using the
// AuthenticationRequired method).
//
// The "secret" parameter should be the binary content read from the Varnish
// secret file.
//
// If you call this method then the server does not require authentication, it
// will simply do nothing and return "nil".
func (c *Client) Authenticate(secret []byte) error {
	if !c.authenticationRequired {
		return nil
	}

	input := string(c.authChallenge) + "\n" + string(secret) + string(c.authChallenge) + "\n"
	response := sha256.Sum256([]byte(input))
	responseHex := hex.EncodeToString(response[:])

	resp, err := c.sendRequest("auth", responseHex)
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("response code was %d, expected %d", resp.Code, ResponseOK)
	}

	c.authenticated = true

	return nil
}

// AuthenticationRequired tells you if the server requires the client to
// authenticate itself before doing anything.
//
// If this returns "true", the next call should be to "Authenticate".
func (c *Client) AuthenticationRequired() bool {
	return c.authenticationRequired
}
