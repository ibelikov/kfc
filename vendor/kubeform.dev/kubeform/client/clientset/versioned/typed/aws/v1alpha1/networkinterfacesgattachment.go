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

// NetworkInterfaceSgAttachmentsGetter has a method to return a NetworkInterfaceSgAttachmentInterface.
// A group's client should implement this interface.
type NetworkInterfaceSgAttachmentsGetter interface {
	NetworkInterfaceSgAttachments(namespace string) NetworkInterfaceSgAttachmentInterface
}

// NetworkInterfaceSgAttachmentInterface has methods to work with NetworkInterfaceSgAttachment resources.
type NetworkInterfaceSgAttachmentInterface interface {
	Create(*v1alpha1.NetworkInterfaceSgAttachment) (*v1alpha1.NetworkInterfaceSgAttachment, error)
	Update(*v1alpha1.NetworkInterfaceSgAttachment) (*v1alpha1.NetworkInterfaceSgAttachment, error)
	UpdateStatus(*v1alpha1.NetworkInterfaceSgAttachment) (*v1alpha1.NetworkInterfaceSgAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkInterfaceSgAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkInterfaceSgAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceSgAttachment, err error)
	NetworkInterfaceSgAttachmentExpansion
}

// networkInterfaceSgAttachments implements NetworkInterfaceSgAttachmentInterface
type networkInterfaceSgAttachments struct {
	client rest.Interface
	ns     string
}

// newNetworkInterfaceSgAttachments returns a NetworkInterfaceSgAttachments
func newNetworkInterfaceSgAttachments(c *AwsV1alpha1Client, namespace string) *networkInterfaceSgAttachments {
	return &networkInterfaceSgAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkInterfaceSgAttachment, and returns the corresponding networkInterfaceSgAttachment object, and an error if there is any.
func (c *networkInterfaceSgAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceSgAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceSgAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkInterfaceSgAttachments that match those selectors.
func (c *networkInterfaceSgAttachments) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceSgAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkInterfaceSgAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkInterfaceSgAttachments.
func (c *networkInterfaceSgAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkInterfaceSgAttachment and creates it.  Returns the server's representation of the networkInterfaceSgAttachment, and an error, if there is any.
func (c *networkInterfaceSgAttachments) Create(networkInterfaceSgAttachment *v1alpha1.NetworkInterfaceSgAttachment) (result *v1alpha1.NetworkInterfaceSgAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceSgAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		Body(networkInterfaceSgAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkInterfaceSgAttachment and updates it. Returns the server's representation of the networkInterfaceSgAttachment, and an error, if there is any.
func (c *networkInterfaceSgAttachments) Update(networkInterfaceSgAttachment *v1alpha1.NetworkInterfaceSgAttachment) (result *v1alpha1.NetworkInterfaceSgAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceSgAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		Name(networkInterfaceSgAttachment.Name).
		Body(networkInterfaceSgAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkInterfaceSgAttachments) UpdateStatus(networkInterfaceSgAttachment *v1alpha1.NetworkInterfaceSgAttachment) (result *v1alpha1.NetworkInterfaceSgAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceSgAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		Name(networkInterfaceSgAttachment.Name).
		SubResource("status").
		Body(networkInterfaceSgAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkInterfaceSgAttachment and deletes it. Returns an error if one occurs.
func (c *networkInterfaceSgAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkInterfaceSgAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkInterfaceSgAttachment.
func (c *networkInterfaceSgAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceSgAttachment, err error) {
	result = &v1alpha1.NetworkInterfaceSgAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkinterfacesgattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
