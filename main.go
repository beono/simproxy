package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

const port = 8080

func main() {
	logger := log.New(os.Stdout, "DEBUG: ", log.Lshortfile)

	fallbackURL, err := url.Parse("https://www.googleapis.com/")
	if err != nil {
		logger.Fatal(err)
	}

	revProxy := httputil.NewSingleHostReverseProxy(fallbackURL)

	serverMux := http.NewServeMux()

	// This is our 404 handler. It sends all requests to `fallbackURL`
	serverMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("request to: %s", "404")
		writer.Header().Add("x-proxy", "true")
		request.Host = fallbackURL.Host
		revProxy.ServeHTTP(writer, request)
	})

	// This handler just replies `hello`.
	serverMux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("request to: %s", "hello")
		writer.Header().Add("x-proxy", "false")
		writer.Write([]byte("hello"))
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
		Handler:      serverMux,
	}

	fmt.Printf("server has stopped: %+v", server.ListenAndServe())
}
