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

// RdsGlobalClustersGetter has a method to return a RdsGlobalClusterInterface.
// A group's client should implement this interface.
type RdsGlobalClustersGetter interface {
	RdsGlobalClusters(namespace string) RdsGlobalClusterInterface
}

// RdsGlobalClusterInterface has methods to work with RdsGlobalCluster resources.
type RdsGlobalClusterInterface interface {
	Create(*v1alpha1.RdsGlobalCluster) (*v1alpha1.RdsGlobalCluster, error)
	Update(*v1alpha1.RdsGlobalCluster) (*v1alpha1.RdsGlobalCluster, error)
	UpdateStatus(*v1alpha1.RdsGlobalCluster) (*v1alpha1.RdsGlobalCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RdsGlobalCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.RdsGlobalClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsGlobalCluster, err error)
	RdsGlobalClusterExpansion
}

// rdsGlobalClusters implements RdsGlobalClusterInterface
type rdsGlobalClusters struct {
	client rest.Interface
	ns     string
}

// newRdsGlobalClusters returns a RdsGlobalClusters
func newRdsGlobalClusters(c *AwsV1alpha1Client, namespace string) *rdsGlobalClusters {
	return &rdsGlobalClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rdsGlobalCluster, and returns the corresponding rdsGlobalCluster object, and an error if there is any.
func (c *rdsGlobalClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.RdsGlobalCluster, err error) {
	result = &v1alpha1.RdsGlobalCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RdsGlobalClusters that match those selectors.
func (c *rdsGlobalClusters) List(opts v1.ListOptions) (result *v1alpha1.RdsGlobalClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RdsGlobalClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rdsGlobalClusters.
func (c *rdsGlobalClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rdsGlobalCluster and creates it.  Returns the server's representation of the rdsGlobalCluster, and an error, if there is any.
func (c *rdsGlobalClusters) Create(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (result *v1alpha1.RdsGlobalCluster, err error) {
	result = &v1alpha1.RdsGlobalCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		Body(rdsGlobalCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rdsGlobalCluster and updates it. Returns the server's representation of the rdsGlobalCluster, and an error, if there is any.
func (c *rdsGlobalClusters) Update(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (result *v1alpha1.RdsGlobalCluster, err error) {
	result = &v1alpha1.RdsGlobalCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		Name(rdsGlobalCluster.Name).
		Body(rdsGlobalCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rdsGlobalClusters) UpdateStatus(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (result *v1alpha1.RdsGlobalCluster, err error) {
	result = &v1alpha1.RdsGlobalCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		Name(rdsGlobalCluster.Name).
		SubResource("status").
		Body(rdsGlobalCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the rdsGlobalCluster and deletes it. Returns an error if one occurs.
func (c *rdsGlobalClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rdsGlobalClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rdsGlobalCluster.
func (c *rdsGlobalClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsGlobalCluster, err error) {
	result = &v1alpha1.RdsGlobalCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rdsglobalclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
