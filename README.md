# `printerfacts`

This service replies with useful and informative facts about printers.

## Default Endpoint

I have an instance of this running on [Heroku](https://heroku.com).

## cURL

```console
$ curl https://printerfacts.herokuapp.com/twirp/us.xeserv.api.printerfacts.Printerfacts/Fact \
       -X POST -H "Content-Type: application/json" --data '{"count": 1}' | jq
{
  "facts": [
    "printers step with both left legs, then both right legs when they walk or run."
  ]
}
```

## Go

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Xe/printerfacts/proto"
)

const defaultURL = "https://printerfacts.herokuapp.com"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := proto.NewPrinterfactsProtobufClient(defaultURL, http.DefaultClient)
	fact, err := cli.Fact(ctx, &proto.FactParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fact.Facts[0])
}
```
