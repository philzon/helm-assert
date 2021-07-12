package manifest

import (
	"path/filepath"

	yaml "github.com/goccy/go-yaml"
)

// k8sKind is a specific key in the Kubernetes manifest which determines the
// resource's kind or type.
type k8sKind struct {
	Kind string `yaml:"kind"`
}

// k8sAPIVersion is a specific key in the Kubernetes manifest which determines
// the Kubernetes API version to create the resource.
type k8sAPIVersion struct {
	APIVersion string `yaml:"apiVersion"`
}

// GetManifestsByPaths returns manifests that matches by provided base
// filenames. This only works for manifests which actually has a filename
// associated.
//
// Anonymous manifests will never match by names.
func GetManifestsByFiles(manifests []Manifest, files []string) []Manifest {
	if len(files) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, file := range files {
			// Match only by filename and extension - ignore parents.
			if filepath.Base(file) == filepath.Base(manifest.Base()) {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}

// GetManifestsByKinds returns manifests that matches the values of key `kind`.
//
// Manifest data that contains YAML errors will be silently ignored.
func GetManifestsByKinds(manifests []Manifest, kinds []string) []Manifest {
	if len(kinds) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, kind := range kinds {
			cfg := &k8sKind{}
			err := yaml.Unmarshal(manifest.Data, cfg)

			// Fail silently on YAML error. This should be checked before selecting by
			// kind(s).
			if err != nil {
				continue
			}

			if cfg.Kind == kind {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}

// GetManifestsByAPIVersions returns manifests that matches the values of key
// `apiVersion`.
//
// Manifest data that contains YAML errors will be silently ignored.
func GetManifestsByAPIVersions(manifests []Manifest, apiVersions []string) []Manifest {
	if len(apiVersions) == 0 {
		return manifests
	}

	var selected []Manifest

	for _, manifest := range manifests {
		for _, apiVersion := range apiVersions {
			cfg := &k8sAPIVersion{}
			err := yaml.Unmarshal(manifest.Data, cfg)

			// Fail silently on YAML error. This should be checked before selecting by API version(s).
			if err != nil {
				continue
			}

			if cfg.APIVersion == apiVersion {
				selected = append(selected, manifest)
			}
		}
	}

	return selected
}
