package report

// Test contains test information and results.
type Test struct {
	Name      string     `json:"name"`
	Summary   string     `json:"summary"`
	Passed    bool       `json:"passed"`
	Skipped   bool       `json:"skipped"`
	Score     Score      `json:"score"`
	Manifests []Manifest `json:"manifests"`
	Asserts   []Assert   `json:"asserts"`
}

// NewTest returns a new instance of Test.
func NewTest() Test {
	return Test{
		Manifests: make([]Manifest, 0),
		Asserts: make([]Assert, 0),
	}
}
