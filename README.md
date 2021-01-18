# Unlaunch Go SDK

| main                                                                                                                | development                                                                                                                | Go Report Card |
|---------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|----------------|
| [![Build Status](https://travis-ci.com/unlaunch/go-sdk.svg?branch=main)](https://travis-ci.com/unlaunch/go-sdk)| [![Build Status](https://travis-ci.com/unlaunch/go-sdk.svg?branch=development)](https://travis-ci.com/unlaunch/go-sdk) | [![Go Report Card](https://goreportcard.com/badge/github.com/unlaunch/go-sdk)](https://goreportcard.com/report/github.com/unlaunch/go-sdk) |

## Overview
The Unlaunch Go SDK provides a Go API to access Unlaunch feature flags and other features. 
Using the SDK, you can easily build Java applications that can evaluate feature flags, dynamic configurations, and more.

### Important Links

- To create feature flags to use with Java SDK, login to your Unlaunch Console at [https://app.unlaunch.io](https://app.unlaunch.io)
- [Official Guide](https://docs.unlaunch.io/docs/sdks/go-sdk)
- [Documentation](https://pkg.go.dev/github.com/unlaunch/go-sdk#section-documentation)
- [Example Project](https://github.com/unlaunch/hello-go)

## Getting Started
Here is a simple example.

Run `go get github.com/unlaunch/go-sdk` add it to your `go.mod` file.

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
		fmt.Printf("Error initializing Unlaunch client %s\n", err)
		return
	}

	unlaunchClient := factory.Client()
	if err = unlaunchClient.AwaitUntilReady(3 * time.Second); err != nil {
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
## Build Instructions

To run all tests

```shell
go test ./...
```

## How to Publish
Create a new tag on GitHub in the following format vx.y.z e.g. v0.0.1

## Submitting issues
If you run into any problems, bugs, or have any questions or feedback, please report them using the [issues feature](https://github.com/unlaunch/go-sdk/issues). We'll respond to all issues usually within 24 to 48 hours.

## Contributing
Please see [CONTRIBUTING](CONTRIBUTING.md) to find how you can contribute.

## License
Licensed under the Apache License, Version 2.0. See: [Apache License](LICENSE.md).

## About Unlaunch
Unlaunch is a Feature Release Platform for engineering teams. Our mission is allow engineering teams of all
sizes to release features safely and quickly to delight their customers. To learn more about Unlaunch, please visit
[unlaunch.io](https://unlaunch.io). You can sign up to get started for free at [https://app.unlaunch.io/signup](https://app.unlaunch.io/signup).