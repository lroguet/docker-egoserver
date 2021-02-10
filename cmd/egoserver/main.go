package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
)

func about(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "About:")
	fmt.Fprintf(w, "\tSimple HTTP echo server written in Go (%v)\n", runtime.Version())
	fmt.Fprintf(w, "\thttps://github.com/lroguet/docker-egoserver\n")
	fmt.Fprintln(w)

}

func headers(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Request Headers:")

	headers := make([]string, 0, len(req.Header))
	for header := range req.Header {
		headers = append(headers, header)
	}
	sort.Strings(headers)

	for _, item := range headers {
		fmt.Fprintf(w, "\t%v: %v\n", item, strings.Join(req.Header[item], ", "))
	}

	fmt.Fprintln(w)

}

func information(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Request Information:")
	fmt.Fprintf(w, "\tHost: %v\n", req.Host)
	fmt.Fprintf(w, "\tMethod: %v\n", req.Method)
	fmt.Fprintf(w, "\tProtocol: %v\n", req.Proto)
	fmt.Fprintln(w)

}

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		about(w, req)
		information(w, req)
		headers(w, req)
	})

	doNothing := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {})

	http.Handle("/", handler)
	// Capturing web-browsers default HTTP request to favicon.ico for cleaner demo logs and traces
	http.Handle("/favicon.ico", doNothing)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
