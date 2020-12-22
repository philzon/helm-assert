package config

// Select contains lists to be used to select manifests from different methods.
type Select struct {
	Files    []string `yaml:"files"`
	Kinds    []string `yaml:"kinds"`
	Versions []string `yaml:"versions"`
}
