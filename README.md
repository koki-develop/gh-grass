# gh-grass

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/gh-grass)](https://github.com/koki-develop/gh-grass/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/gh-grass/ci.yml?logo=github)](https://github.com/koki-develop/gh-grass/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/gh-grass?style=flat&logo=codeclimate)](https://codeclimate.com/github/koki-develop/gh-grass/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/gh-grass)](https://goreportcard.com/report/github.com/koki-develop/gh-grass)
[![LICENSE](https://img.shields.io/github/license/koki-develop/gh-grass)](./LICENSE)

Grow github grass to console.

![demo](./docs/demo.gif)

- [Installation](#installation)
- [Usage](#usage)
- [LICENSE](#license)

## Installation

```console
$ gh extension install koki-develop/gh-grass
```

## Usage

```console
$ gh grass --help
Grow github grass to console.

Usage:
  gh grass [flags]

Flags:
  -u, --user string    github username
  -t, --theme string   grass theme (dark|light) (default "dark")
      --total          print total contributions
  -h, --help           help for gh
```

### Basic

```console
$ gh grass
```

![demo](./docs/demo.gif)

### Specify a User

You can specify a user with the `-u` or `--user` flag.

```console
$ gh grass --user <USERNAME>
```

![](./docs/user.gif)

## LICENSE

[MIT](./LICENSE)
