// +build mage

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
)

// Generate runs all relevant code generation tasks.
func Generate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "statik", "-src", "./facts", "-f")
	shouldWork(ctx, nil, filepath.Join(wd, "rpc", "printerfacts"), "sh", "./regen.sh")
	fmt.Println("reran code generation")

	fin, err := os.Open("./rpc/printerfacts/printerfacts_twirp.js")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create("./facts/printerfacts_twirp.js")
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	io.Copy(fout, fin)
}

// Build creates a binary of printerfacts and pfact in a new directory named bin
func Build() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mg.Deps(Generate)
	os.Mkdir("bin", 0777)

	outd := filepath.Join(wd, "bin")
	cmds := []string{"printerfacts", "pfact"}

	for _, c := range cmds {
		shouldWork(ctx, nil, outd, "go", "build", "../cmd/"+c)
		fmt.Println("built ./bin/" + c)
	}
}

// Docker creates the docker image xena/printerfacts with the printerfacts server.
func Docker() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "docker", "pull", "xena/go-mini:1.9")
	shouldWork(ctx, nil, wd, "docker", "build", "-t", "xena/printerfacts", ".")
}

// Heroku deploys this to the heroku app printerfacts.
func Heroku() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldWork(ctx, nil, wd, "docker", "pull", "xena/go-mini:1.9")
	shouldWork(ctx, nil, wd, "heroku", "container:push", "-a", "printerfacts", "web")
}
