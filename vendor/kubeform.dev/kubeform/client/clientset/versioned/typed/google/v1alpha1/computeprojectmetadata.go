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

// ComputeProjectMetadatasGetter has a method to return a ComputeProjectMetadataInterface.
// A group's client should implement this interface.
type ComputeProjectMetadatasGetter interface {
	ComputeProjectMetadatas(namespace string) ComputeProjectMetadataInterface
}

// ComputeProjectMetadataInterface has methods to work with ComputeProjectMetadata resources.
type ComputeProjectMetadataInterface interface {
	Create(*v1alpha1.ComputeProjectMetadata) (*v1alpha1.ComputeProjectMetadata, error)
	Update(*v1alpha1.ComputeProjectMetadata) (*v1alpha1.ComputeProjectMetadata, error)
	UpdateStatus(*v1alpha1.ComputeProjectMetadata) (*v1alpha1.ComputeProjectMetadata, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeProjectMetadata, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeProjectMetadataList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeProjectMetadata, err error)
	ComputeProjectMetadataExpansion
}

// computeProjectMetadatas implements ComputeProjectMetadataInterface
type computeProjectMetadatas struct {
	client rest.Interface
	ns     string
}

// newComputeProjectMetadatas returns a ComputeProjectMetadatas
func newComputeProjectMetadatas(c *GoogleV1alpha1Client, namespace string) *computeProjectMetadatas {
	return &computeProjectMetadatas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeProjectMetadata, and returns the corresponding computeProjectMetadata object, and an error if there is any.
func (c *computeProjectMetadatas) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeProjectMetadata, err error) {
	result = &v1alpha1.ComputeProjectMetadata{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeProjectMetadatas that match those selectors.
func (c *computeProjectMetadatas) List(opts v1.ListOptions) (result *v1alpha1.ComputeProjectMetadataList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeProjectMetadataList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeProjectMetadatas.
func (c *computeProjectMetadatas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeProjectMetadata and creates it.  Returns the server's representation of the computeProjectMetadata, and an error, if there is any.
func (c *computeProjectMetadatas) Create(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (result *v1alpha1.ComputeProjectMetadata, err error) {
	result = &v1alpha1.ComputeProjectMetadata{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		Body(computeProjectMetadata).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeProjectMetadata and updates it. Returns the server's representation of the computeProjectMetadata, and an error, if there is any.
func (c *computeProjectMetadatas) Update(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (result *v1alpha1.ComputeProjectMetadata, err error) {
	result = &v1alpha1.ComputeProjectMetadata{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		Name(computeProjectMetadata.Name).
		Body(computeProjectMetadata).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeProjectMetadatas) UpdateStatus(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (result *v1alpha1.ComputeProjectMetadata, err error) {
	result = &v1alpha1.ComputeProjectMetadata{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		Name(computeProjectMetadata.Name).
		SubResource("status").
		Body(computeProjectMetadata).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeProjectMetadata and deletes it. Returns an error if one occurs.
func (c *computeProjectMetadatas) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeProjectMetadatas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeProjectMetadata.
func (c *computeProjectMetadatas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeProjectMetadata, err error) {
	result = &v1alpha1.ComputeProjectMetadata{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeprojectmetadatas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
