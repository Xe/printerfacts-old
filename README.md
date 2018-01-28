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

	"github.com/Xe/printerfacts/rpc/printerfacts"
)

const defaultURL = "https://printerfacts.herokuapp.com"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := printerfacts.NewPrinterfactsProtobufClient(defaultURL, http.DefaultClient)
	fact, err := cli.Fact(ctx, &printerfacts.FactParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fact.Facts[0])
}
```

## Build

```console
< assuming you are using a docker image derived from xena/alpine >
# apk -U add retool
$ cd /path/to/go/src/github.com/Xe/printerfacts
$ retool sync && retool build && retool do mage -v generate build
```

## Run locally

```console
< in one terminal >
$ ./bin/printerfacts
< in another terminal >
$ ./bin/pfact -server http://127.0.0.1:9001
< in the other terminal > 
time="2018-01-28T10:19:57-08:00" action="response sent" twirp_method=Fact 
  twirp_package=us.xeserv.api twirp_service=Printerfacts remote_ip=127.0.0.1 
  x_forwarded_for= path=/twirp/us.xeserv.api.Printerfacts/Fact 
  response_time=191.671µs
```
