package cli

import (
	"fmt"
	"log"

	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
	"github.com/spf13/cobra"
)

var (
	populationSize int
	numGenerations int
	numElites      int
	mutationRate   float64
	numLocations   int
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the genetic algorithm",
	Long:  `Run the genetic algorithm with specified parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		locations := generateMockLocations(numLocations)
		startingPoint, coreLocations, err := locations.ToCoreLocation()
		if err != nil {
			log.Fatalf("Error converting to core locations: %v", err)
		}
		config := algorithm.NewConfig(populationSize, numGenerations, numElites, mutationRate)
		algorithmInstance := algorithm.NewAlgorithm(config, startingPoint, coreLocations)
		algorithmInstance.Run()
		algorithmInstance.RenderChart()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVarP(&populationSize, "max-population", "p", 7000, "Size of the population")
	runCmd.Flags().IntVarP(&numGenerations, "max-generations", "g", 300, "Number of generations")
	runCmd.Flags().IntVarP(&numElites, "elites", "e", 4, "Number of elite individuals")
	runCmd.Flags().Float64VarP(&mutationRate, "mutation", "m", 0.4, "Mutation rate")
	runCmd.Flags().IntVarP(&numLocations, "locations", "l", 10, "Number of locations")
}

func generateMockLocations(size int) api.LocationRequest {
	locations := make(api.LocationRequest, size)
	for i := 0; i < size; i++ {
		address := fmt.Sprintf("any_address_%d", i+1)
		isStarting := (i == 0)
		locations[i] = api.Location{Address: address, IsStarting: isStarting}
	}
	return locations
}
