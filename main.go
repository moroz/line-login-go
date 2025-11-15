package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", handleHome)

	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
