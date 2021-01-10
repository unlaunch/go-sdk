package api

import (
	"container/list"
	"encoding/json"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sync"
	"time"
)

type EventsRecorder struct {
	logger        logger.Interface
	url           string
	httpClient    *util.HTTPClient
	queue         *list.List
	mutexQ        *sync.Mutex
	queueSize     int
	name          string
	shutdown      chan bool
}

const itemsToSendBatch = 100

func (e *EventsRecorder) postMetrics() error {
	e.mutexQ.Lock()
	defer e.mutexQ.Unlock()

	if e.queue.Len() == 0 {
		return nil
	}

	var total int

	if e.queue.Len() >= itemsToSendBatch {
		total = itemsToSendBatch
	} else {
		total = e.queue.Len()
	}

	result := make([]*dtos.Event, total)

	for i := 0; i < total; i++ {
		result[i] = e.queue.Remove(e.queue.Front()).(*dtos.Event)
	}

	data, _ := json.Marshal(result)
	e.httpClient.Post(e.url, data)

	return nil
}

func (e *EventsRecorder) Shutdown() {
	e.logger.Debug("Sending shutdown signal to ", e.name)
	e.shutdown <- true
}


func (e *EventsRecorder) Record(event *dtos.Event) error {
	if event == nil {
		return nil
	}

	e.mutexQ.Lock()
	defer e.mutexQ.Unlock()

	e.queue.PushBack(event)

	return nil
}

func NewHTTPEventsRecorder(
	sdkKey string,
	host string,
	url string,
	httpTimeout int,
	flushInterval int,
	queueSize int,
	name string,
	logger logger.Interface) *EventsRecorder {
	er := &EventsRecorder{
		logger:        logger,
		url:           url,
		queue:         list.New(),
		mutexQ:        &sync.Mutex{},
		queueSize:     queueSize,
		name:          name,
		httpClient:    util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
	}
	er.shutdown = util.Schedule(er.postMetrics, time.Duration(flushInterval)*time.Millisecond)
	return er
}