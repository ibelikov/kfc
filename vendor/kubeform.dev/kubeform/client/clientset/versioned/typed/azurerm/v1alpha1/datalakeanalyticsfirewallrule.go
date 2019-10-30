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

// DataLakeAnalyticsFirewallRulesGetter has a method to return a DataLakeAnalyticsFirewallRuleInterface.
// A group's client should implement this interface.
type DataLakeAnalyticsFirewallRulesGetter interface {
	DataLakeAnalyticsFirewallRules(namespace string) DataLakeAnalyticsFirewallRuleInterface
}

// DataLakeAnalyticsFirewallRuleInterface has methods to work with DataLakeAnalyticsFirewallRule resources.
type DataLakeAnalyticsFirewallRuleInterface interface {
	Create(*v1alpha1.DataLakeAnalyticsFirewallRule) (*v1alpha1.DataLakeAnalyticsFirewallRule, error)
	Update(*v1alpha1.DataLakeAnalyticsFirewallRule) (*v1alpha1.DataLakeAnalyticsFirewallRule, error)
	UpdateStatus(*v1alpha1.DataLakeAnalyticsFirewallRule) (*v1alpha1.DataLakeAnalyticsFirewallRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataLakeAnalyticsFirewallRule, error)
	List(opts v1.ListOptions) (*v1alpha1.DataLakeAnalyticsFirewallRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error)
	DataLakeAnalyticsFirewallRuleExpansion
}

// dataLakeAnalyticsFirewallRules implements DataLakeAnalyticsFirewallRuleInterface
type dataLakeAnalyticsFirewallRules struct {
	client rest.Interface
	ns     string
}

// newDataLakeAnalyticsFirewallRules returns a DataLakeAnalyticsFirewallRules
func newDataLakeAnalyticsFirewallRules(c *AzurermV1alpha1Client, namespace string) *dataLakeAnalyticsFirewallRules {
	return &dataLakeAnalyticsFirewallRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataLakeAnalyticsFirewallRule, and returns the corresponding dataLakeAnalyticsFirewallRule object, and an error if there is any.
func (c *dataLakeAnalyticsFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	result = &v1alpha1.DataLakeAnalyticsFirewallRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataLakeAnalyticsFirewallRules that match those selectors.
func (c *dataLakeAnalyticsFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.DataLakeAnalyticsFirewallRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataLakeAnalyticsFirewallRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataLakeAnalyticsFirewallRules.
func (c *dataLakeAnalyticsFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataLakeAnalyticsFirewallRule and creates it.  Returns the server's representation of the dataLakeAnalyticsFirewallRule, and an error, if there is any.
func (c *dataLakeAnalyticsFirewallRules) Create(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	result = &v1alpha1.DataLakeAnalyticsFirewallRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		Body(dataLakeAnalyticsFirewallRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataLakeAnalyticsFirewallRule and updates it. Returns the server's representation of the dataLakeAnalyticsFirewallRule, and an error, if there is any.
func (c *dataLakeAnalyticsFirewallRules) Update(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	result = &v1alpha1.DataLakeAnalyticsFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		Name(dataLakeAnalyticsFirewallRule.Name).
		Body(dataLakeAnalyticsFirewallRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataLakeAnalyticsFirewallRules) UpdateStatus(dataLakeAnalyticsFirewallRule *v1alpha1.DataLakeAnalyticsFirewallRule) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	result = &v1alpha1.DataLakeAnalyticsFirewallRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		Name(dataLakeAnalyticsFirewallRule.Name).
		SubResource("status").
		Body(dataLakeAnalyticsFirewallRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataLakeAnalyticsFirewallRule and deletes it. Returns an error if one occurs.
func (c *dataLakeAnalyticsFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataLakeAnalyticsFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataLakeAnalyticsFirewallRule.
func (c *dataLakeAnalyticsFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataLakeAnalyticsFirewallRule, err error) {
	result = &v1alpha1.DataLakeAnalyticsFirewallRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datalakeanalyticsfirewallrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
