package client

import (
	"errors"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"testing"
	"time"
)

const (
	featureNotInStore = "feature-not-in-store"
	feature1_on       = "feature1_on"
	feature2_error    = "feature2_error"
)

// Create mock Evaluator that returns variations or errors based on flag key
type mockEvaluator struct{}
func (e *mockEvaluator) Evaluate(
	feature *dtos.Feature,
	identity string,
	attributes *map[string]interface{})(*dtos.UnlaunchFeature, error) {

	switch feature.Key {
	case feature1_on:
		return &dtos.UnlaunchFeature{
			Feature:                "feature1_on",
			Variation:              "on",
			VariationConfiguration: nil,
			EvaluationReason:       "test 1",
		}, nil
	case feature2_error:
		return nil, errors.New("ignore")

	default:
		return nil, nil
	}
}

var mockEventsRecorderCalls = make(map[string]bool) // to record function calls
type mockEventsRecorder struct{}
func (e *mockEventsRecorder) Shutdown() {
	mockEventsRecorderCalls["Shutdown"] = true
}
func (e *mockEventsRecorder) Record(event *dtos.Event) error {
	mockEventsRecorderCalls[event.Impression.FlagKey] = true
	return nil
}

var mockEventsCountAggregatorRecordCalls = make(map[string]string) // to record "Record()" calls with variation key
var mockEventsCountAggregatorCalls = make(map[string]bool) // to record function calls
type mockEventsCountAggregator struct{}
func (e *mockEventsCountAggregator) Shutdown() {
	mockEventsCountAggregatorCalls["Shutdown"] = true
}
func (e *mockEventsCountAggregator) Record(flagKey string, variationKey string) error {
	mockEventsCountAggregatorRecordCalls[flagKey] = variationKey
	return nil
}

var mockFeatureStoreCalls = make(map[string]bool) // to record function calls
type mockFeatureStore struct {
	ready bool
}
func (e *mockFeatureStore) GetFeature(key string) (*dtos.Feature, error) {
	switch key {
	case feature1_on:
		return &dtos.Feature{
			Key:          feature1_on,
			Name:         "",
			State:        "",
			Variations:   nil,
			OffVariation: 0,
			Rules:        nil,
		}, nil
	case feature2_error:
		return &dtos.Feature{
			Key:          feature2_error,
			Name:         "",
			State:        "",
			Variations:   nil,
			OffVariation: 0,
			Rules:        nil,
		}, nil
	default:
		return nil, nil
	}
}
func (e *mockFeatureStore) Ready(timeout time.Duration) {
	return
}
func (e *mockFeatureStore) Shutdown() {
	mockFeatureStoreCalls["Shutdown"] = true
}
func (e *mockFeatureStore) IsReady() bool {
	return e.ready
}

var mfs = &mockFeatureStore{}
func clientWithMocks() *UnlaunchClient {
	mfs.ready = true

	return &UnlaunchClient{
		sdkKey:          "prod-server-abc",
		pollingInterval: 2000,
		httpTimeout:     3000,
		FeatureStore: mfs,
		eventsRecorder: &mockEventsRecorder{},
		eventsCountAggregator: &mockEventsCountAggregator{},
		logger: logger.NewLogger(nil),
		evaluator: &mockEvaluator{},
	}
}

// Reset all mock recorders that we assert on to default values
func reset() {
	mfs.ready = true // reset
	mockEventsCountAggregatorRecordCalls = make(map[string]string)
	mockFeatureStoreCalls = make(map[string]bool)
	mockEventsCountAggregatorCalls = make(map[string]bool)
	mockEventsRecorderCalls = make(map[string]bool)
}

func TestWhen_FeatureKeyIsEmpty_Then_ControlIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Variation("", "u123", nil)

	if v != "control" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "control", v))
	}
}

func TestWhen_IdentityIsEmpty_Then_ControlIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Variation(feature1_on, "", nil)

	if v != "control" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "control", v))
	}
}

func TestWhen_FeatureNotInStore_Then_ControlIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Variation(featureNotInStore, "u123", nil)

	if v != "control" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "control", v))
	}
}

func TestWhen_FeatureIsInStore_Then_VariationIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Variation(feature1_on, "u123", nil)

	if v != "on" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "on", v))
	}

	if !mockEventsRecorderCalls[feature1_on] {
		t.Error("event recorder should have been called")
	}

	if mockEventsCountAggregatorRecordCalls[feature1_on] != "on" {
		t.Error("event count recorder should have been called")
	}
}

func TestWhen_ClientIsReady_BlockUntilReadyFunctionReturnsImmediately(t *testing.T) {
	reset()
	c := clientWithMocks()

	c.BlockUntilReady(50 * time.Hour)
}

func TestWhen_FeatureIsInStore_Then_UnlaunchFeatureIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Feature(feature1_on, "u123", nil)

	if v.Variation != "on" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "on", v))
	}

	if v.Feature != feature1_on {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", feature1_on, v.Feature))
	}

	if !mockEventsRecorderCalls[feature1_on] {
		t.Error("event recorder should have been called")
	}

	if mockEventsCountAggregatorRecordCalls[feature1_on] != "on" {
		t.Error("event count recorder should have been called")
	}
}

func TestWhen_SDKIsNotReady_Then_ControlIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()
	mfs.ready = false

	v := c.Variation(feature1_on, "u123", nil)

	if v != "control" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "on", v))
	}

	if mockEventsRecorderCalls[feature1_on] {
		t.Error("event recorder shouldn't have been called")
	}

	if mockEventsCountAggregatorRecordCalls[feature1_on] != "" {
		t.Error("event count recorder shouldn't have been called")
	}
}

func TestWhen_FeatureStoreReturnsError_Then_ControlIsReturned(t *testing.T) {
	reset()
	c := clientWithMocks()

	v := c.Variation(feature2_error, "u123", nil)

	if v != "control" {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", "on", v))
	}

	if mockEventsRecorderCalls[feature1_on] {
		t.Error("event recorder shouldn't have been called")
	}

	if mockEventsCountAggregatorRecordCalls[feature1_on] != "" {
		t.Error("event count recorder shouldn't have been called")
	}
}

func TestWhen_ShutDownIsCalled_Then_AllSubRoutinesAreShutdown(t *testing.T) {
	reset()
	c := clientWithMocks()

	c.Shutdown()

	if !c.IsShutdown() {
		t.Error("client not shutdown")
	}

	if !mockFeatureStoreCalls["Shutdown"] {
		t.Error("feature store shutdown not called")
	}

	if !mockEventsCountAggregatorCalls["Shutdown"] {
		t.Error("count aggregator store shutdown not called")
	}

	if !mockEventsRecorderCalls["Shutdown"] {
		t.Error("count recorder store shutdown not called")
	}
}