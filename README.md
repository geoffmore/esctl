# esctl-go

A command line client built on [go-elasticsearch](https://github.com/elastic/go-elasticsearch) for interacting with elasticsearch clusters.

[![GoDoc](https://godoc.org/github.com/geoffmore/esctl-go?status.svg)](http://godoc.org/github.com/geoffmore/esctl-go)
[![CircleCI](https://circleci.com/gh/geoffmore/esctl-go.svg?style=svg)](https://circleci.com/gh/geoffmore/esctl-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/geoffmore/esctl-go)](https://goreportcard.com/report/github.com/geoffmore/esctl-go)

## Usage
For each of these commands talks to a specific elasticsearch endpoint. Rather
than explaining each command, the appropriate endpoint will instead be shown.

[This link](https://github.com/elastic/elasticsearch/tree/master/rest-api-spec/src/main/resources/rest-api-spec/api) may be a valuable resource in seeing which api endpoints are avaiable.

| Command | Endpoint |
| --- | --- |
| esctl get cluster-info | / |
| esctl get cluster-health | /_cluster/health |

## TODO
See [this](TODO.md)
