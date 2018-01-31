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

HOST=${1:-localhost:9000}
ROOT=$(cd "$(dirname "$0")/../.."; pwd)
PART="${ROOT}/store/testdata/node.tar.gz"

BLOB=$(base64 -i ${PART})

curl -sSL -X POST \
  ${HOST}/api/v1/packages/${NAMESPACE}/${PACKAGE} \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d "{
  \"blob\": \"${BLOB}\",
  \"release\": \"1.1\"
}
"
