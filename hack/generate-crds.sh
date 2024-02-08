#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0


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
