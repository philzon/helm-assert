package config

// Selection contains lists to be used to select sources from different methods.
type Selection struct {
	Kinds    []string `yaml:"kinds"`
	Files    []string `yaml:"files"`
	Versions []string `yaml:"versions"`
}
