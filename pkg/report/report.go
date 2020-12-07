package report

// Report contains the overall test report.
type Report struct {
	Chart   string `json:"chart"`
	Version string `json:"version"`
	Date    string `json:"date"`
	Score   Score  `json:"score"`
	Tests   []Test `json:"tests"`
}

// NewReport returns a new instance of Report.
func NewReport() Report {
	return Report{
		Tests: make([]Test, 0),
	}
}
