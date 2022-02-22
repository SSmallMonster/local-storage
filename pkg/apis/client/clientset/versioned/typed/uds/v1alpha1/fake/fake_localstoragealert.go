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

// FakeLocalStorageAlerts implements LocalStorageAlertInterface
type FakeLocalStorageAlerts struct {
	Fake *FakeUdsV1alpha1
}

var localstoragealertsResource = schema.GroupVersionResource{Group: "localstorage.hwameistor.io", Version: "v1alpha1", Resource: "localstoragealerts"}

var localstoragealertsKind = schema.GroupVersionKind{Group: "localstorage.hwameistor.io", Version: "v1alpha1", Kind: "LocalStorageAlert"}

// Get takes name of the localStorageAlert, and returns the corresponding localStorageAlert object, and an error if there is any.
func (c *FakeLocalStorageAlerts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocalStorageAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localstoragealertsResource, name), &v1alpha1.LocalStorageAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalStorageAlert), err
}

// List takes label and field selectors, and returns the list of LocalStorageAlerts that match those selectors.
func (c *FakeLocalStorageAlerts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocalStorageAlertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localstoragealertsResource, localstoragealertsKind, opts), &v1alpha1.LocalStorageAlertList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocalStorageAlertList{ListMeta: obj.(*v1alpha1.LocalStorageAlertList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocalStorageAlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localStorageAlerts.
func (c *FakeLocalStorageAlerts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localstoragealertsResource, opts))
}

// Create takes the representation of a localStorageAlert and creates it.  Returns the server's representation of the localStorageAlert, and an error, if there is any.
func (c *FakeLocalStorageAlerts) Create(ctx context.Context, localStorageAlert *v1alpha1.LocalStorageAlert, opts v1.CreateOptions) (result *v1alpha1.LocalStorageAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localstoragealertsResource, localStorageAlert), &v1alpha1.LocalStorageAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalStorageAlert), err
}

// Update takes the representation of a localStorageAlert and updates it. Returns the server's representation of the localStorageAlert, and an error, if there is any.
func (c *FakeLocalStorageAlerts) Update(ctx context.Context, localStorageAlert *v1alpha1.LocalStorageAlert, opts v1.UpdateOptions) (result *v1alpha1.LocalStorageAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localstoragealertsResource, localStorageAlert), &v1alpha1.LocalStorageAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalStorageAlert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalStorageAlerts) UpdateStatus(ctx context.Context, localStorageAlert *v1alpha1.LocalStorageAlert, opts v1.UpdateOptions) (*v1alpha1.LocalStorageAlert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localstoragealertsResource, "status", localStorageAlert), &v1alpha1.LocalStorageAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalStorageAlert), err
}

// Delete takes name of the localStorageAlert and deletes it. Returns an error if one occurs.
func (c *FakeLocalStorageAlerts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localstoragealertsResource, name), &v1alpha1.LocalStorageAlert{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalStorageAlerts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localstoragealertsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocalStorageAlertList{})
	return err
}

// Patch applies the patch and returns the patched localStorageAlert.
func (c *FakeLocalStorageAlerts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocalStorageAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localstoragealertsResource, name, pt, data, subresources...), &v1alpha1.LocalStorageAlert{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocalStorageAlert), err
}
