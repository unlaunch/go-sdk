package client

import (
	"testing"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)


func offlineClientWithMocks() *OfflineClient {
	mfs.ready = true

	return &OfflineClient{
		FeatureStore:          mfs,
		eventsRecorder:        nil,
		eventsCountAggregator: nil,
		logger:                logger.NewLogger(nil),
		evaluator:             &mockEvaluator{},
	}
}

func TestWhen_OfflineMode(t *testing.T) {
	reset()
	c := offlineClientWithMocks()

	v := c.Variation("flagKey", "u123", nil)

	if v != "control" {
		t.Errorf("Expected '%s'. Got '%s'", "control", v)
	}
}
