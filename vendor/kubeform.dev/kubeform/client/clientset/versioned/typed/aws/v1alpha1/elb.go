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

// ElbsGetter has a method to return a ElbInterface.
// A group's client should implement this interface.
type ElbsGetter interface {
	Elbs(namespace string) ElbInterface
}

// ElbInterface has methods to work with Elb resources.
type ElbInterface interface {
	Create(*v1alpha1.Elb) (*v1alpha1.Elb, error)
	Update(*v1alpha1.Elb) (*v1alpha1.Elb, error)
	UpdateStatus(*v1alpha1.Elb) (*v1alpha1.Elb, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Elb, error)
	List(opts v1.ListOptions) (*v1alpha1.ElbList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elb, err error)
	ElbExpansion
}

// elbs implements ElbInterface
type elbs struct {
	client rest.Interface
	ns     string
}

// newElbs returns a Elbs
func newElbs(c *AwsV1alpha1Client, namespace string) *elbs {
	return &elbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the elb, and returns the corresponding elb object, and an error if there is any.
func (c *elbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Elb, err error) {
	result = &v1alpha1.Elb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Elbs that match those selectors.
func (c *elbs) List(opts v1.ListOptions) (result *v1alpha1.ElbList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("elbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elbs.
func (c *elbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("elbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elb and creates it.  Returns the server's representation of the elb, and an error, if there is any.
func (c *elbs) Create(elb *v1alpha1.Elb) (result *v1alpha1.Elb, err error) {
	result = &v1alpha1.Elb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("elbs").
		Body(elb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elb and updates it. Returns the server's representation of the elb, and an error, if there is any.
func (c *elbs) Update(elb *v1alpha1.Elb) (result *v1alpha1.Elb, err error) {
	result = &v1alpha1.Elb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elbs").
		Name(elb.Name).
		Body(elb).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elbs) UpdateStatus(elb *v1alpha1.Elb) (result *v1alpha1.Elb, err error) {
	result = &v1alpha1.Elb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("elbs").
		Name(elb.Name).
		SubResource("status").
		Body(elb).
		Do().
		Into(result)
	return
}

// Delete takes name of the elb and deletes it. Returns an error if one occurs.
func (c *elbs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("elbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elb.
func (c *elbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Elb, err error) {
	result = &v1alpha1.Elb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("elbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
