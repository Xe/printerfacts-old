package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Xe/ln"
	"github.com/Xe/printerfacts/proto"
)

const defaultURL = "https://printerfacts.herokuapp.com"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := proto.NewPrinterfactsProtobufClient(defaultURL, http.DefaultClient)
	fact, err := cli.Fact(ctx, &proto.FactParams{})
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("fetching printer facts"))
	}

	fmt.Println(fact.Facts[0])
}
