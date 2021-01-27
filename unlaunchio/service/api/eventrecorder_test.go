package api

import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"strconv"
	"testing"
	"time"
)

func TestWhen_FlushIntervalIsHit_Then_ImpressionsArePosted(t *testing.T) {
	reset()
	flushInterval :=  500 * time.Millisecond // this relies on timing; don't change

	he := NewHTTPEventsRecorder(true, &mockHTTPClient{}, "bs", flushInterval, util.MaxInt, "impression", logger.NewLogger(nil))

	// Send 10 events, All of same type
	for i := 0; i < 10; i++ {
		err := he.Record(&dtos.Event{
			CreatedTime: time.Now().UTC().UnixNano() / int64(time.Millisecond), // java time
			Key:         "feature" + strconv.Itoa(i),
			Type:        "IMPRESSION",
			Properties:  nil,
			Sdk:         "Go",
			SdkVersion:  "0.0.1",
			Impression: dtos.Impression{
				FlagKey:          "feature" + strconv.Itoa(i),
				UserID:           "identity" + strconv.Itoa(i),
				VariationKey:     "variation" + strconv.Itoa(i),
				EvaluationReason: "eval",
				MachineName:      "UNKNOWN",
			},
		})

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(250 * time.Millisecond) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] > 0 {
		t.Error("POST shouldn't have been called")
	}

	time.Sleep(500 * time.Millisecond)

	if mockHTTPClientCalls["Post"] != 1 { // because we only send 1 flag and 1 variation, we're expecting 1 event
		t.Errorf("POST should have been called %d times. It was called %d", 1, mockHTTPClientCalls["Post"])
	}

	if len(eventsList) != 10 {
		t.Error("there should have been 10 events, one per flag")
	}

	if eventsList[0].Impression.FlagKey == "feature1" {
		t.Error("fail")
	}
}

func TestWhen_MaxQueueSizeIsReached_Then_ImpressionsArePosted(t *testing.T) {
	reset()
	he := NewHTTPEventsRecorder(true, &mockHTTPClient{}, "bs", 100_000_000, 4, "impression", logger.NewLogger(nil))

	// Send 10 events, All of same type
	for i := 0; i < 10; i++ {
		err := he.Record(&dtos.Event{
			CreatedTime: time.Now().UTC().UnixNano() / int64(time.Millisecond), // java time
			Key:         "feature" + strconv.Itoa(i),
			Type:        "IMPRESSION",
			Properties:  nil,
			Sdk:         "Go",
			SdkVersion:  "0.0.1",
			Impression: dtos.Impression{
				FlagKey:          "feature" + strconv.Itoa(i),
				UserID:           "identity" + strconv.Itoa(i),
				VariationKey:     "variation" + strconv.Itoa(i),
				EvaluationReason: "eval",
				MachineName:      "UNKNOWN",
			},
		})

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(200 * time.Millisecond) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] != 2 {
		t.Errorf("POST should have been called %d times. It was called %d", 2, mockHTTPClientCalls["Post"])
	}
}


func TestWhen_ImpressionsAreDisabled_Then_NothingShouldBePosted(t *testing.T) {
	reset()
	he := NewHTTPEventsRecorder(false, &mockHTTPClient{}, "bs", 100_000_000, 4, "impression", logger.NewLogger(nil))

	// Send 10 events, All of same type
	for i := 0; i < 10; i++ {
		err := he.Record(&dtos.Event{
			CreatedTime: time.Now().UTC().UnixNano() / int64(time.Millisecond), // java time
			Key:         "feature" + strconv.Itoa(i),
			Type:        "IMPRESSION",
			Properties:  nil,
			Sdk:         "Go",
			SdkVersion:  "0.0.1",
			Impression: dtos.Impression{
				FlagKey:          "feature" + strconv.Itoa(i),
				UserID:           "identity" + strconv.Itoa(i),
				VariationKey:     "variation" + strconv.Itoa(i),
				EvaluationReason: "eval",
				MachineName:      "UNKNOWN",
			},
		})

		if err != nil {
			t.Error("fail")
		}
	}

	time.Sleep(200 * time.Millisecond) // give it some time, events shouldn't fire

	if mockHTTPClientCalls["Post"] != 0 {
		t.Errorf("POST should have been called %d times. It was called %d", 0, mockHTTPClientCalls["Post"])
	}
}