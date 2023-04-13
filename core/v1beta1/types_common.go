package v1beta1

import corev1 "k8s.io/api/core/v1"

// TLSConfig hold the TLS configuration.
type TLSConfig struct {
	ServerSecretRef corev1.SecretReference
	ClientSecretRef corev1.SecretReference
}

type EtcdObjectMeta struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}
