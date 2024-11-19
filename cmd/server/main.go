package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		requestData := new(api.LocationRequest)

		if err := json.NewDecoder(r.Body).Decode(requestData); err != nil {
			http.Error(w, "failed to decode request body", http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(200)
	})

	log.Println("Server running on port", 3333)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "3333"), mainMux); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
