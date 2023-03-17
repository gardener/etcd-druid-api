#!/usr/bin/env bash
# Copyright 2023 SAP SE or an SAP affiliate company
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


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
