package main

import (
	"log"
	"net/http"

	"github.com/Nezent/Go-Pay/cmd/api"
)

func main() {
	r := api.NewRouter()

	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
