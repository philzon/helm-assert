package output

import (
	"strings"

	"github.com/philzon/helm-assert/pkg/report"
	"github.com/sirupsen/logrus"
)

const (
	StatusFailed  = "FAIL"
	StatusPassed  = "PASS"
	StatusSkipped = "SKIP"
	StatusNone    = "NONE"
)

// ConsoleSimple uses the log framework to print the test results given.
// This is used to get direct output from the test to the terminal.
func ConsoleSimple(testReport report.Test) {
	if testReport.Skipped {
		logrus.Infof("%s %s", StatusSkipped, testReport.Name)
		return
	}

	if testReport.Score.Failed != 0 {
		logrus.Errorf("%s %s", StatusFailed, testReport.Name)

		if testReport.Description != "" {
			logrus.Errorf("")

			for _, line := range strings.Split(strings.TrimSpace(testReport.Description), "\n") {
				logrus.Errorf("    %s", line)
			}
		}

		if len(testReport.Asserts) > 0 {
			logrus.Errorf("")

			for _, assert := range testReport.Asserts {
				if !assert.Passed {
					logrus.Errorf("    Assert %d: %s", assert.Index, assert.Message)
				}
			}

			logrus.Errorf("")
		}

		// Failures should be printed and then move on.
		return
	}

	// If nothing passed, check if it was skipped or never run.
	if testReport.Score.Passed == 0 {
		logrus.Infof("%s %s", StatusNone, testReport.Name)
	} else {
		logrus.Infof("%s %s", StatusPassed, testReport.Name)

		// This is code duplication from above error output.
		//
		// TODO: remove duplication and find a better way to present in the same manner
		// as error but using this log level.
		if testReport.Description != "" {
			logrus.Debugf("")

			for _, line := range strings.Split(strings.TrimSpace(testReport.Description), "\n") {
				logrus.Debugf("    %s", line)
			}
		}

		if len(testReport.Asserts) > 0 {
			logrus.Debugf("")

			for _, assert := range testReport.Asserts {
				logrus.Debugf("    Assert %d: %s", assert.Index, assert.Message)
			}

			logrus.Debugf("")
		}
	}
}
