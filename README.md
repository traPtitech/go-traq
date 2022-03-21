# go-traq

## example
```go
package main

import (
	"context"
	"fmt"

	traq "github.com/sapphi-red/go-traq"
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

### build
```sh
sh generate.sh
```
