package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sort"
	"time"
)

type HTTPFeatureStore struct {
	httpClient          util.HTTPClient
	logger              logger.Interface
	features            map[string]dtos.Feature
	initialSyncComplete bool
	shutdownCh          chan bool
}

func (h *HTTPFeatureStore) Shutdown() {
	h.logger.Debug("Sending shutdownCh signal to feature store")
	h.shutdownCh <- true
}

func (h *HTTPFeatureStore) fetchFlags() error {
	res, err := h.httpClient.Get("/api/v1/flags")

	if err != nil {
		if httpError, ok := err.(*dtos.HTTPError); ok {
			if httpError.Code == 403 {
				h.logger.Error(
					fmt.Sprintf("The API key you provided was rejected by the server. %s", util.SDKKeyHelpMessage))
			}
		} else {
			h.logger.Error("error fetching flags ", err)
			return err
		}
	}

	if res == nil {
		// No error and empty response means nothing changed
		// most like due to 304; not modified
		return nil
	}

	h.logger.Trace("responseDto ", string(res))

	var responseDto dtos.TopLevelEnvelope
	err = json.Unmarshal(res, &responseDto)

	if err != nil {
		h.logger.Error("error parsing feature flag JSON response ", err)
		return err
	}

	if h.initialSyncComplete == false {
		h.initialSyncComplete = true
	}

	// Todo: Remove this when rules and rollouts are sorted on server
	for _, feature := range responseDto.Data.Features {
		sort.Sort(dtos.ByRulePriority(feature.Rules))

		for _, rule := range feature.Rules {
			sort.Sort(dtos.ByVariationID(rule.Rollout))
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
	}
	return nil, errors.New("flag was not found in local storage")
}

func (h *HTTPFeatureStore) IsReady() bool {
	if h.initialSyncComplete {
		return true
	} else {
		return false
	}
}

func (h *HTTPFeatureStore) Ready(timeout time.Duration) {
	// TODO Find a better way to do this
	if h.initialSyncComplete {
		return
	}

	deadline := time.Now().Add(timeout)

	for {
		if h.initialSyncComplete || time.Now().After(deadline) {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func NewHTTPFeatureStore(
	httpClient util.HTTPClient,
	pollingInterval int,
	logger logger.Interface) FeatureStore {
	httpStore := &HTTPFeatureStore{
		httpClient:          httpClient,
		logger:              logger,
		initialSyncComplete: false,
		features:            nil,
	}

	httpStore.shutdownCh = util.RunImmediatelyAndSchedule(httpStore.fetchFlags, time.Duration(pollingInterval)*time.Millisecond)

	return httpStore
}
