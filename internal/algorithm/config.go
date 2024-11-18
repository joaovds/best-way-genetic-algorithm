package algorithm

type (
	Config struct {
		MaxPopulationSize int
		MaxGenerations    int
	}
)

func NewConfig(maxPopulationSize, maxGenerations int) *Config {
	return &Config{
		MaxPopulationSize: maxPopulationSize,
		MaxGenerations:    maxGenerations,
	}
}
