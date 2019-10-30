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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoggingProjectSinksGetter has a method to return a LoggingProjectSinkInterface.
// A group's client should implement this interface.
type LoggingProjectSinksGetter interface {
	LoggingProjectSinks(namespace string) LoggingProjectSinkInterface
}

// LoggingProjectSinkInterface has methods to work with LoggingProjectSink resources.
type LoggingProjectSinkInterface interface {
	Create(*v1alpha1.LoggingProjectSink) (*v1alpha1.LoggingProjectSink, error)
	Update(*v1alpha1.LoggingProjectSink) (*v1alpha1.LoggingProjectSink, error)
	UpdateStatus(*v1alpha1.LoggingProjectSink) (*v1alpha1.LoggingProjectSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LoggingProjectSink, error)
	List(opts v1.ListOptions) (*v1alpha1.LoggingProjectSinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingProjectSink, err error)
	LoggingProjectSinkExpansion
}

// loggingProjectSinks implements LoggingProjectSinkInterface
type loggingProjectSinks struct {
	client rest.Interface
	ns     string
}

// newLoggingProjectSinks returns a LoggingProjectSinks
func newLoggingProjectSinks(c *GoogleV1alpha1Client, namespace string) *loggingProjectSinks {
	return &loggingProjectSinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loggingProjectSink, and returns the corresponding loggingProjectSink object, and an error if there is any.
func (c *loggingProjectSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.LoggingProjectSink, err error) {
	result = &v1alpha1.LoggingProjectSink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoggingProjectSinks that match those selectors.
func (c *loggingProjectSinks) List(opts v1.ListOptions) (result *v1alpha1.LoggingProjectSinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoggingProjectSinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loggingProjectSinks.
func (c *loggingProjectSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loggingProjectSink and creates it.  Returns the server's representation of the loggingProjectSink, and an error, if there is any.
func (c *loggingProjectSinks) Create(loggingProjectSink *v1alpha1.LoggingProjectSink) (result *v1alpha1.LoggingProjectSink, err error) {
	result = &v1alpha1.LoggingProjectSink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		Body(loggingProjectSink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loggingProjectSink and updates it. Returns the server's representation of the loggingProjectSink, and an error, if there is any.
func (c *loggingProjectSinks) Update(loggingProjectSink *v1alpha1.LoggingProjectSink) (result *v1alpha1.LoggingProjectSink, err error) {
	result = &v1alpha1.LoggingProjectSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		Name(loggingProjectSink.Name).
		Body(loggingProjectSink).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loggingProjectSinks) UpdateStatus(loggingProjectSink *v1alpha1.LoggingProjectSink) (result *v1alpha1.LoggingProjectSink, err error) {
	result = &v1alpha1.LoggingProjectSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		Name(loggingProjectSink.Name).
		SubResource("status").
		Body(loggingProjectSink).
		Do().
		Into(result)
	return
}

// Delete takes name of the loggingProjectSink and deletes it. Returns an error if one occurs.
func (c *loggingProjectSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loggingProjectSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loggingProjectSink.
func (c *loggingProjectSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoggingProjectSink, err error) {
	result = &v1alpha1.LoggingProjectSink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loggingprojectsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
