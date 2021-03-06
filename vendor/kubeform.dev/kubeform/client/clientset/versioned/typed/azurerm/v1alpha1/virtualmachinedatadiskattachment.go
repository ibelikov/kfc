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

// VirtualMachineDataDiskAttachmentsGetter has a method to return a VirtualMachineDataDiskAttachmentInterface.
// A group's client should implement this interface.
type VirtualMachineDataDiskAttachmentsGetter interface {
	VirtualMachineDataDiskAttachments(namespace string) VirtualMachineDataDiskAttachmentInterface
}

// VirtualMachineDataDiskAttachmentInterface has methods to work with VirtualMachineDataDiskAttachment resources.
type VirtualMachineDataDiskAttachmentInterface interface {
	Create(*v1alpha1.VirtualMachineDataDiskAttachment) (*v1alpha1.VirtualMachineDataDiskAttachment, error)
	Update(*v1alpha1.VirtualMachineDataDiskAttachment) (*v1alpha1.VirtualMachineDataDiskAttachment, error)
	UpdateStatus(*v1alpha1.VirtualMachineDataDiskAttachment) (*v1alpha1.VirtualMachineDataDiskAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineDataDiskAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineDataDiskAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error)
	VirtualMachineDataDiskAttachmentExpansion
}

// virtualMachineDataDiskAttachments implements VirtualMachineDataDiskAttachmentInterface
type virtualMachineDataDiskAttachments struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineDataDiskAttachments returns a VirtualMachineDataDiskAttachments
func newVirtualMachineDataDiskAttachments(c *AzurermV1alpha1Client, namespace string) *virtualMachineDataDiskAttachments {
	return &virtualMachineDataDiskAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineDataDiskAttachment, and returns the corresponding virtualMachineDataDiskAttachment object, and an error if there is any.
func (c *virtualMachineDataDiskAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	result = &v1alpha1.VirtualMachineDataDiskAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineDataDiskAttachments that match those selectors.
func (c *virtualMachineDataDiskAttachments) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineDataDiskAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineDataDiskAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineDataDiskAttachments.
func (c *virtualMachineDataDiskAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineDataDiskAttachment and creates it.  Returns the server's representation of the virtualMachineDataDiskAttachment, and an error, if there is any.
func (c *virtualMachineDataDiskAttachments) Create(virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	result = &v1alpha1.VirtualMachineDataDiskAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		Body(virtualMachineDataDiskAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineDataDiskAttachment and updates it. Returns the server's representation of the virtualMachineDataDiskAttachment, and an error, if there is any.
func (c *virtualMachineDataDiskAttachments) Update(virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	result = &v1alpha1.VirtualMachineDataDiskAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		Name(virtualMachineDataDiskAttachment.Name).
		Body(virtualMachineDataDiskAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineDataDiskAttachments) UpdateStatus(virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	result = &v1alpha1.VirtualMachineDataDiskAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		Name(virtualMachineDataDiskAttachment.Name).
		SubResource("status").
		Body(virtualMachineDataDiskAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineDataDiskAttachment and deletes it. Returns an error if one occurs.
func (c *virtualMachineDataDiskAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineDataDiskAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineDataDiskAttachment.
func (c *virtualMachineDataDiskAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	result = &v1alpha1.VirtualMachineDataDiskAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinedatadiskattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
