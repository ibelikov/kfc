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

// GlueCrawlersGetter has a method to return a GlueCrawlerInterface.
// A group's client should implement this interface.
type GlueCrawlersGetter interface {
	GlueCrawlers(namespace string) GlueCrawlerInterface
}

// GlueCrawlerInterface has methods to work with GlueCrawler resources.
type GlueCrawlerInterface interface {
	Create(*v1alpha1.GlueCrawler) (*v1alpha1.GlueCrawler, error)
	Update(*v1alpha1.GlueCrawler) (*v1alpha1.GlueCrawler, error)
	UpdateStatus(*v1alpha1.GlueCrawler) (*v1alpha1.GlueCrawler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GlueCrawler, error)
	List(opts v1.ListOptions) (*v1alpha1.GlueCrawlerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCrawler, err error)
	GlueCrawlerExpansion
}

// glueCrawlers implements GlueCrawlerInterface
type glueCrawlers struct {
	client rest.Interface
	ns     string
}

// newGlueCrawlers returns a GlueCrawlers
func newGlueCrawlers(c *AwsV1alpha1Client, namespace string) *glueCrawlers {
	return &glueCrawlers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the glueCrawler, and returns the corresponding glueCrawler object, and an error if there is any.
func (c *glueCrawlers) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueCrawler, err error) {
	result = &v1alpha1.GlueCrawler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluecrawlers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlueCrawlers that match those selectors.
func (c *glueCrawlers) List(opts v1.ListOptions) (result *v1alpha1.GlueCrawlerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GlueCrawlerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluecrawlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested glueCrawlers.
func (c *glueCrawlers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gluecrawlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a glueCrawler and creates it.  Returns the server's representation of the glueCrawler, and an error, if there is any.
func (c *glueCrawlers) Create(glueCrawler *v1alpha1.GlueCrawler) (result *v1alpha1.GlueCrawler, err error) {
	result = &v1alpha1.GlueCrawler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gluecrawlers").
		Body(glueCrawler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a glueCrawler and updates it. Returns the server's representation of the glueCrawler, and an error, if there is any.
func (c *glueCrawlers) Update(glueCrawler *v1alpha1.GlueCrawler) (result *v1alpha1.GlueCrawler, err error) {
	result = &v1alpha1.GlueCrawler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluecrawlers").
		Name(glueCrawler.Name).
		Body(glueCrawler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *glueCrawlers) UpdateStatus(glueCrawler *v1alpha1.GlueCrawler) (result *v1alpha1.GlueCrawler, err error) {
	result = &v1alpha1.GlueCrawler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluecrawlers").
		Name(glueCrawler.Name).
		SubResource("status").
		Body(glueCrawler).
		Do().
		Into(result)
	return
}

// Delete takes name of the glueCrawler and deletes it. Returns an error if one occurs.
func (c *glueCrawlers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluecrawlers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *glueCrawlers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluecrawlers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched glueCrawler.
func (c *glueCrawlers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueCrawler, err error) {
	result = &v1alpha1.GlueCrawler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gluecrawlers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
