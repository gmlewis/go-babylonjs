// docs2go convert the Babylon.js API docs to Go files for
// the WebAssembly binding using syscall/js.
package main

import (
	"flag"
	"log"
	"path/filepath"
)

var (
	apiDir  = flag.String("api", "/home/glenn/src/github.com/BabylonJS/Documentation/public/html/api", "Directory where Babylon.js API docs reside")
	destDir = flag.String("dest", "/home/glenn/go/src/github.com/gmlewis/go-babylonjs/babylon", "Destination directory for generated Go code")
	verbose = flag.String("v", "", "Debug with verbose log messages for this class")

	// blacklist is a list of filenames to skip.
	blacklist = map[string]bool{}
)

func main() {
	flag.Parse()

	c := &classes{m: map[string]*ClassHTML{}}
	check("Walk: %v", filepath.Walk(*apiDir, c.walker()))

	check("writeGo: %v", c.writeGo(*destDir))

	log.Println("Done.")
}

func check(fmtStr string, args ...interface{}) {
	if err := args[len(args)-1]; err != nil {
		log.Fatalf(fmtStr, args...)
	}
}

func logf(fmt string, args ...interface{}) {
	if *verbose != "" {
		log.Printf(fmt, args...)
	}
}
