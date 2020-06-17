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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeSecurityPoliciesGetter has a method to return a ComputeSecurityPolicyInterface.
// A group's client should implement this interface.
type ComputeSecurityPoliciesGetter interface {
	ComputeSecurityPolicies(namespace string) ComputeSecurityPolicyInterface
}

// ComputeSecurityPolicyInterface has methods to work with ComputeSecurityPolicy resources.
type ComputeSecurityPolicyInterface interface {
	Create(*v1alpha1.ComputeSecurityPolicy) (*v1alpha1.ComputeSecurityPolicy, error)
	Update(*v1alpha1.ComputeSecurityPolicy) (*v1alpha1.ComputeSecurityPolicy, error)
	UpdateStatus(*v1alpha1.ComputeSecurityPolicy) (*v1alpha1.ComputeSecurityPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSecurityPolicy, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSecurityPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSecurityPolicy, err error)
	ComputeSecurityPolicyExpansion
}

// computeSecurityPolicies implements ComputeSecurityPolicyInterface
type computeSecurityPolicies struct {
	client rest.Interface
	ns     string
}

// newComputeSecurityPolicies returns a ComputeSecurityPolicies
func newComputeSecurityPolicies(c *GoogleV1alpha1Client, namespace string) *computeSecurityPolicies {
	return &computeSecurityPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeSecurityPolicy, and returns the corresponding computeSecurityPolicy object, and an error if there is any.
func (c *computeSecurityPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	result = &v1alpha1.ComputeSecurityPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSecurityPolicies that match those selectors.
func (c *computeSecurityPolicies) List(opts v1.ListOptions) (result *v1alpha1.ComputeSecurityPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSecurityPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSecurityPolicies.
func (c *computeSecurityPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSecurityPolicy and creates it.  Returns the server's representation of the computeSecurityPolicy, and an error, if there is any.
func (c *computeSecurityPolicies) Create(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	result = &v1alpha1.ComputeSecurityPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		Body(computeSecurityPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSecurityPolicy and updates it. Returns the server's representation of the computeSecurityPolicy, and an error, if there is any.
func (c *computeSecurityPolicies) Update(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	result = &v1alpha1.ComputeSecurityPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		Name(computeSecurityPolicy.Name).
		Body(computeSecurityPolicy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSecurityPolicies) UpdateStatus(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	result = &v1alpha1.ComputeSecurityPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		Name(computeSecurityPolicy.Name).
		SubResource("status").
		Body(computeSecurityPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSecurityPolicy and deletes it. Returns an error if one occurs.
func (c *computeSecurityPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSecurityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSecurityPolicy.
func (c *computeSecurityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	result = &v1alpha1.ComputeSecurityPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computesecuritypolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
