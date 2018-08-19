package main

import (
	"context"
	"encoding/json"
	"expvar"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"

	"gopkg.in/segmentio/analytics-go.v3"
	"github.com/Xe/ln"
	"github.com/Xe/printerfacts/internal/printerfactsserver"
	"github.com/Xe/printerfacts/rpc/printerfacts"
	_ "github.com/Xe/printerfacts/statik"
	"github.com/go-kit/kit/metrics/provider"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/rakyll/statik/fs"
)

func init() {
	rand.Seed(time.Now().Unix())

	if p := os.Getenv("PORT"); p == "" {
		os.Setenv("PORT", "9001")
	}
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

	s := &printerfactsserver.Impl{Facts: facts}
	handler := printerfacts.NewPrinterfactsServer(
		printerfacts.NewPrinterfactsLogging(
			printerfacts.NewPrinterfactsMetrics(printerfacts.NewPrinterfactsAnalytics(
				s,
				analytics.New(os.Getenv("SEGMENT_WRITE_KEY")),
			), provider.NewExpvarProvider()),
		), nil)
	mux := http.NewServeMux()

	mux.Handle(printerfacts.PrinterfactsPathPrefix, handler)
	mux.Handle("/", http.FileServer(sfs))
	mux.Handle("/metrics", expvar.Handler())

	ln.Log(ctx, ln.F{"port": os.Getenv("PORT")}, ln.Action("Listening on http"))
	ln.FatalErr(ctx, http.ListenAndServe(":"+os.Getenv("PORT"), metaInfo(mux)), ln.Action("http server stopped for some reason"))
}

func metaInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		f := ln.F{
			"remote_ip":       host,
			"x_forwarded_for": r.Header.Get("X-Forwarded-For"),
			"path":            r.URL.Path,
		}
		ctx := ln.WithF(r.Context(), f)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
