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

type EventsRecorder interface {
	Shutdown()
	Record(event *dtos.Event) error
}

type SimpleEventsRecorder struct {
	logger     logger.LoggerInterface
	url        string
	httpClient *util.SimpleHTTPClient
	queue      *list.List
	queueMu    *sync.Mutex
	queueSize  int
	name       string
	shutdown   chan bool
}

const itemsToSendBatch = 100

func (e *SimpleEventsRecorder) postMetrics() error {
	e.logger.Debug("er RUNNING")
	e.queueMu.Lock()
	defer e.queueMu.Unlock()

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

func (e *SimpleEventsRecorder) Shutdown() {
	e.logger.Debug("Flushing and sending shutdown signal to ", e.name)
	e.flush()
	e.shutdown <- true
}

func (e *SimpleEventsRecorder) flush() {
	e.postMetrics()
}

func (e *SimpleEventsRecorder) Record(event *dtos.Event) error {
	if event == nil {
		return nil
	}

	e.queueMu.Lock()
	defer e.queueMu.Unlock()

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
	logger logger.LoggerInterface) *SimpleEventsRecorder {
	er := &SimpleEventsRecorder{
		logger:     logger,
		url:        url,
		queue:      list.New(),
		queueMu:    &sync.Mutex{},
		queueSize:  queueSize,
		name:       name,
		httpClient: util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
	}
	er.shutdown = util.Schedule(er.postMetrics, time.Duration(flushInterval)*time.Millisecond)
	return er
}
