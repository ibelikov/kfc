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

// MqConfigurationsGetter has a method to return a MqConfigurationInterface.
// A group's client should implement this interface.
type MqConfigurationsGetter interface {
	MqConfigurations(namespace string) MqConfigurationInterface
}

// MqConfigurationInterface has methods to work with MqConfiguration resources.
type MqConfigurationInterface interface {
	Create(*v1alpha1.MqConfiguration) (*v1alpha1.MqConfiguration, error)
	Update(*v1alpha1.MqConfiguration) (*v1alpha1.MqConfiguration, error)
	UpdateStatus(*v1alpha1.MqConfiguration) (*v1alpha1.MqConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MqConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.MqConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MqConfiguration, err error)
	MqConfigurationExpansion
}

// mqConfigurations implements MqConfigurationInterface
type mqConfigurations struct {
	client rest.Interface
	ns     string
}

// newMqConfigurations returns a MqConfigurations
func newMqConfigurations(c *AwsV1alpha1Client, namespace string) *mqConfigurations {
	return &mqConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mqConfiguration, and returns the corresponding mqConfiguration object, and an error if there is any.
func (c *mqConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.MqConfiguration, err error) {
	result = &v1alpha1.MqConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mqconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MqConfigurations that match those selectors.
func (c *mqConfigurations) List(opts v1.ListOptions) (result *v1alpha1.MqConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MqConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mqconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mqConfigurations.
func (c *mqConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mqconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mqConfiguration and creates it.  Returns the server's representation of the mqConfiguration, and an error, if there is any.
func (c *mqConfigurations) Create(mqConfiguration *v1alpha1.MqConfiguration) (result *v1alpha1.MqConfiguration, err error) {
	result = &v1alpha1.MqConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mqconfigurations").
		Body(mqConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mqConfiguration and updates it. Returns the server's representation of the mqConfiguration, and an error, if there is any.
func (c *mqConfigurations) Update(mqConfiguration *v1alpha1.MqConfiguration) (result *v1alpha1.MqConfiguration, err error) {
	result = &v1alpha1.MqConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mqconfigurations").
		Name(mqConfiguration.Name).
		Body(mqConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mqConfigurations) UpdateStatus(mqConfiguration *v1alpha1.MqConfiguration) (result *v1alpha1.MqConfiguration, err error) {
	result = &v1alpha1.MqConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mqconfigurations").
		Name(mqConfiguration.Name).
		SubResource("status").
		Body(mqConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the mqConfiguration and deletes it. Returns an error if one occurs.
func (c *mqConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mqconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mqConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mqconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mqConfiguration.
func (c *mqConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MqConfiguration, err error) {
	result = &v1alpha1.MqConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mqconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
