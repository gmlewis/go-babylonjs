// docs2go convert the Babylon.js API docs to Go files for
// the WebAssembly binding using syscall/js.
package main

import (
	"flag"
	"log"
	"path/filepath"
)

var (
	apiDir = flag.String("api", "/home/glenn/src/github.com/BabylonJS/Documentation/public/html/api", "Directory where Babylon.js API docs reside")

	// blacklist is a list of filenames to skip.
	blacklist = map[string]bool{}
)

func main() {
	flag.Parse()

	dir := filepath.Join(*apiDir, "classes")
	c := &classes{}
	check("Walk: %v", filepath.Walk(dir, c.walker()))
}

func check(fmtStr string, args ...interface{}) {
	if err := args[len(args)-1]; err != nil {
		log.Fatalf(fmtStr, args...)
	}
}
