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
	logger       logger.LoggerInterface
	url          string
	httpClient   *util.SimpleHTTPClient
	queue        *list.List
	queueMu      *sync.Mutex
	maxQueueSize int
	name         string
	shutdown     chan bool
}

const itemsToSendInBatch = 500

func (e *SimpleEventsRecorder) copyAndEmptyQueue() *list.List {
	e.queueMu.Lock()
	defer e.queueMu.Unlock()

	if e.queue.Len() <= 0 {
		return nil
	}

	r := e.queue

	// empty out the original list so we don't double count
	e.queue = list.New()

	return r
}

func (e *SimpleEventsRecorder) postMetrics() error {
	r := e.copyAndEmptyQueue()

	if r == nil || r.Len() == 0 {
		return nil
	}

	var total int

	if r.Len() >= itemsToSendInBatch {
		total = itemsToSendInBatch
	} else {
		total = r.Len()
	}

	result := make([]*dtos.Event, total)

	for i := 0; i < total; i++ {
		result[i] = r.Remove(r.Front()).(*dtos.Event)
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
	defer func() {
		s  := e.queue.Len()
		e.queueMu.Unlock()
		if s > e.maxQueueSize {
			e.flush()
		}
	}()

	e.queue.PushBack(event)

	return nil
}

func NewHTTPEventsRecorder(
	httpClient *util.SimpleHTTPClient,
	url string,
	flushInterval int,
	maxQueueSize int,
	name string,
	logger logger.LoggerInterface) *SimpleEventsRecorder {
	er := &SimpleEventsRecorder{
		logger:     logger,
		url:        url,
		queue:      list.New(),
		queueMu:    &sync.Mutex{},
		name:       name,
		maxQueueSize: maxQueueSize,
		httpClient: httpClient,
	}
	er.shutdown = util.RunImmediatelyAndSchedule(er.postMetrics, time.Duration(flushInterval)*time.Millisecond)
	return er
}
