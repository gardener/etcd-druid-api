package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackupRestore struct {
	Container *Container
	// TLS
	TLS *TLSConfig `json:"tls,omitempty"`

	Snapshots *Snapshots

	// EtcdLeaderMonitoringConfig contains configuration which backup-restore uses to track leadership changes for its peer etcd container.
	EtcdLeaderMonitoringConfig *EtcdLeaderMonitoring `json:"etcdLeaderMonitoringConfig,omitempty"`
}

// CompressionAlgorithm is the algorithm used for compression of snapshots.
type CompressionAlgorithm string

const (
	GZipCompression CompressionAlgorithm = "gzip"
	LzwCompression  CompressionAlgorithm = "lzw"
	ZlibCompression CompressionAlgorithm = "zlib"
)

type Compression struct {
	Enabled *bool `json:"enabled"`
	// Algorithm determines how the algorithm that will be used for compressing snapshots.
	// Supported algorithms are GZipCompression, LzwCompression and ZlibCompression. If compression is enabled
	// and algorithm has not been set then a default value of GZipCompression will be used.
	Algorithm *CompressionAlgorithm
}

type DeltaSnapshot struct {
	Enabled               *bool
	Interval              *metav1.Duration
	MemoryBufferThreshold *resource.Quantity
}

type FullSnapshot struct {
	Cron    *string
	Timeout *metav1.Duration
}

type Snapshots struct {
	// Enabled indicates whether backup-restore should take snapshots.
	// If it is not set then a default value of `true` will be used. If it is disabled then all other configuration
	// will be ignored.
	Enabled *bool
	// Compression
	Compression *Compression `json:"compression,omitempty"`
	// Store defines the specification of a store for backing up snapshots (delta and full).
	// +optional
	Store *Store
	// FullSnapshot
	FullSnapshot  *FullSnapshot
	DeltaSnapshot *DeltaSnapshot
}

// EtcdLeaderMonitoring contains configuration for backup-restore to determine etcd leadership changes.
// As long as backup-restore is used as a side-car in every etcd pod, this configuration will be required. For setups
// where there is a separate backup-restore pod for the entire etcd cluster, this configuration is not required to be set.
type EtcdLeaderMonitoring struct {
	// PollInterval is the interval at which the backup-restore should poll etcd for its leadership status.
	// +optional
	PollInterval *metav1.Duration `json:"pollInterval,omitempty"`
	// ConnectionTimeout is the timeout for etcd client connection.
	ConnectionTimeout *metav1.Duration `json:"connectionTimeout,omitempty"`
}

// StorageProvider defines the type of store provider for storing backup snapshots.
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

type GarbageCollectionPolicy string

type GarbageCollectionConfig struct {
	Policy   *GarbageCollectionPolicy `json:"policy,omitempty"`
	Interval *metav1.Duration         `json:"interval,omitempty"`
}

// Store defines parameters related to ObjectStore persisting backups
type Store struct {
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
	// GarbageCollection is the configuration to manage garbage collection of stored (delta and full) snapshots in the store.
	GarbageCollection *GarbageCollectionConfig `json:"garbageCollection,omitempty"`
}