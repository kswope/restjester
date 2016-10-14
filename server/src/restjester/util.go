package main

import "fmt"
import "bytes"

// import "encoding/json"
import "net/http"
import "net/url"

// func prettyPrintJSON(str string) {
// 	byt := []byte(str)
// 	var out bytes.Buffer
// 	err := json.Indent(&out, byt, "", "  ")
// 	_ = err
// 	fmt.Printf("%s", out.Bytes())
// }

func formatEndpoint(ep endpoint) string {

	// use http.NewRequest to build up a unused request that we just use to spit
	// out a url string
	var request, _ = http.NewRequest(ep.Method, "", bytes.NewBufferString(""))
	request.URL.Host = fmt.Sprintf("%s:%d", "localhost", gPort)
	request.URL.Scheme = "http"
	request.URL.Path = ep.Path
	request.URL.RawQuery = url.Values.Encode(ep.Query)

	str := ""
	str = str + request.Method
	str = str + " "
	str = str + fmt.Sprintf("%d", ep.Status)
	str = str + " "
	str = str + request.URL.String()
	return str

}
