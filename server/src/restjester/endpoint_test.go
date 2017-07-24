package main

import "testing"
import "reflect"

// import "fmt"

//-----------------------------------------------------------------------------
// grab bag of endpoints for our testing
//-----------------------------------------------------------------------------
func getSampleEndpoint(index int) endpoint {

	var samples []endpoint
	var ep endpoint
	var query map[string][]string

	query = map[string][]string{"a": []string{"1"}}
	ep = createEndpoint("/a/b/c", query, "POST", 200, []byte("data a"), defaultHeader)
	samples = append(samples, ep)

	query = map[string][]string{"b": []string{"2"}}
	ep = createEndpoint("/d/e/f", query, "GET", 200, []byte("data b"), defaultHeader)
	samples = append(samples, ep)

	query = map[string][]string{"c": []string{"3"}}
	ep = createEndpoint("/g/h/i", query, "POST", 200, []byte("data c"), defaultHeader)
	samples = append(samples, ep)

	query = map[string][]string{"d": []string{"4"}}
	ep = createEndpoint("/j/k/l", query, "POST", 200, []byte("data d"), defaultHeader)
	samples = append(samples, ep)

	return samples[index]

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestCreateEndpoint(t *testing.T) {

	query := map[string][]string{"a": []string{"1"}}
	endpoint := createEndpoint("/a/b/c", query, "POST", 200, []byte("data a"), defaultHeader)

	// equality checks
	pathCheck := endpoint.Path == "/a/b/c"
	queryCheck := endpoint.Query["a"][0] == "1"
	methodCheck := endpoint.Method == "POST"
	statusCheck := endpoint.Status == 200
	dataCheck := string(endpoint.Data) == "data a"

	if !(pathCheck && queryCheck && methodCheck && statusCheck && dataCheck) {
		t.Errorf("Failed with path:%q query:%s method:%s data:%s",
			endpoint.Path, endpoint.Query, endpoint.Method, endpoint.Data)
	}

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestEndpointIndex(t *testing.T) {

	endpoints := createEndpoints()
	endpoints = append(endpoints, getSampleEndpoint(0))
	endpoints = append(endpoints, getSampleEndpoint(1))
	endpoints = append(endpoints, getSampleEndpoint(2))

	// find second endpoint we just appended
	index := endpointIndex(endpoints, getSampleEndpoint(1))
	if index != 1 {
		t.Errorf("Expected 1, got %d", index)
	}

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestEndpointIndexNotFound(t *testing.T) {

	endpoints := createEndpoints()
	endpoints = append(endpoints, getSampleEndpoint(0))
	endpoints = append(endpoints, getSampleEndpoint(1))
	endpoints = append(endpoints, getSampleEndpoint(2))

	// find endpoint not added
	index := endpointIndex(endpoints, getSampleEndpoint(3))
	if index != -1 {
		t.Errorf("Expected -1, got %d", index)
	}

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestEndpointGet(t *testing.T) {

	endpoints := createEndpoints()
	endpoints = append(endpoints, getSampleEndpoint(0))
	endpoints = append(endpoints, getSampleEndpoint(1))
	endpoints = append(endpoints, getSampleEndpoint(2))

	// pick one and get it
	endpoint, _ := endpointGet(endpoints, getSampleEndpoint(1))

	if !reflect.DeepEqual(getSampleEndpoint(1), endpoint) {
		t.Errorf("Failed %q != %q", getSampleEndpoint(1), endpoint)
	}

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestEndpointDelete(t *testing.T) {

	endpoints := createEndpoints()
	endpoints = append(endpoints, getSampleEndpoint(0))
	endpoints = append(endpoints, getSampleEndpoint(1))

	if len(endpoints) != 2 {
		t.Errorf("Unexpected endpoints length of %d expected 2", len(endpoints))
	}

	// shouldn't delete anything because endpoint is not in endpoints
	endpoints = endpointDelete(endpoints, getSampleEndpoint(2))

	if len(endpoints) != 2 {
		t.Errorf("Unexpected endpoints length of %d expected 2", len(endpoints))
	}

	endpoints = endpointDelete(endpoints, getSampleEndpoint(1))
	if len(endpoints) != 1 {
		t.Errorf("Unexpected endpoints length of %d expected 1", len(endpoints))
	}

	// delete last endpoints
	endpoints = endpointDelete(endpoints, getSampleEndpoint(0))
	if len(endpoints) != 0 {
		t.Errorf("Unexpected endpoints length of %d expected 0", len(endpoints))
	}

}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func TestEndpointPutIsIdempotent(t *testing.T) {

	endpoints := createEndpoints()

	// put two different endpoints
	endpoints = endpointPut(endpoints, getSampleEndpoint(0))
	endpoints = endpointPut(endpoints, getSampleEndpoint(1))

	if len(endpoints) != 2 {
		t.Errorf("Unexpected endpoints length of %d, expected 2", len(endpoints))
	}

	// add one identical endpoint
	endpoints = endpointPut(endpoints, getSampleEndpoint(1))

	if len(endpoints) != 2 {
		t.Errorf("Unexpected endpoints length of %d expected 2", len(endpoints))
	}

	endpointMod := getSampleEndpoint(0)
	if endpointMod.Method == "GET" {
		endpointMod.Method = "POST"
	} else {
		endpointMod.Method = "GET"
	}

	// add modified endpoint, changed only method, which is enough
	endpoints = endpointPut(endpoints, endpointMod)

	if len(endpoints) != 3 {
		t.Errorf("Unexpected endpoints length of %d expected 3", len(endpoints))
	}

}
