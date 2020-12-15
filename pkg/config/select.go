package config

// Select contains lists to be used to select manifests from different methods.
type Select struct {
	Kinds    []string `yaml:"kinds"`
	Files    []string `yaml:"files"`
	Versions []string `yaml:"versions"`
}
