package main

import "fmt"
import "bytes"

// import "encoding/json"
import "net/http"

// func prettyPrintJSON(str string) {
// 	byt := []byte(str)
// 	var out bytes.Buffer
// 	err := json.Indent(&out, byt, "", "  ")
// 	_ = err
// 	fmt.Printf("%s", out.Bytes())
// }

func formatEndpoint(ep endpoint) string {

	// make request just to format URL for output
	var request, _ = http.NewRequest(ep.Method, "", bytes.NewBufferString(""))
	request.URL.Host = fmt.Sprintf("%s:%d", "localhost", gPort)
	request.URL.Scheme = "http"
	request.URL.Path = ep.Path

	str := ""
	str = str + request.Method
	str = str + " "
	str = str + fmt.Sprintf("%d", ep.Status)
	str = str + " "
	str = str + request.URL.String()
	return str

}
