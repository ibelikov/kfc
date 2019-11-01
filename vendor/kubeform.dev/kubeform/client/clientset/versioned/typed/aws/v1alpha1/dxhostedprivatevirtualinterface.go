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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DxHostedPrivateVirtualInterfacesGetter has a method to return a DxHostedPrivateVirtualInterfaceInterface.
// A group's client should implement this interface.
type DxHostedPrivateVirtualInterfacesGetter interface {
	DxHostedPrivateVirtualInterfaces(namespace string) DxHostedPrivateVirtualInterfaceInterface
}

// DxHostedPrivateVirtualInterfaceInterface has methods to work with DxHostedPrivateVirtualInterface resources.
type DxHostedPrivateVirtualInterfaceInterface interface {
	Create(*v1alpha1.DxHostedPrivateVirtualInterface) (*v1alpha1.DxHostedPrivateVirtualInterface, error)
	Update(*v1alpha1.DxHostedPrivateVirtualInterface) (*v1alpha1.DxHostedPrivateVirtualInterface, error)
	UpdateStatus(*v1alpha1.DxHostedPrivateVirtualInterface) (*v1alpha1.DxHostedPrivateVirtualInterface, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DxHostedPrivateVirtualInterface, error)
	List(opts v1.ListOptions) (*v1alpha1.DxHostedPrivateVirtualInterfaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error)
	DxHostedPrivateVirtualInterfaceExpansion
}

// dxHostedPrivateVirtualInterfaces implements DxHostedPrivateVirtualInterfaceInterface
type dxHostedPrivateVirtualInterfaces struct {
	client rest.Interface
	ns     string
}

// newDxHostedPrivateVirtualInterfaces returns a DxHostedPrivateVirtualInterfaces
func newDxHostedPrivateVirtualInterfaces(c *AwsV1alpha1Client, namespace string) *dxHostedPrivateVirtualInterfaces {
	return &dxHostedPrivateVirtualInterfaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dxHostedPrivateVirtualInterface, and returns the corresponding dxHostedPrivateVirtualInterface object, and an error if there is any.
func (c *dxHostedPrivateVirtualInterfaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPrivateVirtualInterface{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DxHostedPrivateVirtualInterfaces that match those selectors.
func (c *dxHostedPrivateVirtualInterfaces) List(opts v1.ListOptions) (result *v1alpha1.DxHostedPrivateVirtualInterfaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DxHostedPrivateVirtualInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dxHostedPrivateVirtualInterfaces.
func (c *dxHostedPrivateVirtualInterfaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dxHostedPrivateVirtualInterface and creates it.  Returns the server's representation of the dxHostedPrivateVirtualInterface, and an error, if there is any.
func (c *dxHostedPrivateVirtualInterfaces) Create(dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPrivateVirtualInterface{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		Body(dxHostedPrivateVirtualInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dxHostedPrivateVirtualInterface and updates it. Returns the server's representation of the dxHostedPrivateVirtualInterface, and an error, if there is any.
func (c *dxHostedPrivateVirtualInterfaces) Update(dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPrivateVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		Name(dxHostedPrivateVirtualInterface.Name).
		Body(dxHostedPrivateVirtualInterface).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dxHostedPrivateVirtualInterfaces) UpdateStatus(dxHostedPrivateVirtualInterface *v1alpha1.DxHostedPrivateVirtualInterface) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPrivateVirtualInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		Name(dxHostedPrivateVirtualInterface.Name).
		SubResource("status").
		Body(dxHostedPrivateVirtualInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the dxHostedPrivateVirtualInterface and deletes it. Returns an error if one occurs.
func (c *dxHostedPrivateVirtualInterfaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dxHostedPrivateVirtualInterfaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dxHostedPrivateVirtualInterface.
func (c *dxHostedPrivateVirtualInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxHostedPrivateVirtualInterface, err error) {
	result = &v1alpha1.DxHostedPrivateVirtualInterface{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dxhostedprivatevirtualinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}