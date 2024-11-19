package algorithm

type (
	Config struct {
		MaxPopulationSize int
		MaxGenerations    int
		ElitismNumber     int
		MutationRate      float64
	}
)

func NewConfig(maxPopulationSize, maxGenerations, elitismNumber int, mutationRate float64) *Config {
	return &Config{
		MaxPopulationSize: maxPopulationSize,
		MaxGenerations:    maxGenerations,
		ElitismNumber:     elitismNumber,
		MutationRate:      mutationRate,
	}
}
