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

// ApiGatewayDomainNamesGetter has a method to return a ApiGatewayDomainNameInterface.
// A group's client should implement this interface.
type ApiGatewayDomainNamesGetter interface {
	ApiGatewayDomainNames(namespace string) ApiGatewayDomainNameInterface
}

// ApiGatewayDomainNameInterface has methods to work with ApiGatewayDomainName resources.
type ApiGatewayDomainNameInterface interface {
	Create(*v1alpha1.ApiGatewayDomainName) (*v1alpha1.ApiGatewayDomainName, error)
	Update(*v1alpha1.ApiGatewayDomainName) (*v1alpha1.ApiGatewayDomainName, error)
	UpdateStatus(*v1alpha1.ApiGatewayDomainName) (*v1alpha1.ApiGatewayDomainName, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiGatewayDomainName, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiGatewayDomainNameList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayDomainName, err error)
	ApiGatewayDomainNameExpansion
}

// apiGatewayDomainNames implements ApiGatewayDomainNameInterface
type apiGatewayDomainNames struct {
	client rest.Interface
	ns     string
}

// newApiGatewayDomainNames returns a ApiGatewayDomainNames
func newApiGatewayDomainNames(c *AwsV1alpha1Client, namespace string) *apiGatewayDomainNames {
	return &apiGatewayDomainNames{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiGatewayDomainName, and returns the corresponding apiGatewayDomainName object, and an error if there is any.
func (c *apiGatewayDomainNames) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayDomainName, err error) {
	result = &v1alpha1.ApiGatewayDomainName{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayDomainNames that match those selectors.
func (c *apiGatewayDomainNames) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayDomainNameList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayDomainNameList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayDomainNames.
func (c *apiGatewayDomainNames) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiGatewayDomainName and creates it.  Returns the server's representation of the apiGatewayDomainName, and an error, if there is any.
func (c *apiGatewayDomainNames) Create(apiGatewayDomainName *v1alpha1.ApiGatewayDomainName) (result *v1alpha1.ApiGatewayDomainName, err error) {
	result = &v1alpha1.ApiGatewayDomainName{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		Body(apiGatewayDomainName).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiGatewayDomainName and updates it. Returns the server's representation of the apiGatewayDomainName, and an error, if there is any.
func (c *apiGatewayDomainNames) Update(apiGatewayDomainName *v1alpha1.ApiGatewayDomainName) (result *v1alpha1.ApiGatewayDomainName, err error) {
	result = &v1alpha1.ApiGatewayDomainName{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		Name(apiGatewayDomainName.Name).
		Body(apiGatewayDomainName).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiGatewayDomainNames) UpdateStatus(apiGatewayDomainName *v1alpha1.ApiGatewayDomainName) (result *v1alpha1.ApiGatewayDomainName, err error) {
	result = &v1alpha1.ApiGatewayDomainName{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		Name(apiGatewayDomainName.Name).
		SubResource("status").
		Body(apiGatewayDomainName).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiGatewayDomainName and deletes it. Returns an error if one occurs.
func (c *apiGatewayDomainNames) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayDomainNames) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiGatewayDomainName.
func (c *apiGatewayDomainNames) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayDomainName, err error) {
	result = &v1alpha1.ApiGatewayDomainName{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apigatewaydomainnames").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
