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

// SpotInstanceRequestsGetter has a method to return a SpotInstanceRequestInterface.
// A group's client should implement this interface.
type SpotInstanceRequestsGetter interface {
	SpotInstanceRequests(namespace string) SpotInstanceRequestInterface
}

// SpotInstanceRequestInterface has methods to work with SpotInstanceRequest resources.
type SpotInstanceRequestInterface interface {
	Create(*v1alpha1.SpotInstanceRequest) (*v1alpha1.SpotInstanceRequest, error)
	Update(*v1alpha1.SpotInstanceRequest) (*v1alpha1.SpotInstanceRequest, error)
	UpdateStatus(*v1alpha1.SpotInstanceRequest) (*v1alpha1.SpotInstanceRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpotInstanceRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.SpotInstanceRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotInstanceRequest, err error)
	SpotInstanceRequestExpansion
}

// spotInstanceRequests implements SpotInstanceRequestInterface
type spotInstanceRequests struct {
	client rest.Interface
	ns     string
}

// newSpotInstanceRequests returns a SpotInstanceRequests
func newSpotInstanceRequests(c *AwsV1alpha1Client, namespace string) *spotInstanceRequests {
	return &spotInstanceRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spotInstanceRequest, and returns the corresponding spotInstanceRequest object, and an error if there is any.
func (c *spotInstanceRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.SpotInstanceRequest, err error) {
	result = &v1alpha1.SpotInstanceRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpotInstanceRequests that match those selectors.
func (c *spotInstanceRequests) List(opts v1.ListOptions) (result *v1alpha1.SpotInstanceRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpotInstanceRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spotInstanceRequests.
func (c *spotInstanceRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spotInstanceRequest and creates it.  Returns the server's representation of the spotInstanceRequest, and an error, if there is any.
func (c *spotInstanceRequests) Create(spotInstanceRequest *v1alpha1.SpotInstanceRequest) (result *v1alpha1.SpotInstanceRequest, err error) {
	result = &v1alpha1.SpotInstanceRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		Body(spotInstanceRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spotInstanceRequest and updates it. Returns the server's representation of the spotInstanceRequest, and an error, if there is any.
func (c *spotInstanceRequests) Update(spotInstanceRequest *v1alpha1.SpotInstanceRequest) (result *v1alpha1.SpotInstanceRequest, err error) {
	result = &v1alpha1.SpotInstanceRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		Name(spotInstanceRequest.Name).
		Body(spotInstanceRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spotInstanceRequests) UpdateStatus(spotInstanceRequest *v1alpha1.SpotInstanceRequest) (result *v1alpha1.SpotInstanceRequest, err error) {
	result = &v1alpha1.SpotInstanceRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		Name(spotInstanceRequest.Name).
		SubResource("status").
		Body(spotInstanceRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the spotInstanceRequest and deletes it. Returns an error if one occurs.
func (c *spotInstanceRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spotInstanceRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spotinstancerequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spotInstanceRequest.
func (c *spotInstanceRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotInstanceRequest, err error) {
	result = &v1alpha1.SpotInstanceRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spotinstancerequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}