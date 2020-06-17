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

// CloudwatchEventPermissionsGetter has a method to return a CloudwatchEventPermissionInterface.
// A group's client should implement this interface.
type CloudwatchEventPermissionsGetter interface {
	CloudwatchEventPermissions(namespace string) CloudwatchEventPermissionInterface
}

// CloudwatchEventPermissionInterface has methods to work with CloudwatchEventPermission resources.
type CloudwatchEventPermissionInterface interface {
	Create(*v1alpha1.CloudwatchEventPermission) (*v1alpha1.CloudwatchEventPermission, error)
	Update(*v1alpha1.CloudwatchEventPermission) (*v1alpha1.CloudwatchEventPermission, error)
	UpdateStatus(*v1alpha1.CloudwatchEventPermission) (*v1alpha1.CloudwatchEventPermission, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudwatchEventPermission, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudwatchEventPermissionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchEventPermission, err error)
	CloudwatchEventPermissionExpansion
}

// cloudwatchEventPermissions implements CloudwatchEventPermissionInterface
type cloudwatchEventPermissions struct {
	client rest.Interface
	ns     string
}

// newCloudwatchEventPermissions returns a CloudwatchEventPermissions
func newCloudwatchEventPermissions(c *AwsV1alpha1Client, namespace string) *cloudwatchEventPermissions {
	return &cloudwatchEventPermissions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudwatchEventPermission, and returns the corresponding cloudwatchEventPermission object, and an error if there is any.
func (c *cloudwatchEventPermissions) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchEventPermission, err error) {
	result = &v1alpha1.CloudwatchEventPermission{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudwatchEventPermissions that match those selectors.
func (c *cloudwatchEventPermissions) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchEventPermissionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudwatchEventPermissionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudwatchEventPermissions.
func (c *cloudwatchEventPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudwatchEventPermission and creates it.  Returns the server's representation of the cloudwatchEventPermission, and an error, if there is any.
func (c *cloudwatchEventPermissions) Create(cloudwatchEventPermission *v1alpha1.CloudwatchEventPermission) (result *v1alpha1.CloudwatchEventPermission, err error) {
	result = &v1alpha1.CloudwatchEventPermission{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		Body(cloudwatchEventPermission).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudwatchEventPermission and updates it. Returns the server's representation of the cloudwatchEventPermission, and an error, if there is any.
func (c *cloudwatchEventPermissions) Update(cloudwatchEventPermission *v1alpha1.CloudwatchEventPermission) (result *v1alpha1.CloudwatchEventPermission, err error) {
	result = &v1alpha1.CloudwatchEventPermission{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		Name(cloudwatchEventPermission.Name).
		Body(cloudwatchEventPermission).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudwatchEventPermissions) UpdateStatus(cloudwatchEventPermission *v1alpha1.CloudwatchEventPermission) (result *v1alpha1.CloudwatchEventPermission, err error) {
	result = &v1alpha1.CloudwatchEventPermission{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		Name(cloudwatchEventPermission.Name).
		SubResource("status").
		Body(cloudwatchEventPermission).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudwatchEventPermission and deletes it. Returns an error if one occurs.
func (c *cloudwatchEventPermissions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudwatchEventPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudwatchEventPermission.
func (c *cloudwatchEventPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchEventPermission, err error) {
	result = &v1alpha1.CloudwatchEventPermission{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudwatcheventpermissions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
