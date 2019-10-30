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

// VpnGatewayAttachmentsGetter has a method to return a VpnGatewayAttachmentInterface.
// A group's client should implement this interface.
type VpnGatewayAttachmentsGetter interface {
	VpnGatewayAttachments(namespace string) VpnGatewayAttachmentInterface
}

// VpnGatewayAttachmentInterface has methods to work with VpnGatewayAttachment resources.
type VpnGatewayAttachmentInterface interface {
	Create(*v1alpha1.VpnGatewayAttachment) (*v1alpha1.VpnGatewayAttachment, error)
	Update(*v1alpha1.VpnGatewayAttachment) (*v1alpha1.VpnGatewayAttachment, error)
	UpdateStatus(*v1alpha1.VpnGatewayAttachment) (*v1alpha1.VpnGatewayAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VpnGatewayAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.VpnGatewayAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnGatewayAttachment, err error)
	VpnGatewayAttachmentExpansion
}

// vpnGatewayAttachments implements VpnGatewayAttachmentInterface
type vpnGatewayAttachments struct {
	client rest.Interface
	ns     string
}

// newVpnGatewayAttachments returns a VpnGatewayAttachments
func newVpnGatewayAttachments(c *AwsV1alpha1Client, namespace string) *vpnGatewayAttachments {
	return &vpnGatewayAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpnGatewayAttachment, and returns the corresponding vpnGatewayAttachment object, and an error if there is any.
func (c *vpnGatewayAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.VpnGatewayAttachment, err error) {
	result = &v1alpha1.VpnGatewayAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpnGatewayAttachments that match those selectors.
func (c *vpnGatewayAttachments) List(opts v1.ListOptions) (result *v1alpha1.VpnGatewayAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpnGatewayAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpnGatewayAttachments.
func (c *vpnGatewayAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vpnGatewayAttachment and creates it.  Returns the server's representation of the vpnGatewayAttachment, and an error, if there is any.
func (c *vpnGatewayAttachments) Create(vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment) (result *v1alpha1.VpnGatewayAttachment, err error) {
	result = &v1alpha1.VpnGatewayAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		Body(vpnGatewayAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vpnGatewayAttachment and updates it. Returns the server's representation of the vpnGatewayAttachment, and an error, if there is any.
func (c *vpnGatewayAttachments) Update(vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment) (result *v1alpha1.VpnGatewayAttachment, err error) {
	result = &v1alpha1.VpnGatewayAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		Name(vpnGatewayAttachment.Name).
		Body(vpnGatewayAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vpnGatewayAttachments) UpdateStatus(vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment) (result *v1alpha1.VpnGatewayAttachment, err error) {
	result = &v1alpha1.VpnGatewayAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		Name(vpnGatewayAttachment.Name).
		SubResource("status").
		Body(vpnGatewayAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the vpnGatewayAttachment and deletes it. Returns an error if one occurs.
func (c *vpnGatewayAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpnGatewayAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vpnGatewayAttachment.
func (c *vpnGatewayAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnGatewayAttachment, err error) {
	result = &v1alpha1.VpnGatewayAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpngatewayattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
