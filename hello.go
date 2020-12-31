package main

import (
	"fmt"

	"github.com/unlaunch/go-sdk/unlaunchio/client"
)

func main() {
	config := client.DefaultConfig()
	factory, err := client.NewUnlaunchClientFactory("<apikey>", config)

	if err != nil {
		fmt.Printf("Unable to initialize Unlaunch Client because there was an error %s\n", err)
		return
	}

	unlaunchClient := factory.Client()

	if err = unlaunchClient.BlockUntilReady(5); err != nil {
		fmt.Printf("Unlaunch Client isn't ready %s\n", err)
	}

	variation := unlaunchClient.Variation("adadadada-hi", "user123", nil)
	fmt.Printf("The variation for feature is %s", variation)

}
