#!/usr/bin/env bash

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/../"
OUT_D=${OUT_D:-${DIR}/builds}
CMD_D=${CMD_D:-"${DIR}/cmd/cftsmt"}
mkdir -p "$OUT_D"

re='^[0-9]+([.][0-9]+)?$'
if ! [[ $GOARM =~ $re ]] ; then
  (GOOS=${GOOS:-linux} GOARCH=${GOARCH:-amd64} CGO_ENABLED=0 go build -ldflags "-extldflags '-static'" -o "${OUT_D}/cftsmt_${GOOS}_${GOARCH}" ${CMD_D})
fi
