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

// DbSecurityGroupsGetter has a method to return a DbSecurityGroupInterface.
// A group's client should implement this interface.
type DbSecurityGroupsGetter interface {
	DbSecurityGroups(namespace string) DbSecurityGroupInterface
}

// DbSecurityGroupInterface has methods to work with DbSecurityGroup resources.
type DbSecurityGroupInterface interface {
	Create(*v1alpha1.DbSecurityGroup) (*v1alpha1.DbSecurityGroup, error)
	Update(*v1alpha1.DbSecurityGroup) (*v1alpha1.DbSecurityGroup, error)
	UpdateStatus(*v1alpha1.DbSecurityGroup) (*v1alpha1.DbSecurityGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DbSecurityGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.DbSecurityGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbSecurityGroup, err error)
	DbSecurityGroupExpansion
}

// dbSecurityGroups implements DbSecurityGroupInterface
type dbSecurityGroups struct {
	client rest.Interface
	ns     string
}

// newDbSecurityGroups returns a DbSecurityGroups
func newDbSecurityGroups(c *AwsV1alpha1Client, namespace string) *dbSecurityGroups {
	return &dbSecurityGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbSecurityGroup, and returns the corresponding dbSecurityGroup object, and an error if there is any.
func (c *dbSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DbSecurityGroup, err error) {
	result = &v1alpha1.DbSecurityGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbSecurityGroups that match those selectors.
func (c *dbSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.DbSecurityGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DbSecurityGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbSecurityGroups.
func (c *dbSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dbSecurityGroup and creates it.  Returns the server's representation of the dbSecurityGroup, and an error, if there is any.
func (c *dbSecurityGroups) Create(dbSecurityGroup *v1alpha1.DbSecurityGroup) (result *v1alpha1.DbSecurityGroup, err error) {
	result = &v1alpha1.DbSecurityGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		Body(dbSecurityGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dbSecurityGroup and updates it. Returns the server's representation of the dbSecurityGroup, and an error, if there is any.
func (c *dbSecurityGroups) Update(dbSecurityGroup *v1alpha1.DbSecurityGroup) (result *v1alpha1.DbSecurityGroup, err error) {
	result = &v1alpha1.DbSecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		Name(dbSecurityGroup.Name).
		Body(dbSecurityGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dbSecurityGroups) UpdateStatus(dbSecurityGroup *v1alpha1.DbSecurityGroup) (result *v1alpha1.DbSecurityGroup, err error) {
	result = &v1alpha1.DbSecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		Name(dbSecurityGroup.Name).
		SubResource("status").
		Body(dbSecurityGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the dbSecurityGroup and deletes it. Returns an error if one occurs.
func (c *dbSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dbSecurityGroup.
func (c *dbSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DbSecurityGroup, err error) {
	result = &v1alpha1.DbSecurityGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbsecuritygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
