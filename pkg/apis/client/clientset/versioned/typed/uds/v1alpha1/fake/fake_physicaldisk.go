// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/HwameiStor/local-storage/pkg/apis/uds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePhysicalDisks implements PhysicalDiskInterface
type FakePhysicalDisks struct {
	Fake *FakeUdsV1alpha1
}

var physicaldisksResource = schema.GroupVersionResource{Group: "localstorage.hwameistor.io", Version: "v1alpha1", Resource: "physicaldisks"}

var physicaldisksKind = schema.GroupVersionKind{Group: "localstorage.hwameistor.io", Version: "v1alpha1", Kind: "PhysicalDisk"}

// Get takes name of the physicalDisk, and returns the corresponding physicalDisk object, and an error if there is any.
func (c *FakePhysicalDisks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PhysicalDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(physicaldisksResource, name), &v1alpha1.PhysicalDisk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PhysicalDisk), err
}

// List takes label and field selectors, and returns the list of PhysicalDisks that match those selectors.
func (c *FakePhysicalDisks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PhysicalDiskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(physicaldisksResource, physicaldisksKind, opts), &v1alpha1.PhysicalDiskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PhysicalDiskList{ListMeta: obj.(*v1alpha1.PhysicalDiskList).ListMeta}
	for _, item := range obj.(*v1alpha1.PhysicalDiskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested physicalDisks.
func (c *FakePhysicalDisks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(physicaldisksResource, opts))
}

// Create takes the representation of a physicalDisk and creates it.  Returns the server's representation of the physicalDisk, and an error, if there is any.
func (c *FakePhysicalDisks) Create(ctx context.Context, physicalDisk *v1alpha1.PhysicalDisk, opts v1.CreateOptions) (result *v1alpha1.PhysicalDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(physicaldisksResource, physicalDisk), &v1alpha1.PhysicalDisk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PhysicalDisk), err
}

// Update takes the representation of a physicalDisk and updates it. Returns the server's representation of the physicalDisk, and an error, if there is any.
func (c *FakePhysicalDisks) Update(ctx context.Context, physicalDisk *v1alpha1.PhysicalDisk, opts v1.UpdateOptions) (result *v1alpha1.PhysicalDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(physicaldisksResource, physicalDisk), &v1alpha1.PhysicalDisk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PhysicalDisk), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePhysicalDisks) UpdateStatus(ctx context.Context, physicalDisk *v1alpha1.PhysicalDisk, opts v1.UpdateOptions) (*v1alpha1.PhysicalDisk, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(physicaldisksResource, "status", physicalDisk), &v1alpha1.PhysicalDisk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PhysicalDisk), err
}

// Delete takes name of the physicalDisk and deletes it. Returns an error if one occurs.
func (c *FakePhysicalDisks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(physicaldisksResource, name), &v1alpha1.PhysicalDisk{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePhysicalDisks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(physicaldisksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PhysicalDiskList{})
	return err
}

// Patch applies the patch and returns the patched physicalDisk.
func (c *FakePhysicalDisks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PhysicalDisk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(physicaldisksResource, name, pt, data, subresources...), &v1alpha1.PhysicalDisk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PhysicalDisk), err
}
