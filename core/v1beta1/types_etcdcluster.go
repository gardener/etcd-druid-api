package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EtcdCluster is the Schema for the etcds API
type EtcdCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdClusterSpec   `json:"spec,omitempty"`
	Status EtcdClusterStatus `json:"status,omitempty"`
}

// EtcdClusterList is a list of EtcdCluster.
type EtcdClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdCluster `json:"items"`
}

type EtcdClusterSpec struct {
	// Replicas is the number of members in an etcd cluster.
	// See [etcd-recommendation]: https://etcd.io/docs/v3.5/faq/#why-an-odd-number-of-cluster-members
	// and [etcd-failure-tolerance]: https://etcd.io/docs/v3.5/faq/#what-is-failure-tolerance
	Replicas                    int32              `json:"replicas"`
	MemberTemplate              EtcdMemberTemplate `json:"memberTemplateSpec"`
	BackupRestore               BackupRestore
	Services                    EtcdServices
	EtcdDBCompactionConfig      *EtcdDBCompactionConfig
	SnapshotCompactionJobConfig *SnapshotCompactionJobConfig
	SchedulingConstraints       *SchedulingConstraints
}

type EtcdClusterStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

type EtcdServices struct {
	ClientTemplate        EtcdServiceTemplate
	PeerTemplate          EtcdServiceTemplate
	BackupRestoreTemplate EtcdServiceTemplate
}

// EtcdDBCompactionMode defines the auto-compaction-mode: 'periodic' or 'revision'.
// 'periodic' for duration based retention and 'revision' for revision number based retention.
// +kubebuilder:validation:Enum=periodic;revision
type EtcdDBCompactionMode string

const (
	// Periodic is a constant to set auto-compaction-mode 'periodic' for duration based retention.
	Periodic EtcdDBCompactionMode = "periodic"
	// Revision is a constant to set auto-compaction-mode 'revision' for revision number based retention.
	Revision EtcdDBCompactionMode = "revision"
)

type EtcdDBCompactionConfig struct {
	// Mode defines how auto-compaction is done by etcd.
	// See [etcd-auto-compaction]: https://etcd.io/docs/v3.4/op-guide/maintenance/#auto-compaction
	Mode EtcdDBCompactionMode
	// RetentionDuration defines the retention window which is used by etcd when Mode is Periodic.
	RetentionDuration *metav1.Duration
	// RevisionDelta is used by etcd when Mode is Revision to compute the revision till which compaction needs to be done.
	// Revision to compact on = <latest-revision> - RevisionDelta.
	RevisionDelta *int32
}

type EtcdServiceTemplate struct {
	Metadata *EtcdObjectMeta `json:"metadata,omitempty"`
	Selector metav1.LabelSelector
	Ports    []corev1.ServicePort
}

type EtcdMemberTemplate struct {
	Metadata *EtcdObjectMeta      `json:"metadata,omitempty"`
	Selector metav1.LabelSelector `json:"selector"`
	// PriorityClassName is the name of the PriorityClass for each etcd member pod. If a PriorityClass
	// has been created then it is recommended that it has a high priority value as it impacts preemption of etcd pods.
	PriorityClassName   *string                        `json:"priorityClassName,omitempty"`
	VolumeClaimTemplate *EtcdMemberVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
	HeartbeatRenewal    *metav1.Duration
	Etcd                Etcd
}

type EtcdMemberVolumeClaimTemplate struct {
	Name            string            `json:"name"`
	StorageClass    string            `json:"storageClass"`
	StorageCapacity resource.Quantity `json:"storageCapacity"`
}

// SchedulingConstraints defines the different scheduling constraints that must be applied to the
// pod spec in the etcd StatefulSet.
// Currently supported constraints are Affinity and TopologySpreadConstraints.
// Specifying scheduling constraints would allow druid to add these scheduling constraints when it creates the StatefulSet.
// There are other ways in which these constraints can be specified:
// 1. One could define a cluster wide constraint. See https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#cluster-level-default-constraints
// 2. Another way is to use a mutating webhook to inject scheduling constraints to StatefulSet based on several factors like failure tolerance etc. Gardener uses this method at present.
// It is a conscious choice to not use k8s types (e.g. corev1.Affinity) and instead use `runtime.RawExtension`. Following are the reasons:
// 1. Topology Spread Constraints has been changing for some time now. It would become difficult for us to continuously change the CRD definition to reflect the changes to the upstream feature.
// 2. Having a complete spec for Affinity and TopologySpreadConstraints duplicated in CRD makes it extremely-big and hard to maintain.
type SchedulingConstraints struct {
	// Affinity are the simplest way to constraint pods to nodes. See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
	Affinity *runtime.RawExtension
	// TopologySpreadConstraints are a set of constraints that determine how the pods are spread across a cluster of nodes.
	// See https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints
	TopologySpreadConstraints []*runtime.RawExtension
}
