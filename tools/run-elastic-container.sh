#!/bin/bash

# https://bugs.centos.org/view.php?id=16570
podman run -d -p 9200:9200 \
           -p 9300:9300 \
           -e "discovery.type=single-node" \
           -e "xpack.ml.enabled=false" \
           docker.elastic.co/elasticsearch/elasticsearch:7.3.0

