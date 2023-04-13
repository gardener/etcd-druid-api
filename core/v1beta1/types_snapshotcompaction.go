package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SnapshotCompactionJobConfig struct {
	Enabled                 *bool
	Resources               *corev1.ResourceRequirements
	EtcdEventsThreshold     *int64
	ActiveDeadlineDuration  *metav1.Duration
	TTLSecondsAfterFinished *int32
}
