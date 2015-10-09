package main

import (
    "log"
    "net/http"
    
    "github.com/mattlaver/marketdata"
)

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}