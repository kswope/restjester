package main

import "net/http"
import "net/url"
import "strconv"

func extractPathFromRequest(req *http.Request) string {
	if len(req.Form["path"]) > 0 {
		url, _ := url.Parse(req.Form["path"][0])
		return url.Path
	} else {
		return ""
	}
}

func extractQueryFromRequest(req *http.Request) map[string][]string {
	if len(req.Form["path"]) > 0 {
		url, _ := url.Parse(req.Form["path"][0])
		return url.Query()
	} else {
		return map[string][]string{}
	}
}

func extractMethodFromRequest(req *http.Request) string {
	if len(req.Form["method"]) > 0 {
		return req.Form["method"][0]
	} else {
		return "GET" // default
	}
}

func extractDataFromRequest(req *http.Request) []byte {
	if len(req.Form["data"]) > 0 {
		return []byte(req.Form["data"][0])
	} else {
		return []byte("") // default
	}
}

func extractStatusFromRequest(req *http.Request) int {
	if len(req.Form["status"]) > 0 {
		// convert str to number
		if i, err := strconv.Atoi(req.Form["status"][0]); err == nil {
			return i
		} else {
			return 200 // convert failed
		}
	} else {
		return 200 // default
	}
}
