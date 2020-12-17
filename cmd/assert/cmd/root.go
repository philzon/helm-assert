package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	yaml "github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/engine"
	"helm.sh/helm/v3/pkg/getter"

	"github.com/philzon/helm-assert/internal/app"
	"github.com/philzon/helm-assert/internal/log"
	"github.com/philzon/helm-assert/internal/manifest"
	"github.com/philzon/helm-assert/internal/output"
	"github.com/philzon/helm-assert/internal/runner"
	v3 "github.com/philzon/helm-assert/internal/v3"
	"github.com/philzon/helm-assert/pkg/config"
	"github.com/philzon/helm-assert/pkg/report"
)

var (
	// Required positional arguments.
	chartPath  string
	configPath string

	// Command flags.
	json     bool
	logLevel string
	out      string
	password string
	skips    []string
	repo     string
	username string
	version  string
)

// Execute runs root command.
func Execute() {
	newRootCmd().Execute()
}

func dieOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: fmt.Sprintf("%s [%s] [%s]", app.Name, "CONFIG", "CHART"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				cmd.Usage()
				os.Exit(1)
			}

			configPath = args[0]
			chartPath = args[1]

			cfg, err := initConfig(configPath, skips)
			dieOnError(err)

			chart, err := getChart(chartPath)
			dieOnError(err)

			logrus.SetFormatter(&log.SimpleFormatter{})
			logrus.SetLevel(logrus.InfoLevel)

			switch strings.ToLower(logLevel) {
			case "verbose":
				logrus.SetLevel(logrus.DebugLevel)
			case "standard":
				logrus.SetLevel(logrus.InfoLevel)
			case "quiet":
				logrus.SetLevel(logrus.ErrorLevel)
			case "none":
				logrus.SetLevel(logrus.PanicLevel)
			default:
				fmt.Printf("Unknown log level '%s'\n", logLevel)
				os.Exit(1)
			}

			err = runAssert(cfg, chart)
			dieOnError(err)
		},
	}

	addRootFlags(cmd)

	return cmd
}

func addRootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&json, "json", "", false, "report should be saved in JSON format")
	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "standard", "severity level to log (\"verbose\"|\"standard\"|\"quiet\"|\"none\")")
	cmd.PersistentFlags().StringVarP(&out, "output", "o", "report", "path to store reports to")
	cmd.PersistentFlags().StringArrayVarP(&skips, "skip", "", []string{}, "skip test by name (can specify multiple)")

	// Replicated Helm flags.
	cmd.PersistentFlags().StringVarP(&password, "password", "", "", "chart repository password where to locate the requested chart")
	cmd.PersistentFlags().StringVarP(&username, "username", "", "", "chart repository username where to locate the requested chart")
	cmd.PersistentFlags().StringVarP(&username, "repo", "", "", "chart repository url where to locate the requested chart")
	cmd.PersistentFlags().StringVarP(&version, "version", "", "", "specify the exact chart version to use. If this is not specified, the latest version is used")
}

func initConfig(configPath string, skips []string) (*config.Config, error) {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, fmt.Errorf("could not open '%s': %s", configPath, err.Error())
	}

	cfg := &config.Config{}
	err = yaml.Unmarshal(data, cfg)

	if err != nil {
		return nil, fmt.Errorf("could not parse config YAML: %s", err.Error())
	}

	// Put all skip strings in a map and iterate all tests and enable
	// skip option for matching tests.
	skipMap := make(map[string]bool, 0)

	for _, skip := range skips {
		skipMap[skip] = true
	}

	// Perform changes to all test configurations.
	for i := range cfg.Tests {
		// Enable skip boolean on tests on matching skip flags.
		_, ok := skipMap[cfg.Tests[i].Name]

		if ok {
			cfg.Tests[i].Skip = true
		}

		// Apply global sets and values to tests.
		cfg.Tests[i].Sets = append(cfg.Tests[i].Sets, cfg.Sets...)
		cfg.Tests[i].Values = append(cfg.Tests[i].Values, cfg.Values...)
	}

	return cfg, nil
}

