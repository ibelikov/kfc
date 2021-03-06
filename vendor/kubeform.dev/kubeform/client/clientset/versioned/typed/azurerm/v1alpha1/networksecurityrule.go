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

// NetworkSecurityRulesGetter has a method to return a NetworkSecurityRuleInterface.
// A group's client should implement this interface.
type NetworkSecurityRulesGetter interface {
	NetworkSecurityRules(namespace string) NetworkSecurityRuleInterface
}

// NetworkSecurityRuleInterface has methods to work with NetworkSecurityRule resources.
type NetworkSecurityRuleInterface interface {
	Create(*v1alpha1.NetworkSecurityRule) (*v1alpha1.NetworkSecurityRule, error)
	Update(*v1alpha1.NetworkSecurityRule) (*v1alpha1.NetworkSecurityRule, error)
	UpdateStatus(*v1alpha1.NetworkSecurityRule) (*v1alpha1.NetworkSecurityRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkSecurityRule, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkSecurityRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkSecurityRule, err error)
	NetworkSecurityRuleExpansion
}

// networkSecurityRules implements NetworkSecurityRuleInterface
type networkSecurityRules struct {
	client rest.Interface
	ns     string
}

// newNetworkSecurityRules returns a NetworkSecurityRules
func newNetworkSecurityRules(c *AzurermV1alpha1Client, namespace string) *networkSecurityRules {
	return &networkSecurityRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkSecurityRule, and returns the corresponding networkSecurityRule object, and an error if there is any.
func (c *networkSecurityRules) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkSecurityRule, err error) {
	result = &v1alpha1.NetworkSecurityRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networksecurityrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkSecurityRules that match those selectors.
func (c *networkSecurityRules) List(opts v1.ListOptions) (result *v1alpha1.NetworkSecurityRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkSecurityRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networksecurityrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkSecurityRules.
func (c *networkSecurityRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networksecurityrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkSecurityRule and creates it.  Returns the server's representation of the networkSecurityRule, and an error, if there is any.
func (c *networkSecurityRules) Create(networkSecurityRule *v1alpha1.NetworkSecurityRule) (result *v1alpha1.NetworkSecurityRule, err error) {
	result = &v1alpha1.NetworkSecurityRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networksecurityrules").
		Body(networkSecurityRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkSecurityRule and updates it. Returns the server's representation of the networkSecurityRule, and an error, if there is any.
func (c *networkSecurityRules) Update(networkSecurityRule *v1alpha1.NetworkSecurityRule) (result *v1alpha1.NetworkSecurityRule, err error) {
	result = &v1alpha1.NetworkSecurityRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networksecurityrules").
		Name(networkSecurityRule.Name).
		Body(networkSecurityRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkSecurityRules) UpdateStatus(networkSecurityRule *v1alpha1.NetworkSecurityRule) (result *v1alpha1.NetworkSecurityRule, err error) {
	result = &v1alpha1.NetworkSecurityRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networksecurityrules").
		Name(networkSecurityRule.Name).
		SubResource("status").
		Body(networkSecurityRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkSecurityRule and deletes it. Returns an error if one occurs.
func (c *networkSecurityRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networksecurityrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkSecurityRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networksecurityrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkSecurityRule.
func (c *networkSecurityRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkSecurityRule, err error) {
	result = &v1alpha1.NetworkSecurityRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networksecurityrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
