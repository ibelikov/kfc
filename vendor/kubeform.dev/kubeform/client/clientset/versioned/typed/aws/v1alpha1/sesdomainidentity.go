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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SesDomainIdentitiesGetter has a method to return a SesDomainIdentityInterface.
// A group's client should implement this interface.
type SesDomainIdentitiesGetter interface {
	SesDomainIdentities(namespace string) SesDomainIdentityInterface
}

// SesDomainIdentityInterface has methods to work with SesDomainIdentity resources.
type SesDomainIdentityInterface interface {
	Create(*v1alpha1.SesDomainIdentity) (*v1alpha1.SesDomainIdentity, error)
	Update(*v1alpha1.SesDomainIdentity) (*v1alpha1.SesDomainIdentity, error)
	UpdateStatus(*v1alpha1.SesDomainIdentity) (*v1alpha1.SesDomainIdentity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesDomainIdentity, error)
	List(opts v1.ListOptions) (*v1alpha1.SesDomainIdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesDomainIdentity, err error)
	SesDomainIdentityExpansion
}

// sesDomainIdentities implements SesDomainIdentityInterface
type sesDomainIdentities struct {
	client rest.Interface
	ns     string
}

// newSesDomainIdentities returns a SesDomainIdentities
func newSesDomainIdentities(c *AwsV1alpha1Client, namespace string) *sesDomainIdentities {
	return &sesDomainIdentities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sesDomainIdentity, and returns the corresponding sesDomainIdentity object, and an error if there is any.
func (c *sesDomainIdentities) Get(name string, options v1.GetOptions) (result *v1alpha1.SesDomainIdentity, err error) {
	result = &v1alpha1.SesDomainIdentity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesDomainIdentities that match those selectors.
func (c *sesDomainIdentities) List(opts v1.ListOptions) (result *v1alpha1.SesDomainIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesDomainIdentityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesDomainIdentities.
func (c *sesDomainIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesDomainIdentity and creates it.  Returns the server's representation of the sesDomainIdentity, and an error, if there is any.
func (c *sesDomainIdentities) Create(sesDomainIdentity *v1alpha1.SesDomainIdentity) (result *v1alpha1.SesDomainIdentity, err error) {
	result = &v1alpha1.SesDomainIdentity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		Body(sesDomainIdentity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesDomainIdentity and updates it. Returns the server's representation of the sesDomainIdentity, and an error, if there is any.
func (c *sesDomainIdentities) Update(sesDomainIdentity *v1alpha1.SesDomainIdentity) (result *v1alpha1.SesDomainIdentity, err error) {
	result = &v1alpha1.SesDomainIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		Name(sesDomainIdentity.Name).
		Body(sesDomainIdentity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesDomainIdentities) UpdateStatus(sesDomainIdentity *v1alpha1.SesDomainIdentity) (result *v1alpha1.SesDomainIdentity, err error) {
	result = &v1alpha1.SesDomainIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		Name(sesDomainIdentity.Name).
		SubResource("status").
		Body(sesDomainIdentity).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesDomainIdentity and deletes it. Returns an error if one occurs.
func (c *sesDomainIdentities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesDomainIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesdomainidentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesDomainIdentity.
func (c *sesDomainIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesDomainIdentity, err error) {
	result = &v1alpha1.SesDomainIdentity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sesdomainidentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
