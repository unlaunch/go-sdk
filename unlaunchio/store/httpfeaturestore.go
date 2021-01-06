package store

import (
	"encoding/json"
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sync"
	"time"
)

type HttpFeatureStore struct {
	service  util.HTTPService
	logger   logger.Interface
	features map[string]dtos.Feature
	initialSyncComplete bool

}

func (h *HttpFeatureStore) Stop()  {
	stop <- true
}


func (h *HttpFeatureStore) fetchFlags() ([]byte, error) {
	if h.initialSyncComplete == false {
		defer wg.Done()
		h.initialSyncComplete = true
	}

	res, err := h.service.Get("/api/v1/flags")

	if err != nil {
		h.logger.Error("error fetching flags ", err)
	}

	//h.logger.Debug("responseDto ", string(res))

	var responseDto dtos.TopLevelEnvelope
	err = json.Unmarshal(res, &responseDto)

	if err != nil {
		h.logger.Error("Error parsing split changes JSON ", err)
		return nil, err
	}

	h.logger.Debug("responseDto ", responseDto)

	// Store features in the store/map
	for _, feature := range responseDto.Data.Features {
		h.features[feature.Key] = feature
	}

	return res, nil
}

func (h *HttpFeatureStore) GetFeature(key string) (*dtos.Feature, error) {
	if feature, ok := h.features[key]; ok {
		return &feature, nil
	} else {
		return nil, errors.New("flag was not found in local storage")
	}
}

func (h *HttpFeatureStore) Ready()  {
	if h.initialSyncComplete {
		return
	}
	wg.Wait()
}

var wg sync.WaitGroup
var stop chan bool


func NewHTTPFeatureStore(
	sdkKey string,
	host string,
	httpTimeout int,
	pollingInterval int,
	logger logger.Interface,
) FeatureStore {

	httpStore := &HttpFeatureStore{
		service:  util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
		logger:   logger,
		initialSyncComplete: false,
		features: make(map[string]dtos.Feature),
	}

	wg.Add(1)
	stop = util.Schedule(httpStore.fetchFlags, time.Duration(pollingInterval)*time.Millisecond)

	return httpStore
}
