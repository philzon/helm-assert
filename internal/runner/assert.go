package runner

import (
	"fmt"
	"strings"

	yaml "github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"

	"github.com/philzon/helm-assert/internal/manifest"
	"github.com/philzon/helm-assert/pkg/config"
	"github.com/philzon/helm-assert/pkg/report"
)

// RunAssert executes all asserts against all sources.
func RunAssert(manifest *manifest.Manifest, assert *config.Assert) report.Assert {
	assertReport := report.Assert{}
	assertReport.Manifest = manifest.Path

	if len(strings.TrimSpace(assert.Exist.Key)) > 0 {
		assertReport.Output, assertReport.Passed = AssertExist(assert.Exist.Key, manifest.Data)
	}

	if len(strings.TrimSpace(assert.Equal.Key)) > 0 {
		assertReport.Output, assertReport.Passed = AssertEqual(assert.Equal.Key, assert.Equal.Value, manifest.Data)
	}

	return assertReport
}

// AssertExist returns true if the given key exists.
func AssertExist(key string, data []byte) (string, bool) {
	path, err := yaml.PathString("$." + key)

	if err != nil {
		return fmt.Sprintf("parse error - %s", yaml.FormatError(err, false, false)), false
	}

	node, _ := path.ReadNode(strings.NewReader(string(data)))

	if node != nil {
		return fmt.Sprintf("key exist '%s'", key), true
	}

	return fmt.Sprintf("key does not exist '%s'", key), false
}

// AssertEqual returns true if the given key matches the value provided.
func AssertEqual(key, value string, data []byte) (string, bool) {
	path, err := yaml.PathString("$." + key)

	if err != nil {
		return fmt.Sprintf("parse error - %s", yaml.FormatError(err, false, false)), false
	}

	result := ""
	node, _ := path.ReadNode(strings.NewReader(string(data)))

	if node != nil {
		switch node.Type() {
		case ast.MappingType:
			result = "{...}"
		case ast.SequenceType:
			result = "[...]"
		default:
			result = node.String()
		}
	}

	if result == value {
		return fmt.Sprintf("got '%s' in key '%s'", result, key), true
	}

	return fmt.Sprintf("expected '%s', but got '%s' in key '%s'", value, result, key), false
}
