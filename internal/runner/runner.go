package runner

import (
	"github.com/philzon/helm-assert/internal/manifest"
	"github.com/philzon/helm-assert/pkg/config"
	"github.com/philzon/helm-assert/pkg/report"
)

// RunTest applies the test configuration to the templates and executes
// all defined asserts.
func RunTest(manifests []manifest.Manifest, test *config.Test) report.Test {
	testReport := report.NewTest()
	testReport.Name = test.Name
	testReport.Summary = test.Summary
	testReport.Score.Total = len(test.Asserts)
	testReport.Passed = true

	// Fail test cases when no sources are presented.
	if len(manifests) == 0 || test.Skip {
		testReport.Skipped = test.Skip
		testReport.Score.Skipped = len(test.Asserts)
		return testReport
	}

	for _, manifest := range manifests {
		// Add selected manifest name and data to test report.
		manifestReport := report.Manifest{
			Path: manifest.Path,
			Data: string(manifest.Data),
		}
		testReport.Manifests = append(testReport.Manifests, manifestReport)

		for index, assert := range test.Asserts {
			assertReport := RunAssert(&manifest, &assert)
			assertReport.Index = index

			if assertReport.Passed {
				testReport.Score.Passed++
			} else {
				testReport.Score.Failed++
			}

			testReport.Asserts = append(testReport.Asserts, assertReport)
		}
	}

	if testReport.Score.Failed > 0 {
		testReport.Passed = false
	}

	return testReport
}
