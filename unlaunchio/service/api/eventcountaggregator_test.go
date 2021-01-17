package api

import (
	"encoding/json"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"strconv"
	"testing"
	"time"
)

type mockHTTPClient struct {}
var mockHTTPClientCalls = make(map[string]int) // to record function calls
var eventsList = make([]*dtos.Event, 10)
func (h *mockHTTPClient) Get(path string) ([]byte, error) {
	mockHTTPClientCalls["Get"] = mockHTTPClientCalls["Get"] + 1
	return nil, nil
}

func (h *mockHTTPClient) Post(path string, body []byte) error {
	_ = json.Unmarshal(body, &eventsList)

	mockHTTPClientCalls["Post"] = mockHTTPClientCalls["Post"] + 1
	return nil
}

func TestWhen_FlushIntervalIsHit_Then_AggregatedCountsArePosted(t *testing.T) {
	reset()
	flushInterval := 500 // this relies on timing; don't change
	ce := NewEventsCountAggregator(&mockHTTPClient{}, "bs", flushInterval, util.MAX_INT, logger.NewLogger(nil))

	// Send 10 events, All of same type
	for i := 0; i<10; i++ {
		err := ce.Record("abc", "xyz")

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(250 * time.Millisecond ) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] > 0 {
		t.Error("POST shouldn't have been called")
	}

	time.Sleep(600 * time.Millisecond ) // give it more time than flush interval

	if mockHTTPClientCalls["Post"] != 1 { // because we only send 1 flag and 1 variation, we're expecting 1 event
		t.Errorf("POST should have been called %d times. It was called %d", 1, mockHTTPClientCalls["Post"])
	}
}

func TestWhen_FlushIntervalIsHit_Then_AggregatedCountsArePosted_OnePerFlag(t *testing.T) {
	reset()
	flushInterval := 500 // this relies on timing; don't change
	ce := NewEventsCountAggregator(&mockHTTPClient{}, "bs", flushInterval, util.MAX_INT, logger.NewLogger(nil))

	// Send 10 events for different flags
	for i := 0; i<10; i++ {
		err := ce.Record("abc" + strconv.Itoa(i), "xyz")

		// Send 2 variations for 1 of the flags. both should be combined
		if i == 8 {
			ce.Record("abc" + strconv.Itoa(i), "xyz")
		}

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(250 * time.Millisecond ) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] > 0 {
		t.Error("POST shouldn't have been called")
	}

	time.Sleep(600 * time.Millisecond ) // give it more time than flush interval


	if mockHTTPClientCalls["Post"] != 1 { // because we only send 1 flag and 1 variation, we're expecting 1 event
		t.Errorf("POST should have been called %d times. It was called %d", 10, mockHTTPClientCalls["Post"])
	}

	if len(eventsList) != 10 {
		t.Error("there should have been 10 events, one per flag")
	}

}

func TestWhen_MaxQueueSizeIsReached_Then_AggregatedCountsArePosted(t *testing.T) {
	reset()
	ce := NewEventsCountAggregator(&mockHTTPClient{}, "bs", 100_000_000, 4, logger.NewLogger(nil))

	// Send 10 events, All of same type
	for i := 0; i<10; i++ {
		err := ce.Record("abc" +  strconv.Itoa(i), "xyz")

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(200 * time.Millisecond ) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] != 2 {
		t.Errorf("POST should have been called %d times. It was called %d", 2, mockHTTPClientCalls["Post"])
	}
}

func reset() {
	mockHTTPClientCalls = make(map[string]int)
	eventsList = make([]*dtos.Event, 10)
}