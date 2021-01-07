package main

import (
	"fmt"
	"time"

	"github.com/unlaunch/go-sdk/unlaunchio/client"
)

func main() {
	config := client.DefaultConfig()
	factory, err := client.NewUnlaunchClientFactory("prod-sdk-e40d9c6a-8bfb-414f-8737-353c5bec2db8", config)

	if err != nil {
		fmt.Printf("Unable to initialize Unlaunch Client because there was an error %s\n", err)
		return
	}

	unlaunchClient := factory.Client()

	if err = unlaunchClient.BlockUntilReady(5); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	attributes := make(map[string]interface{})
	attributes["showBalance"] = true

	variation := unlaunchClient.Variation("adadadada-hi", "user123", &attributes)
	fmt.Printf("The variation for feature is %s\n", variation)

	go func() {
		variation := unlaunchClient.Variation("adadadada-hi", "user123631", nil)
		fmt.Printf("- The variation for feature is %s\n", variation)

	}()

	time.Sleep(1 * time.Second)

}
