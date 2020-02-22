# ARCHIVED on 2020-02-22

This was a fun experiment, but I have stopped using Babylon.JS in
favor of the awesome [entity-component-system](https://wikipedia.org/wiki/Entity_component_system)
[aframe.io](https://aframe.io/). Therefore, I've decided to archive
this repo and move on.

# go-babylonjs

[![GoDoc](https://godoc.org/github.com/gmlewis/go-babylonjs/babylon?status.svg)](https://godoc.org/github.com/gmlewis/go-babylonjs/babylon)
[![Test Status](https://github.com/gmlewis/go-babylonjs/workflows/tests/badge.svg)](https://github.com/gmlewis/go-babylonjs/actions?query=workflow%3Atests)

Experimental Go binding to Babylon.js using WebAssembly.

## Dependencies

- [Go](https://golang.org) version 1.13.4 or later
- [TinyGo](https://tinygo.org) version 0.10.0 or later
- [go-wasm-cli](https://github.com/mfrachet/go-wasm-cli)

## Running an example

You can use the `run.sh` bash script and the example number found in
the [examples directory](/examples) to locally build and serve an
example in your browser (using `go-wasm start`):

```sh
$ ./run.sh 11
(ctrl-c to quit)
```

----------------------------------------------------------------------

Enjoy!

----------------------------------------------------------------------

# License

Copyright 2019 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
