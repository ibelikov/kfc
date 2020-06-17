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

// ServicebusQueueAuthorizationRulesGetter has a method to return a ServicebusQueueAuthorizationRuleInterface.
// A group's client should implement this interface.
type ServicebusQueueAuthorizationRulesGetter interface {
	ServicebusQueueAuthorizationRules(namespace string) ServicebusQueueAuthorizationRuleInterface
}

// ServicebusQueueAuthorizationRuleInterface has methods to work with ServicebusQueueAuthorizationRule resources.
type ServicebusQueueAuthorizationRuleInterface interface {
	Create(*v1alpha1.ServicebusQueueAuthorizationRule) (*v1alpha1.ServicebusQueueAuthorizationRule, error)
	Update(*v1alpha1.ServicebusQueueAuthorizationRule) (*v1alpha1.ServicebusQueueAuthorizationRule, error)
	UpdateStatus(*v1alpha1.ServicebusQueueAuthorizationRule) (*v1alpha1.ServicebusQueueAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServicebusQueueAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.ServicebusQueueAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error)
	ServicebusQueueAuthorizationRuleExpansion
}

// servicebusQueueAuthorizationRules implements ServicebusQueueAuthorizationRuleInterface
type servicebusQueueAuthorizationRules struct {
	client rest.Interface
	ns     string
}

// newServicebusQueueAuthorizationRules returns a ServicebusQueueAuthorizationRules
func newServicebusQueueAuthorizationRules(c *AzurermV1alpha1Client, namespace string) *servicebusQueueAuthorizationRules {
	return &servicebusQueueAuthorizationRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the servicebusQueueAuthorizationRule, and returns the corresponding servicebusQueueAuthorizationRule object, and an error if there is any.
func (c *servicebusQueueAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusQueueAuthorizationRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServicebusQueueAuthorizationRules that match those selectors.
func (c *servicebusQueueAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.ServicebusQueueAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServicebusQueueAuthorizationRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested servicebusQueueAuthorizationRules.
func (c *servicebusQueueAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a servicebusQueueAuthorizationRule and creates it.  Returns the server's representation of the servicebusQueueAuthorizationRule, and an error, if there is any.
func (c *servicebusQueueAuthorizationRules) Create(servicebusQueueAuthorizationRule *v1alpha1.ServicebusQueueAuthorizationRule) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusQueueAuthorizationRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		Body(servicebusQueueAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a servicebusQueueAuthorizationRule and updates it. Returns the server's representation of the servicebusQueueAuthorizationRule, and an error, if there is any.
func (c *servicebusQueueAuthorizationRules) Update(servicebusQueueAuthorizationRule *v1alpha1.ServicebusQueueAuthorizationRule) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusQueueAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		Name(servicebusQueueAuthorizationRule.Name).
		Body(servicebusQueueAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *servicebusQueueAuthorizationRules) UpdateStatus(servicebusQueueAuthorizationRule *v1alpha1.ServicebusQueueAuthorizationRule) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusQueueAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		Name(servicebusQueueAuthorizationRule.Name).
		SubResource("status").
		Body(servicebusQueueAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the servicebusQueueAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *servicebusQueueAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *servicebusQueueAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched servicebusQueueAuthorizationRule.
func (c *servicebusQueueAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	result = &v1alpha1.ServicebusQueueAuthorizationRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicebusqueueauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
