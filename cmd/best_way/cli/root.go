package cli

import (
	"os"

	"github.com/joaovds/best-way-genetic-algorithm/internal/algorithm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "genetic-algorithm",
	Short: "Genetic Algorithm CLI Application",
	Long:  `A command-line application to run a genetic algorithm for optimizing routes.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	algorithm.LoadEnv()
	rootCmd.Flags().BoolP("help", "h", false, "best_way help commands")
}
