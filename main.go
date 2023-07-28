package main

import (
	"log"
	"net/http"
)


func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":7080", server))
}
