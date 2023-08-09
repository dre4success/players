package main

import (
	"github.com/dre4success/players"
	"log"
	"net/http"
)

func main() {
	server := poker.NewPlayerServer(poker.NewInMemoryPlayerStore())

	if err := http.ListenAndServe(":7080", server); err != nil {
		log.Fatalf("could not listen on port 7080 %v", err)
	}
}
