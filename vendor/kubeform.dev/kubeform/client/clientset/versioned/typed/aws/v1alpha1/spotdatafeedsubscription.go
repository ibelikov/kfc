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

// SpotDatafeedSubscriptionsGetter has a method to return a SpotDatafeedSubscriptionInterface.
// A group's client should implement this interface.
type SpotDatafeedSubscriptionsGetter interface {
	SpotDatafeedSubscriptions(namespace string) SpotDatafeedSubscriptionInterface
}

// SpotDatafeedSubscriptionInterface has methods to work with SpotDatafeedSubscription resources.
type SpotDatafeedSubscriptionInterface interface {
	Create(*v1alpha1.SpotDatafeedSubscription) (*v1alpha1.SpotDatafeedSubscription, error)
	Update(*v1alpha1.SpotDatafeedSubscription) (*v1alpha1.SpotDatafeedSubscription, error)
	UpdateStatus(*v1alpha1.SpotDatafeedSubscription) (*v1alpha1.SpotDatafeedSubscription, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SpotDatafeedSubscription, error)
	List(opts v1.ListOptions) (*v1alpha1.SpotDatafeedSubscriptionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotDatafeedSubscription, err error)
	SpotDatafeedSubscriptionExpansion
}

// spotDatafeedSubscriptions implements SpotDatafeedSubscriptionInterface
type spotDatafeedSubscriptions struct {
	client rest.Interface
	ns     string
}

// newSpotDatafeedSubscriptions returns a SpotDatafeedSubscriptions
func newSpotDatafeedSubscriptions(c *AwsV1alpha1Client, namespace string) *spotDatafeedSubscriptions {
	return &spotDatafeedSubscriptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the spotDatafeedSubscription, and returns the corresponding spotDatafeedSubscription object, and an error if there is any.
func (c *spotDatafeedSubscriptions) Get(name string, options v1.GetOptions) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	result = &v1alpha1.SpotDatafeedSubscription{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SpotDatafeedSubscriptions that match those selectors.
func (c *spotDatafeedSubscriptions) List(opts v1.ListOptions) (result *v1alpha1.SpotDatafeedSubscriptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SpotDatafeedSubscriptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spotDatafeedSubscriptions.
func (c *spotDatafeedSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a spotDatafeedSubscription and creates it.  Returns the server's representation of the spotDatafeedSubscription, and an error, if there is any.
func (c *spotDatafeedSubscriptions) Create(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	result = &v1alpha1.SpotDatafeedSubscription{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		Body(spotDatafeedSubscription).
		Do().
		Into(result)
	return
}

// Update takes the representation of a spotDatafeedSubscription and updates it. Returns the server's representation of the spotDatafeedSubscription, and an error, if there is any.
func (c *spotDatafeedSubscriptions) Update(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	result = &v1alpha1.SpotDatafeedSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		Name(spotDatafeedSubscription.Name).
		Body(spotDatafeedSubscription).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *spotDatafeedSubscriptions) UpdateStatus(spotDatafeedSubscription *v1alpha1.SpotDatafeedSubscription) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	result = &v1alpha1.SpotDatafeedSubscription{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		Name(spotDatafeedSubscription.Name).
		SubResource("status").
		Body(spotDatafeedSubscription).
		Do().
		Into(result)
	return
}

// Delete takes name of the spotDatafeedSubscription and deletes it. Returns an error if one occurs.
func (c *spotDatafeedSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spotDatafeedSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched spotDatafeedSubscription.
func (c *spotDatafeedSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotDatafeedSubscription, err error) {
	result = &v1alpha1.SpotDatafeedSubscription{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("spotdatafeedsubscriptions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
