#!/bin/bash

set -e

OPTS=$(getopt --long host:,namespace:,package:,:release -n 'parse-options' -- "$@")

HOST="http://localhost:9000"
NAMESPACE="ksonnet"
PACKAGE="deployed-service"
RELEASE="1.1.0"

while true; do
  case "$1" in
    --host ) HOST="$2"; shift 2; ;;
    --namespace ) NAMESPACE="$2"; shift 2; ;;
    --package ) PACKAGE="$2"; shift 2; ;;
    --release ) PACKAGE="$2"; shift 2; ;;
    * ) break ;;
  esac
done

curl -sSL -XGET \
  ${HOST}/api/v1/packages/${NAMESPACE}/${PACKAGE}/${RELEASE}