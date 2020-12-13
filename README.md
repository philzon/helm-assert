# Assert

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/philzon/helm-assert/CI?style=flat)](https://github.com/philzon/helm-assert/actions?query=workflow%3ACI)
[![Fossa report](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/releases)
[![GitHub](https://img.shields.io/github/license/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/blob/dev/LICENSE)

**assert** is a plugin for [Helm](https://github.com/helm/helm).

## Table of Contents

- [Installing](#installing)
- [Usage](#usage)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Installing

```txt
$ helm plugin install https://github.com/philzon/helm-assert.git
```

## Usage

```txt
$ helm assert
```

```txt
Usage:
  assert CONFIG CHART [flags]

Flags:
  -l, --log-level string     severity level to log ("verbose"|"standard"|"quiet"|"none") (default "normal")
  -h, --help                 help for assert
      --json                 Report should be saved in JSON format
  -o, --output string        Path to store reports to (default "report")
      --password string      chart repository password where to locate the requested chart
      --repo string          chart repository url where to locate the requested chart
      --skip stringArray     Skip test by name (can specify multiple)
      --username string      chart repository username where to locate the requested chart
      --version string       specify the exact chart version to use. If this is not specified, the latest version is used
```

## Documentation

See [DOCUMENTATION.md](./docs/DOCUMENTATION.md) for documentation.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) on how to contribute.

## License

See [LICENSE](LICENSE) for project license.
