package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// The httpbin service is a tool for testing HTTP clients.
// The /post endpoint returns data that was posted
// to it along with information on the client.
const endpoint = "https://httpbin.org/post"

func main() {
	data := strings.NewReader(`{ "activity": "Learn to play a new instrument",
  "type": "music" }`)
  	if res, err := http.Post(endpoint, "application/json", data); err != nil {
		log.Fatal(err)
	} else {
		defer res.Body.Close()
		contents, _ := io.ReadAll(res.Body)
		fmt.Printf("%s", contents)
	}
}