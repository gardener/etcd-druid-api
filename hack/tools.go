// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:build tools
// +build tools

package tools

import (
	_ "github.com/ahmetb/gen-crd-api-reference-docs"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
