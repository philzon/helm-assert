package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/philzon/helm-assert/pkg/report"
)

// JSON takes a path and the test report and writes it to a file as JSON.
func JSON(path string, rep *report.Report) error {
	file, err := os.Create(fmt.Sprintf("%s.json", path))

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.Marshal(rep)

	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}
