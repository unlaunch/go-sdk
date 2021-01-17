package client

import "testing"

func TestWhen_NullConfigIsPassedForProdSDKKey_Then_ValuesAreInitializedToReasonableDefaults(t *testing.T) {
	f, _ := NewUnlaunchClientFactory("prod-server-abc", nil) // prod SDK key begins with "prod" prefix

	if f.cfg.Host != prodConfigMinValues.host {
		t.Error("unexpected config value")
	}

	if f.cfg.PollingInterval != prodConfigMinValues.minPollingInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsFlushInterval != prodConfigMinValues.minMetricsFlushInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsQueueSize != prodConfigMinValues.minMetricsQueueSize {
		t.Error("unexpected config value")
	}
}

func TestWhen_NullConfigIsPassedForNonProdSDKKey_Then_ValuesAreInitializedToReasonableDefaults(t *testing.T) {
	f, _ := NewUnlaunchClientFactory("xyz", nil) // pnon-rod SDK key don't have "prod" prefix

	if f.cfg.Host != debugConfigMinValues.host {
		t.Error("unexpected config value")
	}

	if f.cfg.PollingInterval != debugConfigMinValues.minPollingInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsFlushInterval != debugConfigMinValues.minMetricsFlushInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsQueueSize != debugConfigMinValues.minMetricsQueueSize {
		t.Error("unexpected config value")
	}
}

func TestWhen_InvalidValuesArePassed_Then_TheyAreResetToEnvironmentDefaults(t *testing.T) {
	cfg := &UnlaunchClientConfig{
		PollingInterval:      1000,   // too aggressive; will be reset
		MetricsFlushInterval: 120000, // ok
		MetricsQueueSize:     1,      // too aggressive; will be reset
		HTTPTimeout:          1000,   // ok
	}

	f, _ := NewUnlaunchClientFactory("prod-server-abc", cfg)

	if f.cfg.PollingInterval != prodConfigMinValues.minPollingInterval {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsFlushInterval != 120000 {
		t.Error("unexpected config value")
	}

	if f.cfg.MetricsQueueSize != prodConfigMinValues.minMetricsQueueSize {
		t.Error("unexpected config value")
	}
}
