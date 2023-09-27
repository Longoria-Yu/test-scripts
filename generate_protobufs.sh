#!/bin/bash

PROTOGEN_VERSION="v0.2.4"
GOLANG_PROTOBUF_VERSION="v1.3.2"
BRANCH="master"
SERVICE=""
BOOL_FLAGS=()

while [[ $# -gt 0 ]]
do
key="$1"
case $key in
    -b|--branch)
        BRANCH="$2"
        shift
        shift
        ;;
    --service)
        SERVICE="$2"
        shift
        shift
        ;;
    --use-local-image)
        [ "$2" == "true" ] && BOOL_FLAGS+=("--use-local-image")
        shift
        shift
        ;;
    *)
        shift
        ;;
esac
done

SHARED_PROTO_DIR="$HOME/.shared-proto"
rm -rf $SHARED_PROTO_DIR || true
git clone -b ${BRANCH} --depth 1 git@github.com:carousell/shared-proto.git $SHARED_PROTO_DIR

$SHARED_PROTO_DIR/generate_protobufs.sh \
    --protogen-version      "${PROTOGEN_VERSION}" \
    --go-protobuf-version   "${GOLANG_PROTOBUF_VERSION}" \
    --service               "${SERVICE}" \
    ${BOOL_FLAGS[@]}
