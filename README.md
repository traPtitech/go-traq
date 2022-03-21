# go-traq
[![Go Reference](https://pkg.go.dev/badge/github.com/traPtitech/go-traq.svg)](https://pkg.go.dev/github.com/traPtitech/go-traq)
[![CI](https://github.com/traPtitech/go-traq/actions/workflows/main.yaml/badge.svg)](https://github.com/traPtitech/go-traq/actions/workflows/main.yaml)

A client library for the traQ API.

This package is updated automatically.

## How to use
```shell
$ go get github.com/traPtitech/go-traq
```

```go
package main

import (
	"context"
	"fmt"

	traq "github.com/traPtitech/go-traq"
)

const TOKEN = "/* your token */"

func main() {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

	v, _, _ := client.ChannelApi.
		GetChannels(auth).
		IncludeDm(true).
		Execute()
	fmt.Printf("%#v", v)
}
```
