![Logo](docs/images/terrain-logo-600.png)

[![Build Status](https://img.shields.io/travis/com/NickLarsenNZ/terrain.svg)](https://travis-ci.com/NickLarsenNZ/terrain)
[![Coveralls](https://img.shields.io/coveralls/github/NickLarsenNZ/terrain.svg)](https://coveralls.io/github/NickLarsenNZ/terrain)
[![GitHub](https://img.shields.io/github/license/NickLarsenNZ/terrain.svg)](https://github.com/NickLarsenNZ/terrain)

# Terrain
_Scans [Hashicorp Terraform](https://www.terraform.io/) files and lists the resources and datasources used in alphabetical order for documentation purposes._

## Installation
`todo`

## Usage
`todo`

## Features
- [ ] Parses a Terraform Module configuration in a given directory
- [ ] Diff two configurations (by directory, historical Terrain output, or Git commits)
- [ ] Flags to choose what to show (resources, or datasources)
- [ ] Flag to choose the output formats such as JSON, Markdown Table, etc...
- [ ] Simple digraph or mermaid/plantuml output

## Build
- `go build github.com/nicklarsennz/terrain/cmd/terrain`

## References
- [`hashicorp/terraform-config-inspect`](https://github.com/hashicorp/terraform-config-inspect)
- [`urfav/cli`](https://github.com/urfave/cli)
- [`olekukonko/tablewriter`](github.com/olekukonko/tablewriter)
