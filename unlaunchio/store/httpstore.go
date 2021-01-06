package store

import (
	"encoding/json"
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"sync"
)

type HttpFeatureStore struct {
	service  util.HttpClient
	logger   logger.Interface
	features map[string]dtos.Feature
}

func (h *HttpFeatureStore) fetchFlags() ([]byte, error) {
	defer wg.Done()
	res, err := h.service.Get("/api/v1/flags")

	if err != nil {
		h.logger.Error("error fetching flags ", err)
	}

	h.logger.Info("responseDto ", string(res))

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
	wg.Wait()
}


var wg sync.WaitGroup

func NewHTTPStore(
	sdkKey string,
	host string,
	httpTimeout int,
	pollingInterval int,
	logger logger.Interface,
) FeatureStore {

	httpStore := &HttpFeatureStore{
		service:  util.NewHTTPClient(sdkKey, host, httpTimeout, logger),
		logger:   logger,
		features: make(map[string]dtos.Feature),
	}

	wg.Add(1)
	go httpStore.fetchFlags()

	return httpStore
}
