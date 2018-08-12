#!/bin/sh

export PATH=$PATH:$HOME/go/bin

export GO111MODULE=off
go get -u github.com/Xe/twirp-codegens/cmd/protoc-gen-twirp_jsbrowser
go get -u github.com/Xe/twirp-codegens/cmd/protoc-gen-twirp_ln
go get -u github.com/twitchtv/twirp/protoc-gen-twirp
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/rakyll/statik

(cd ./rpc/printerfacts && sh ./regen.sh)

cp ./rpc/printerfacts/printerfacts_twirp.js ./facts/printerfacts_twirp.js
statik -src ./facts -f

