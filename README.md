# JQ in Go

A command line JSON processor in Go. Trying to replicate [stedolan/jq](https://github.com/stedolan/jq) in Go and learning along.

[![CircleCI](https://circleci.com/gh/amitkr0201/jq-in-go.svg?style=svg)](https://circleci.com/gh/amitkr0201/jq-in-go) [![Build Status](https://travis-ci.com/amitkr0201/jq-in-go.svg?branch=master)](https://travis-ci.com/amitkr0201/jq-in-go) [![codecov](https://codecov.io/gh/amitkr0201/jq-in-go/branch/master/graph/badge.svg)](https://codecov.io/gh/amitkr0201/jq-in-go) [![BCH compliance](https://bettercodehub.com/edge/badge/amitkr0201/jq-in-go?branch=master)](https://bettercodehub.com/) [![CodeFactor](https://www.codefactor.io/repository/github/amitkr0201/jq-in-go/badge)](https://www.codefactor.io/repository/github/amitkr0201/jq-in-go) [![Go Report Card](https://goreportcard.com/badge/github.com/amitkr0201/jq-in-go)](https://goreportcard.com/report/github.com/amitkr0201/jq-in-go)

## Usage

```sh
  jq-in-go '<selector>' <input-json-file> [flags]
```

Examples:

``` sh
jq-in-go -c '.' ~/input.json
```

Flags:

```sh
  -c, --compact   Print JSON in compact
  -h, --help      help for jq-in-go
      --version   version for jq-in-go
```

## Build

```sh
go get -u github.com/spf13/cobra/cobra
go build -o jqingo main.go
```
