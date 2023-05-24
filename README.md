# Agoraio

# Installation

```
go get github.com/umtaktpe/agoraio
```

# Agora RestAPI Documentation

https://docs.agora.io/en/interactive-live-streaming/reference/agora-console-rest-api

https://docs.agora.io/en/interactive-live-streaming/reference/channel-management-rest-api

# Usage

```go
package main

import (
	"fmt"

	"github.com/umtaktpe/agoraio"
)

const (
	appID          = ""
	customerKey    = ""
	customerSecret = ""
)

func main() {
	client := agoraio.NewClient(appID, customerKey, customerSecret)

	params := &agoraio.GetProjectUsageParameters{
		ProjectID: "",
		FromDate:  "2022-01-01",
		ToDate:    "2023-01-01",
		Business:  "default",
	}

	resp, err := client.GetProjectUsage(params)
	if err != nil {
		panic(err)
	}

	fmt.Println("Resp:", resp)
}
```
