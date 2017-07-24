package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
)

var endpoints = createEndpoints()
var gPort int
var proxyURL string
var defaultHeader = http.Header{"Content-Type": []string{"application/json"}}
var httpClient = http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

func handler(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/" {

		// endpoint creation request
		if req.Method == "POST" {
			handleRootPost(w, req)
		} else if req.Method == "DELETE" {
			handleRootDelete(w, req)
		} else if req.Method == "GET" {
			handleRootGet(w, req)
		} else if req.Method == "PUT" {
			handleRootPut(w, req)
		}

	} else {
		// serve up endpoint
		handlerEndpoint(w, req)
	}

}

func main() {

	flagPort := flag.Int("p", 5351, "the server port")
	flagURL := flag.String("f", "", "url to forward the request")
	flag.Parse()

	gPort = *flagPort
	proxyURL = *flagURL

	fmt.Printf("Starting server at port %d\n", gPort)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", gPort), nil)

}
