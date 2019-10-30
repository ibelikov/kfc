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

// GlobalacceleratorListenersGetter has a method to return a GlobalacceleratorListenerInterface.
// A group's client should implement this interface.
type GlobalacceleratorListenersGetter interface {
	GlobalacceleratorListeners(namespace string) GlobalacceleratorListenerInterface
}

// GlobalacceleratorListenerInterface has methods to work with GlobalacceleratorListener resources.
type GlobalacceleratorListenerInterface interface {
	Create(*v1alpha1.GlobalacceleratorListener) (*v1alpha1.GlobalacceleratorListener, error)
	Update(*v1alpha1.GlobalacceleratorListener) (*v1alpha1.GlobalacceleratorListener, error)
	UpdateStatus(*v1alpha1.GlobalacceleratorListener) (*v1alpha1.GlobalacceleratorListener, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GlobalacceleratorListener, error)
	List(opts v1.ListOptions) (*v1alpha1.GlobalacceleratorListenerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlobalacceleratorListener, err error)
	GlobalacceleratorListenerExpansion
}

// globalacceleratorListeners implements GlobalacceleratorListenerInterface
type globalacceleratorListeners struct {
	client rest.Interface
	ns     string
}

// newGlobalacceleratorListeners returns a GlobalacceleratorListeners
func newGlobalacceleratorListeners(c *AwsV1alpha1Client, namespace string) *globalacceleratorListeners {
	return &globalacceleratorListeners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the globalacceleratorListener, and returns the corresponding globalacceleratorListener object, and an error if there is any.
func (c *globalacceleratorListeners) Get(name string, options v1.GetOptions) (result *v1alpha1.GlobalacceleratorListener, err error) {
	result = &v1alpha1.GlobalacceleratorListener{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalacceleratorListeners that match those selectors.
func (c *globalacceleratorListeners) List(opts v1.ListOptions) (result *v1alpha1.GlobalacceleratorListenerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GlobalacceleratorListenerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalacceleratorListeners.
func (c *globalacceleratorListeners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a globalacceleratorListener and creates it.  Returns the server's representation of the globalacceleratorListener, and an error, if there is any.
func (c *globalacceleratorListeners) Create(globalacceleratorListener *v1alpha1.GlobalacceleratorListener) (result *v1alpha1.GlobalacceleratorListener, err error) {
	result = &v1alpha1.GlobalacceleratorListener{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		Body(globalacceleratorListener).
		Do().
		Into(result)
	return
}

// Update takes the representation of a globalacceleratorListener and updates it. Returns the server's representation of the globalacceleratorListener, and an error, if there is any.
func (c *globalacceleratorListeners) Update(globalacceleratorListener *v1alpha1.GlobalacceleratorListener) (result *v1alpha1.GlobalacceleratorListener, err error) {
	result = &v1alpha1.GlobalacceleratorListener{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		Name(globalacceleratorListener.Name).
		Body(globalacceleratorListener).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *globalacceleratorListeners) UpdateStatus(globalacceleratorListener *v1alpha1.GlobalacceleratorListener) (result *v1alpha1.GlobalacceleratorListener, err error) {
	result = &v1alpha1.GlobalacceleratorListener{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		Name(globalacceleratorListener.Name).
		SubResource("status").
		Body(globalacceleratorListener).
		Do().
		Into(result)
	return
}

// Delete takes name of the globalacceleratorListener and deletes it. Returns an error if one occurs.
func (c *globalacceleratorListeners) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalacceleratorListeners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched globalacceleratorListener.
func (c *globalacceleratorListeners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlobalacceleratorListener, err error) {
	result = &v1alpha1.GlobalacceleratorListener{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("globalacceleratorlisteners").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
