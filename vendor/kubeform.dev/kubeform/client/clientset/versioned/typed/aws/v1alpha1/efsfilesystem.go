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

// EfsFileSystemsGetter has a method to return a EfsFileSystemInterface.
// A group's client should implement this interface.
type EfsFileSystemsGetter interface {
	EfsFileSystems(namespace string) EfsFileSystemInterface
}

// EfsFileSystemInterface has methods to work with EfsFileSystem resources.
type EfsFileSystemInterface interface {
	Create(*v1alpha1.EfsFileSystem) (*v1alpha1.EfsFileSystem, error)
	Update(*v1alpha1.EfsFileSystem) (*v1alpha1.EfsFileSystem, error)
	UpdateStatus(*v1alpha1.EfsFileSystem) (*v1alpha1.EfsFileSystem, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EfsFileSystem, error)
	List(opts v1.ListOptions) (*v1alpha1.EfsFileSystemList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EfsFileSystem, err error)
	EfsFileSystemExpansion
}

// efsFileSystems implements EfsFileSystemInterface
type efsFileSystems struct {
	client rest.Interface
	ns     string
}

// newEfsFileSystems returns a EfsFileSystems
func newEfsFileSystems(c *AwsV1alpha1Client, namespace string) *efsFileSystems {
	return &efsFileSystems{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the efsFileSystem, and returns the corresponding efsFileSystem object, and an error if there is any.
func (c *efsFileSystems) Get(name string, options v1.GetOptions) (result *v1alpha1.EfsFileSystem, err error) {
	result = &v1alpha1.EfsFileSystem{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("efsfilesystems").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EfsFileSystems that match those selectors.
func (c *efsFileSystems) List(opts v1.ListOptions) (result *v1alpha1.EfsFileSystemList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EfsFileSystemList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("efsfilesystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested efsFileSystems.
func (c *efsFileSystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("efsfilesystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a efsFileSystem and creates it.  Returns the server's representation of the efsFileSystem, and an error, if there is any.
func (c *efsFileSystems) Create(efsFileSystem *v1alpha1.EfsFileSystem) (result *v1alpha1.EfsFileSystem, err error) {
	result = &v1alpha1.EfsFileSystem{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("efsfilesystems").
		Body(efsFileSystem).
		Do().
		Into(result)
	return
}

// Update takes the representation of a efsFileSystem and updates it. Returns the server's representation of the efsFileSystem, and an error, if there is any.
func (c *efsFileSystems) Update(efsFileSystem *v1alpha1.EfsFileSystem) (result *v1alpha1.EfsFileSystem, err error) {
	result = &v1alpha1.EfsFileSystem{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("efsfilesystems").
		Name(efsFileSystem.Name).
		Body(efsFileSystem).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *efsFileSystems) UpdateStatus(efsFileSystem *v1alpha1.EfsFileSystem) (result *v1alpha1.EfsFileSystem, err error) {
	result = &v1alpha1.EfsFileSystem{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("efsfilesystems").
		Name(efsFileSystem.Name).
		SubResource("status").
		Body(efsFileSystem).
		Do().
		Into(result)
	return
}

// Delete takes name of the efsFileSystem and deletes it. Returns an error if one occurs.
func (c *efsFileSystems) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("efsfilesystems").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *efsFileSystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("efsfilesystems").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched efsFileSystem.
func (c *efsFileSystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EfsFileSystem, err error) {
	result = &v1alpha1.EfsFileSystem{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("efsfilesystems").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
