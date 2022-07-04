[![Go Reference](https://pkg.go.dev/badge/github.com/rrobrms/w3w.svg)](https://pkg.go.dev/github.com/rrobrms/w3w)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrobrms/w3w)](https://goreportcard.com/report/github.com/rrobrms/w3w)
[![Coverage Status](https://coveralls.io/repos/github/rrobrms/w3w/badge.svg?branch=main)](https://coveralls.io/github/rrobrms/w3w?branch=main)
<!-- [![Latest Release](https://img.shields.io/github/v/release/rrobrms/w3w)](https://github.com/rrobrms/w3w/releases) -->

# W3W - eth wallet CLI

> Simple CLI to generate ether wallets from `go-ethereum`

## Install

#### Packages
Direct downloads are available through the [releases page](https://github.com/rrobrms/w3w/releases/latest).

#### GO
> `go install github.com/rrobrms/w3w@latest`

#### Manual

> You'll need to [install Go](https://golang.org/doc/install)
>```sh
> git clone https://github.com/rrobrms/w3w.git
> cd w3w
> go install
>```

## usage

```sh
$ w3w
$ 23ed07206ce891e433a0193879caa12b6f50b69qbe374ecrd5fed1c7658a90b6 0xC3c7902A11f79184D5138B047FD9a993633b642d
```

```sh
$ w3w -v
$ ╭─────────────┬──────────────────────────────────────────────────────────────────╮
$ │ private key │ 23ed07206ce891e433a0193879caa12b6f50b69qbe374ecrd5fed1c7658a90b6 │
$ ├─────────────├──────────────────────────────────────────────────────────────────┤
$ │ address     │ 0xC3c7902A11f79184D5138B047FD9a993633b642d                       │
$ ╰─────────────┴──────────────────────────────────────────────────────────────────╯
$   valid   https://etherscan.io/address/0xC3c7902A11f79184D5138B047FD9a993633b642d
```