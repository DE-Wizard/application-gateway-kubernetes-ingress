/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceEntries implements ServiceEntryInterface
type FakeServiceEntries struct {
	Fake *FakeNetworkingV1alpha3
	ns   string
}

var serviceentriesResource = v1alpha3.SchemeGroupVersion.WithResource("serviceentries")

var serviceentriesKind = v1alpha3.SchemeGroupVersion.WithKind("ServiceEntry")

// Get takes name of the serviceEntry, and returns the corresponding serviceEntry object, and an error if there is any.
func (c *FakeServiceEntries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.ServiceEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceentriesResource, c.ns, name), &v1alpha3.ServiceEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceEntry), err
}

// List takes label and field selectors, and returns the list of ServiceEntries that match those selectors.
func (c *FakeServiceEntries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.ServiceEntryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceentriesResource, serviceentriesKind, c.ns, opts), &v1alpha3.ServiceEntryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.ServiceEntryList{ListMeta: obj.(*v1alpha3.ServiceEntryList).ListMeta}
	for _, item := range obj.(*v1alpha3.ServiceEntryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceEntries.
func (c *FakeServiceEntries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceentriesResource, c.ns, opts))

}

// Create takes the representation of a serviceEntry and creates it.  Returns the server's representation of the serviceEntry, and an error, if there is any.
func (c *FakeServiceEntries) Create(ctx context.Context, serviceEntry *v1alpha3.ServiceEntry, opts v1.CreateOptions) (result *v1alpha3.ServiceEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceentriesResource, c.ns, serviceEntry), &v1alpha3.ServiceEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceEntry), err
}

// Update takes the representation of a serviceEntry and updates it. Returns the server's representation of the serviceEntry, and an error, if there is any.
func (c *FakeServiceEntries) Update(ctx context.Context, serviceEntry *v1alpha3.ServiceEntry, opts v1.UpdateOptions) (result *v1alpha3.ServiceEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceentriesResource, c.ns, serviceEntry), &v1alpha3.ServiceEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceEntry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceEntries) UpdateStatus(ctx context.Context, serviceEntry *v1alpha3.ServiceEntry, opts v1.UpdateOptions) (*v1alpha3.ServiceEntry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceentriesResource, "status", c.ns, serviceEntry), &v1alpha3.ServiceEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceEntry), err
}

// Delete takes name of the serviceEntry and deletes it. Returns an error if one occurs.
func (c *FakeServiceEntries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(serviceentriesResource, c.ns, name, opts), &v1alpha3.ServiceEntry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceEntries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceentriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.ServiceEntryList{})
	return err
}

// Patch applies the patch and returns the patched serviceEntry.
func (c *FakeServiceEntries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.ServiceEntry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceentriesResource, c.ns, name, pt, data, subresources...), &v1alpha3.ServiceEntry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceEntry), err
}
