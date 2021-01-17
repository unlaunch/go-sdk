package api

import (
	"encoding/json"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"strings"
	"sync"
	"time"
)

type EventsCountAggregator interface {
	Shutdown()
	Record(flagKey string, variationKey string) error
}

type SimpleEventsCountAggregator struct {
	logger         logger.LoggerInterface
	queueMu        *sync.Mutex
	queue          map[string]int
	maxQueueSize   int
	url            string
	HTTPClient     *util.SimpleHTTPClient
	eventsRecorder *SimpleEventsRecorder
	shutdown       chan bool
}

func (e *SimpleEventsCountAggregator) Shutdown() {
	e.logger.Debug("Sending shutdown signal to count aggregator and flushing")
	e.flush()
	e.shutdown <- true
}

func (e *SimpleEventsCountAggregator) flush() {
	e.postMetrics()
}

func (e *SimpleEventsCountAggregator) copyAndEmptyQueue() map[string]int {
	e.queueMu.Lock()
	defer e.queueMu.Unlock()

	r := make(map[string]int)
	for k, v := range e.queue {
		r[k] = v
	}

	// empty out the original map so we don't double count
	e.queue = make(map[string]int)

	return r
}

func (e *SimpleEventsCountAggregator) postMetrics() error {
	rawEvents := e.copyAndEmptyQueue()

	// nothing to do
	if len(rawEvents) == 0 {
		return nil
	}

	eventsList := make([]*dtos.Event, len(rawEvents))
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

	data, _ := json.Marshal(eventsList)
	e.HTTPClient.Post(e.url, data)
	return nil
}

func (e *SimpleEventsCountAggregator) Record(flagKey string, variationKey string) error {
	if flagKey == "" || variationKey == "" {
		return nil
	}

	e.queueMu.Lock()

	defer func() {
		s  := len(e.queue)
		e.queueMu.Unlock()
		if s > e.maxQueueSize {
			e.flush()
		}
	}()

	e.queue[flagKey + "," + variationKey]+= 1

	return nil
}

func NewEventsCountAggregator(HTTPClient *util.SimpleHTTPClient, url string, flushInterval int, maxQueueSize int, logger logger.LoggerInterface) *SimpleEventsCountAggregator {
	ec := &SimpleEventsCountAggregator{
		logger:         logger,
		queueMu:        &sync.Mutex{},
		url:            url,
		HTTPClient:     HTTPClient,
		queue:          make(map[string]int),
		maxQueueSize: maxQueueSize,
		eventsRecorder: nil,
	}

	ec.shutdown = util.Schedule(ec.postMetrics, time.Duration(flushInterval)*time.Millisecond)
	return ec
}
