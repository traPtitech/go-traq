# go-traq

## example
```go
package main

import (
    "context"
    "fmt"
    "github.com/antihax/optional"
    api "github.com/sapphi-red/go-traq"
)

const TOKEN = "/* your token */";

func main() {
  client := api.NewAPIClient(api.NewConfiguration())
  auth := context.WithValue(context.Background(), api.ContextAccessToken, TOKEN)

  v, _, _ := client.ChannelApi.GetChannels(auth, &api.ChannelApiGetChannelsOpts{
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
