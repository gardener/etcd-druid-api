package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MetricsLevel string

const (
	Basic     MetricsLevel = "basic"
	Extensive MetricsLevel = "extensive"
)

type Etcd struct {
	Container       *Container
	Quota           *resource.Quantity `json:"quota,omitempty"`
	Defragmentation *Defragmentation
	ClientTLS       *TLSConfig
	ClientPort      *int32
	PeerTLS         *TLSConfig
	PeerPort        *int32
	Metrics         *MetricsLevel `json:"metrics,omitempty"`
}

type Defragmentation struct {
	Cron            *string
	Timeout         *metav1.Duration
	DBSizeThreshold *resource.Quantity
}
