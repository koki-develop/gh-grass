<p align="center">
<img src="./assets/logo-light.svg#gh-light-mode-only" />
<img src="./assets/logo-dark.svg#gh-dark-mode-only" />
</p>

<p align="center">
<a href="https://github.com/koki-develop/gh-grass/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/gh-grass" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/gh-grass/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/gh-grass/ci.yml?logo=github" alt="GitHub Workflow Status"></a>
<a href="https://codeclimate.com/github/koki-develop/gh-grass/maintainability"><img src="https://img.shields.io/codeclimate/maintainability/koki-develop/gh-grass?style=flat&amp;logo=codeclimate" alt="Maintainability"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/gh-grass"><img src="https://goreportcard.com/badge/github.com/koki-develop/gh-grass" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/gh-grass" alt="LICENSE"></a>
</p>

<p align="center">
Grow github grass to console.
</p>

<p align="center">
<img src="./assets/demo.gif" alt="demo">
</p>

# gh-grass

- [Installation](#installation)
- [Usage](#usage)
- [LICENSE](#license)

## Installation

```sh
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
  -g, --grass string   grass string (default "■")
      --total          print total contributions
  -h, --help           help for gh
```

### Basic

```sh
$ gh grass
```

![demo](./assets/demo.gif)

### Specify a User

```sh
$ gh grass --user <USERNAME>
# e.g.
$ gh grass --user koki-develop
```

![](./assets/user.gif)

### Change Theme

```sh
$ gh grass --theme <dark or light>
# e.g.
$ gh grass --theme light
```

![](./assets/theme.gif)

### Print Total Contributions

```sh
$ gh grass --total
```

![](./assets/total.gif)

### Custom Grass

```sh
$ gh grass --grass "<GRASS STRING>"
# e.g.
$ gh grass --grass "●"
```

![](./assets/grass.gif)

## LICENSE

[MIT](./LICENSE)
