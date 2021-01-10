package api

import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"strings"
	"sync"
	"time"
)

type EventsCountAggregator struct {
	logger logger.Interface
	mutexM *sync.Mutex
	store  map[string]int
	eventsRecorder *EventsRecorder
}

func (e *EventsCountAggregator) copyAndEmptyMap() map[string]int {
	e.mutexM.Lock()
	defer e.mutexM.Unlock()

	r := make(map[string]int)
	for k, v := range e.store {
		r[k] = v
	}

	// empty out the original map so we don't double count
	e.store = make(map[string]int)

	return r
}

func (e *EventsCountAggregator) postMetrics() error {
	rawEvents := e.copyAndEmptyMap()

	// nothing to do
	if len(e.store) == 0 {
		return nil
	}

	eventsList := make([]*dtos.Event, len(e.store))
	for k, v := range rawEvents {
		d := strings.Split(k, ",")
		f := d[0]
		varKey := d[1]

		p := make(map[string]interface{})
		p[varKey] = v

		event := &dtos.Event{
			CreatedTime:  time.Now().UTC().Unix() * 1000,
			Key:          f,
			Type: "VARIATIONS_COUNT_EVENT",
			Properties:   p,
			Sdk:          "Go",
			SdkVersion:   "0.0.1",
		}

		eventsList = append(eventsList, event)
	}

	//data, _ := json.Marshal(eventsList)

	return nil



	return nil
}

func (e *EventsCountAggregator) Record(flagKey string, variationKey string) error {
	if flagKey == "" || variationKey == "" {
		return nil
	}

	e.mutexM.Lock()
	defer e.mutexM.Unlock()

	e.store[flagKey + "," + variationKey]+= 1

	return nil
}

func NewEventsCountAggregator(logger logger.Interface) *EventsCountAggregator {
	return &EventsCountAggregator {
		logger: logger,
		mutexM: &sync.Mutex{},
		store: make(map[string]int),
		eventsRecorder: nil,
	}
}
