# esctl

[![GoDoc](https://godoc.org/github.com/geoffmore/esctl?status.svg)](http://godoc.org/github.com/geoffmore/esctl)
[![CircleCI](https://circleci.com/gh/geoffmore/esctl.svg?style=svg)](https://circleci.com/gh/geoffmore/esctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/geoffmore/esctl)](https://goreportcard.com/report/github.com/geoffmore/esctl)

----

esctl is a command line client built on [go-elasticsearch](https://github.com/elastic/go-elasticsearch) for interacting with elasticsearch clusters.

## Getting started
### Install
Run one of 
* `./build/make-build.sh` to build only core features
* `./build/make-buildext.sh` to build core features and extensions
* `./build/make-build-opt.sh` to build core features and all optional features
  (including extensions)
You will end up with a binary `esctl` that can then be put into your $PATH
### Initializing
Run `esctl config generate` to generate a config file at `~/.elastic/config` with a similar structure to
[kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig) 

## Usage
Every command has a list of its subcommands and usually correspond to an elastic
endpoint. For example, the `cat` command corresponds to `/_cat/` and `cat
pending-tasks` corresponds to `/_cat/pending_tasks`

### Available commands
#### Core
```
  cat         Endpoints under /_cat
  config      Interact with a config file
  get         Get a resource
  help        Help about any command
  version     Print the version number of esctl client/server
```
#### Extensions
```
  kibana (planned)  Kibana endpoints
  admin (planned)   Useful commands for elasticsearch administrators
```
## Contributing
In the case of contributing to the core tool or existing extensions, you should
raise an issue,
fork the repo, 
and submit a PR referencing the issue

In the case of extending the tool, you should
create an extension by doing the following:
Run `./build/gen-extension.sh \<your-extension/package name\>`
Build!

## List of current and future features
See [this](TODO.md)

[This link](https://github.com/elastic/elasticsearch/tree/master/rest-api-spec/src/main/resources/rest-api-spec/api) may be a valuable resource in seeing which api endpoints are avaiable.
