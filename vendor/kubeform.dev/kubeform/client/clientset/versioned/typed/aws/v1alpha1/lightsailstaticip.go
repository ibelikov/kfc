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

// LightsailStaticIPsGetter has a method to return a LightsailStaticIPInterface.
// A group's client should implement this interface.
type LightsailStaticIPsGetter interface {
	LightsailStaticIPs(namespace string) LightsailStaticIPInterface
}

// LightsailStaticIPInterface has methods to work with LightsailStaticIP resources.
type LightsailStaticIPInterface interface {
	Create(*v1alpha1.LightsailStaticIP) (*v1alpha1.LightsailStaticIP, error)
	Update(*v1alpha1.LightsailStaticIP) (*v1alpha1.LightsailStaticIP, error)
	UpdateStatus(*v1alpha1.LightsailStaticIP) (*v1alpha1.LightsailStaticIP, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LightsailStaticIP, error)
	List(opts v1.ListOptions) (*v1alpha1.LightsailStaticIPList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIP, err error)
	LightsailStaticIPExpansion
}

// lightsailStaticIPs implements LightsailStaticIPInterface
type lightsailStaticIPs struct {
	client rest.Interface
	ns     string
}

// newLightsailStaticIPs returns a LightsailStaticIPs
func newLightsailStaticIPs(c *AwsV1alpha1Client, namespace string) *lightsailStaticIPs {
	return &lightsailStaticIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lightsailStaticIP, and returns the corresponding lightsailStaticIP object, and an error if there is any.
func (c *lightsailStaticIPs) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailStaticIP, err error) {
	result = &v1alpha1.LightsailStaticIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LightsailStaticIPs that match those selectors.
func (c *lightsailStaticIPs) List(opts v1.ListOptions) (result *v1alpha1.LightsailStaticIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LightsailStaticIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lightsailStaticIPs.
func (c *lightsailStaticIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lightsailStaticIP and creates it.  Returns the server's representation of the lightsailStaticIP, and an error, if there is any.
func (c *lightsailStaticIPs) Create(lightsailStaticIP *v1alpha1.LightsailStaticIP) (result *v1alpha1.LightsailStaticIP, err error) {
	result = &v1alpha1.LightsailStaticIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		Body(lightsailStaticIP).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lightsailStaticIP and updates it. Returns the server's representation of the lightsailStaticIP, and an error, if there is any.
func (c *lightsailStaticIPs) Update(lightsailStaticIP *v1alpha1.LightsailStaticIP) (result *v1alpha1.LightsailStaticIP, err error) {
	result = &v1alpha1.LightsailStaticIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		Name(lightsailStaticIP.Name).
		Body(lightsailStaticIP).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lightsailStaticIPs) UpdateStatus(lightsailStaticIP *v1alpha1.LightsailStaticIP) (result *v1alpha1.LightsailStaticIP, err error) {
	result = &v1alpha1.LightsailStaticIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		Name(lightsailStaticIP.Name).
		SubResource("status").
		Body(lightsailStaticIP).
		Do().
		Into(result)
	return
}

// Delete takes name of the lightsailStaticIP and deletes it. Returns an error if one occurs.
func (c *lightsailStaticIPs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lightsailStaticIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lightsailstaticips").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lightsailStaticIP.
func (c *lightsailStaticIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIP, err error) {
	result = &v1alpha1.LightsailStaticIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lightsailstaticips").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
