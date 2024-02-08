// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:generate gen-crd-api-reference-docs -api-dir . -config ../../hack/api-reference/druid-config.json -template-dir ../../hack/api-reference/template -out-file ../../docs/api-reference/druid.md

// Package v1alpha1 is the v1alpha1 version of the etcd-druid API.
// +groupName=druid.gardener.cloud
package v1alpha1 // import "github.com/gardener/etcd-druid-api/core/v1alpha1"
