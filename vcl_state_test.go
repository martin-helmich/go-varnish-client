package varnishclient

var exampleClient Client

func ExampleClient_SetVCLState() {
	if err := exampleClient.SetVCLState("boot", VCLStateCold); err != nil {
		// handle error
	}
}
