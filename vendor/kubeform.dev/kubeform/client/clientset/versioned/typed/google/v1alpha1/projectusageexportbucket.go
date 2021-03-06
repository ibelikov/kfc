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

// ProjectUsageExportBucketsGetter has a method to return a ProjectUsageExportBucketInterface.
// A group's client should implement this interface.
type ProjectUsageExportBucketsGetter interface {
	ProjectUsageExportBuckets(namespace string) ProjectUsageExportBucketInterface
}

// ProjectUsageExportBucketInterface has methods to work with ProjectUsageExportBucket resources.
type ProjectUsageExportBucketInterface interface {
	Create(*v1alpha1.ProjectUsageExportBucket) (*v1alpha1.ProjectUsageExportBucket, error)
	Update(*v1alpha1.ProjectUsageExportBucket) (*v1alpha1.ProjectUsageExportBucket, error)
	UpdateStatus(*v1alpha1.ProjectUsageExportBucket) (*v1alpha1.ProjectUsageExportBucket, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ProjectUsageExportBucket, error)
	List(opts v1.ListOptions) (*v1alpha1.ProjectUsageExportBucketList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectUsageExportBucket, err error)
	ProjectUsageExportBucketExpansion
}

// projectUsageExportBuckets implements ProjectUsageExportBucketInterface
type projectUsageExportBuckets struct {
	client rest.Interface
	ns     string
}

// newProjectUsageExportBuckets returns a ProjectUsageExportBuckets
func newProjectUsageExportBuckets(c *GoogleV1alpha1Client, namespace string) *projectUsageExportBuckets {
	return &projectUsageExportBuckets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the projectUsageExportBucket, and returns the corresponding projectUsageExportBucket object, and an error if there is any.
func (c *projectUsageExportBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	result = &v1alpha1.ProjectUsageExportBucket{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProjectUsageExportBuckets that match those selectors.
func (c *projectUsageExportBuckets) List(opts v1.ListOptions) (result *v1alpha1.ProjectUsageExportBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProjectUsageExportBucketList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested projectUsageExportBuckets.
func (c *projectUsageExportBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a projectUsageExportBucket and creates it.  Returns the server's representation of the projectUsageExportBucket, and an error, if there is any.
func (c *projectUsageExportBuckets) Create(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	result = &v1alpha1.ProjectUsageExportBucket{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		Body(projectUsageExportBucket).
		Do().
		Into(result)
	return
}

// Update takes the representation of a projectUsageExportBucket and updates it. Returns the server's representation of the projectUsageExportBucket, and an error, if there is any.
func (c *projectUsageExportBuckets) Update(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	result = &v1alpha1.ProjectUsageExportBucket{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		Name(projectUsageExportBucket.Name).
		Body(projectUsageExportBucket).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *projectUsageExportBuckets) UpdateStatus(projectUsageExportBucket *v1alpha1.ProjectUsageExportBucket) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	result = &v1alpha1.ProjectUsageExportBucket{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		Name(projectUsageExportBucket.Name).
		SubResource("status").
		Body(projectUsageExportBucket).
		Do().
		Into(result)
	return
}

// Delete takes name of the projectUsageExportBucket and deletes it. Returns an error if one occurs.
func (c *projectUsageExportBuckets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *projectUsageExportBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched projectUsageExportBucket.
func (c *projectUsageExportBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectUsageExportBucket, err error) {
	result = &v1alpha1.ProjectUsageExportBucket{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("projectusageexportbuckets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
