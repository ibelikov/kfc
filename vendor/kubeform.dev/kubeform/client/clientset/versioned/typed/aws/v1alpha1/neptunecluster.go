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

// NeptuneClustersGetter has a method to return a NeptuneClusterInterface.
// A group's client should implement this interface.
type NeptuneClustersGetter interface {
	NeptuneClusters(namespace string) NeptuneClusterInterface
}

// NeptuneClusterInterface has methods to work with NeptuneCluster resources.
type NeptuneClusterInterface interface {
	Create(*v1alpha1.NeptuneCluster) (*v1alpha1.NeptuneCluster, error)
	Update(*v1alpha1.NeptuneCluster) (*v1alpha1.NeptuneCluster, error)
	UpdateStatus(*v1alpha1.NeptuneCluster) (*v1alpha1.NeptuneCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NeptuneCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.NeptuneClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneCluster, err error)
	NeptuneClusterExpansion
}

// neptuneClusters implements NeptuneClusterInterface
type neptuneClusters struct {
	client rest.Interface
	ns     string
}

// newNeptuneClusters returns a NeptuneClusters
func newNeptuneClusters(c *AwsV1alpha1Client, namespace string) *neptuneClusters {
	return &neptuneClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the neptuneCluster, and returns the corresponding neptuneCluster object, and an error if there is any.
func (c *neptuneClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.NeptuneCluster, err error) {
	result = &v1alpha1.NeptuneCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NeptuneClusters that match those selectors.
func (c *neptuneClusters) List(opts v1.ListOptions) (result *v1alpha1.NeptuneClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NeptuneClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested neptuneClusters.
func (c *neptuneClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a neptuneCluster and creates it.  Returns the server's representation of the neptuneCluster, and an error, if there is any.
func (c *neptuneClusters) Create(neptuneCluster *v1alpha1.NeptuneCluster) (result *v1alpha1.NeptuneCluster, err error) {
	result = &v1alpha1.NeptuneCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("neptuneclusters").
		Body(neptuneCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a neptuneCluster and updates it. Returns the server's representation of the neptuneCluster, and an error, if there is any.
func (c *neptuneClusters) Update(neptuneCluster *v1alpha1.NeptuneCluster) (result *v1alpha1.NeptuneCluster, err error) {
	result = &v1alpha1.NeptuneCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusters").
		Name(neptuneCluster.Name).
		Body(neptuneCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *neptuneClusters) UpdateStatus(neptuneCluster *v1alpha1.NeptuneCluster) (result *v1alpha1.NeptuneCluster, err error) {
	result = &v1alpha1.NeptuneCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusters").
		Name(neptuneCluster.Name).
		SubResource("status").
		Body(neptuneCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the neptuneCluster and deletes it. Returns an error if one occurs.
func (c *neptuneClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *neptuneClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched neptuneCluster.
func (c *neptuneClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneCluster, err error) {
	result = &v1alpha1.NeptuneCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("neptuneclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
