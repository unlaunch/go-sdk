package client

import (
	"testing"
	"time"
)

func TestWhen_NullConfigIsPassedForProdSDKKey_Then_ValuesAreInitializedToReasonableDefaults(t *testing.T) {
	f, _ := NewUnlaunchClientFactory("prod-server-abc", nil) // prod SDK key begins with "prod" prefix

	cfg := DefaultConfig()
	if f.cfg.Host != cfg.Host {
		t.Error("unexpected config value")
	}

	if f.cfg.PollingInterval != cfg.PollingInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsFlushInterval != cfg.MetricsFlushInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsQueueSize != cfg.MetricsQueueSize {
		t.Error("unexpected config value")
	}
}

func TestWhen_NullConfigIsPassedForNonProdSDKKey_Then_ValuesAreInitializedToReasonableDefaults(t *testing.T) {
	f, _ := NewUnlaunchClientFactory("xyz", nil) // non-prod SDK key don't have "prod" prefix

	if f.cfg.Host != "https://api.unlaunch.io" {
		t.Error("unexpected config value")
	}
}

func TestWhen_InvalidValuesArePassed_Then_TheyAreResetToEnvironmentDefaults(t *testing.T) {
	cfg := &UnlaunchClientConfig{
		PollingInterval:      1 * time.Second,   // too aggressive; will be reset
		MetricsFlushInterval: 120 * time.Second, // ok
		MetricsQueueSize:     1,      // too aggressive; will be reset
		HTTPTimeout:          1 * time.Second,   // ok
	}

	f, _ := NewUnlaunchClientFactory("prod-server-abc", cfg)

	if f.cfg.PollingInterval == 1 * time.Second {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsFlushInterval != 120 * time.Second {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsQueueSize == 1 {
		t.Error("unexpected config value")
	}
}
