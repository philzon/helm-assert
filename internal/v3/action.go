package v3

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

var (
	Settings = cli.New()
)

func debug(format string, v ...interface{}) {
	if Settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}

func GetActionConfig() (*action.Configuration, error) {
	actionConfig := new(action.Configuration)

	err := actionConfig.Init(Settings.RESTClientGetter(), Settings.Namespace(), os.Getenv("HELM_DRIVER"), debug)

	if err != nil {
		return nil, err
	}

	return actionConfig, nil
}
