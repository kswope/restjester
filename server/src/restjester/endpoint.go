package main

// import "fmt"
import "reflect"

type endpoint struct {
	Path   string              // /a/b/c
	Query  map[string][]string // ?a=b&c=d
	Method string              // POST, GET
	Status int                 //
	Data   string              // endpoint data
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func createEndpoints() []endpoint {
	return make([]endpoint, 0)
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func createEndpoint(
	path string,
	query map[string][]string,
	method string,
	status int,
	data string,
) endpoint {
	return endpoint{
		Path:   path,
		Query:  query,
		Method: method,
		Status: status,
		Data:   data,
	}
}

//-----------------------------------------------------------------------------
// Idempotent: add or replace endpoint
//-----------------------------------------------------------------------------
func endpointPut(endpoints []endpoint, ep endpoint) []endpoint {

	endpoints = endpointDelete(endpoints, ep)
	endpoints = append(endpoints, ep)
	return endpoints

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func endpointDelete(endpoints []endpoint, ep endpoint) []endpoint {

	index := endpointIndex(endpoints, ep)

	if index >= 0 {
		// Delete without potential memory leak.
		// https://github.com/golang/go/wiki/SliceTricks
		endpoints[index] = endpoints[len(endpoints)-1]
		endpoints[len(endpoints)-1] = endpoint{} // nil endpoint because of GC
		endpoints = endpoints[:len(endpoints)-1]
	}

	return endpoints

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func endpointGet(endpoints []endpoint, ep endpoint) (endpoint, bool) {

	index := endpointIndex(endpoints, ep)
	if index >= 0 {
		return endpoints[index], true
	} else {
		return endpoint{}, false
	}

}

//-----------------------------------------------------------------------------
// Find first (should be only) endpoint index in endpoints (if its there). We
// match ONLY on endpoint.path, endpoint.query and endpoint.method, Return
// index, or -1
//-----------------------------------------------------------------------------
func endpointIndex(endpoints []endpoint, ep endpoint) int {

	for i, x := range endpoints {

		// three equality checks
		path := x.Path == ep.Path
		query := reflect.DeepEqual(x.Query, ep.Query)
		method := x.Method == ep.Method

		if path && query && method {
			return i
		}

	}

	// nothing matched
	return -1

}
