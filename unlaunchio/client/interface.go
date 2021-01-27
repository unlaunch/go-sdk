package client

 import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"time"
 )

// Client is the main interface that all Unlaunch clients must implement
 type Client interface {
	Feature(featureKey string, identity string, attributes map[string]interface{},) *dtos.UnlaunchFeature
	IsShutdown() bool
	Variation(featureKey string, identity string, attributes map[string]interface{},) string
	AwaitUntilReady(timeout time.Duration) error
	Shutdown()
}
