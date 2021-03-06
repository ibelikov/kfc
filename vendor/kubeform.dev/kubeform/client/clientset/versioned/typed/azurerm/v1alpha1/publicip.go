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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// PublicIPsGetter has a method to return a PublicIPInterface.
// A group's client should implement this interface.
type PublicIPsGetter interface {
	PublicIPs(namespace string) PublicIPInterface
}

// PublicIPInterface has methods to work with PublicIP resources.
type PublicIPInterface interface {
	Create(*v1alpha1.PublicIP) (*v1alpha1.PublicIP, error)
	Update(*v1alpha1.PublicIP) (*v1alpha1.PublicIP, error)
	UpdateStatus(*v1alpha1.PublicIP) (*v1alpha1.PublicIP, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PublicIP, error)
	List(opts v1.ListOptions) (*v1alpha1.PublicIPList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIP, err error)
	PublicIPExpansion
}

// publicIPs implements PublicIPInterface
type publicIPs struct {
	client rest.Interface
	ns     string
}

// newPublicIPs returns a PublicIPs
func newPublicIPs(c *AzurermV1alpha1Client, namespace string) *publicIPs {
	return &publicIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the publicIP, and returns the corresponding publicIP object, and an error if there is any.
func (c *publicIPs) Get(name string, options v1.GetOptions) (result *v1alpha1.PublicIP, err error) {
	result = &v1alpha1.PublicIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("publicips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PublicIPs that match those selectors.
func (c *publicIPs) List(opts v1.ListOptions) (result *v1alpha1.PublicIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PublicIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("publicips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested publicIPs.
func (c *publicIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("publicips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a publicIP and creates it.  Returns the server's representation of the publicIP, and an error, if there is any.
func (c *publicIPs) Create(publicIP *v1alpha1.PublicIP) (result *v1alpha1.PublicIP, err error) {
	result = &v1alpha1.PublicIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("publicips").
		Body(publicIP).
		Do().
		Into(result)
	return
}

// Update takes the representation of a publicIP and updates it. Returns the server's representation of the publicIP, and an error, if there is any.
func (c *publicIPs) Update(publicIP *v1alpha1.PublicIP) (result *v1alpha1.PublicIP, err error) {
	result = &v1alpha1.PublicIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("publicips").
		Name(publicIP.Name).
		Body(publicIP).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *publicIPs) UpdateStatus(publicIP *v1alpha1.PublicIP) (result *v1alpha1.PublicIP, err error) {
	result = &v1alpha1.PublicIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("publicips").
		Name(publicIP.Name).
		SubResource("status").
		Body(publicIP).
		Do().
		Into(result)
	return
}

// Delete takes name of the publicIP and deletes it. Returns an error if one occurs.
func (c *publicIPs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("publicips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *publicIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("publicips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched publicIP.
func (c *publicIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIP, err error) {
	result = &v1alpha1.PublicIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("publicips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
