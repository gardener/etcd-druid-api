#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0



set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"

echo "> Generate..."
go generate "$PROJECT_DIR/..."

echo "> Controller Generator..."

if ! command -v controller-gen &>/dev/null; then
  >&2 echo "controller-gen is not available"
  exit 1
fi

cd "$PROJECT_DIR/core" && controller-gen "object:headerFile=$SCRIPT_DIR/license_boilerplate.txt" paths=./...

# needed as long as https://github.com/kubernetes-sigs/controller-tools/issues/559 is not fixed
find "$PROJECT_DIR/core" -type f -name "zz_*.go" -exec goimports -w '{}' \;
