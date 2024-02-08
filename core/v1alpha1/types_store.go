// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import corev1 "k8s.io/api/core/v1"

// StorageProvider defines the type of object store provider for storing backups.
type StorageProvider string

const (
	// StorageProviderAWS_S3 represents an AWS S3 storage provider.
	StorageProviderAWS_S3 StorageProvider = "S3"
	// StorageProviderAliCloud_OSS represents an AliCloud OSS storage provider.
	StorageProviderAliCloud_OSS StorageProvider = "OSS"
	// StorageProviderAzure_ABS represents an Azure ABS storage provider.
	StorageProviderAzure_ABS StorageProvider = "ABS"
	// StorageProviderGCP_GCS represents a GCP GCS storage provider.
	StorageProviderGCP_GCS StorageProvider = "GCS"
	// StorageProviderOpenStack_Swift represents an OpenStack Swift storage provider.
	StorageProviderOpenStack_Swift StorageProvider = "Swift"
	// StorageProviderDell_ECS represents a Dell ECS storage provider
	StorageProviderDell_ECS StorageProvider = "ECS"
	// StorageProviderOpenShift_OCS represents an OpenShift OCS storage provider.
	StorageProviderOpenShift_OCS StorageProvider = "OCS"
	// StorageProviderLocal represents a local storage provider.
	StorageProviderLocal StorageProvider = "Local"
)

// StoreSpec defines parameters related to ObjectStore persisting backups
type StoreSpec struct {
	// Container is the name of the container the backup is stored at.
	// +optional
	Container *string `json:"container,omitempty"`
	// Prefix is the prefix used for the store.
	// +required
	Prefix string `json:"prefix"`
	// Provider is the name of the backup provider.
	// +optional
	Provider *StorageProvider `json:"provider,omitempty"`
	// SecretRef is the reference to the secret which used to connect to the backup store.
	// +optional
	SecretRef *corev1.SecretReference `json:"secretRef,omitempty"`
}
