package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EtcdCluster is the Schema for the etcds API
type EtcdCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdClusterSpec   `json:"spec,omitempty"`
	Status EtcdClusterStatus `json:"status,omitempty"`
}

type EtcdClusterSpec struct {
	// Replicas is the number of members in an etcd cluster.
	// See [etcd-recommendation]: https://etcd.io/docs/v3.5/faq/#why-an-odd-number-of-cluster-members
	// and [etcd-failure-tolerance]: https://etcd.io/docs/v3.5/faq/#what-is-failure-tolerance
	Replicas                    int32              `json:"replicas"`
	MemberTemplate              EtcdMemberTemplate `json:"memberTemplateSpec"`
	Services                    EtcdServices
	EtcdDBCompactionConfig      *EtcdDBCompactionConfig
	SnapshotCompactionJobConfig *SnapshotCompactionJobConfig
}

type EtcdClusterStatus struct {
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
	Mode      EtcdDBCompactionMode
	Retention string
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
}

type EtcdMemberVolumeClaimTemplate struct {
	Name            string            `json:"name"`
	StorageClass    string            `json:"storageClass"`
	StorageCapacity resource.Quantity `json:"storageCapacity"`
}
