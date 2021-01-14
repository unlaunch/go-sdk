package service

import (
	"encoding/json"
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sort"
	"time"
)

type HTTPFeatureStore struct {
	httpClient          *util.HTTPClient
	logger              logger.Interface
	features            map[string]dtos.Feature
	initialSyncComplete bool
	shutdown            chan bool
	readyChannel 		chan bool
}

func (h *HTTPFeatureStore) Stop() {
	h.logger.Debug("Sending shutdown signal to feature store")
	h.shutdown <- true
}

func (h *HTTPFeatureStore) fetchFlags()  error {

	res, err := h.httpClient.Get("/api/v1/flags")

	if err != nil {
		h.logger.Error("error fetching flags ", err)
		return err
	}

	if res == nil {
		// No error and empty response means nothing changed
		// most like due to 304; not modified
		return nil
	}

	// h.logger.Trace("responseDto ", string(res))

	var responseDto dtos.TopLevelEnvelope
	err = json.Unmarshal(res, &responseDto)

	if err != nil {
		h.logger.Error("error parsing feature flag JSON response ", err)
		return err
	}

	if h.initialSyncComplete == false {
		h.readyChannel <- true
		h.initialSyncComplete = true
	}

	// Todo: Remove this when rules and rollouts are sorted on server
	for _, feature := range responseDto.Data.Features {
		sort.Sort(dtos.ByRulePriority(feature.Rules))

		for _, rule := range feature.Rules {
			sort.Sort(dtos.ByVariationId(rule.Rollout))
		}
	}

	// Store features in the service/map
	temp := make(map[string]dtos.Feature)
	for _, feature := range responseDto.Data.Features {
		temp[feature.Key] = feature

	}
	h.features = temp

	h.logger.Debug("Downloaded: ", len(h.features))
	return nil
}

func (h *HTTPFeatureStore) GetFeature(key string) (*dtos.Feature, error) {
	if feature, ok := h.features[key]; ok {
		return &feature, nil
	} else {
		return nil, errors.New("flag was not found in local storage")
	}
}

func (h *HTTPFeatureStore) Ready(timeout time.Duration) {
	if h.initialSyncComplete {
		return
	}

	select {
	case <- h.readyChannel:
	case <-time.After(timeout):
		h.logger.Error("The client wasn't ready in the required time")
	}




}



func NewHTTPFeatureStore(
	sdkKey string,
	host string,
	httpTimeout int,
	pollingInterval int,
	logger logger.Interface,
) FeatureStore {
	httpStore := &HTTPFeatureStore{
		httpClient:          util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
		logger:              logger,
		initialSyncComplete: false,
		features:            nil,
		readyChannel: make(chan bool),
	}


	httpStore.shutdown = util.Schedule(httpStore.fetchFlags, time.Duration(pollingInterval)*time.Millisecond)

	return httpStore
}
