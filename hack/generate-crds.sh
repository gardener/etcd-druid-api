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

if ! command -v controller-gen &>/dev/null; then
  echo >&2 "controller-gen is not available"
  exit 1
fi

generate_etcd_crds() {
  local output_dir="$(pwd)"
  local package="github.com/gardener/etcd-druid-api/core/v1alpha1"

  local package_path="$(go list -f '{{ .Dir }}' "$package")"
  if [ -z "$package_path" ]; then
    echo "Could not locate the directory for package: $package"
    exit 1
  fi

  # clean all generated crd files
  if ls "$output_dir"/*.yaml >/dev/null 2>&1; then
    rm -f "$output_dir"/*.yaml
  fi

  controller-gen crd paths="$package_path" output:crd:dir="$output_dir" output:stdout
}

echo "> Generating CRDs for etcd and etcd-copy-backup-task..."
generate_etcd_crds
