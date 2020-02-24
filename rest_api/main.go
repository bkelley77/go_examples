package main

import (
	"log"
	"net/http"
)

// to test the web server: curl http://localhost:8090/todos

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8090", router))
}
