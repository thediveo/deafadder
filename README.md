# deafadder

[![PkgGoDev](https://img.shields.io/badge/-reference-blue?logo=go&logoColor=white&labelColor=505050)](https://pkg.go.dev/github.com/thediveo/deafadder)
[![GitHub](https://img.shields.io/github/license/thediveo/deafadder)](https://img.shields.io/github/license/thediveo/deafadder)
![build and test](https://github.com/thediveo/deafadder/actions/workflows/buildandtest.yaml/badge.svg?branch=master)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/deafadder)](https://goreportcard.com/report/github.com/thediveo/deafadder)

Package `deafadder` wraps [Koanf](https://pkg.go.dev/github.com/knadh/koanf/v2)
configuration data objects, adding convenience accessor functions for
[pflag](https://pkg.go.dev/github.com/spf13/pflag) configuration data types.

Please refer to the [module 
documentation](https://pkg.go.dev/github.com/thediveo/deafadder) for usage and
details.

## Tinkering

When tinkering with the `deafadder` source code base, the recommended way is a
devcontainer environment. The devcontainer specified in this repository
contains:

- `gocover` command to run all tests with coverage, updating the README coverage
  badge automatically after successful runs.
- Go package documentation is served in the background on port TCP/HTTP `6060`
  of the devcontainer.
- [`go-mod-upgrade`](https://github.com/oligot/go-mod-upgrade)
- [`goreportcard-cli`](https://github.com/gojp/goreportcard).
- [`pin-github-action`](https://github.com/mheap/pin-github-action) for
  maintaining Github Actions.

## Copyright and License

`deafadder` is Copyright 2025 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
