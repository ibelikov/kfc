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

// NeptuneClusterParameterGroupsGetter has a method to return a NeptuneClusterParameterGroupInterface.
// A group's client should implement this interface.
type NeptuneClusterParameterGroupsGetter interface {
	NeptuneClusterParameterGroups(namespace string) NeptuneClusterParameterGroupInterface
}

// NeptuneClusterParameterGroupInterface has methods to work with NeptuneClusterParameterGroup resources.
type NeptuneClusterParameterGroupInterface interface {
	Create(*v1alpha1.NeptuneClusterParameterGroup) (*v1alpha1.NeptuneClusterParameterGroup, error)
	Update(*v1alpha1.NeptuneClusterParameterGroup) (*v1alpha1.NeptuneClusterParameterGroup, error)
	UpdateStatus(*v1alpha1.NeptuneClusterParameterGroup) (*v1alpha1.NeptuneClusterParameterGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NeptuneClusterParameterGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.NeptuneClusterParameterGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneClusterParameterGroup, err error)
	NeptuneClusterParameterGroupExpansion
}

// neptuneClusterParameterGroups implements NeptuneClusterParameterGroupInterface
type neptuneClusterParameterGroups struct {
	client rest.Interface
	ns     string
}

// newNeptuneClusterParameterGroups returns a NeptuneClusterParameterGroups
func newNeptuneClusterParameterGroups(c *AwsV1alpha1Client, namespace string) *neptuneClusterParameterGroups {
	return &neptuneClusterParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the neptuneClusterParameterGroup, and returns the corresponding neptuneClusterParameterGroup object, and an error if there is any.
func (c *neptuneClusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NeptuneClusterParameterGroup, err error) {
	result = &v1alpha1.NeptuneClusterParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NeptuneClusterParameterGroups that match those selectors.
func (c *neptuneClusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.NeptuneClusterParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NeptuneClusterParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested neptuneClusterParameterGroups.
func (c *neptuneClusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a neptuneClusterParameterGroup and creates it.  Returns the server's representation of the neptuneClusterParameterGroup, and an error, if there is any.
func (c *neptuneClusterParameterGroups) Create(neptuneClusterParameterGroup *v1alpha1.NeptuneClusterParameterGroup) (result *v1alpha1.NeptuneClusterParameterGroup, err error) {
	result = &v1alpha1.NeptuneClusterParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		Body(neptuneClusterParameterGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a neptuneClusterParameterGroup and updates it. Returns the server's representation of the neptuneClusterParameterGroup, and an error, if there is any.
func (c *neptuneClusterParameterGroups) Update(neptuneClusterParameterGroup *v1alpha1.NeptuneClusterParameterGroup) (result *v1alpha1.NeptuneClusterParameterGroup, err error) {
	result = &v1alpha1.NeptuneClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		Name(neptuneClusterParameterGroup.Name).
		Body(neptuneClusterParameterGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *neptuneClusterParameterGroups) UpdateStatus(neptuneClusterParameterGroup *v1alpha1.NeptuneClusterParameterGroup) (result *v1alpha1.NeptuneClusterParameterGroup, err error) {
	result = &v1alpha1.NeptuneClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		Name(neptuneClusterParameterGroup.Name).
		SubResource("status").
		Body(neptuneClusterParameterGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the neptuneClusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *neptuneClusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *neptuneClusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched neptuneClusterParameterGroup.
func (c *neptuneClusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NeptuneClusterParameterGroup, err error) {
	result = &v1alpha1.NeptuneClusterParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("neptuneclusterparametergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
