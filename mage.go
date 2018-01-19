// +build mage

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
)

func Tools() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tools := []string{
		"github.com/golang/protobuf/protoc-gen-go",
		"github.com/twitchtv/twirp/protoc-gen-twirp",
		"github.com/rakyll/statik",
	}

	for _, t := range tools {
		fmt.Printf("installing: %s\n", t)
		shouldWork(ctx, nil, wd, "go", "get", "-u", t)
	}
}

func Generate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mg.Deps(Tools)

	if Contained() {
		p := os.Getenv("PATH") + ":/root/go/bin"
		log.Printf("setting path to %s", p)
		os.Setenv("PATH", p)
	}

	shouldWork(ctx, nil, wd, "statik", "-src", "./facts", "-f")
	shouldWork(ctx, nil, filepath.Join(wd, "proto"), "sh", "./regen.sh")
}

func Build() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mg.Deps(Generate)
	os.Mkdir("bin", 0777)

	outd := filepath.Join(wd, "bin")
	cmds := []string{"printerfacts", "pfact"}

	for _, c := range cmds {
		shouldWork(ctx, nil, outd, "go", "build", "../cmd/"+c)
	}
}
