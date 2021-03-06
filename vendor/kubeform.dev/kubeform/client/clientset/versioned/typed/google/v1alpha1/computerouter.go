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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeRoutersGetter has a method to return a ComputeRouterInterface.
// A group's client should implement this interface.
type ComputeRoutersGetter interface {
	ComputeRouters(namespace string) ComputeRouterInterface
}

// ComputeRouterInterface has methods to work with ComputeRouter resources.
type ComputeRouterInterface interface {
	Create(*v1alpha1.ComputeRouter) (*v1alpha1.ComputeRouter, error)
	Update(*v1alpha1.ComputeRouter) (*v1alpha1.ComputeRouter, error)
	UpdateStatus(*v1alpha1.ComputeRouter) (*v1alpha1.ComputeRouter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRouter, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRouterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouter, err error)
	ComputeRouterExpansion
}

// computeRouters implements ComputeRouterInterface
type computeRouters struct {
	client rest.Interface
	ns     string
}

// newComputeRouters returns a ComputeRouters
func newComputeRouters(c *GoogleV1alpha1Client, namespace string) *computeRouters {
	return &computeRouters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeRouter, and returns the corresponding computeRouter object, and an error if there is any.
func (c *computeRouters) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRouter, err error) {
	result = &v1alpha1.ComputeRouter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computerouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRouters that match those selectors.
func (c *computeRouters) List(opts v1.ListOptions) (result *v1alpha1.ComputeRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computerouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRouters.
func (c *computeRouters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computerouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRouter and creates it.  Returns the server's representation of the computeRouter, and an error, if there is any.
func (c *computeRouters) Create(computeRouter *v1alpha1.ComputeRouter) (result *v1alpha1.ComputeRouter, err error) {
	result = &v1alpha1.ComputeRouter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computerouters").
		Body(computeRouter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRouter and updates it. Returns the server's representation of the computeRouter, and an error, if there is any.
func (c *computeRouters) Update(computeRouter *v1alpha1.ComputeRouter) (result *v1alpha1.ComputeRouter, err error) {
	result = &v1alpha1.ComputeRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computerouters").
		Name(computeRouter.Name).
		Body(computeRouter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRouters) UpdateStatus(computeRouter *v1alpha1.ComputeRouter) (result *v1alpha1.ComputeRouter, err error) {
	result = &v1alpha1.ComputeRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computerouters").
		Name(computeRouter.Name).
		SubResource("status").
		Body(computeRouter).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRouter and deletes it. Returns an error if one occurs.
func (c *computeRouters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computerouters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRouters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computerouters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRouter.
func (c *computeRouters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRouter, err error) {
	result = &v1alpha1.ComputeRouter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computerouters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
