#!/bin/bash -ex

cat $(go env GOROOT)/misc/wasm/wasm_exec.js |
    sed -e 's/global.fs = require.*$/try { & } catch (err) {}/' \
    > js/wasm_exec.js

# tinygo works great on this example!
# https://tinygo.org/
tinygo build -o main.wasm -target wasm .

# but if you don't have tinygo, use Go:
# GOARCH=wasm GOOS=js go build -o main.wasm
