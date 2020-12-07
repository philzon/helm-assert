package runner

import (
	"fmt"

	"github.com/philzon/helm-assert/internal/yaml"
	"github.com/smallfish/simpleyaml"
)

// AssertExist returns true if the given key exists.
func AssertExist(key string, tree *simpleyaml.Yaml) (string, bool) {
	result, err := yaml.Exist(key, tree)

	if err != nil {
		return fmt.Sprintf("parse error - %s", err.Error()), false
	}

	if result {
		return fmt.Sprintf("key exist '%s'", key), true
	}

	return fmt.Sprintf("key does not exist '%s'", key), false
}

// AssertEqual returns true if the given key matches the value provided.
func AssertEqual(key, value string, tree *simpleyaml.Yaml) (string, bool) {
	result, err := yaml.Get(key, tree)

	if err != nil {
		return fmt.Sprintf("parse error - %s", err.Error()), false
	}

	if value == result {
		return fmt.Sprintf("got '%s' in key '%s'", result, key), true
	}

	return fmt.Sprintf("expected '%s', but got '%s' in key '%s'", value, result, key), false
}
