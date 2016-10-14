package main

import "fmt"
import "net/http"
import "encoding/json"

func handleRootPut(w http.ResponseWriter, req *http.Request) {
	errMsg := "PUT / not implemented yet\n"
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, errMsg)
	fmt.Printf(errMsg)
}

func handleRootGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("dumping endpoints")
	dumped, _ := json.Marshal(endpoints)
	w.Write([]byte(dumped))
}

func handleRootPost(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	var endpoint = createEndpoint(
		extractPathFromRequest(req),
		extractQueryFromRequest(req),
		extractMethodFromRequest(req),
		extractStatusFromRequest(req),
		extractDataFromRequest(req),
	)

	if endpoint.Path == "" {
		fmt.Printf("endpoint NOT ADDED ( path missing ) %s", formatEndpoint(endpoint))
		fmt.Printf(" ( to create a endpoint a path is required )\n")
		return
	}

	endpoints = endpointPut(endpoints, endpoint)

	fmt.Printf("endpoint ADD %s\n", formatEndpoint(endpoint))

}

func handleRootDelete(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "clearing endpoints")
	endpoints = nil // correct way to clear array for GC
	// nilling array leaves null, which looks unsightly on dump in browser
	endpoints = createEndpoints() // create empty endpoints

}

func handlerEndpoint(w http.ResponseWriter, req *http.Request) {

	// create a 'close enough' endpoint to search with
	var endpointApprox = createEndpoint(req.URL.Path, req.URL.Query(), req.Method, 200, "")

	if endpoint, found := endpointGet(endpoints, endpointApprox); found {
		fmt.Printf("endpoint HIT %s\n", formatEndpoint(endpoint))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(endpoint.Status)
		w.Write([]byte(endpoint.Data))
	} else {
		fmt.Printf("endpoint MISS, %s\n", formatEndpoint(endpointApprox))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 endpoint not found")
	}

}
