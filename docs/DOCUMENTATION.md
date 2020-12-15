# Documentation

## Table of Contents

- [Configuration](#configuration)

## Configuration

This is a full example of the configuration and its usage:

```yaml
# Provide global level sets which are applied on all tests.
sets:
  - "global.foo=baz"

# Provide global level sets from files applied on all tests.
values:
  - "global-values.yaml"

# Specify all test cases.
tests:
    # Test case name.
  - name: "TC_001"

    # Summary of what the test case does.
    summary: "Test example"

    # Test case level sets.
    sets:
      - "foo.bar=baz"

    # Test case level sets from files.
    values:
      - tc-001-values.yaml

    select:
      # Select manifests by filename (unamed manifests will not be included if used).
      files:
        - "deployment.yaml"
      # Select manifests based on its kind.
      kinds:
        - Deployment
      # Select manifests based on its API version.
      version:
        - apps/v1

    asserts:
      # Test if key exists.
      - exist:
          key: "foo.bar[0].baz"
      # Test if key has value of "123".
      - equal:
          key: "foo.bar[0].baz"
          value: "123"
```
