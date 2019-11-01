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

// DataLakeStoreFirewallRulesGetter has a method to return a DataLakeStoreFirewallRuleInterface.
// A group's client should implement this interface.
type DataLakeStoreFirewallRulesGetter interface {
	DataLakeStoreFirewallRules(namespace string) DataLakeStoreFirewallRuleInterface
}

// DataLakeStoreFirewallRuleInterface has methods to work with DataLakeStoreFirewallRule resources.
type DataLakeStoreFirewallRuleInterface interface {
	Create(*v1alpha1.DataLakeStoreFirewallRule) (*v1alpha1.DataLakeStoreFirewallRule, error)
	Update(*v1alpha1.DataLakeStoreFirewallRule) (*v1alpha1.DataLakeStoreFirewallRule, error)
	UpdateStatus(*v1alpha1.DataLakeStoreFirewallRule) (*v1alpha1.DataLakeStoreFirewallRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataLakeStoreFirewallRule, error)
	List(opts v1.ListOptions) (*v1alpha1.DataLakeStoreFirewallRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeStoreFirewallRule, err error)
	DataLakeStoreFirewallRuleExpansion
}

// dataLakeStoreFirewallRules implements DataLakeStoreFirewallRuleInterface
type dataLakeStoreFirewallRules struct {
	client rest.Interface
	ns     string
}

// newDataLakeStoreFirewallRules returns a DataLakeStoreFirewallRules
func newDataLakeStoreFirewallRules(c *AzurermV1alpha1Client, namespace string) *dataLakeStoreFirewallRules {
	return &dataLakeStoreFirewallRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataLakeStoreFirewallRule, and returns the corresponding dataLakeStoreFirewallRule object, and an error if there is any.
func (c *dataLakeStoreFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.DataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.DataLakeStoreFirewallRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataLakeStoreFirewallRules that match those selectors.
func (c *dataLakeStoreFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.DataLakeStoreFirewallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataLakeStoreFirewallRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataLakeStoreFirewallRules.
func (c *dataLakeStoreFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataLakeStoreFirewallRule and creates it.  Returns the server's representation of the dataLakeStoreFirewallRule, and an error, if there is any.
func (c *dataLakeStoreFirewallRules) Create(dataLakeStoreFirewallRule *v1alpha1.DataLakeStoreFirewallRule) (result *v1alpha1.DataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.DataLakeStoreFirewallRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		Body(dataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataLakeStoreFirewallRule and updates it. Returns the server's representation of the dataLakeStoreFirewallRule, and an error, if there is any.
func (c *dataLakeStoreFirewallRules) Update(dataLakeStoreFirewallRule *v1alpha1.DataLakeStoreFirewallRule) (result *v1alpha1.DataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.DataLakeStoreFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		Name(dataLakeStoreFirewallRule.Name).
		Body(dataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataLakeStoreFirewallRules) UpdateStatus(dataLakeStoreFirewallRule *v1alpha1.DataLakeStoreFirewallRule) (result *v1alpha1.DataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.DataLakeStoreFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		Name(dataLakeStoreFirewallRule.Name).
		SubResource("status").
		Body(dataLakeStoreFirewallRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataLakeStoreFirewallRule and deletes it. Returns an error if one occurs.
func (c *dataLakeStoreFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataLakeStoreFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataLakeStoreFirewallRule.
func (c *dataLakeStoreFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeStoreFirewallRule, err error) {
	result = &v1alpha1.DataLakeStoreFirewallRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datalakestorefirewallrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}