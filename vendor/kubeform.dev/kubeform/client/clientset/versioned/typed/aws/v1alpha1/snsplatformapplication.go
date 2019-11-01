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

// SnsPlatformApplicationsGetter has a method to return a SnsPlatformApplicationInterface.
// A group's client should implement this interface.
type SnsPlatformApplicationsGetter interface {
	SnsPlatformApplications(namespace string) SnsPlatformApplicationInterface
}

// SnsPlatformApplicationInterface has methods to work with SnsPlatformApplication resources.
type SnsPlatformApplicationInterface interface {
	Create(*v1alpha1.SnsPlatformApplication) (*v1alpha1.SnsPlatformApplication, error)
	Update(*v1alpha1.SnsPlatformApplication) (*v1alpha1.SnsPlatformApplication, error)
	UpdateStatus(*v1alpha1.SnsPlatformApplication) (*v1alpha1.SnsPlatformApplication, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SnsPlatformApplication, error)
	List(opts v1.ListOptions) (*v1alpha1.SnsPlatformApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsPlatformApplication, err error)
	SnsPlatformApplicationExpansion
}

// snsPlatformApplications implements SnsPlatformApplicationInterface
type snsPlatformApplications struct {
	client rest.Interface
	ns     string
}

// newSnsPlatformApplications returns a SnsPlatformApplications
func newSnsPlatformApplications(c *AwsV1alpha1Client, namespace string) *snsPlatformApplications {
	return &snsPlatformApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the snsPlatformApplication, and returns the corresponding snsPlatformApplication object, and an error if there is any.
func (c *snsPlatformApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.SnsPlatformApplication, err error) {
	result = &v1alpha1.SnsPlatformApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SnsPlatformApplications that match those selectors.
func (c *snsPlatformApplications) List(opts v1.ListOptions) (result *v1alpha1.SnsPlatformApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SnsPlatformApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested snsPlatformApplications.
func (c *snsPlatformApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a snsPlatformApplication and creates it.  Returns the server's representation of the snsPlatformApplication, and an error, if there is any.
func (c *snsPlatformApplications) Create(snsPlatformApplication *v1alpha1.SnsPlatformApplication) (result *v1alpha1.SnsPlatformApplication, err error) {
	result = &v1alpha1.SnsPlatformApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		Body(snsPlatformApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a snsPlatformApplication and updates it. Returns the server's representation of the snsPlatformApplication, and an error, if there is any.
func (c *snsPlatformApplications) Update(snsPlatformApplication *v1alpha1.SnsPlatformApplication) (result *v1alpha1.SnsPlatformApplication, err error) {
	result = &v1alpha1.SnsPlatformApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		Name(snsPlatformApplication.Name).
		Body(snsPlatformApplication).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *snsPlatformApplications) UpdateStatus(snsPlatformApplication *v1alpha1.SnsPlatformApplication) (result *v1alpha1.SnsPlatformApplication, err error) {
	result = &v1alpha1.SnsPlatformApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		Name(snsPlatformApplication.Name).
		SubResource("status").
		Body(snsPlatformApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the snsPlatformApplication and deletes it. Returns an error if one occurs.
func (c *snsPlatformApplications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *snsPlatformApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("snsplatformapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched snsPlatformApplication.
func (c *snsPlatformApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SnsPlatformApplication, err error) {
	result = &v1alpha1.SnsPlatformApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("snsplatformapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}