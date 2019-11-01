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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IothubConsumerGroupsGetter has a method to return a IothubConsumerGroupInterface.
// A group's client should implement this interface.
type IothubConsumerGroupsGetter interface {
	IothubConsumerGroups(namespace string) IothubConsumerGroupInterface
}

// IothubConsumerGroupInterface has methods to work with IothubConsumerGroup resources.
type IothubConsumerGroupInterface interface {
	Create(*v1alpha1.IothubConsumerGroup) (*v1alpha1.IothubConsumerGroup, error)
	Update(*v1alpha1.IothubConsumerGroup) (*v1alpha1.IothubConsumerGroup, error)
	UpdateStatus(*v1alpha1.IothubConsumerGroup) (*v1alpha1.IothubConsumerGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IothubConsumerGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.IothubConsumerGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubConsumerGroup, err error)
	IothubConsumerGroupExpansion
}

// iothubConsumerGroups implements IothubConsumerGroupInterface
type iothubConsumerGroups struct {
	client rest.Interface
	ns     string
}

// newIothubConsumerGroups returns a IothubConsumerGroups
func newIothubConsumerGroups(c *AzurermV1alpha1Client, namespace string) *iothubConsumerGroups {
	return &iothubConsumerGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iothubConsumerGroup, and returns the corresponding iothubConsumerGroup object, and an error if there is any.
func (c *iothubConsumerGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.IothubConsumerGroup, err error) {
	result = &v1alpha1.IothubConsumerGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IothubConsumerGroups that match those selectors.
func (c *iothubConsumerGroups) List(opts v1.ListOptions) (result *v1alpha1.IothubConsumerGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IothubConsumerGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iothubConsumerGroups.
func (c *iothubConsumerGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iothubConsumerGroup and creates it.  Returns the server's representation of the iothubConsumerGroup, and an error, if there is any.
func (c *iothubConsumerGroups) Create(iothubConsumerGroup *v1alpha1.IothubConsumerGroup) (result *v1alpha1.IothubConsumerGroup, err error) {
	result = &v1alpha1.IothubConsumerGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		Body(iothubConsumerGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iothubConsumerGroup and updates it. Returns the server's representation of the iothubConsumerGroup, and an error, if there is any.
func (c *iothubConsumerGroups) Update(iothubConsumerGroup *v1alpha1.IothubConsumerGroup) (result *v1alpha1.IothubConsumerGroup, err error) {
	result = &v1alpha1.IothubConsumerGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		Name(iothubConsumerGroup.Name).
		Body(iothubConsumerGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iothubConsumerGroups) UpdateStatus(iothubConsumerGroup *v1alpha1.IothubConsumerGroup) (result *v1alpha1.IothubConsumerGroup, err error) {
	result = &v1alpha1.IothubConsumerGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		Name(iothubConsumerGroup.Name).
		SubResource("status").
		Body(iothubConsumerGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the iothubConsumerGroup and deletes it. Returns an error if one occurs.
func (c *iothubConsumerGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iothubConsumerGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iothubConsumerGroup.
func (c *iothubConsumerGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubConsumerGroup, err error) {
	result = &v1alpha1.IothubConsumerGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iothubconsumergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}