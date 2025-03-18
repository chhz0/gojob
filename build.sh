#!/bin/bash

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")
OUTPUT_DIR=${PROJ_ROOT_DIR}/_output
VERSION_PACKAGE=github.com/chhz0/gotasks/pkg/version

if [[ -z "${VERSION}" ]]; then
    VERSION=$(git describe --tags --always --match='v*')
fi

GIT_TREE_STAT="dirt"
is_clean=$(git status --porcelain 2>/dev/null)
if [[ -z "${is_clean}" ]]; then
    GIT_TREE_STAT="clean"
fi

GIT_COMMIT=$(git rev-parse HEAD)

GO_LDFLAGS="-X ${VERSION_PACKAGE}.gitVersion=${VERSION} \
    -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
    -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STAT} \
    -X ${VERSION_PACKAGE}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

go build -v -ldflags "${GO_LDFLAGS}" -o ${OUTPUT_DIR}/tasks -v cmd/tasks/main.go