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

// SecurityGroupsGetter has a method to return a SecurityGroupInterface.
// A group's client should implement this interface.
type SecurityGroupsGetter interface {
	SecurityGroups(namespace string) SecurityGroupInterface
}

// SecurityGroupInterface has methods to work with SecurityGroup resources.
type SecurityGroupInterface interface {
	Create(*v1alpha1.SecurityGroup) (*v1alpha1.SecurityGroup, error)
	Update(*v1alpha1.SecurityGroup) (*v1alpha1.SecurityGroup, error)
	UpdateStatus(*v1alpha1.SecurityGroup) (*v1alpha1.SecurityGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecurityGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.SecurityGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityGroup, err error)
	SecurityGroupExpansion
}

// securityGroups implements SecurityGroupInterface
type securityGroups struct {
	client rest.Interface
	ns     string
}

// newSecurityGroups returns a SecurityGroups
func newSecurityGroups(c *AwsV1alpha1Client, namespace string) *securityGroups {
	return &securityGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the securityGroup, and returns the corresponding securityGroup object, and an error if there is any.
func (c *securityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.SecurityGroup, err error) {
	result = &v1alpha1.SecurityGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securitygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityGroups that match those selectors.
func (c *securityGroups) List(opts v1.ListOptions) (result *v1alpha1.SecurityGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecurityGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securitygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityGroups.
func (c *securityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securitygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a securityGroup and creates it.  Returns the server's representation of the securityGroup, and an error, if there is any.
func (c *securityGroups) Create(securityGroup *v1alpha1.SecurityGroup) (result *v1alpha1.SecurityGroup, err error) {
	result = &v1alpha1.SecurityGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securitygroups").
		Body(securityGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a securityGroup and updates it. Returns the server's representation of the securityGroup, and an error, if there is any.
func (c *securityGroups) Update(securityGroup *v1alpha1.SecurityGroup) (result *v1alpha1.SecurityGroup, err error) {
	result = &v1alpha1.SecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securitygroups").
		Name(securityGroup.Name).
		Body(securityGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *securityGroups) UpdateStatus(securityGroup *v1alpha1.SecurityGroup) (result *v1alpha1.SecurityGroup, err error) {
	result = &v1alpha1.SecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securitygroups").
		Name(securityGroup.Name).
		SubResource("status").
		Body(securityGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the securityGroup and deletes it. Returns an error if one occurs.
func (c *securityGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securitygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securitygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched securityGroup.
func (c *securityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecurityGroup, err error) {
	result = &v1alpha1.SecurityGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securitygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
