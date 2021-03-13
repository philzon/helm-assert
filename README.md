# Assert

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/philzon/helm-assert/CI?style=flat)](https://github.com/philzon/helm-assert/actions?query=workflow%3ACI)
[![Fossa report](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fphilzon%2Fhelm-assert)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/releases)
[![GitHub](https://img.shields.io/github/license/philzon/helm-assert?style=flat)](https://github.com/philzon/helm-assert/blob/master/LICENSE)

**assert** is a plugin for [Helm](https://github.com/helm/helm) (or a standalone tool) to verify the Helm chart's rendered manifests.

## Features

- YAML-based test configurations.
- Render Helm charts from repositories or locally from the file system.
- Override values per test case (or globally) using sets from a list, or files, when rendering.
- Sort manifests by files, Kubernetes resource kinds or API versions for test cases.
- Verify the rendered manifest's YAML by checking for key's values or if they exists using asserts.
- Write test report to file as JSON to be processed.

## Table of Contents

- [Install](#install)
- [Build](#build)
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
This requires that the source has been built, see [Build](#build) section.

The installation path is set to be installed in `/usr/local/bin` by default.
You can invoke `make INSDIR="/new/install/path" ...` to override its path.

```txt
# make install
```

To uninstall the binary:

```txt
# make uninstall
```

## Build

This project is using Golang to both build the project and manage dependencies using Go modules.

To build the source using `make`:

```txt
$ make clean all
```

Built artifacts will be placed in the `bin/` directory.

To build for different systems, the following targets are available:

- `build-linux-amd64` x86 64-bit GNU/Linux systems (most).
- `build-linux-arm64` ARM based 64-bit GNU/Linux systems (most).
- `build-windows-amd64` x86 64-bit Windows systems.
- `build-darwin-amd64` x86 64-bit OSX systems.

Using default target `build` will always build based on the current system.

## Usage

Without providing arguments, or adding flag `-h, --help`, will output its usage:

```txt
Usage:
  assert [CONFIG] [CHART] [flags]

Flags:
  -h, --help               help for assert
      --json string        write report to a file in JSON format
  -l, --log-level string   severity level to log ("verbose"|"standard"|"quiet"|"none") (default "standard")
      --password string    chart repository password where to locate the requested chart
      --repo string        chart repository url where to locate the requested chart
      --skip stringArray   skip test by name (can specify multiple)
      --username string    chart repository username where to locate the requested chart
      --version string     specify the exact chart version to use. If this is not specified, the latest version is used
```

## Documentation

For a detailed overview of the tool, see [docs/DOCUMENTATION.md](./docs/DOCUMENTATION.md) page.

### Quick Start

Test configurations are written in YAML.
Each test case defines what values to override, which manifests to select, and then, which keys to check and how.

Below configuration tests if manifests, with resource kind `Deployment`, has its **first** container's image changed based on the values overriden using `sets`.

```yaml
tests:
  - name: TC_001
    summary: Test if image is being set
    sets:
      - image.repository=nginx
      - image.tag=latest
    select:
      kinds:
        - Deployment
    asserts:
      - equal:
          key: spec.template.spec.containers[0].image
          value: nginx:latest
```

## Contributing

Feel free to contribute in the form of bug reports or feature requests.
Use Github's [issue](https://github.com/philzon/helm-assert/issues) system to create and detail them.

Since the project's design is not finalized, or until a stable, major release has been performed, anything added could be changed or removed later on without further notice.

## License

This project licensed under Apache License 2.0.
See [LICENSE](LICENSE) for the full license details.
