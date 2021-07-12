package runner

import (
	"time"

	"helm.sh/helm/v3/pkg/chart"

	"github.com/philzon/helm-assert/internal/manifest"
	"github.com/philzon/helm-assert/internal/output"
	"github.com/philzon/helm-assert/pkg/config"
	"github.com/philzon/helm-assert/pkg/report"
)

// Run takes a config and Helm chart as input, runs through all tests, and returns
// the report object.
//
// On YAML or template rendering errors, it returns an error object without the report.
func Run(cfg *config.Config, chrt *chart.Chart) (*report.Report, error) {
	rep := report.NewReport()
	rep.Chart.Name = chrt.Name()
	rep.Chart.Version = chrt.Metadata.Version
	rep.Chart.Icon = chrt.Metadata.Icon
	rep.Chart.Path = chrt.ChartFullPath()
	rep.Date = time.Now().Format(time.RFC3339)
	rep.Score.Total = len(cfg.Tests)

	for _, test := range cfg.Tests {
		// Render all manifests with values from test configuration.
		manifests, err := renderManifests(chrt, &test)

		// Any failures when rendering should put the application to a stop.
		if err != nil {
			return nil, err
		}

		// Any failures when verifying the YAML data should put the application to a stop.
		err = validateManifests(manifests)

		if err != nil {
			return nil, err
		}

		// Select manifests based on test requirements.
		manifests = manifest.GetManifestsByFiles(manifests, test.Select.Files)
		manifests = manifest.GetManifestsByKinds(manifests, test.Select.Kinds)
		manifests = manifest.GetManifestsByAPIVersions(manifests, test.Select.Versions)

		testReport := RunTest(manifests, &test)

		if testReport.Skipped {
			rep.Score.Skipped++
		} else {
			if testReport.Passed {
				rep.Score.Passed++
			} else {
				rep.Score.Failed++
			}
		}

		output.ConsoleSimple(testReport)

		rep.Tests = append(rep.Tests, testReport)
	}

	return &rep, nil
}

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
