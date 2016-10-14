package main

import "fmt"
import "net/http"

var endpoints = createEndpoints()
var gPort int

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

	gPort = 5351
	fmt.Printf("Starting server at port %d\n", gPort)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", gPort), nil)

}
