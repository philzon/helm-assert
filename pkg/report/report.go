package report

// Report contains the overall test report.
type Report struct {
	Date    string `json:"date"`
	Chart   Chart  `json:"chart"`
	Score   Score  `json:"score"`
	Tests   []Test `json:"tests"`
}

// NewReport returns a new instance of Report.
func NewReport() Report {
	return Report{
		Tests: make([]Test, 0),
	}
}
