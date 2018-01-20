package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/Xe/ln"
	"github.com/Xe/printerfacts/proto"
)

const defaultURL = "https://printerfacts.herokuapp.com"

var (
	serverURL = flag.String("server", "https://printerfacts.herokuapp.com", "the api server the printer facts api is running on")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	cli := proto.NewPrinterfactsProtobufClient(*serverURL, http.DefaultClient)
	fact, err := cli.Fact(ctx, &proto.FactParams{})
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("fetching printer facts"))
	}

	fmt.Println(fact.Facts[0])
}
