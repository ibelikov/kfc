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

// LicensemanagerLicenseConfigurationsGetter has a method to return a LicensemanagerLicenseConfigurationInterface.
// A group's client should implement this interface.
type LicensemanagerLicenseConfigurationsGetter interface {
	LicensemanagerLicenseConfigurations(namespace string) LicensemanagerLicenseConfigurationInterface
}

// LicensemanagerLicenseConfigurationInterface has methods to work with LicensemanagerLicenseConfiguration resources.
type LicensemanagerLicenseConfigurationInterface interface {
	Create(*v1alpha1.LicensemanagerLicenseConfiguration) (*v1alpha1.LicensemanagerLicenseConfiguration, error)
	Update(*v1alpha1.LicensemanagerLicenseConfiguration) (*v1alpha1.LicensemanagerLicenseConfiguration, error)
	UpdateStatus(*v1alpha1.LicensemanagerLicenseConfiguration) (*v1alpha1.LicensemanagerLicenseConfiguration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LicensemanagerLicenseConfiguration, error)
	List(opts v1.ListOptions) (*v1alpha1.LicensemanagerLicenseConfigurationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error)
	LicensemanagerLicenseConfigurationExpansion
}

// licensemanagerLicenseConfigurations implements LicensemanagerLicenseConfigurationInterface
type licensemanagerLicenseConfigurations struct {
	client rest.Interface
	ns     string
}

// newLicensemanagerLicenseConfigurations returns a LicensemanagerLicenseConfigurations
func newLicensemanagerLicenseConfigurations(c *AwsV1alpha1Client, namespace string) *licensemanagerLicenseConfigurations {
	return &licensemanagerLicenseConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the licensemanagerLicenseConfiguration, and returns the corresponding licensemanagerLicenseConfiguration object, and an error if there is any.
func (c *licensemanagerLicenseConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	result = &v1alpha1.LicensemanagerLicenseConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LicensemanagerLicenseConfigurations that match those selectors.
func (c *licensemanagerLicenseConfigurations) List(opts v1.ListOptions) (result *v1alpha1.LicensemanagerLicenseConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LicensemanagerLicenseConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested licensemanagerLicenseConfigurations.
func (c *licensemanagerLicenseConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a licensemanagerLicenseConfiguration and creates it.  Returns the server's representation of the licensemanagerLicenseConfiguration, and an error, if there is any.
func (c *licensemanagerLicenseConfigurations) Create(licensemanagerLicenseConfiguration *v1alpha1.LicensemanagerLicenseConfiguration) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	result = &v1alpha1.LicensemanagerLicenseConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		Body(licensemanagerLicenseConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a licensemanagerLicenseConfiguration and updates it. Returns the server's representation of the licensemanagerLicenseConfiguration, and an error, if there is any.
func (c *licensemanagerLicenseConfigurations) Update(licensemanagerLicenseConfiguration *v1alpha1.LicensemanagerLicenseConfiguration) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	result = &v1alpha1.LicensemanagerLicenseConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		Name(licensemanagerLicenseConfiguration.Name).
		Body(licensemanagerLicenseConfiguration).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *licensemanagerLicenseConfigurations) UpdateStatus(licensemanagerLicenseConfiguration *v1alpha1.LicensemanagerLicenseConfiguration) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	result = &v1alpha1.LicensemanagerLicenseConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		Name(licensemanagerLicenseConfiguration.Name).
		SubResource("status").
		Body(licensemanagerLicenseConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the licensemanagerLicenseConfiguration and deletes it. Returns an error if one occurs.
func (c *licensemanagerLicenseConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *licensemanagerLicenseConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched licensemanagerLicenseConfiguration.
func (c *licensemanagerLicenseConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	result = &v1alpha1.LicensemanagerLicenseConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("licensemanagerlicenseconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}