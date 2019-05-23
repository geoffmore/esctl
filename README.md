# Esctl

The github.com/geoffmore/esctl` repo is intended to provide the ability to
interact with elasticsearch clusters in a similar way that kubectl interacts
with kubernetes clusters.

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

