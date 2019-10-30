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

// CloudfrontOriginAccessIdentitiesGetter has a method to return a CloudfrontOriginAccessIdentityInterface.
// A group's client should implement this interface.
type CloudfrontOriginAccessIdentitiesGetter interface {
	CloudfrontOriginAccessIdentities(namespace string) CloudfrontOriginAccessIdentityInterface
}

// CloudfrontOriginAccessIdentityInterface has methods to work with CloudfrontOriginAccessIdentity resources.
type CloudfrontOriginAccessIdentityInterface interface {
	Create(*v1alpha1.CloudfrontOriginAccessIdentity) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	Update(*v1alpha1.CloudfrontOriginAccessIdentity) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	UpdateStatus(*v1alpha1.CloudfrontOriginAccessIdentity) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudfrontOriginAccessIdentity, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudfrontOriginAccessIdentityList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error)
	CloudfrontOriginAccessIdentityExpansion
}

// cloudfrontOriginAccessIdentities implements CloudfrontOriginAccessIdentityInterface
type cloudfrontOriginAccessIdentities struct {
	client rest.Interface
	ns     string
}

// newCloudfrontOriginAccessIdentities returns a CloudfrontOriginAccessIdentities
func newCloudfrontOriginAccessIdentities(c *AwsV1alpha1Client, namespace string) *cloudfrontOriginAccessIdentities {
	return &cloudfrontOriginAccessIdentities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudfrontOriginAccessIdentity, and returns the corresponding cloudfrontOriginAccessIdentity object, and an error if there is any.
func (c *cloudfrontOriginAccessIdentities) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudfrontOriginAccessIdentities that match those selectors.
func (c *cloudfrontOriginAccessIdentities) List(opts v1.ListOptions) (result *v1alpha1.CloudfrontOriginAccessIdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudfrontOriginAccessIdentityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudfrontOriginAccessIdentities.
func (c *cloudfrontOriginAccessIdentities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudfrontOriginAccessIdentity and creates it.  Returns the server's representation of the cloudfrontOriginAccessIdentity, and an error, if there is any.
func (c *cloudfrontOriginAccessIdentities) Create(cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Body(cloudfrontOriginAccessIdentity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudfrontOriginAccessIdentity and updates it. Returns the server's representation of the cloudfrontOriginAccessIdentity, and an error, if there is any.
func (c *cloudfrontOriginAccessIdentities) Update(cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(cloudfrontOriginAccessIdentity.Name).
		Body(cloudfrontOriginAccessIdentity).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudfrontOriginAccessIdentities) UpdateStatus(cloudfrontOriginAccessIdentity *v1alpha1.CloudfrontOriginAccessIdentity) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(cloudfrontOriginAccessIdentity.Name).
		SubResource("status").
		Body(cloudfrontOriginAccessIdentity).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudfrontOriginAccessIdentity and deletes it. Returns an error if one occurs.
func (c *cloudfrontOriginAccessIdentities) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudfrontOriginAccessIdentities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudfrontOriginAccessIdentity.
func (c *cloudfrontOriginAccessIdentities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfrontOriginAccessIdentity, err error) {
	result = &v1alpha1.CloudfrontOriginAccessIdentity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudfrontoriginaccessidentities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
