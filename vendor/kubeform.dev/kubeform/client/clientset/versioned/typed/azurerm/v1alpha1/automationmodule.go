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

// AutomationModulesGetter has a method to return a AutomationModuleInterface.
// A group's client should implement this interface.
type AutomationModulesGetter interface {
	AutomationModules(namespace string) AutomationModuleInterface
}

// AutomationModuleInterface has methods to work with AutomationModule resources.
type AutomationModuleInterface interface {
	Create(*v1alpha1.AutomationModule) (*v1alpha1.AutomationModule, error)
	Update(*v1alpha1.AutomationModule) (*v1alpha1.AutomationModule, error)
	UpdateStatus(*v1alpha1.AutomationModule) (*v1alpha1.AutomationModule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AutomationModule, error)
	List(opts v1.ListOptions) (*v1alpha1.AutomationModuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationModule, err error)
	AutomationModuleExpansion
}

// automationModules implements AutomationModuleInterface
type automationModules struct {
	client rest.Interface
	ns     string
}

// newAutomationModules returns a AutomationModules
func newAutomationModules(c *AzurermV1alpha1Client, namespace string) *automationModules {
	return &automationModules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the automationModule, and returns the corresponding automationModule object, and an error if there is any.
func (c *automationModules) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationModule, err error) {
	result = &v1alpha1.AutomationModule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationmodules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AutomationModules that match those selectors.
func (c *automationModules) List(opts v1.ListOptions) (result *v1alpha1.AutomationModuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AutomationModuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("automationmodules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested automationModules.
func (c *automationModules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("automationmodules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a automationModule and creates it.  Returns the server's representation of the automationModule, and an error, if there is any.
func (c *automationModules) Create(automationModule *v1alpha1.AutomationModule) (result *v1alpha1.AutomationModule, err error) {
	result = &v1alpha1.AutomationModule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("automationmodules").
		Body(automationModule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a automationModule and updates it. Returns the server's representation of the automationModule, and an error, if there is any.
func (c *automationModules) Update(automationModule *v1alpha1.AutomationModule) (result *v1alpha1.AutomationModule, err error) {
	result = &v1alpha1.AutomationModule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationmodules").
		Name(automationModule.Name).
		Body(automationModule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *automationModules) UpdateStatus(automationModule *v1alpha1.AutomationModule) (result *v1alpha1.AutomationModule, err error) {
	result = &v1alpha1.AutomationModule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("automationmodules").
		Name(automationModule.Name).
		SubResource("status").
		Body(automationModule).
		Do().
		Into(result)
	return
}

// Delete takes name of the automationModule and deletes it. Returns an error if one occurs.
func (c *automationModules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationmodules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *automationModules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("automationmodules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched automationModule.
func (c *automationModules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationModule, err error) {
	result = &v1alpha1.AutomationModule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("automationmodules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
