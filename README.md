# go-traq

[![Go Reference][godoc-badge]][godoc] [![CI][ci-badge]][ci] [![Update][update-badge]][update]

[godoc]: https://pkg.go.dev/github.com/traPtitech/go-traq
[godoc-badge]: https://pkg.go.dev/badge/github.com/traPtitech/go-traq.svg
[ci]: https://github.com/traPtitech/go-traq/actions/workflows/main.yaml
[ci-badge]: https://github.com/traPtitech/go-traq/actions/workflows/main.yaml/badge.svg
[update]: https://github.com/traPtitech/go-traq/actions/workflows/release.yaml
[update-badge]: https://github.com/traPtitech/go-traq/actions/workflows/release.yaml/badge.svg

A client library for the traQ API.

This package is updated automatically.

## How to use

See [client.md](./client.md) and [docs/](./docs/) for more information.

```shell
go get github.com/traPtitech/go-traq
```

<!-- markdownlint-disable MD010 -->
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
