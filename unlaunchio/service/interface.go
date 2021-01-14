package service

import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"time"
)

type FeatureStore interface {
	GetFeature(key string) (*dtos.Feature, error)
	Ready(timeout time.Duration)
	Stop()
}