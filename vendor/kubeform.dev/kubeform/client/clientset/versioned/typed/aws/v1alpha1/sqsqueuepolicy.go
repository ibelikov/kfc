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

// SqsQueuePoliciesGetter has a method to return a SqsQueuePolicyInterface.
// A group's client should implement this interface.
type SqsQueuePoliciesGetter interface {
	SqsQueuePolicies(namespace string) SqsQueuePolicyInterface
}

// SqsQueuePolicyInterface has methods to work with SqsQueuePolicy resources.
type SqsQueuePolicyInterface interface {
	Create(*v1alpha1.SqsQueuePolicy) (*v1alpha1.SqsQueuePolicy, error)
	Update(*v1alpha1.SqsQueuePolicy) (*v1alpha1.SqsQueuePolicy, error)
	UpdateStatus(*v1alpha1.SqsQueuePolicy) (*v1alpha1.SqsQueuePolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SqsQueuePolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.SqsQueuePolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqsQueuePolicy, err error)
	SqsQueuePolicyExpansion
}

// sqsQueuePolicies implements SqsQueuePolicyInterface
type sqsQueuePolicies struct {
	client rest.Interface
	ns     string
}

// newSqsQueuePolicies returns a SqsQueuePolicies
func newSqsQueuePolicies(c *AwsV1alpha1Client, namespace string) *sqsQueuePolicies {
	return &sqsQueuePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sqsQueuePolicy, and returns the corresponding sqsQueuePolicy object, and an error if there is any.
func (c *sqsQueuePolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SqsQueuePolicy, err error) {
	result = &v1alpha1.SqsQueuePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SqsQueuePolicies that match those selectors.
func (c *sqsQueuePolicies) List(opts v1.ListOptions) (result *v1alpha1.SqsQueuePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SqsQueuePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sqsQueuePolicies.
func (c *sqsQueuePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sqsQueuePolicy and creates it.  Returns the server's representation of the sqsQueuePolicy, and an error, if there is any.
func (c *sqsQueuePolicies) Create(sqsQueuePolicy *v1alpha1.SqsQueuePolicy) (result *v1alpha1.SqsQueuePolicy, err error) {
	result = &v1alpha1.SqsQueuePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		Body(sqsQueuePolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sqsQueuePolicy and updates it. Returns the server's representation of the sqsQueuePolicy, and an error, if there is any.
func (c *sqsQueuePolicies) Update(sqsQueuePolicy *v1alpha1.SqsQueuePolicy) (result *v1alpha1.SqsQueuePolicy, err error) {
	result = &v1alpha1.SqsQueuePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		Name(sqsQueuePolicy.Name).
		Body(sqsQueuePolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sqsQueuePolicies) UpdateStatus(sqsQueuePolicy *v1alpha1.SqsQueuePolicy) (result *v1alpha1.SqsQueuePolicy, err error) {
	result = &v1alpha1.SqsQueuePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		Name(sqsQueuePolicy.Name).
		SubResource("status").
		Body(sqsQueuePolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the sqsQueuePolicy and deletes it. Returns an error if one occurs.
func (c *sqsQueuePolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sqsQueuePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sqsQueuePolicy.
func (c *sqsQueuePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqsQueuePolicy, err error) {
	result = &v1alpha1.SqsQueuePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqsqueuepolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
