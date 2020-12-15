# Assert

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/philzon/helm-assert/CI?style=flat)](https://github.com/philzon/helm-assert/actions?query=workflow%3ACI)
[![Fossa report](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/releases)
[![GitHub](https://img.shields.io/github/license/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/blob/master/LICENSE)

**assert** is a plugin for [Helm](https://github.com/helm/helm).

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Install

Pre-built binaries for various systems can be found in [Releases](https://github.com/philzon/helm-assert/releases) section.

### Helm

The plugin can be installed using Helm's plugin command:

```txt
$ helm plugin install https://github.com/philzon/helm-assert.git
```

Using Helm to update the plugin will always fetch the latest version published:

```txt
$ helm plugin update assert
```

### Standalone

It is possible to install as a standalone tool using `make` without requiring the use of Helm.
This requires that the source has been built.

The installation path is set to be installed in `/usr/local/bin` by default.
You can invoke `make INSDIR="/new/install/path"` to override its path.

```txt
# make install
```

To uninstall the binary:

```txt
# make uninstall
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
