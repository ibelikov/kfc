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

// NotificationHubAuthorizationRulesGetter has a method to return a NotificationHubAuthorizationRuleInterface.
// A group's client should implement this interface.
type NotificationHubAuthorizationRulesGetter interface {
	NotificationHubAuthorizationRules(namespace string) NotificationHubAuthorizationRuleInterface
}

// NotificationHubAuthorizationRuleInterface has methods to work with NotificationHubAuthorizationRule resources.
type NotificationHubAuthorizationRuleInterface interface {
	Create(*v1alpha1.NotificationHubAuthorizationRule) (*v1alpha1.NotificationHubAuthorizationRule, error)
	Update(*v1alpha1.NotificationHubAuthorizationRule) (*v1alpha1.NotificationHubAuthorizationRule, error)
	UpdateStatus(*v1alpha1.NotificationHubAuthorizationRule) (*v1alpha1.NotificationHubAuthorizationRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NotificationHubAuthorizationRule, error)
	List(opts v1.ListOptions) (*v1alpha1.NotificationHubAuthorizationRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHubAuthorizationRule, err error)
	NotificationHubAuthorizationRuleExpansion
}

// notificationHubAuthorizationRules implements NotificationHubAuthorizationRuleInterface
type notificationHubAuthorizationRules struct {
	client rest.Interface
	ns     string
}

// newNotificationHubAuthorizationRules returns a NotificationHubAuthorizationRules
func newNotificationHubAuthorizationRules(c *AzurermV1alpha1Client, namespace string) *notificationHubAuthorizationRules {
	return &notificationHubAuthorizationRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notificationHubAuthorizationRule, and returns the corresponding notificationHubAuthorizationRule object, and an error if there is any.
func (c *notificationHubAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	result = &v1alpha1.NotificationHubAuthorizationRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NotificationHubAuthorizationRules that match those selectors.
func (c *notificationHubAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.NotificationHubAuthorizationRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NotificationHubAuthorizationRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notificationHubAuthorizationRules.
func (c *notificationHubAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a notificationHubAuthorizationRule and creates it.  Returns the server's representation of the notificationHubAuthorizationRule, and an error, if there is any.
func (c *notificationHubAuthorizationRules) Create(notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	result = &v1alpha1.NotificationHubAuthorizationRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		Body(notificationHubAuthorizationRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a notificationHubAuthorizationRule and updates it. Returns the server's representation of the notificationHubAuthorizationRule, and an error, if there is any.
func (c *notificationHubAuthorizationRules) Update(notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	result = &v1alpha1.NotificationHubAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		Name(notificationHubAuthorizationRule.Name).
		Body(notificationHubAuthorizationRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *notificationHubAuthorizationRules) UpdateStatus(notificationHubAuthorizationRule *v1alpha1.NotificationHubAuthorizationRule) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	result = &v1alpha1.NotificationHubAuthorizationRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		Name(notificationHubAuthorizationRule.Name).
		SubResource("status").
		Body(notificationHubAuthorizationRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the notificationHubAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *notificationHubAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notificationHubAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched notificationHubAuthorizationRule.
func (c *notificationHubAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHubAuthorizationRule, err error) {
	result = &v1alpha1.NotificationHubAuthorizationRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notificationhubauthorizationrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
