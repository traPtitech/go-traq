# go-traq

## example
```go
package main

import (
  "context"
  api "go-traq"
  "fmt"
)

const TOKEN = "/* your token */";

func main() {
  client := api.NewAPIClient(api.NewConfiguration())
  auth := context.WithValue(context.Background(), api.ContextAccessToken, TOKEN)

  v, _, _ := client.ChannelApi.ChannelsGet(auth)
  fmt.Printf("%#v", v)
}
```

### build
```sh
sh generate.sh
```
その後手動で修正
