package main

import (
	"log"
	"net/http"

	"github.com/alexshd/total-coder-w19/services/auction/api"
)

func main() {
	handler := http.HandlerFunc(api.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
