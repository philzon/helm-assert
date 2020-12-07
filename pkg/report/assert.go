package report

// Assert contains the test result from an assert.
type Assert struct {
	Index   int    `json:"index"`
	Passed  bool   `json:"passed"`
	Message string `json:"message"`
}
