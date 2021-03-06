/*
Copyright The Kubeform Authors.

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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// RamResourceAssociationsGetter has a method to return a RamResourceAssociationInterface.
// A group's client should implement this interface.
type RamResourceAssociationsGetter interface {
	RamResourceAssociations(namespace string) RamResourceAssociationInterface
}

// RamResourceAssociationInterface has methods to work with RamResourceAssociation resources.
type RamResourceAssociationInterface interface {
	Create(*v1alpha1.RamResourceAssociation) (*v1alpha1.RamResourceAssociation, error)
	Update(*v1alpha1.RamResourceAssociation) (*v1alpha1.RamResourceAssociation, error)
	UpdateStatus(*v1alpha1.RamResourceAssociation) (*v1alpha1.RamResourceAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RamResourceAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.RamResourceAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RamResourceAssociation, err error)
	RamResourceAssociationExpansion
}

// ramResourceAssociations implements RamResourceAssociationInterface
type ramResourceAssociations struct {
	client rest.Interface
	ns     string
}

// newRamResourceAssociations returns a RamResourceAssociations
func newRamResourceAssociations(c *AwsV1alpha1Client, namespace string) *ramResourceAssociations {
	return &ramResourceAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ramResourceAssociation, and returns the corresponding ramResourceAssociation object, and an error if there is any.
func (c *ramResourceAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.RamResourceAssociation, err error) {
	result = &v1alpha1.RamResourceAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RamResourceAssociations that match those selectors.
func (c *ramResourceAssociations) List(opts v1.ListOptions) (result *v1alpha1.RamResourceAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RamResourceAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ramResourceAssociations.
func (c *ramResourceAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ramResourceAssociation and creates it.  Returns the server's representation of the ramResourceAssociation, and an error, if there is any.
func (c *ramResourceAssociations) Create(ramResourceAssociation *v1alpha1.RamResourceAssociation) (result *v1alpha1.RamResourceAssociation, err error) {
	result = &v1alpha1.RamResourceAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		Body(ramResourceAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ramResourceAssociation and updates it. Returns the server's representation of the ramResourceAssociation, and an error, if there is any.
func (c *ramResourceAssociations) Update(ramResourceAssociation *v1alpha1.RamResourceAssociation) (result *v1alpha1.RamResourceAssociation, err error) {
	result = &v1alpha1.RamResourceAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		Name(ramResourceAssociation.Name).
		Body(ramResourceAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ramResourceAssociations) UpdateStatus(ramResourceAssociation *v1alpha1.RamResourceAssociation) (result *v1alpha1.RamResourceAssociation, err error) {
	result = &v1alpha1.RamResourceAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		Name(ramResourceAssociation.Name).
		SubResource("status").
		Body(ramResourceAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the ramResourceAssociation and deletes it. Returns an error if one occurs.
func (c *ramResourceAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ramResourceAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ramresourceassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ramResourceAssociation.
func (c *ramResourceAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RamResourceAssociation, err error) {
	result = &v1alpha1.RamResourceAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ramresourceassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