func getChart(chart string) (*chart.Chart, error) {
	actionConfig, err := v3.GetActionConfig()

	if err != nil {
		return nil, err
	}

	// Install client is needed in case the chart need to be fetched from a URL.
	client := action.NewInstall(actionConfig)
	client.ChartPathOptions.Password = password
	client.ChartPathOptions.Username = username
	client.ChartPathOptions.Version = version
	client.ChartPathOptions.RepoURL = repo

	chartPath, err := client.ChartPathOptions.LocateChart(chart, v3.Settings)

	if err != nil {
		return nil, err
	}

	chrt, err := loader.Load(chartPath)

	if err != nil {
		return nil, err
	}

	switch chrt.Metadata.Type {
	case "", "application":
	default:
		return nil, fmt.Errorf("chart is not of type 'application'")
	}

	return chrt, nil
}

func runAssert(cfg *config.Config, chrt *chart.Chart) error {
	rep := report.NewReport()
	rep.Chart = chrt.Name()
	rep.Version = chrt.Metadata.Version
	rep.Date = time.Now().Format(time.RFC3339)
	rep.Score.Total = len(cfg.Tests)

	for _, test := range cfg.Tests {
		// Render all manifests with values from test configuration.
		manifests, err := renderManifests(chrt, &test)

		// FIXME: possible invalid manifests should be part of test failure
		// and not fail the whole run.
		if err != nil {
			return err
		}

		// Select manifests based on test requirements.
		manifests = manifest.GetManifestsByNames(manifests, test.Select.Files)
		manifests = manifest.GetManifestsByKinds(manifests, test.Select.Kinds)
		manifests = manifest.GetManifestsByAPIVersions(manifests, test.Select.Versions)

		testReport := runner.RunTest(manifests, &test)

		if testReport.Skipped {
			rep.Score.Skipped++
		} else {
			if testReport.Score.Failed > 0 {
				rep.Score.Failed++
			} else {
				if testReport.Score.Passed > 0 {
					rep.Score.Passed++
				}
			}
		}

		if logLevel != "none" {
			output.ConsoleSimple(testReport)
		}

		rep.Tests = append(rep.Tests, testReport)
	}

	if json {
		return output.JSON(out, &rep)
	}

	return nil
}

func renderManifests(chrt *chart.Chart, test *config.Test) ([]manifest.Manifest, error) {
	manifests := make([]manifest.Manifest, 0)

	// Generate values for the rendering engine to use.
	// This covers for both sets (--set) and values (--values).
	valueOpts := values.Options{}
	valueOpts.Values = append(valueOpts.StringValues, test.Sets...)
	valueOpts.ValueFiles = append(valueOpts.ValueFiles, test.Values...)

	vals, err := valueOpts.MergeValues(getter.All(v3.Settings))

	if err != nil {
		return manifests, err
	}

	// Set default release options. This will then result in both values from
	// Release.Name and Release.Namespace to be set.
	options := chartutil.ReleaseOptions{
		Name:      "RELEASE",
		Namespace: "NAMESPACE",
	}
	capabilities := &chartutil.Capabilities{}
	valuesToRender, err := chartutil.ToRenderValues(chrt, vals, options, capabilities)

	if err != nil {
		return manifests, err
	}

	renderedManifests, err := engine.Render(chrt, valuesToRender)

	if err != nil {
		return manifests, err
	}

	for name, data := range renderedManifests {
		// Only accept manifests from files that has YAML extension. This is to also
		// prevent non-template files, like NOTES.txt, from being parsed as YAML later.
		ext := filepath.Ext(name)

		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		if len(data) != 0 {
			// Create manifests for each document present in a rendered manifest. These
			// manifests will still use the same filename from the file they were split
			// from.
			//
			// If there are multiple documents in a single manifest, and contain different
			// Kubernetes resource kinds, selecting by file will fail the assert, and so
			// must also be selected by resource kind.
			documents := manifest.SplitDocument(data)

			for _, document := range documents {
				manifests = append(manifests, manifest.NewManifestFromData(name, []byte(document)))
			}
		}
	}

	return manifests, nil
}
