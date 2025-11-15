package main

import (
	"log"
	"net/http"

	"github.com/moroz/line-login-go/handlers"
)

func main() {
	r := handlers.Router()
	log.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
