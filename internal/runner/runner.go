package runner

import (
	"strings"

	"github.com/smallfish/simpleyaml"

	"github.com/philzon/helm-assert/internal/manifest"
	"github.com/philzon/helm-assert/pkg/config"
	"github.com/philzon/helm-assert/pkg/report"
)

// RunTest applies the test configuration to the templates and executes
// all defined asserts.
func RunTest(manifests []manifest.Manifest, test *config.Test) report.Test {
	testReport := report.NewTest()
	testReport.Name = test.Name
	testReport.Description = test.Description
	testReport.Score.Total = len(test.Asserts)

	// Fail test cases when no sources are presented.
	if len(manifests) == 0 || test.Skip {
		testReport.Skipped = test.Skip
		testReport.Score.Skipped = len(test.Asserts)
		return testReport
	}

	for _, manifest := range manifests {
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

	return testReport
}

// RunAssert executes all asserts against all sources.
func RunAssert(manifest *manifest.Manifest, assert *config.Assert) report.Assert {
	tree, err := simpleyaml.NewYaml(manifest.Data)

	if err != nil {
		return report.Assert{
			Message: "Could not parse below YAML:\n\n%s\n",
			Passed:  false,
		}
	}

	assertReport := report.Assert{}

	var message string
	var passed bool

	if len(strings.TrimSpace(assert.Exist.Key)) > 0 {
		message, passed = AssertExist(assert.Exist.Key, tree)
	}

	if len(strings.TrimSpace(assert.Equal.Key)) > 0 {
		message, passed = AssertEqual(assert.Equal.Key, assert.Equal.Value, tree)
	}

	assertReport.Message = message
	assertReport.Passed = passed

	return assertReport
}
