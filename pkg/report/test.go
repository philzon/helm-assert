package report

// Test contains test information and results.
type Test struct {
	Name      string     `json:"name"`
	Manifests []Manifest `json:"manifests"`
	Passed    bool       `json:"passed"`
	Summary   string     `json:"summary"`
	Skipped   bool       `json:"skipped"`
	Asserts   []Assert   `json:"asserts"`
	Score     Score      `json:"score"`
}

// NewTest returns a new instance of Test.
func NewTest() Test {
	return Test{
		Manifests: make([]Manifest, 0),
		Asserts: make([]Assert, 0),
	}
}
