package report

// Manifest contains the manifest path and data selected for the test report.
type Manifest struct {
	Path string `json:"path"`
	Data string `json:"data"`
}
