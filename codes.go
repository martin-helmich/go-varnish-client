package varnishclient

// These constants define the usual response codes returned by the Varnish
// admin server.
const (
	ResponseSyntaxError            = 100
	ResponseUnknownCommand         = 101
	ResponseUnimplemented          = 102
	ResponseTooFewArguments        = 104
	ResponseTooManyArguments       = 105
	ResponseParams                 = 106
	ResponseAuthenticationRequired = 107
	ResponseOK                     = 200
	ResponseCant                   = 300
	ResponseComms                  = 400
	ResponseClose                  = 500
)
