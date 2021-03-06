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

// SesEventDestinationsGetter has a method to return a SesEventDestinationInterface.
// A group's client should implement this interface.
type SesEventDestinationsGetter interface {
	SesEventDestinations(namespace string) SesEventDestinationInterface
}

// SesEventDestinationInterface has methods to work with SesEventDestination resources.
type SesEventDestinationInterface interface {
	Create(*v1alpha1.SesEventDestination) (*v1alpha1.SesEventDestination, error)
	Update(*v1alpha1.SesEventDestination) (*v1alpha1.SesEventDestination, error)
	UpdateStatus(*v1alpha1.SesEventDestination) (*v1alpha1.SesEventDestination, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesEventDestination, error)
	List(opts v1.ListOptions) (*v1alpha1.SesEventDestinationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesEventDestination, err error)
	SesEventDestinationExpansion
}

// sesEventDestinations implements SesEventDestinationInterface
type sesEventDestinations struct {
	client rest.Interface
	ns     string
}

// newSesEventDestinations returns a SesEventDestinations
func newSesEventDestinations(c *AwsV1alpha1Client, namespace string) *sesEventDestinations {
	return &sesEventDestinations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sesEventDestination, and returns the corresponding sesEventDestination object, and an error if there is any.
func (c *sesEventDestinations) Get(name string, options v1.GetOptions) (result *v1alpha1.SesEventDestination, err error) {
	result = &v1alpha1.SesEventDestination{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("seseventdestinations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesEventDestinations that match those selectors.
func (c *sesEventDestinations) List(opts v1.ListOptions) (result *v1alpha1.SesEventDestinationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesEventDestinationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("seseventdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesEventDestinations.
func (c *sesEventDestinations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("seseventdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesEventDestination and creates it.  Returns the server's representation of the sesEventDestination, and an error, if there is any.
func (c *sesEventDestinations) Create(sesEventDestination *v1alpha1.SesEventDestination) (result *v1alpha1.SesEventDestination, err error) {
	result = &v1alpha1.SesEventDestination{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("seseventdestinations").
		Body(sesEventDestination).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesEventDestination and updates it. Returns the server's representation of the sesEventDestination, and an error, if there is any.
func (c *sesEventDestinations) Update(sesEventDestination *v1alpha1.SesEventDestination) (result *v1alpha1.SesEventDestination, err error) {
	result = &v1alpha1.SesEventDestination{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("seseventdestinations").
		Name(sesEventDestination.Name).
		Body(sesEventDestination).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesEventDestinations) UpdateStatus(sesEventDestination *v1alpha1.SesEventDestination) (result *v1alpha1.SesEventDestination, err error) {
	result = &v1alpha1.SesEventDestination{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("seseventdestinations").
		Name(sesEventDestination.Name).
		SubResource("status").
		Body(sesEventDestination).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesEventDestination and deletes it. Returns an error if one occurs.
func (c *sesEventDestinations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("seseventdestinations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesEventDestinations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("seseventdestinations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesEventDestination.
func (c *sesEventDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesEventDestination, err error) {
	result = &v1alpha1.SesEventDestination{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("seseventdestinations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
