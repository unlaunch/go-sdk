package api

import (
	"encoding/json"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

type EventsRecorder struct {
	logger     logger.Interface
	url        string
	httpClient *util.HTTPClient

}

func (e *EventsRecorder) Record(impressions [1]*dtos.Event) error {

	data, _ := json.Marshal(impressions)
	return e.httpClient.Post(e.url, data)
}

func NewHTTPEventsRecorder(
	sdkKey string,
	host string,
	url string,
	httpTimeout int,
	logger logger.Interface) *EventsRecorder {
	return &EventsRecorder{
		logger:     logger,
		url:        url,
		httpClient: util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
	}

}
