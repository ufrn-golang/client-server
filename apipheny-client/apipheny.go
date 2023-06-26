package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://www.boredapi.com/api/activity"

type Activity struct {
    Description     string `json:"activity"`
    Type            string
    Participants    int
}

func main() {
    if res, err := http.Get(url); err != nil {
        log.Fatal(err)
    } else {
        contents, _ := io.ReadAll(res.Body)
        defer res.Body.Close()

        var a Activity
        err := json.Unmarshal(contents, &a)
        if err != nil {
            log.Fatal(err)
        } else {
            fmt.Printf("What? %s\n", a.Description)
            fmt.Printf("What kind? %s\n", a.Type)
            fmt.Printf("How many? %d\n", a.Participants)
        }
    }
} 
