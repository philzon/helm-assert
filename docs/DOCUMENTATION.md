# Documentation

**INCOMPLETE**

## Table of Contents

- [Root](#root)
- [Test](#test)
- [Assert](#assert)

## Root

**Config**

- `sets` - global, override values in the chart when rendering (string array)
- `values` - global, override values in the chart from values files when rendering (string array)
- `tests` - list of tests, see [Test](#test) section (object array)

```yaml
sets:
  - foo.bar=baz
values:
  - test-values.yaml
tests:
  - ...
```

**Report**

The generated report will be in JSON format.

- `date` - test execution date in RFC3339 format (string)
- `chart` - chart metadata (object)
- `chart.path` - chart location (string)
- `chart.name` - chart name (string)
- `chart.version` - chart version (string)
- `tests` - individual test case reports, see [Test](#test) section (object array)

```json
{
    "date": "1970-01-01T00:00:00Z00:00",
    "chart": {
        "path": "./test/chart-example",
        "name": "chart-example",
        "version": "0.1.0"
    },
    "score": {
        "total": 1,
        "passed": 1,
        "failed": 0,
        "skipped": 0
    },
    "tests": [
        {...}
    ]
}
```

## Test

YAML or template rendering issues are considered to be fatal to the test.
The application will stop and print the issue on screen and exit in a failure state.
No report will be generated.

**Config**

- `name` - test case name (string)
- `summary` - test case description (string)
- `sets` - override values in the chart when rendering (string array)
- `values` - override values in the chart from values files when rendering (string array)
- `select` - select manifests through different methods (object)
- `select.files` - select manifests by file name (string array)
- `select.kinds` - select manifests by resource `kind` (string array)
- `select.versions` - select manifests by resource `apiVersion` (string array)
- `asserts` - list of asserts, see [Assert](#assert) section (object array)

```yaml
- name: TC_001
  summary: Test case description
  sets:
    - foo.bar=baz
  values:
    - test-values.yaml
  select:
    files:
      - deployment.yaml
    kinds:
      - Deployment
    versions:
      - apps/v1
  asserts:
    - ...
```

**Report**

- `name` - test case name (string)
- `summary` - test case description (string)
- `passed` - if the test case passed (boolean)
- `skipped` - if the test case was skipped (boolean)
- `score` - results from test case assert(s) (object)
- `score.total` - number of asserts (integer)
- `score.passed` - number of asserts that passed (integer)
- `score.failed` - number of asserts that failed (integer)
- `score.skipped` - number of asserts that were skipped (integer)
- `manifests` - manifests selected for the test case (object array)
- `manifests` - manifests full file name in the chart (string)
- `manifests` - manifests YAML data (string)
- `asserts` - assert reports, see [Assert](#assert) section (object array)

```json
{
    "name": "TC_001",
    "summary": "Test case description",
    "passed": true,
    "skipped": false,
    "score": {
        "total": 1,
        "passed": 1,
        "failed": 0,
        "skipped": 0
    },
    "manifests": [
        {
            "path": "chart-example/templates/deployment.yaml",
            "data": "..."
        }
    ],
    "asserts": [
        {...}
    ]
}
```

## Assert

If the test case is skipped, all asserts are marked as skipped in the score report.

The supported method to access keys are written in dot notation:

- `foo.bar.baz` - access key from maps
- `foo.bar[2].baz` - access from array elements

**Report**

All reports generated from asserts will be in the same structure.

- `index` - which assert from the array in the test case configuration (integer)
- `passed` - if the assert passed (boolean)
- `output` - the generated output from the assert when executed (string)
- `manifest` - the manifest the assert was checked on (string)

```json
{
    "asserts": [
        {
            "index": 0,
            "passed": true,
            "output": "got 'Deployment' from key 'kind'",
            "manifest": "chart-example/templates/deployment.yaml"
        }
    ]
}
```

### Exist

Returns `true` if the defined YAML key exists, otherwise `false`.

- `key` - which key to check using dot notation (string)

**Config**

```yaml
- exist:
    key: spec.template.spec.containers[0].name
```

### Equal

Exist returns `true` if the selected YAML key's value matches the given value, otherwise `false`.

- `key` - which key to retrieve the value from using dot notation (string)
- `value` - the value to compare (string)

If the `key` is pointing to an object or an array, it will return the following values:

- `[...]` - arrays
- `{...}` - objects
- `null` - null

If a key does not exist, it returns an empty string `''`.

**Config**

```yaml
- equal:
    key: spec.template.spec.containers[0].name
    value: example-container
```
