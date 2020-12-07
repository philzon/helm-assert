package report

// Test contains test information and results.
type Test struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Skipped     bool     `json:"skipped"`
	Asserts     []Assert `json:"asserts"`
	Score       Score    `json:"score"`
}

// NewTest returns a new instance of Test.
func NewTest() Test {
	return Test{
		Asserts: make([]Assert, 0),
	}
}
