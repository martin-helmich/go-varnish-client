package varnishclient

func ExampleClient_SetVCLState() {
	if err := exampleClient.SetVCLState(ctx, "boot", VCLStateCold); err != nil {
		// handle error
	}
}
