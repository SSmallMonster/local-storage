// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/HwameiStor/local-storage/pkg/apis/client/clientset/versioned/typed/uds/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeUdsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeUdsV1alpha1) LocalStorageAlerts() v1alpha1.LocalStorageAlertInterface {
	return &FakeLocalStorageAlerts{c}
}

func (c *FakeUdsV1alpha1) LocalStorageNodes() v1alpha1.LocalStorageNodeInterface {
	return &FakeLocalStorageNodes{c}
}

func (c *FakeUdsV1alpha1) LocalVolumes() v1alpha1.LocalVolumeInterface {
	return &FakeLocalVolumes{c}
}

func (c *FakeUdsV1alpha1) LocalVolumeExpands() v1alpha1.LocalVolumeExpandInterface {
	return &FakeLocalVolumeExpands{c}
}

func (c *FakeUdsV1alpha1) LocalVolumeReplicas() v1alpha1.LocalVolumeReplicaInterface {
	return &FakeLocalVolumeReplicas{c}
}

func (c *FakeUdsV1alpha1) PhysicalDisks() v1alpha1.PhysicalDiskInterface {
	return &FakePhysicalDisks{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeUdsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
