package config

// Test contains test configurations.
type Test struct {
	Name    string   `yaml:"name"`
	Summary string   `yaml:"summary"`
	Skip    bool     `yaml:"skip"`
	Sets    []string `yaml:"sets"`
	Values  []string `yaml:"values"`
	Select  Select   `yaml:"select"`
	Asserts []Assert `yaml:"asserts"`
}

// NewTest returns a new instance of Test.
func NewTest() Test {
	return Test{
		Sets:    make([]string, 0),
		Values:  make([]string, 0),
		Asserts: make([]Assert, 0),
	}
}
