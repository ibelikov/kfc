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

// ApiManagementAuthorizationServersGetter has a method to return a ApiManagementAuthorizationServerInterface.
// A group's client should implement this interface.
type ApiManagementAuthorizationServersGetter interface {
	ApiManagementAuthorizationServers(namespace string) ApiManagementAuthorizationServerInterface
}

// ApiManagementAuthorizationServerInterface has methods to work with ApiManagementAuthorizationServer resources.
type ApiManagementAuthorizationServerInterface interface {
	Create(*v1alpha1.ApiManagementAuthorizationServer) (*v1alpha1.ApiManagementAuthorizationServer, error)
	Update(*v1alpha1.ApiManagementAuthorizationServer) (*v1alpha1.ApiManagementAuthorizationServer, error)
	UpdateStatus(*v1alpha1.ApiManagementAuthorizationServer) (*v1alpha1.ApiManagementAuthorizationServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementAuthorizationServer, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementAuthorizationServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAuthorizationServer, err error)
	ApiManagementAuthorizationServerExpansion
}

// apiManagementAuthorizationServers implements ApiManagementAuthorizationServerInterface
type apiManagementAuthorizationServers struct {
	client rest.Interface
	ns     string
}

// newApiManagementAuthorizationServers returns a ApiManagementAuthorizationServers
func newApiManagementAuthorizationServers(c *AzurermV1alpha1Client, namespace string) *apiManagementAuthorizationServers {
	return &apiManagementAuthorizationServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementAuthorizationServer, and returns the corresponding apiManagementAuthorizationServer object, and an error if there is any.
func (c *apiManagementAuthorizationServers) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.ApiManagementAuthorizationServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementAuthorizationServers that match those selectors.
func (c *apiManagementAuthorizationServers) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementAuthorizationServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementAuthorizationServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementAuthorizationServers.
func (c *apiManagementAuthorizationServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementAuthorizationServer and creates it.  Returns the server's representation of the apiManagementAuthorizationServer, and an error, if there is any.
func (c *apiManagementAuthorizationServers) Create(apiManagementAuthorizationServer *v1alpha1.ApiManagementAuthorizationServer) (result *v1alpha1.ApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.ApiManagementAuthorizationServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		Body(apiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementAuthorizationServer and updates it. Returns the server's representation of the apiManagementAuthorizationServer, and an error, if there is any.
func (c *apiManagementAuthorizationServers) Update(apiManagementAuthorizationServer *v1alpha1.ApiManagementAuthorizationServer) (result *v1alpha1.ApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.ApiManagementAuthorizationServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		Name(apiManagementAuthorizationServer.Name).
		Body(apiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementAuthorizationServers) UpdateStatus(apiManagementAuthorizationServer *v1alpha1.ApiManagementAuthorizationServer) (result *v1alpha1.ApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.ApiManagementAuthorizationServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		Name(apiManagementAuthorizationServer.Name).
		SubResource("status").
		Body(apiManagementAuthorizationServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementAuthorizationServer and deletes it. Returns an error if one occurs.
func (c *apiManagementAuthorizationServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementAuthorizationServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementAuthorizationServer.
func (c *apiManagementAuthorizationServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementAuthorizationServer, err error) {
	result = &v1alpha1.ApiManagementAuthorizationServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementauthorizationservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
