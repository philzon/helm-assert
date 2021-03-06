package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"

	"github.com/philzon/helm-assert/internal/log"
	"github.com/philzon/helm-assert/internal/output"
	"github.com/philzon/helm-assert/internal/runner"
	v3 "github.com/philzon/helm-assert/internal/v3"
	"github.com/philzon/helm-assert/pkg/config"
)

var (
	// Required positional arguments.
	chartPath  string
	configPath string

	// Command flags.
	json     string
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
		fmt.Printf("Fatal: %s\n", err.Error())
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "assert [OPTIONS] CONFIG CHART",
		DisableFlagsInUseLine: true,
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
	cmd.PersistentFlags().StringVarP(&json, "json", "", "", "write report to a file in JSON format")
	cmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "standard", "severity level to log (\"verbose\"|\"standard\"|\"quiet\"|\"none\")")
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
		return nil, err
	}

	cfg := &config.Config{}
	err = yaml.Unmarshal(data, cfg)

	if err != nil {
		return nil, fmt.Errorf("YAML parse error\n%s %s", configPath, yaml.FormatError(err, false, true))
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

		// Apply global sets and values to tests.
		cfg.Tests[i].Skip = ok
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
	rep, err := runner.Run(cfg, chrt)

	if err != nil {
		return err
	}

	// The value set is just a fallback value. This should be overriden to use
	// whatever path waas used to reference the Helm chart.
	rep.Chart.Path = chartPath

	if json != "" {
		return output.JSON(json, rep)
	}

	if rep.Score.Failed > 0 {
		os.Exit(1)
	}

	return nil
}
