package v1beta1

import corev1 "k8s.io/api/core/v1"

type BackupRestoreSpec struct {
	// Image defines the etcd container image and tag in the form of <image>:<tag>
	// In gardener context there are three ways in which druid can determine the image for backup-restore container:
	// 1. Image-Vector-Override (See https://github.com/gardener/gardener/blob/master/docs/deployment/image_vector.md) for additional information.
	// 2. Image that is specified as part of BackupRestoreSpec
	// 3. Druid maintains default image for backup-restore container.
	// Druid will retrieve the image by looking at the above places in the order these have been specified.
	// As soon as it finds the image it takes that image location and ignores the rest of the places.
	// +optional
	Image *string `json:"image,omitempty"`
	// Resources defines compute Resources required by backup-restore container.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Resources *corev1.ResourceRequirements
	// TLS <TODO>
	TLS *TLSConfig
}

// StorageProvider defines the type of store provider for storing backups.
type StorageProvider string

const (
	// StorageProviderAWSS3 represents an AWS S3 storage provider.
	StorageProviderAWSS3 StorageProvider = "S3"
	// StorageProviderAlicloudOSS represents an AliCloud OSS storage provider.
	StorageProviderAlicloudOSS StorageProvider = "OSS"
	// StorageProviderAzureABS represents an Azure ABS storage provider.
	StorageProviderAzureABS StorageProvider = "ABS"
	// StorageProviderGCPGCS represents a GCP GCS storage provider.
	StorageProviderGCPGCS StorageProvider = "GCS"
	// StorageProviderOpenStackSwift represents an OpenStack Swift storage provider.
	StorageProviderOpenStackSwift StorageProvider = "Swift"
	// StorageProviderDellECS represents a Dell ECS storage provider
	StorageProviderDellECS StorageProvider = "ECS"
	// StorageProviderOpenShiftOCS represents an OpenShift OCS storage provider.
	StorageProviderOpenShiftOCS StorageProvider = "OCS"
	// StorageProviderLocal represents a local storage provider.
	StorageProviderLocal StorageProvider = "Local"
)

// StoreSpec defines parameters related to ObjectStore persisting backups
type StoreSpec struct {
	// Container is the name of the container the backup is stored at.
	// In case of an object store (StorageProvider in [S3, OSS, ABS, GCS, Swift, ECS, OCS]) container should not be nil,
	// however in case of StorageProvider = Local it can be nil.
	// +optional
	Container *string `json:"container,omitempty"`
	// Prefix is the prefix used for the store.
	// +required
	Prefix string `json:"prefix"`
	// Provider is the name of the backup provider.
	// +required
	Provider StorageProvider `json:"provider"`
	// SecretRef is the reference to the secret which used to connect to the backup store.
	// In case of an object store (StorageProvider in [S3, OSS, ABS, GCS, Swift, ECS, OCS]) container should not be nil
	// as it is required to connect to a remote object store. However, in case of StorageProvider = Local it can be nil.
	// +optional
	SecretRef *corev1.SecretReference `json:"secretRef,omitempty"`
}
