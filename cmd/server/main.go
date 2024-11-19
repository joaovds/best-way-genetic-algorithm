package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
	"github.com/joaovds/best-way-genetic-algorithm/internal/distance"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var requestData api.LocationRequest

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		if err := requestData.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		startingPoint, coreLocations, err := requestData.ToCoreLocation()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error converting to core locations: %v", err), http.StatusBadRequest)
			return
		}
		config := algorithm.NewConfig(7000, 500, 4, 0.03)
		distanceCalculator := distance.NewInBatchCalculator()
		algorithmInstance := algorithm.NewAlgorithm(config, startingPoint, coreLocations, distanceCalculator)
		algorithmRes := algorithmInstance.Run()
		go algorithmInstance.RenderChart()

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(api.AlgorithmResponseToApiResponse(algorithmRes))
	})

	srv := &http.Server{
		Addr:         ":3333",
		Handler:      mainMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  100 * time.Second,
	}

	log.Println("Server running on port", 3333)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
