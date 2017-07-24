package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleRootPut(w http.ResponseWriter, req *http.Request) {
	errMsg := "PUT / not implemented yet\n"
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, errMsg)
	fmt.Printf(errMsg)
}

func handleRootGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("dumping endpoints")
	dumped, _ := json.Marshal(endpoints)
	w.Header().Set("Content-Type", "application/json")
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
		defaultHeader,
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
	var endpointApprox = createEndpoint(req.URL.Path, req.URL.Query(), req.Method, 200, []byte{}, defaultHeader)

	if endpoint, found := endpointGet(endpoints, endpointApprox); found {
		doResponse(w, endpoint)
	} else {

		if len(proxyURL) > 0 {
			urlEndPoint := fmt.Sprintf("%s%s", proxyURL, req.URL)
			fmt.Printf("Proxying url, %s\n", urlEndPoint)
			proxyRequest, err := http.NewRequest(endpointApprox.Method, urlEndPoint, req.Body)
			if err != nil {
				errorHandle(w, err)
				return
			}
			proxyRequest.Header = req.Header

			proxyResponse, err := httpClient.Do(proxyRequest)
			if err != nil {
				errorHandle(w, err)
				return
			}
			proxyResponseData, err := ioutil.ReadAll(proxyResponse.Body)
			if err != nil {
				errorHandle(w, err)
				return
			}
			endpointApprox.Data = proxyResponseData
			endpointApprox.Header = proxyResponse.Header
			fmt.Println(endpointApprox.Header)
			endpoints = endpointPut(endpoints, endpointApprox)

			doResponse(w, endpointApprox)
		} else {
			fmt.Printf("endpoint MISS, %s\n", formatEndpoint(endpointApprox))
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 endpoint not found")
		}
	}

}

func errorHandle(w http.ResponseWriter, err error) {
	fmt.Printf("proxy MISS, %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "500 Internal Server Error")
}

func doResponse(w http.ResponseWriter, endpoint endpoint) {
	fmt.Printf("endpoint HIT %s\n", formatEndpoint(endpoint))
	for k, v := range endpoint.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(endpoint.Status)
	w.Write(endpoint.Data)
}
