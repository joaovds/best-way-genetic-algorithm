package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/joaovds/best-way-genetic-algorithm/internal/api"
	"github.com/joaovds/best-way-genetic-algorithm/internal/distance"
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
		distanceCalculator := distance.NewInBatchCalculator()
		algorithmInstance := algorithm.NewAlgorithm(config, startingPoint, coreLocations, distanceCalculator)

		start := time.Now()
		algorithmRes := algorithmInstance.Run()

		algorithmInstance.RenderChart()
		chartsHTML, err := algorithmInstance.ChartHTML()
		if err != nil {
			chartsHTML = "Error when making charts"
		}
		printOutput(algorithmRes, chartsHTML, start)
	},
}

func printOutput(response *algorithm.AlgorithmResponse, chartsHtml string, start time.Time) {
	resString, _ := json.MarshalIndent(api.AlgorithmResponseToApiResponse(response, chartsHtml, start), "", "  ")
	fmt.Println(string(resString))
	fmt.Println("----- ... -----")
	fmt.Println("ID:", response.BestWay.StartingPoint.GetID(), "=>", response.BestWay.StartingPoint.Address)
	for _, point := range response.BestWay.Genes {
		fmt.Println("ID:", point.GetID(), "=>", point.Address)
	}
	fmt.Println("----- ... -----")
	fmt.Println("\033[36mRoute")
	fmt.Print(response.BestWay.StartingPoint.GetID())
	for _, point := range response.BestWay.Genes {
		fmt.Print(" -> ", point.GetID())
	}
	fmt.Println("\033[0m")
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
	locations := make([]api.Location, size)
	for i := 0; i < size; i++ {
		address := fmt.Sprintf("any_address_%d", i+1)
		isStarting := (i == 0)
		locations[i] = api.Location{Address: address, IsStarting: isStarting}
	}
	return api.LocationRequest{Locations: locations}
}
