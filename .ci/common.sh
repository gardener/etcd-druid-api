#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0



# For the test step concourse will set the following environment variables:
# SOURCE_PATH - path to component repository root directory.
if [[ $(uname) == 'Darwin' ]]; then
  READLINK_BIN="greadlink"
else
  READLINK_BIN="readlink"
fi

if [[ -z "${SOURCE_PATH}" ]]; then
  export SOURCE_PATH="$(${READLINK_BIN} -f "$(dirname ${0})/..")"
else
  export SOURCE_PATH="$(${READLINK_BIN} -f "${SOURCE_PATH}")"
fi

VCS="github.com"
ORGANIZATION="gardener"
PROJECT="etcd-druid-api"
REPOSITORY=${VCS}/${ORGANIZATION}/${PROJECT}

# The `go <cmd>` commands requires to see the target repository to be part of a
# Go workspace. Thus, if we are not yet in a Go workspace, let's create one
# temporarily by using symbolic links.
if [[ "${SOURCE_PATH}" != *"src/${REPOSITORY}" ]]; then
  SOURCE_SYMLINK_PATH="${SOURCE_PATH}/tmp/src/${REPOSITORY}"
  if [[ -d "${SOURCE_PATH}/tmp" ]]; then
    rm -rf "${SOURCE_PATH}/tmp"
  fi
  mkdir -p "${SOURCE_PATH}/tmp/src/${VCS}/${ORGANIZATION}"
  ln -s "${SOURCE_PATH}" "${SOURCE_SYMLINK_PATH}"
  cd "${SOURCE_SYMLINK_PATH}"

  export GOPATH="${SOURCE_PATH}/tmp"
  export GOBIN="${SOURCE_PATH}/tmp/bin"
  export PATH="${GOBIN}:${PATH}"
  export GO111MODULE=on
fi