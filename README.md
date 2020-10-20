# go-traq

## example
```go
package main

import (
    "context"
    "fmt"
    "github.com/antihax/optional"
    traq "github.com/sapphi-red/go-traq"
)

const TOKEN = "/* your token */";

func main() {
  client := traq.NewAPIClient(traq.NewConfiguration())
  auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

  v, _, _ := client.ChannelApi.GetChannels(auth, &traq.ChannelApiGetChannelsOpts{
      IncludeDm: optional.NewBool(true),
  })
  fmt.Printf("%#v", v)
}
```

### build
```sh
sh generate.sh
```
その後手動で修正
