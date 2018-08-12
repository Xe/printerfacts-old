#!/bin/sh

mkdir bin
export GOBIN="$(pwd)"/bin
go install ./cmd/...
