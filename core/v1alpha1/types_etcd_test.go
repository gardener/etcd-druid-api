package v1alpha1

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
)

const (
	defaultNs       = "default"
	defaultEtcdName = "etcd-test"
)

func TestGetPeerServiceName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	g.Expect(e.GetPeerServiceName()).To(Equal(fmt.Sprintf("%s-peer", defaultEtcdName)))
}

func TestGetClientServiceName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	g.Expect(e.GetClientServiceName()).To(Equal(fmt.Sprintf("%s-client", defaultEtcdName)))
}

func TestGetServiceAccountName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	g.Expect(e.GetServiceAccountName()).To(Equal(defaultEtcdName))
}

func TestGetConfigMapName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	expectedConfigMapName := fmt.Sprintf("etcd-bootstrap-%s", string(e.UID[:6]))
	g.Expect(e.GetConfigmapName()).To(Equal(expectedConfigMapName))
}

func TestGetCompactionJobName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	expectedCompactionJobName := fmt.Sprintf("%s-compact-job", string(e.UID[:6]))
	g.Expect(e.GetCompactionJobName()).To(Equal(expectedCompactionJobName))
}

func TestGetOrdinalPodName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	const ordinal = 1
	g.Expect(e.GetOrdinalPodName(ordinal)).To(Equal(fmt.Sprintf("%s-%d", defaultEtcdName, ordinal)))
}

func TestDeltaSnapshotLeaseName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	g.Expect(e.GetDeltaSnapshotLeaseName()).To(Equal(fmt.Sprintf("%s-delta-snap", defaultEtcdName)))
}

func TestFullSnapshotLeaseName(t *testing.T) {
	g := NewWithT(t)
	e := createEtcd(defaultEtcdName)
	g.Expect(e.GetFullSnapshotLeaseName()).To(Equal(fmt.Sprintf("%s-full-snap", defaultEtcdName)))
}

func createEtcd(name string) *Etcd {
	return &Etcd{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: defaultNs,
			UID:       uuid.NewUUID(),
		},
		Spec: EtcdSpec{
			Etcd:                  EtcdConfig{},
			Backup:                BackupSpec{},
			Common:                SharedConfig{},
			SchedulingConstraints: SchedulingConstraints{},
			Replicas:              0,
		},
	}
}
