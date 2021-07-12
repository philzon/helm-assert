package manifest

import (
	"path/filepath"
	"strings"
)

// Manifest cotains fields for manifests.
type Manifest struct {
	Path string
	Data []byte
}

// NewManifest returns a new instance of Manifest.
func NewManifest() Manifest {
	return Manifest{
		Data: make([]byte, 0),
	}
}

// NewManifestFromData creates and returns a new instance of Manifest with its
// path and data.
func NewManifestFromData(path string, data []byte) Manifest {
	return Manifest{
		Path: path,
		Data: data,
	}
}

// NewManifestsFromData creates and returns a slice of new instances of Manifest
// if the data contains multiple documents.
func NewManifestsFromData(path string, data []byte) []Manifest {
	var manifests []Manifest

	documents := GetDocuments(string(data))

	for _, document := range documents {
		manifests = append(manifests, NewManifestFromData(path, []byte(document)))
	}

	return manifests
}

// GetDocuments takes a string from a YAML document and attempts to return a
// slice of strings of the split content if there are directive(s) included.
//
// https://yaml.org/spec/1.2/spec.html#id2760395
func GetDocuments(data string) []string {
	var documents []string
	var buffer string

	lines := strings.Split(data, "\n")

	for count, line := range lines {
		if count == len(lines)-1 || strings.HasPrefix(line, "---") {
			buffer = strings.TrimSpace(buffer)

			if len(buffer) != 0 {
				documents = append(documents, buffer)
			}

			buffer = ""
		} else {
			buffer += line + "\n"
		}
	}

	return documents
}

// Base returns the manifest's base path.
func (m Manifest) Base() string {
	return filepath.Base(m.Path)
}
