package runner

import (
	"fmt"
	"path/filepath"

	yaml "github.com/goccy/go-yaml"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/engine"
	"helm.sh/helm/v3/pkg/getter"

	"github.com/philzon/helm-assert/internal/manifest"
	v3 "github.com/philzon/helm-assert/internal/v3"
	"github.com/philzon/helm-assert/pkg/config"
)

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
			manifests = append(manifests, manifest.NewManifestsFromData(name, []byte(data))...)
		}
	}

	return manifests, nil
}

func validateManifests(manifests []manifest.Manifest) error {
	for _, manifest := range manifests {
		err := yaml.Unmarshal(manifest.Data, new(interface{}))

		if err != nil {
			return fmt.Errorf("YAML parse error\n%s %s", manifest.Path, yaml.FormatError(err, false, true))
		}
	}

	return nil
}
