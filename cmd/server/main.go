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

func init() {
	algorithm.LoadEnv()
}

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
		config := algorithm.NewConfig(7000, 500, 4, 0.3)
		distanceCalculator := distance.NewDistanceMatrixGoogle(algorithm.ENV.MAPS_API_KEY)
		algorithmInstance := algorithm.NewAlgorithm(config, startingPoint, coreLocations, distanceCalculator)
		algorithmRes := algorithmInstance.Run()
		go algorithmInstance.RenderChart()

		chartsHTML, err := algorithmInstance.ChartHTML()
		if err != nil {
			chartsHTML = "Error when making charts"
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(api.AlgorithmResponseToApiResponse(algorithmRes, chartsHTML))
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", algorithm.ENV.ServerPort),
		Handler:      mainMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  100 * time.Second,
	}

	log.Println("Server running on port", algorithm.ENV.ServerPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
