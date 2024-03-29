// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GarbageCollectionPolicy defines the type of policy for snapshot garbage collection.
// +kubebuilder:validation:Enum=Exponential;LimitBased
type GarbageCollectionPolicy string

const (
	// GarbageCollectionPolicyExponential defines the exponential policy for garbage collecting old backups
	GarbageCollectionPolicyExponential GarbageCollectionPolicy = "Exponential"
	// GarbageCollectionPolicyLimitBased defines the limit based policy for garbage collecting old backups
	GarbageCollectionPolicyLimitBased GarbageCollectionPolicy = "LimitBased"
)

// CompressionPolicy defines the type of policy for compression of snapshots.
// +kubebuilder:validation:Enum=gzip;lzw;zlib
type CompressionPolicy string

const (
	// GzipCompression is constant for gzip compression policy.
	GzipCompression CompressionPolicy = "gzip"
	// LzwCompression is constant for lzw compression policy.
	LzwCompression CompressionPolicy = "lzw"
	// ZlibCompression is constant for zlib compression policy.
	ZlibCompression CompressionPolicy = "zlib"

	// DefaultCompression is constant for default compression policy(only if compression is enabled).
	DefaultCompression CompressionPolicy = GzipCompression
	// DefaultCompressionEnabled is constant to define whether to compress the snapshots or not.
	DefaultCompressionEnabled = false
)

// CompressionSpec defines parameters related to compression of Snapshots(full as well as delta).
type CompressionSpec struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// +optional
	Policy *CompressionPolicy `json:"policy,omitempty"`
}

// LeaderElectionSpec defines parameters related to the LeaderElection configuration.
type LeaderElectionSpec struct {
	// ReelectionPeriod defines the Period after which leadership status of corresponding etcd is checked.
	// +optional
	ReelectionPeriod *metav1.Duration `json:"reelectionPeriod,omitempty"`
	// EtcdConnectionTimeout defines the timeout duration for etcd client connection during leader election.
	// +optional
	EtcdConnectionTimeout *metav1.Duration `json:"etcdConnectionTimeout,omitempty"`
}

// BackupSpec defines parameters associated with the full and delta snapshots of etcd.
type BackupSpec struct {
	// Port define the port on which etcd-backup-restore server will be exposed.
	// +optional
	Port *int32 `json:"port,omitempty"`
	// +optional
	TLS *TLSConfig `json:"tls,omitempty"`
	// Image defines the etcd container image and tag
	// +optional
	Image *string `json:"image,omitempty"`
	// Store defines the specification of object store provider for storing backups.
	// +optional
	Store *StoreSpec `json:"store,omitempty"`
	// Resources defines compute Resources required by backup-restore container.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// CompactionResources defines compute Resources required by compaction job.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	CompactionResources *corev1.ResourceRequirements `json:"compactionResources,omitempty"`
	// FullSnapshotSchedule defines the cron standard schedule for full snapshots.
	// +optional
	FullSnapshotSchedule *string `json:"fullSnapshotSchedule,omitempty"`
	// GarbageCollectionPolicy defines the policy for garbage collecting old backups
	// +optional
	GarbageCollectionPolicy *GarbageCollectionPolicy `json:"garbageCollectionPolicy,omitempty"`
	// GarbageCollectionPeriod defines the period for garbage collecting old backups
	// +optional
	GarbageCollectionPeriod *metav1.Duration `json:"garbageCollectionPeriod,omitempty"`
	// DeltaSnapshotPeriod defines the period after which delta snapshots will be taken
	// +optional
	DeltaSnapshotPeriod *metav1.Duration `json:"deltaSnapshotPeriod,omitempty"`
	// DeltaSnapshotMemoryLimit defines the memory limit after which delta snapshots will be taken
	// +optional
	DeltaSnapshotMemoryLimit *resource.Quantity `json:"deltaSnapshotMemoryLimit,omitempty"`
	// SnapshotCompression defines the specification for compression of Snapshots.
	// +optional
	SnapshotCompression *CompressionSpec `json:"compression,omitempty"`
	// EnableProfiling defines if profiling should be enabled for the etcd-backup-restore-sidecar
	// +optional
	EnableProfiling *bool `json:"enableProfiling,omitempty"`
	// EtcdSnapshotTimeout defines the timeout duration for etcd FullSnapshot operation
	// +optional
	EtcdSnapshotTimeout *metav1.Duration `json:"etcdSnapshotTimeout,omitempty"`
	// LeaderElection defines parameters related to the LeaderElection configuration.
	// +optional
	LeaderElection *LeaderElectionSpec `json:"leaderElection,omitempty"`
}
