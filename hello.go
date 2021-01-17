package main

import (
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"time"
	"github.com/unlaunch/go-sdk/unlaunchio/client"
)

func main() {
	config := client.DefaultConfig()
	config.LoggerConfig = &logger.LogOptions{
		Level: "TRACE",
		Colorful: true,
	}
	config.PollingInterval = 2000
	config.MetricsFlushInterval = 15000
	factory, err := client.NewUnlaunchClientFactory("prod-server-51028624-eb18-4bc7-986f-5a0de8084589", config)

	if err != nil {
		fmt.Printf("Unable to initialize Unlaunch Client because there was an error %s\n", err)
		return
	}

	unlaunchClient := factory.Client()

	//time.Sleep(2 * time.Second)

	if err = unlaunchClient.BlockUntilReady(4 * time.Second); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	flagKey := "set-attr-type-3"
	attributes := make(map[string]interface{})
	attributes["boolAttr"] = true

	variation := unlaunchClient.Variation(flagKey, "user123", attributes)
	fmt.Printf("The variation for feature is %s\n", variation)


	go func() {
		feature := unlaunchClient.Feature(flagKey, "user123631", nil)
		fmt.Println(fmt.Sprintf("%v", feature.VariationConfiguration))
		fmt.Printf("- The variation for feature is %s\n", feature.Variation)
	}()

	time.Sleep(1 * time.Second)

	unlaunchClient.Shutdown()

	fmt.Println(unlaunchClient.IsShutdown())

	time.Sleep(5 * time.Second)
	fmt.Println("bye")

}
