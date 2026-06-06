package main

import (
	"log"
	"net/http"

	httpapi "public-aquarium-care-board/internal/http"
	"public-aquarium-care-board/internal/seed"
	"public-aquarium-care-board/internal/store"
)

func main() {
	s := store.New(seed.DefaultRecords())
	log.Fatal(http.ListenAndServe(":8080", httpapi.Router(s)))
}
