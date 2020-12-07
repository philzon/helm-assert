package manifest

import (
	"path/filepath"

	"github.com/philzon/helm-assert/internal/yaml"
	"github.com/smallfish/simpleyaml"
)

const (
	keyKind       = "kind"
	keyAPIVersion = "apiVersion"
)

// GetManifestsByNames returns manifests that matches by provided base filenames.
// This only works for manifests which actually has a filename associated.
// Anonymous manifests will never match by names.
func GetManifestsByNames(manifests []Manifest, names []string) []Manifest {
	if len(names) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, name := range names {
			// Match only by filename and extension - ignore parents.
			if filepath.Base(name) == filepath.Base(manifest.Base()) {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}

// GetManifestsByKinds returns manifests that matches the values of key `kind`.
func GetManifestsByKinds(manifests []Manifest, kinds []string) []Manifest {
	if len(kinds) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, kind := range kinds {
			tree, _ := simpleyaml.NewYaml(manifest.Data)

			value, err := yaml.Get(keyKind, tree)

			if err != nil {
				break
			}

			if value == kind {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}

// GetManifestsByAPIVersions returns manifests that matches the values of key `apiVersion`.
func GetManifestsByAPIVersions(manifests []Manifest, apiVersions []string) []Manifest {
	if len(apiVersions) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, apiVersion := range apiVersions {
			tree, _ := simpleyaml.NewYaml(manifest.Data)

			value, err := yaml.Get(keyKind, tree)

			if err != nil {
				break
			}

			if value == apiVersion {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}
