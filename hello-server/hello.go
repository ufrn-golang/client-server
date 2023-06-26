package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sayHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, stranger"))
}

func printNow(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	fmt.Fprint(res, currentTime.Format("02/01/2006 15:04:05"))
}

func main() {
	server := &http.Server{
		Addr: "localhost:4000",
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/now", printNow)
	log.Fatal(server.ListenAndServe())
}