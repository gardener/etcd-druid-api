package v1beta1

import corev1 "k8s.io/api/core/v1"

// TLSConfig hold the TLS configuration.
type TLSConfig struct {
	CASecretRef     SecretReference
	ServerSecretRef corev1.SecretReference
	ClientSecretRef corev1.SecretReference
}

// SecretReference defines a reference to a secret.
type SecretReference struct {
	corev1.SecretReference `json:",inline"`
	// DataKey is the name of the key in the data map containing the credentials.
	// +optional
	DataKey *string `json:"dataKey,omitempty"`
}

type EtcdObjectMeta struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Container struct {
	// Image defines the etcd container image and tag in the form of <image>:<tag>
	// In gardener context there are three ways in which druid can determine the image for backup-restore container:
	// 1. Image-Vector-Override (See https://github.com/gardener/gardener/blob/master/docs/deployment/image_vector.md) for additional information.
	// 2. Image that is specified as part of BackupRestore
	// 3. Druid maintains default image for backup-restore container.
	// Druid will retrieve the image by looking at the above places in the order these have been specified.
	// As soon as it finds the image it takes that image location and ignores the rest of the places.
	// +optional
	Image *string `json:"image,omitempty"`
	// Resources defines compute Resources required by backup-restore container.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}
