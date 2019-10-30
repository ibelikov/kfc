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

// AutomationRunbooksGetter has a method to return a AutomationRunbookInterface.
// A group's client should implement this interface.
type AutomationRunbooksGetter interface {
	AutomationRunbooks(namespace string) AutomationRunbookInterface
}

// AutomationRunbookInterface has methods to work with AutomationRunbook resources.
type AutomationRunbookInterface interface {
	Create(*v1alpha1.AutomationRunbook) (*v1alpha1.AutomationRunbook, error)
	Update(*v1alpha1.AutomationRunbook) (*v1alpha1.AutomationRunbook, error)
	UpdateStatus(*v1alpha1.AutomationRunbook) (*v1alpha1.AutomationRunbook, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutomationRunbook, error)
	List(opts v1.ListOptions) (*v1alpha1.AutomationRunbookList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationRunbook, err error)
	AutomationRunbookExpansion
}

// automationRunbooks implements AutomationRunbookInterface
type automationRunbooks struct {
	client rest.Interface
	ns     string
}

// newAutomationRunbooks returns a AutomationRunbooks
func newAutomationRunbooks(c *AzurermV1alpha1Client, namespace string) *automationRunbooks {
	return &automationRunbooks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the automationRunbook, and returns the corresponding automationRunbook object, and an error if there is any.
func (c *automationRunbooks) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationRunbook, err error) {
	result = &v1alpha1.AutomationRunbook{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationrunbooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutomationRunbooks that match those selectors.
func (c *automationRunbooks) List(opts v1.ListOptions) (result *v1alpha1.AutomationRunbookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutomationRunbookList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationrunbooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested automationRunbooks.
func (c *automationRunbooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("automationrunbooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a automationRunbook and creates it.  Returns the server's representation of the automationRunbook, and an error, if there is any.
func (c *automationRunbooks) Create(automationRunbook *v1alpha1.AutomationRunbook) (result *v1alpha1.AutomationRunbook, err error) {
	result = &v1alpha1.AutomationRunbook{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("automationrunbooks").
		Body(automationRunbook).
		Do().
		Into(result)
	return
}

// Update takes the representation of a automationRunbook and updates it. Returns the server's representation of the automationRunbook, and an error, if there is any.
func (c *automationRunbooks) Update(automationRunbook *v1alpha1.AutomationRunbook) (result *v1alpha1.AutomationRunbook, err error) {
	result = &v1alpha1.AutomationRunbook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationrunbooks").
		Name(automationRunbook.Name).
		Body(automationRunbook).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *automationRunbooks) UpdateStatus(automationRunbook *v1alpha1.AutomationRunbook) (result *v1alpha1.AutomationRunbook, err error) {
	result = &v1alpha1.AutomationRunbook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationrunbooks").
		Name(automationRunbook.Name).
		SubResource("status").
		Body(automationRunbook).
		Do().
		Into(result)
	return
}

// Delete takes name of the automationRunbook and deletes it. Returns an error if one occurs.
func (c *automationRunbooks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationrunbooks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *automationRunbooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationrunbooks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched automationRunbook.
func (c *automationRunbooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationRunbook, err error) {
	result = &v1alpha1.AutomationRunbook{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("automationrunbooks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
