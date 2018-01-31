#!/bin/bash

set -e

OPTS=$(getopt --long host:,namespace:,package: -n 'parse-options' -- "$@")

HOST="http://localhost:9000"
NAMESPACE="ns"
PACKAGE="node"

while true; do
  case "$1" in
    --host ) HOST="$2"; shift; shift; ;;
    --namespace ) NAMESPACE="$2"; shift; shift; ;;
    --package ) PACKAGE="$2"; shift; shift; ;;
    * ) break ;;
  esac
done

DIGEST=$1

if [ -z "${DIGEST}" ]; then
  echo "usage: $0 [options] <digest>"
  exit 1
fi

curl -sSL -XGET -o /tmp/part.tar.gz \
  ${HOST}/api/v1/packages/${NAMESPACE}/${PACKAGE}/blobs/sha256/${DIGEST}