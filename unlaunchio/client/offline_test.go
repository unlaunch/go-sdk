package client

import (
	"testing"
)

func offlineClient() ClientInterface {

	cfg := DefaultConfig()
	cfg.OfflineMode = true
	factory, error := NewUnlaunchClientFactory("sdk-key", cfg)
	if error != nil {
	}
	offline := factory.Client()
	return offline

}

func TestWhen_OfflineMode(t *testing.T) {
	c := offlineClient()

	v := c.Variation("flagKey", "u123", nil)

	if v != "control" {
		t.Errorf("Expected '%s'. Got '%s'", "control", v)
	}
}

func TestWhen_OfflineMode_CallFeature(t *testing.T) {
	c := offlineClient()

	f := c.Feature("flagKey", "u123", nil)

	if f.Variation != "control" {
		t.Errorf("Expected '%s'. Got '%s'", "control", f.Variation)
	}
	expectedReson := "Client is initialized in Offline Mode. Returning 'control' variation for all flags."
	if f.EvaluationReason != expectedReson {
		t.Errorf("Expected '%s'. Got '%s'", expectedReson, f.EvaluationReason)
	}
}

func TestWhen_ShutdownIsCalled(t *testing.T) {
	c := offlineClient()

	c.Shutdown()

	if !c.IsShutdown() {
		t.Error("client not shutdown")
	}

}
