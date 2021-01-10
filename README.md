# Unlaunch Go SDK

## Overview
The Unlaunch Go SDK provides a Go API to access Unlaunch feature flags and other features. 
Using the SDK, you can easily build Java applications that can evaluate feature flags, dynamic configurations, and more.

### Important Links

- To create feature flags to use with Java SDK, login to your Unlaunch Console at [https://app.unlaunch.io](https://app.unlaunch.io)
- [Official Guide](https://docs.unlaunch.io/docs/sdks/go-sdk)

## Getting Started
Here is a simple example.

Run `go get github.com/unlaunch/go-sdk/` or `dep ensure -add go get github.com/unlaunch/go-sdk`

```go
package main

import (
	"fmt"
	"time"
	"github.com/unlaunch/go-sdk/unlaunchio/client"
)

func main() {
	config := client.DefaultConfig()
	factory, err := client.NewUnlaunchClientFactory("YOUR_SERVER_KEY", config)

	if err != nil {
		fmt.Printf("Unable to initialize Unlaunch Client because there was an error %s\n", err)
		return
	}

	unlaunchClient := factory.Client()
	if _ = unlaunchClient.BlockUntilReady(3); err != nil {
		fmt.Printf("Unlaunch Client wasn't ready %s\n", err)
	}

	variation := unlaunchClient.Variation("FLAG_KEY", "USER_ID_123", nil)

	if variation == "on" {
		// show feature
	} else {
		// don't show feature
	}
}
```
## How to use

To run all tests

```shell
go test ./...
```

## Contributing
Please see [CONTRIBUTING](CONTRIBUTING.md) to find how you can contribute.

## License
Licensed under the Apache License, Version 2.0. See: [Apache License](LICENSE.md).

## About Unlaunch
Unlaunch is a Feature Release Platform for engineering teams. Our mission is allow engineering teams of all
sizes to release features safely and quickly to delight their customers. To learn more about Unlaunch, please visit
[unlaunch.io](https://unlaunch.io). You can sign up to get started for free at [https://app.unlaunch.io/signup](https://app.unlaunch.io/signup).