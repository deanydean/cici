package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Received HTTP request for ", request.URL.String())
	writer.Write([]byte("Hello, I'm Cici\n"))
}

// Start the HTTP services listening on the provide listenStr
func Start(listenStr *string) {
	server := &http.Server{
		Addr:           *listenStr,
		Handler:        http.HandlerFunc(handleRequest),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Listening for HTTP on", *listenStr)

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}
