package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Xe/ln"
	"github.com/Xe/printerfacts/proto"
	_ "github.com/Xe/printerfacts/statik"
	"github.com/rakyll/statik/fs"
)

func init() {
	rand.Seed(time.Now().Unix())

	if p := os.Getenv("PORT"); p == "" {
		os.Setenv("PORT", "9001")
	}
}

type server struct {
	facts []string
}

// Fact grabs a random set of printer facts and returns them to the user.
func (s *server) Fact(ctx context.Context, prm *proto.FactParams) (*proto.Facts, error) {
	result := &proto.Facts{}

	if prm.Count == 0 {
		prm.Count = 1
	}

	for range make([]struct{}, prm.Count) {
		result.Facts = append(result.Facts, s.facts[rand.Intn(len(s.facts))])
	}

	return result, nil
}

func main() {
	ctx := context.Background()

	sfs, err := fs.New()
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("statik fs"))
	}

	fin, err := sfs.Open("/facts.json")
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("can't find facts.json"))
	}

	var facts []string

	err = json.NewDecoder(fin).Decode(&facts)
	if err != nil {
		ln.FatalErr(ctx, err, ln.Action("can't read facts as json"))
	}

	s := &server{facts: facts}
	handler := proto.NewPrinterfactsServer(s, nil)
	mux := http.NewServeMux()

	mux.Handle(proto.PrinterfactsPathPrefix, handler)

	ln.Log(ctx, ln.F{"port": os.Getenv("PORT")}, ln.Action("Listening on http"))
	ln.FatalErr(ctx, http.ListenAndServe(":"+os.Getenv("PORT"), mux), ln.Action("http server stopped for some reason"))
}
