package config

// Test contains test configurations.
type Test struct {
	Asserts     []Assert  `yaml:"asserts"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Sets        []string  `yaml:"sets"`
	Selection   Selection `yaml:"selection"`
	Skip        bool      `yaml:"skip"`
	Values      []string  `yaml:"values"`
}

// NewTest returns a new instance of Test.
func NewTest() Test {
	return Test{
		Asserts: make([]Assert, 0),
		Sets:    make([]string, 0),
		Values:  make([]string, 0),
	}
}
