package apollet

import (
	"testing"
)

var (
	fakedAppId         = "SampleApp"
	fakedCluster       = "default"
	fakedNameSpaceName = "application"
	fakedServerIP      = "localhost:8080"
	fakedKey           = "timeout"
)

func TestClientGetString(t *testing.T) {
	// init client
	client := NewClient(
		fakedServerIP,
	)

	// get value
	value := client.GetString(fakedAppId, fakedCluster, fakedNameSpaceName, fakedKey, "")
	if value == "" {
		t.Fatalf("Err get string value!")
	}

	t.Log(fakedKey, "=", value)
	t.Log("Test passed!")
}
