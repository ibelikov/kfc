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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// GoogleServiceAccountsGetter has a method to return a GoogleServiceAccountInterface.
// A group's client should implement this interface.
type GoogleServiceAccountsGetter interface {
	GoogleServiceAccounts(namespace string) GoogleServiceAccountInterface
}

// GoogleServiceAccountInterface has methods to work with GoogleServiceAccount resources.
type GoogleServiceAccountInterface interface {
	Create(*v1alpha1.GoogleServiceAccount) (*v1alpha1.GoogleServiceAccount, error)
	Update(*v1alpha1.GoogleServiceAccount) (*v1alpha1.GoogleServiceAccount, error)
	UpdateStatus(*v1alpha1.GoogleServiceAccount) (*v1alpha1.GoogleServiceAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GoogleServiceAccount, error)
	List(opts v1.ListOptions) (*v1alpha1.GoogleServiceAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccount, err error)
	GoogleServiceAccountExpansion
}

// googleServiceAccounts implements GoogleServiceAccountInterface
type googleServiceAccounts struct {
	client rest.Interface
	ns     string
}

// newGoogleServiceAccounts returns a GoogleServiceAccounts
func newGoogleServiceAccounts(c *ModulesV1alpha1Client, namespace string) *googleServiceAccounts {
	return &googleServiceAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the googleServiceAccount, and returns the corresponding googleServiceAccount object, and an error if there is any.
func (c *googleServiceAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.GoogleServiceAccount, err error) {
	result = &v1alpha1.GoogleServiceAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GoogleServiceAccounts that match those selectors.
func (c *googleServiceAccounts) List(opts v1.ListOptions) (result *v1alpha1.GoogleServiceAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GoogleServiceAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested googleServiceAccounts.
func (c *googleServiceAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a googleServiceAccount and creates it.  Returns the server's representation of the googleServiceAccount, and an error, if there is any.
func (c *googleServiceAccounts) Create(googleServiceAccount *v1alpha1.GoogleServiceAccount) (result *v1alpha1.GoogleServiceAccount, err error) {
	result = &v1alpha1.GoogleServiceAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		Body(googleServiceAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a googleServiceAccount and updates it. Returns the server's representation of the googleServiceAccount, and an error, if there is any.
func (c *googleServiceAccounts) Update(googleServiceAccount *v1alpha1.GoogleServiceAccount) (result *v1alpha1.GoogleServiceAccount, err error) {
	result = &v1alpha1.GoogleServiceAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		Name(googleServiceAccount.Name).
		Body(googleServiceAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *googleServiceAccounts) UpdateStatus(googleServiceAccount *v1alpha1.GoogleServiceAccount) (result *v1alpha1.GoogleServiceAccount, err error) {
	result = &v1alpha1.GoogleServiceAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		Name(googleServiceAccount.Name).
		SubResource("status").
		Body(googleServiceAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the googleServiceAccount and deletes it. Returns an error if one occurs.
func (c *googleServiceAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *googleServiceAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched googleServiceAccount.
func (c *googleServiceAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GoogleServiceAccount, err error) {
	result = &v1alpha1.GoogleServiceAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("googleserviceaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
