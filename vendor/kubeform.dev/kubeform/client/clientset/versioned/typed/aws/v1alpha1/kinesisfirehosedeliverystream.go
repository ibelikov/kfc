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

// KinesisFirehoseDeliveryStreamsGetter has a method to return a KinesisFirehoseDeliveryStreamInterface.
// A group's client should implement this interface.
type KinesisFirehoseDeliveryStreamsGetter interface {
	KinesisFirehoseDeliveryStreams(namespace string) KinesisFirehoseDeliveryStreamInterface
}

// KinesisFirehoseDeliveryStreamInterface has methods to work with KinesisFirehoseDeliveryStream resources.
type KinesisFirehoseDeliveryStreamInterface interface {
	Create(*v1alpha1.KinesisFirehoseDeliveryStream) (*v1alpha1.KinesisFirehoseDeliveryStream, error)
	Update(*v1alpha1.KinesisFirehoseDeliveryStream) (*v1alpha1.KinesisFirehoseDeliveryStream, error)
	UpdateStatus(*v1alpha1.KinesisFirehoseDeliveryStream) (*v1alpha1.KinesisFirehoseDeliveryStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KinesisFirehoseDeliveryStream, error)
	List(opts v1.ListOptions) (*v1alpha1.KinesisFirehoseDeliveryStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error)
	KinesisFirehoseDeliveryStreamExpansion
}

// kinesisFirehoseDeliveryStreams implements KinesisFirehoseDeliveryStreamInterface
type kinesisFirehoseDeliveryStreams struct {
	client rest.Interface
	ns     string
}

// newKinesisFirehoseDeliveryStreams returns a KinesisFirehoseDeliveryStreams
func newKinesisFirehoseDeliveryStreams(c *AwsV1alpha1Client, namespace string) *kinesisFirehoseDeliveryStreams {
	return &kinesisFirehoseDeliveryStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kinesisFirehoseDeliveryStream, and returns the corresponding kinesisFirehoseDeliveryStream object, and an error if there is any.
func (c *kinesisFirehoseDeliveryStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	result = &v1alpha1.KinesisFirehoseDeliveryStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KinesisFirehoseDeliveryStreams that match those selectors.
func (c *kinesisFirehoseDeliveryStreams) List(opts v1.ListOptions) (result *v1alpha1.KinesisFirehoseDeliveryStreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KinesisFirehoseDeliveryStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kinesisFirehoseDeliveryStreams.
func (c *kinesisFirehoseDeliveryStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kinesisFirehoseDeliveryStream and creates it.  Returns the server's representation of the kinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *kinesisFirehoseDeliveryStreams) Create(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	result = &v1alpha1.KinesisFirehoseDeliveryStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		Body(kinesisFirehoseDeliveryStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kinesisFirehoseDeliveryStream and updates it. Returns the server's representation of the kinesisFirehoseDeliveryStream, and an error, if there is any.
func (c *kinesisFirehoseDeliveryStreams) Update(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	result = &v1alpha1.KinesisFirehoseDeliveryStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		Name(kinesisFirehoseDeliveryStream.Name).
		Body(kinesisFirehoseDeliveryStream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kinesisFirehoseDeliveryStreams) UpdateStatus(kinesisFirehoseDeliveryStream *v1alpha1.KinesisFirehoseDeliveryStream) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	result = &v1alpha1.KinesisFirehoseDeliveryStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		Name(kinesisFirehoseDeliveryStream.Name).
		SubResource("status").
		Body(kinesisFirehoseDeliveryStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the kinesisFirehoseDeliveryStream and deletes it. Returns an error if one occurs.
func (c *kinesisFirehoseDeliveryStreams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kinesisFirehoseDeliveryStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kinesisFirehoseDeliveryStream.
func (c *kinesisFirehoseDeliveryStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KinesisFirehoseDeliveryStream, err error) {
	result = &v1alpha1.KinesisFirehoseDeliveryStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kinesisfirehosedeliverystreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
