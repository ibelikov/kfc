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

// EmrClustersGetter has a method to return a EmrClusterInterface.
// A group's client should implement this interface.
type EmrClustersGetter interface {
	EmrClusters(namespace string) EmrClusterInterface
}

// EmrClusterInterface has methods to work with EmrCluster resources.
type EmrClusterInterface interface {
	Create(*v1alpha1.EmrCluster) (*v1alpha1.EmrCluster, error)
	Update(*v1alpha1.EmrCluster) (*v1alpha1.EmrCluster, error)
	UpdateStatus(*v1alpha1.EmrCluster) (*v1alpha1.EmrCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EmrCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.EmrClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EmrCluster, err error)
	EmrClusterExpansion
}

// emrClusters implements EmrClusterInterface
type emrClusters struct {
	client rest.Interface
	ns     string
}

// newEmrClusters returns a EmrClusters
func newEmrClusters(c *AwsV1alpha1Client, namespace string) *emrClusters {
	return &emrClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the emrCluster, and returns the corresponding emrCluster object, and an error if there is any.
func (c *emrClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.EmrCluster, err error) {
	result = &v1alpha1.EmrCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emrclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EmrClusters that match those selectors.
func (c *emrClusters) List(opts v1.ListOptions) (result *v1alpha1.EmrClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EmrClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("emrclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested emrClusters.
func (c *emrClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("emrclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a emrCluster and creates it.  Returns the server's representation of the emrCluster, and an error, if there is any.
func (c *emrClusters) Create(emrCluster *v1alpha1.EmrCluster) (result *v1alpha1.EmrCluster, err error) {
	result = &v1alpha1.EmrCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("emrclusters").
		Body(emrCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a emrCluster and updates it. Returns the server's representation of the emrCluster, and an error, if there is any.
func (c *emrClusters) Update(emrCluster *v1alpha1.EmrCluster) (result *v1alpha1.EmrCluster, err error) {
	result = &v1alpha1.EmrCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("emrclusters").
		Name(emrCluster.Name).
		Body(emrCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *emrClusters) UpdateStatus(emrCluster *v1alpha1.EmrCluster) (result *v1alpha1.EmrCluster, err error) {
	result = &v1alpha1.EmrCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("emrclusters").
		Name(emrCluster.Name).
		SubResource("status").
		Body(emrCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the emrCluster and deletes it. Returns an error if one occurs.
func (c *emrClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("emrclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *emrClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("emrclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched emrCluster.
func (c *emrClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EmrCluster, err error) {
	result = &v1alpha1.EmrCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("emrclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
