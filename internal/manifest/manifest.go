package manifest

import "path/filepath"

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

// NewManifestFromData creates and returns a new instance of Manifest with its path and data.
func NewManifestFromData(path string, data []byte) Manifest {
	return Manifest{
		Path: path,
		Data: data,
	}
}

// TODO: implement splits for anonymous manifests.
//
// SplitDirectives takes a manifest and returns a slice of extra directives and removing them
// from the main manifest itself. If returns nil, no extra directives were found.
//
// A directive is delimited by three dashes ("---") above the content:
// ---
// foo: bar
// ---
// baz: 123
//
// https://yaml.org/spec/1.2/spec.html#id2760395
func SplitDirectives(manifest *Manifest) []Manifest {

	return nil
}

// Base returns the manifest's base path.
func (m Manifest) Base() string {
	return filepath.Base(m.Path)
}
