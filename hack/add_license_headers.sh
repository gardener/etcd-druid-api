#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0



set -e

echo "> Adding SPDX License header to all go files where it is not present"

# Uses the tool https://github.com/google/addlicense
addlicense -v -f hack/license_boilerplate.txt -ignore "vendor/**" -ignore "**/*.md" -ignore "**/*.yaml" -ignore "**/Dockerfile" .
