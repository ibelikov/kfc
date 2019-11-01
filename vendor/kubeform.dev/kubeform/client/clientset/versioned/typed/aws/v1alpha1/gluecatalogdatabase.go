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

// GlueCatalogDatabasesGetter has a method to return a GlueCatalogDatabaseInterface.
// A group's client should implement this interface.
type GlueCatalogDatabasesGetter interface {
	GlueCatalogDatabases(namespace string) GlueCatalogDatabaseInterface
}

// GlueCatalogDatabaseInterface has methods to work with GlueCatalogDatabase resources.
type GlueCatalogDatabaseInterface interface {
	Create(*v1alpha1.GlueCatalogDatabase) (*v1alpha1.GlueCatalogDatabase, error)
	Update(*v1alpha1.GlueCatalogDatabase) (*v1alpha1.GlueCatalogDatabase, error)
	UpdateStatus(*v1alpha1.GlueCatalogDatabase) (*v1alpha1.GlueCatalogDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GlueCatalogDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.GlueCatalogDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCatalogDatabase, err error)
	GlueCatalogDatabaseExpansion
}

// glueCatalogDatabases implements GlueCatalogDatabaseInterface
type glueCatalogDatabases struct {
	client rest.Interface
	ns     string
}

// newGlueCatalogDatabases returns a GlueCatalogDatabases
func newGlueCatalogDatabases(c *AwsV1alpha1Client, namespace string) *glueCatalogDatabases {
	return &glueCatalogDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the glueCatalogDatabase, and returns the corresponding glueCatalogDatabase object, and an error if there is any.
func (c *glueCatalogDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueCatalogDatabase, err error) {
	result = &v1alpha1.GlueCatalogDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlueCatalogDatabases that match those selectors.
func (c *glueCatalogDatabases) List(opts v1.ListOptions) (result *v1alpha1.GlueCatalogDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GlueCatalogDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested glueCatalogDatabases.
func (c *glueCatalogDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a glueCatalogDatabase and creates it.  Returns the server's representation of the glueCatalogDatabase, and an error, if there is any.
func (c *glueCatalogDatabases) Create(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (result *v1alpha1.GlueCatalogDatabase, err error) {
	result = &v1alpha1.GlueCatalogDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		Body(glueCatalogDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a glueCatalogDatabase and updates it. Returns the server's representation of the glueCatalogDatabase, and an error, if there is any.
func (c *glueCatalogDatabases) Update(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (result *v1alpha1.GlueCatalogDatabase, err error) {
	result = &v1alpha1.GlueCatalogDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		Name(glueCatalogDatabase.Name).
		Body(glueCatalogDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *glueCatalogDatabases) UpdateStatus(glueCatalogDatabase *v1alpha1.GlueCatalogDatabase) (result *v1alpha1.GlueCatalogDatabase, err error) {
	result = &v1alpha1.GlueCatalogDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		Name(glueCatalogDatabase.Name).
		SubResource("status").
		Body(glueCatalogDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the glueCatalogDatabase and deletes it. Returns an error if one occurs.
func (c *glueCatalogDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *glueCatalogDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched glueCatalogDatabase.
func (c *glueCatalogDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCatalogDatabase, err error) {
	result = &v1alpha1.GlueCatalogDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gluecatalogdatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}