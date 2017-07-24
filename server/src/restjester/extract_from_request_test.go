package main

import "testing"
import "net/http"
import "bytes"
import "reflect"

// import "fmt"

var data = bytes.NewBufferString("not important")
var request, _ = http.NewRequest("GET", "http://localhost:0000", data)

func TestExtractPath(t *testing.T) {

	request.Form = map[string][]string{
		"path": []string{"/1/2/3?a=1&b=2&c=3"},
	}

	path := extractPathFromRequest(request)

	expected := "/1/2/3"
	if path != expected {
		t.Errorf("expected '%s', got '%s' from request %q", expected, path, request)
	}

}

func TestExtractQuery(t *testing.T) {

	request.Form = map[string][]string{
		"path": []string{"/1/2/3?a=1&b=2&c=3"},
	}

	query := extractQueryFromRequest(request)

	expected := map[string][]string{
		"a": []string{"1"},
		"b": []string{"2"},
		"c": []string{"3"},
	}

	if !reflect.DeepEqual(expected, query) {
		t.Errorf("expected %q, got %q from request %q", expected, query, request)
	}

}

func TestExtractMethod(t *testing.T) {

	request.Form = map[string][]string{
		"method": []string{"GET"},
	}

	method := extractMethodFromRequest(request)

	expected := "GET"

	if method != expected {
		t.Errorf("expected '%s', got '%s' from request %q", expected, method, request)
	}

}

func TestExtractData(t *testing.T) {

	request.Form = map[string][]string{
		"data": []string{"the payload"},
	}

	data := extractDataFromRequest(request)

	expected := "the payload"

	if string(data) != expected {
		t.Errorf("expected '%s', got '%s' from request %q", expected, data, request)
	}

}

func TestExtractStatus(t *testing.T) {

	request.Form = map[string][]string{
		"status": []string{"200"},
	}

	method := extractStatusFromRequest(request)

	expected := 200

	if method != expected {
		t.Errorf("expected '%s', got '%s' from request %q", expected, method, request)
	}

}
