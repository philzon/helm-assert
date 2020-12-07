package config

// Config contains the overall tests and their configurations.
type Config struct {
	Sets   []string `yaml:"sets"`
	Tests  []Test   `yaml:"tests"`
	Values []string `yaml:"values"`
}

// NewConfig returns a new instance of Config.
func NewConfig() Config {
	return Config{
		Sets:   make([]string, 0),
		Tests:  make([]Test, 0),
		Values: make([]string, 0),
	}
}
