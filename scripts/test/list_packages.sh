#!/bin/bash

set -e

OPTS=$(getopt --long host:,namespace:,package: -n 'parse-options' -- "$@")

HOST="http://localhost:9000"
# NAMESPACE="ksonnet"

while true; do
  case "$1" in
    --host ) HOST="$2"; shift; shift; ;;
    # --namespace ) NAMESPACE="$2"; shift; shift; ;;
    * ) break ;;
  esac
done

curl -sSL -XGET \
  ${HOST}/api/v1/packages