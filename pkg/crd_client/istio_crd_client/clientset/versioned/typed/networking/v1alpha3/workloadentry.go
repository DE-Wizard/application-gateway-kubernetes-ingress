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

package v1alpha3

import (
	scheme "/clientset/versioned/scheme"
	"context"
	"time"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkloadEntriesGetter has a method to return a WorkloadEntryInterface.
// A group's client should implement this interface.
type WorkloadEntriesGetter interface {
	WorkloadEntries(namespace string) WorkloadEntryInterface
}

// WorkloadEntryInterface has methods to work with WorkloadEntry resources.
type WorkloadEntryInterface interface {
	Create(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.CreateOptions) (*v1alpha3.WorkloadEntry, error)
	Update(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.UpdateOptions) (*v1alpha3.WorkloadEntry, error)
	UpdateStatus(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.UpdateOptions) (*v1alpha3.WorkloadEntry, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.WorkloadEntry, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.WorkloadEntryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.WorkloadEntry, err error)
	WorkloadEntryExpansion
}

// workloadEntries implements WorkloadEntryInterface
type workloadEntries struct {
	client rest.Interface
	ns     string
}

// newWorkloadEntries returns a WorkloadEntries
func newWorkloadEntries(c *NetworkingV1alpha3Client, namespace string) *workloadEntries {
	return &workloadEntries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workloadEntry, and returns the corresponding workloadEntry object, and an error if there is any.
func (c *workloadEntries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.WorkloadEntry, err error) {
	result = &v1alpha3.WorkloadEntry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloadentries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkloadEntries that match those selectors.
func (c *workloadEntries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.WorkloadEntryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.WorkloadEntryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloadentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workloadEntries.
func (c *workloadEntries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workloadentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workloadEntry and creates it.  Returns the server's representation of the workloadEntry, and an error, if there is any.
func (c *workloadEntries) Create(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.CreateOptions) (result *v1alpha3.WorkloadEntry, err error) {
	result = &v1alpha3.WorkloadEntry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workloadentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadEntry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workloadEntry and updates it. Returns the server's representation of the workloadEntry, and an error, if there is any.
func (c *workloadEntries) Update(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.UpdateOptions) (result *v1alpha3.WorkloadEntry, err error) {
	result = &v1alpha3.WorkloadEntry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloadentries").
		Name(workloadEntry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadEntry).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workloadEntries) UpdateStatus(ctx context.Context, workloadEntry *v1alpha3.WorkloadEntry, opts v1.UpdateOptions) (result *v1alpha3.WorkloadEntry, err error) {
	result = &v1alpha3.WorkloadEntry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloadentries").
		Name(workloadEntry.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadEntry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workloadEntry and deletes it. Returns an error if one occurs.
func (c *workloadEntries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloadentries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workloadEntries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloadentries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workloadEntry.
func (c *workloadEntries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.WorkloadEntry, err error) {
	result = &v1alpha3.WorkloadEntry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workloadentries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
