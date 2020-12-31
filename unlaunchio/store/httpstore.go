package store

import (
	"encoding/json"
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/http"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

type HttpStore struct {
	service http.ServiceClient
	logger  logger.Interface
	features map[string]dtos.Feature
}

func (h *HttpStore) fetchFlags() ([]byte, error) {
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

type FeatureStore interface {
	GetFeature(key string) (*dtos.Feature, error)
}

func (h *HttpStore) GetFeature(key string) (*dtos.Feature, error) {

	if feature, ok := h.features[key]; ok {
		return &feature, nil
	} else {
		return nil, errors.New("flag was not found in local storage")
	}
}

func NewHTTPStore(
	sdkKey string,
	host string,
	httpTimeout int,
	pollingInterval int,
	logger logger.Interface,
) FeatureStore {

	httpStore := &HttpStore{
		service: http.NewHTTPClient(sdkKey, host, httpTimeout, logger),
		logger:  logger,
		features: make(map[string]dtos.Feature),
	}

	httpStore.fetchFlags()

	return httpStore
}
