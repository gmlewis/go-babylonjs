#!/bin/bash -ex
cd examples/$@*
./build.sh
./serve.sh
