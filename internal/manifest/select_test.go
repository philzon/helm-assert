package manifest

import (
	"strings"
	"testing"
)

// Meant to be constant - do not mutate it!
var files = []Manifest{
	{
		Path: "templates/configmap.yaml",
	},
	{
		Path: "templates/cronjob.yaml",
	},
	{
		Path: "templates/daemonset.yaml",
	},
	{
		Path: "templates/deployment.yaml",
	},
	{
		Path: "templates/job.yaml",
	},
	{
		Path: "templates/replicaset.yaml",
	},
	{
		Path: "templates/service.yaml",
	},
	{
		Path: "templates/statefulset.yaml",
	},
}

// Meant to be constant - do not mutate it!
var kinds = []Manifest{
	{
		Data: []byte("kind: ConfigMap"),
	},
	{
		Data: []byte("kind: CronJob"),
	},
	{
		Data: []byte("kind: DaemonSet"),
	},
	{
		Data: []byte("kind: Deployment"),
	},
	{
		Data: []byte("kind: Job"),
	},
	{
		Data: []byte("kind: ReplicaSet"),
	},
	{
		Data: []byte("kind: Service"),
	},
	{
		Data: []byte("kind: StatefulSet"),
	},
}

// Meant to be constant - do not mutate it!
var apiVersions = []Manifest{
	{
		Data: []byte("apiVersion: v1"),
	},
	{
		Data: []byte("apiVersion: v2"),
	},
	{
		Data: []byte("apiVersion: v3"),
	},
	{
		Data: []byte("apiVersion: v4"),
	},
	{
		Data: []byte("apiVersion: v5"),
	},
	{
		Data: []byte("apiVersion: v6"),
	},
	{
		Data: []byte("apiVersion: v7"),
	},
	{
		Data: []byte("apiVersion: v8"),
	},
}

func TestGetManifestsByFiles(t *testing.T) {
	// To avoid having to do a linear search for all resources returned, this list
	// needs to be in the same alphabetical order as the defined test content in
	// 'files' variable.
	expect := []string{
		"templates/configmap.yaml",
		"templates/cronjob.yaml",
		"templates/daemonset.yaml",
		"templates/deployment.yaml",
		"templates/job.yaml",
		"templates/replicaset.yaml",
		"templates/service.yaml",
		"templates/statefulset.yaml",
	}

	got := GetManifestsByFiles(files, expect)

	if len(got) != len(expect) {
		t.Fatalf("expected returned count %d, but got %d", len(expect), len(got))
	}

	for i := 0; i < len(expect); i++ {
		if got[i].Path != expect[i] {
			t.Fatalf("expected '%s', but got '%s'", expect[i], got[i].Path)
		}
	}
}

func TestGetManifestsByKinds(t *testing.T) {
	// To avoid having to do a linear search for all resources returned, this list
	// needs to be in the same alphabetical order as the defined test content in
	// 'kinds' variable.
	expect := []string{
		"ConfigMap",
		"CronJob",
		"DaemonSet",
		"Deployment",
		"Job",
		"ReplicaSet",
		"Service",
		"StatefulSet",
	}

	got := GetManifestsByKinds(kinds, expect)

	if len(got) != len(expect) {
		t.Fatalf("expected returned count %d, but got %d", len(expect), len(got))
	}

	for i := range expect {
		if !strings.Contains(string(got[i].Data), expect[i]) {
			t.Fatalf("expected 'kind: %s', but got '%s'", expect[i], got[i].Data)
		}
	}
}

func TestGetManifestsByAPIVersions(t *testing.T) {
	// To avoid having to do a linear search for all resources returned, this list
	// needs to be in the same alphabetical order as the defined test content in
	// 'apiVersions' variable.
	expect := []string{
		"v1",
		"v2",
		"v3",
		"v4",
		"v5",
		"v6",
		"v7",
		"v8",
	}

	got := GetManifestsByAPIVersions(apiVersions, expect)

	if len(got) != len(expect) {
		t.Fatalf("expected returned count %d, but got %d", len(expect), len(got))
	}

	for i := range expect {
		if !strings.Contains(string(got[i].Data), expect[i]) {
			t.Fatalf("expected 'kind: %s', but got '%s'", expect[i], got[i].Data)
		}
	}
}
