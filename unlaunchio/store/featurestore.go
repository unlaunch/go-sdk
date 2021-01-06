package store

import "github.com/unlaunch/go-sdk/unlaunchio/dtos"

type FeatureStore interface {
	GetFeature(key string) (*dtos.Feature, error)
	Ready()
}