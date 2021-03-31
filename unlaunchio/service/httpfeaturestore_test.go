package service

import (
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sync"
	"testing"
	"time"
)

type mockHTTPClient struct {
	returnValidJSON bool
}

var mockHTTPClientCallsMu = &sync.Mutex{}
var mockHTTPClientCalls = make(map[string]bool) // to record function calls
func (h *mockHTTPClient) Get(path string) ([]byte, error) {
	mockHTTPClientCallsMu.Lock()
	defer mockHTTPClientCallsMu.Unlock()

	mockHTTPClientCalls["Get"] = true

	if h.returnValidJSON {
		return []byte("{}"), nil
	}

	return nil, nil
}

func (h *mockHTTPClient) Post(path string, body []byte) error {
	mockHTTPClientCalls["Post"] = true
	return nil
}

func TestWhen_PollingIntervalIsHit_Then_FetchFlagsIsCalled(t *testing.T) {
	reset()


	_ = getHTTPFeatureStore()

	time.Sleep(1 * time.Second)

	mockHTTPClientCallsMu.Lock()
	if !mockHTTPClientCalls["Get"] {
		t.Error(fmt.Sprintf("Expected HTTP Get to be called"))
	}
	mockHTTPClientCallsMu.Unlock()

	// wait for it to be called again
	reset()
	mockHTTPClientCallsMu.Lock()
	if mockHTTPClientCalls["Get"] {
		t.Error(fmt.Sprintf("Expected HTTP Get NOT to be called"))
	}
	mockHTTPClientCallsMu.Unlock()

	time.Sleep(1 * time.Second)

	if !mockHTTPClientCalls["Get"] {
		t.Error(fmt.Sprintf("Expected HTTP Get to be called"))
	}
}

func TestWhen_DataIsNotReturned_Then_FeatureStoreIsNotReady(t *testing.T) {
	reset()
	fs := getHTTPFeatureStore()
	time.Sleep(1000 * time.Millisecond)
	if fs.IsReady() {
		t.Error(fmt.Sprintf("Expected not ready because valid data was not returned"))
	}
}

func TestWhen_DataIsReturned_Then_FeatureStoreIsReady(t *testing.T) {
	reset()
	h := &mockHTTPClient{}
	h.returnValidJSON = true

	fs := NewHTTPFeatureStore(h, 100000000, logger.NewLogger(nil), false, nil)

	time.Sleep(100 * time.Millisecond)

	if !fs.IsReady() {
		t.Error(fmt.Sprintf("Expected ready because valid data was returned"))
	}

}

func getHTTPFeatureStore() FeatureStore {
	h := NewHTTPFeatureStore(
		&mockHTTPClient{},
		900,
		logger.NewLogger(nil), false, nil)
	return h
}

func reset() {
	mockHTTPClientCallsMu.Lock()
	defer mockHTTPClientCallsMu.Unlock()
	mockHTTPClientCalls = make(map[string]bool)
}
