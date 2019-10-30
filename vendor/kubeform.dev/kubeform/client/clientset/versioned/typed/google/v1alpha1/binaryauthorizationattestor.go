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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BinaryAuthorizationAttestorsGetter has a method to return a BinaryAuthorizationAttestorInterface.
// A group's client should implement this interface.
type BinaryAuthorizationAttestorsGetter interface {
	BinaryAuthorizationAttestors(namespace string) BinaryAuthorizationAttestorInterface
}

// BinaryAuthorizationAttestorInterface has methods to work with BinaryAuthorizationAttestor resources.
type BinaryAuthorizationAttestorInterface interface {
	Create(*v1alpha1.BinaryAuthorizationAttestor) (*v1alpha1.BinaryAuthorizationAttestor, error)
	Update(*v1alpha1.BinaryAuthorizationAttestor) (*v1alpha1.BinaryAuthorizationAttestor, error)
	UpdateStatus(*v1alpha1.BinaryAuthorizationAttestor) (*v1alpha1.BinaryAuthorizationAttestor, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BinaryAuthorizationAttestor, error)
	List(opts v1.ListOptions) (*v1alpha1.BinaryAuthorizationAttestorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BinaryAuthorizationAttestor, err error)
	BinaryAuthorizationAttestorExpansion
}

// binaryAuthorizationAttestors implements BinaryAuthorizationAttestorInterface
type binaryAuthorizationAttestors struct {
	client rest.Interface
	ns     string
}

// newBinaryAuthorizationAttestors returns a BinaryAuthorizationAttestors
func newBinaryAuthorizationAttestors(c *GoogleV1alpha1Client, namespace string) *binaryAuthorizationAttestors {
	return &binaryAuthorizationAttestors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the binaryAuthorizationAttestor, and returns the corresponding binaryAuthorizationAttestor object, and an error if there is any.
func (c *binaryAuthorizationAttestors) Get(name string, options v1.GetOptions) (result *v1alpha1.BinaryAuthorizationAttestor, err error) {
	result = &v1alpha1.BinaryAuthorizationAttestor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BinaryAuthorizationAttestors that match those selectors.
func (c *binaryAuthorizationAttestors) List(opts v1.ListOptions) (result *v1alpha1.BinaryAuthorizationAttestorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BinaryAuthorizationAttestorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested binaryAuthorizationAttestors.
func (c *binaryAuthorizationAttestors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a binaryAuthorizationAttestor and creates it.  Returns the server's representation of the binaryAuthorizationAttestor, and an error, if there is any.
func (c *binaryAuthorizationAttestors) Create(binaryAuthorizationAttestor *v1alpha1.BinaryAuthorizationAttestor) (result *v1alpha1.BinaryAuthorizationAttestor, err error) {
	result = &v1alpha1.BinaryAuthorizationAttestor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		Body(binaryAuthorizationAttestor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a binaryAuthorizationAttestor and updates it. Returns the server's representation of the binaryAuthorizationAttestor, and an error, if there is any.
func (c *binaryAuthorizationAttestors) Update(binaryAuthorizationAttestor *v1alpha1.BinaryAuthorizationAttestor) (result *v1alpha1.BinaryAuthorizationAttestor, err error) {
	result = &v1alpha1.BinaryAuthorizationAttestor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		Name(binaryAuthorizationAttestor.Name).
		Body(binaryAuthorizationAttestor).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *binaryAuthorizationAttestors) UpdateStatus(binaryAuthorizationAttestor *v1alpha1.BinaryAuthorizationAttestor) (result *v1alpha1.BinaryAuthorizationAttestor, err error) {
	result = &v1alpha1.BinaryAuthorizationAttestor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		Name(binaryAuthorizationAttestor.Name).
		SubResource("status").
		Body(binaryAuthorizationAttestor).
		Do().
		Into(result)
	return
}

// Delete takes name of the binaryAuthorizationAttestor and deletes it. Returns an error if one occurs.
func (c *binaryAuthorizationAttestors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *binaryAuthorizationAttestors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched binaryAuthorizationAttestor.
func (c *binaryAuthorizationAttestors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BinaryAuthorizationAttestor, err error) {
	result = &v1alpha1.BinaryAuthorizationAttestor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("binaryauthorizationattestors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
