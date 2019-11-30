#!/bin/bash -ex
PATH=$(go env GOROOT)/misc/wasm:${PATH} GOARCH=wasm GOOS=js go test
