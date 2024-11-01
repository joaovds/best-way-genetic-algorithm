package algorithm

type Config struct {
	MaxGenerations int
	PopulationSize int
}

func NewConfig() *Config {
	return &Config{}
}
