package manifest

import (
	"testing"
)

const (
	document = `---
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

func TestSplitDocument(t *testing.T) {
	docs := SplitDocument(document)

	expected := 4
	got := len(docs)

	if got != expected {
		t.Errorf("expected %d, but got %d", expected, got)
	}
}
