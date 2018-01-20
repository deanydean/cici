package main

import (
	"flag"

	"github.com/deanydean/cici/services/http"
	"github.com/deanydean/cici/services/watch"
)

func main() {

	// Get CLI modes
	httpEnabled := flag.Bool("http", false, "If http should be enabled")
	watchingEnabled := flag.Bool("watch", false, "If watching is enabled")

	// Get CLI settings
	listenStr := flag.String("listen", ":5151", "The address to listen on")
	flag.Parse()

	if *watchingEnabled {
		watch.Start()
	}

	if *httpEnabled {
		http.Start(listenStr)
	}
}
