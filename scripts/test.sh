#!/bin/bash

set -e
set -x

HOST=${1:-localhost:9000}
ROOT=$(cd "$(dirname "$0")/.."; pwd)
PART="${ROOT}/store/testdata/node.tar.gz"

BLOB=$(base64 -i ${PART})

curl -X POST \
  http://${HOST}/api/v1/packages/ns/node \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d "{
  \"blob\": \"${BLOB}\",
  \"release\": \"1.1\"
}
"
