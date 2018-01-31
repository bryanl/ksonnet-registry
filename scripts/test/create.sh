#!/bin/bash

# Upload a part to the registry service. By default, it
# will construct the part.

set -e
set -x

OPTS=$(getopt --long host:,namespace:,package:,part:,part-dir: -n 'parse-options' -- "$@")

HOST="http://localhost:9000"
NAMESPACE="ksonnet"
PACKAGE="deployed-service"
PART_DIR="data/deployed-service"

while true; do
  case "$1" in
    --host ) HOST="$2"; shift 2; ;;
    --namespace ) NAMESPACE="$2"; shift 2; ;;
    --package ) PACKAGE="$2"; shift 2; ;;
    --part ) PART="$2"; shift 2; ;;
    --part-dir ) PART_DIR="$2"; shift 2; ;;
    * ) break ;;
  esac
done

if [ -z "${PART}" ]; then
  name=$(basename $0)
  dir=$(mktemp -d /tmp/${name}.$$)
  if [ $? -ne 0 ]; then
    echo "$0: can't create temp directory, exiting..."
    exit 1
  fi

  tar cvzf ${dir}/part.tar.gz -C ${PART_DIR} .
  PART="${dir}/part.tar.gz"
fi

BLOB=$(base64 -i ${PART})

curl -vsSL -X POST \
  "${HOST}/api/v1/packages/${NAMESPACE}/${PACKAGE}" \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d "{
  \"blob\": \"${BLOB}\",
  \"release\": \"1.1\"
}
"
