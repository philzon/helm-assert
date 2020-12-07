package config

// Assert contains different objects to be used for different assert operations.
// It is only intended that one assert type is used.
type Assert struct {
	Exist Exist `yaml:"exist"`
	Equal Equal `yaml:"equal"`
}

// Exist is an assert type used to check if a key exist.
type Exist struct {
	Key string `yaml:"key"`
}

// Equal is an assert type used to check if a key contains the expected value.
type Equal struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}
