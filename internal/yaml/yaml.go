package yaml

import (
	"strconv"

	"github.com/smallfish/simpleyaml"
)

// Exist returns true if the key exists.
func Exist(key string, yaml *simpleyaml.Yaml) (bool, error) {
	y, err := locate(key, yaml)

	if err != nil {
		return false, err
	}

	return y != nil, nil
}

// Get returns the value from a key, if it exists.
//
// The value of the key is determined by its content:
//
// hello      - the value is a string
// 12         - the value is a number
// 1.2        - the value is a float
// true|false - the value is a boolean
// nil        - the value is nil, not found or its type could not be parsed
// {...}      - the value is an object
// [...]      - the value is an array
func Get(key string, yaml *simpleyaml.Yaml) (string, error) {
	y, err := locate(key, yaml)

	if err != nil {
		return "", err
	}

	if y == nil {
		return "nil", nil
	}

	if y.IsMap() {
		return "{...}", nil
	}

	if y.IsArray() {
		return "[...]", nil
	}

	// TODO: find a library which can grab the value (not an object or array) as a string.
	str, err := y.String()

	if err == nil {
		return str, nil
	}

	number, err := y.Int()

	if err == nil {
		return strconv.Itoa(number), nil
	}

	flo, err := y.Float()

	if err == nil {
		return strconv.FormatFloat(flo, 'f', -1, 32), nil
	}

	boolean, err := y.Bool()

	if err == nil {
		if boolean {
			return "true", nil
		}
		return "false", nil
	}

	return "nil", nil
}

// locate returns the YAML node based on a key if it exists.
// Returns a non-nil simpleyaml.Yaml object if found.
// Returns error if there are parse issues with the provided key.
func locate(key string, yaml *simpleyaml.Yaml) (*simpleyaml.Yaml, error) {
	tokens, err := split(key)

	if err != nil {
		return nil, err
	}

	y := yaml

	for _, token := range tokens {
		switch token.Type {
		case TokenIndex:
			if y.IsArray() {
				size, _ := y.GetArraySize()

				if token.Index >= 0 && token.Index < size {
					y = y.GetIndex(token.Index)
					continue
				}
			}
		case TokenKey:
			y = y.Get(token.Key)

			if y.IsFound() {
				continue
			}
		}

		// If no token matches the criterias, as a map or an array with selected index,
		// then consider the key to be not found.
		return nil, nil
	}

	// All tokens matched the criterias and the key was found.
	return y, nil
}
