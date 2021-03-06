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

// AutoscalingAttachmentsGetter has a method to return a AutoscalingAttachmentInterface.
// A group's client should implement this interface.
type AutoscalingAttachmentsGetter interface {
	AutoscalingAttachments(namespace string) AutoscalingAttachmentInterface
}

// AutoscalingAttachmentInterface has methods to work with AutoscalingAttachment resources.
type AutoscalingAttachmentInterface interface {
	Create(*v1alpha1.AutoscalingAttachment) (*v1alpha1.AutoscalingAttachment, error)
	Update(*v1alpha1.AutoscalingAttachment) (*v1alpha1.AutoscalingAttachment, error)
	UpdateStatus(*v1alpha1.AutoscalingAttachment) (*v1alpha1.AutoscalingAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutoscalingAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.AutoscalingAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingAttachment, err error)
	AutoscalingAttachmentExpansion
}

// autoscalingAttachments implements AutoscalingAttachmentInterface
type autoscalingAttachments struct {
	client rest.Interface
	ns     string
}

// newAutoscalingAttachments returns a AutoscalingAttachments
func newAutoscalingAttachments(c *AwsV1alpha1Client, namespace string) *autoscalingAttachments {
	return &autoscalingAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the autoscalingAttachment, and returns the corresponding autoscalingAttachment object, and an error if there is any.
func (c *autoscalingAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoscalingAttachment, err error) {
	result = &v1alpha1.AutoscalingAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutoscalingAttachments that match those selectors.
func (c *autoscalingAttachments) List(opts v1.ListOptions) (result *v1alpha1.AutoscalingAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutoscalingAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested autoscalingAttachments.
func (c *autoscalingAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a autoscalingAttachment and creates it.  Returns the server's representation of the autoscalingAttachment, and an error, if there is any.
func (c *autoscalingAttachments) Create(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (result *v1alpha1.AutoscalingAttachment, err error) {
	result = &v1alpha1.AutoscalingAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		Body(autoscalingAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a autoscalingAttachment and updates it. Returns the server's representation of the autoscalingAttachment, and an error, if there is any.
func (c *autoscalingAttachments) Update(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (result *v1alpha1.AutoscalingAttachment, err error) {
	result = &v1alpha1.AutoscalingAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		Name(autoscalingAttachment.Name).
		Body(autoscalingAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *autoscalingAttachments) UpdateStatus(autoscalingAttachment *v1alpha1.AutoscalingAttachment) (result *v1alpha1.AutoscalingAttachment, err error) {
	result = &v1alpha1.AutoscalingAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		Name(autoscalingAttachment.Name).
		SubResource("status").
		Body(autoscalingAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the autoscalingAttachment and deletes it. Returns an error if one occurs.
func (c *autoscalingAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *autoscalingAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("autoscalingattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched autoscalingAttachment.
func (c *autoscalingAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoscalingAttachment, err error) {
	result = &v1alpha1.AutoscalingAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("autoscalingattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
