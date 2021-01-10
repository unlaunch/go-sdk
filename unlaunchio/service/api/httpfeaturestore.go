package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sort"
	"sync"
	"time"
)

type HttpFeatureStore struct {
	httpClient          *util.HTTPClient
	logger              logger.Interface
	features            map[string]dtos.Feature
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

	res, err := h.httpClient.Get("/api/v1/flags")

	if err != nil {
		h.logger.Error("error fetching flags ", err)
	}

	h.logger.Debug("responseDto ", string(res))

	var responseDto dtos.TopLevelEnvelope
	err = json.Unmarshal(res, &responseDto)

	if err != nil {
		h.logger.Error("Error parsing split changes JSON ", err)
		return nil, err
	}

	for _, v := range responseDto.Data.Features[0].Variations {
		fmt.Println(v)
	}

	// Todo: Remove this when rules and rollouts are sorted on server
	for _, feature := range responseDto.Data.Features {
		sort.Sort(dtos.ByRulePriority(feature.Rules))

		for _, rule := range feature.Rules {
			sort.Sort(dtos.ByVariationId(rule.Rollout))
		}
	}

	fmt.Println("-")
	for _, v := range responseDto.Data.Features[0].Variations {
		fmt.Println(v)
	}

	h.logger.Debug("responseDto ", responseDto)

	// Store features in the service/map
	temp := make(map[string]dtos.Feature)
	for _, feature := range responseDto.Data.Features {
		temp[feature.Key] = feature

	}
	h.features = temp

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
) service.FeatureStore {
	httpStore := &HttpFeatureStore{
		httpClient:          util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
		logger:              logger,
		initialSyncComplete: false,
		features:            nil,
	}

	wg.Add(1)
	stop = util.Schedule(httpStore.fetchFlags, time.Duration(pollingInterval)*time.Millisecond)

	return httpStore
}
