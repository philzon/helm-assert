package manifest

import (
	"testing"
)

const (
	data = `---
apiVersion: apps/v1
kind: Deployment
---
apiVersion: apps/v1
kind: ConfigMap
---
apiVersion: apps/v1
kind: StatefulSet
---
apiVersion: apps/v1
kind: Service`
)

func TestNewManifest(t *testing.T) {
	manifest := NewManifest()

	if manifest.Path != "" {
		t.Fatalf("expected empty string, but got '%s'", manifest.Path)
	}

	if len(manifest.Data) != 0 {
		t.Fatalf("expected empty byte array, but got %d number of bytes", len(manifest.Data))
	}
}

func TestNewManifestFromData(t *testing.T) {
	file := "test.yaml"
	manifest := NewManifestFromData(file, []byte(data))

	if manifest.Path != file {
		t.Fatalf("expected '%s', but got '%s'", file, manifest.Path)
	}

	if string(manifest.Data) != data {
		t.Fatalf("expected to get same manifest data, but did not")
	}
}

func TestNewManifestsFromData(t *testing.T) {
	file := "test.yaml"
	manifests := NewManifestsFromData(file, []byte(data))

	expect := 4
	got := len(manifests)

	if got != expect {
		t.Fatalf("expected %d, but got %d", expect, got)
	}

	for _, manifest := range manifests {
		if manifest.Path != file {
			t.Fatalf("expected '%s', but got '%s'", file, manifest.Path)
		}
	}
}

func TestGetDocuments(t *testing.T) {
	docs := GetDocuments(data)

	expect := 4
	got := len(docs)

	if got != expect {
		t.Fatalf("expected %d, but got %d", expect, got)
	}
}
