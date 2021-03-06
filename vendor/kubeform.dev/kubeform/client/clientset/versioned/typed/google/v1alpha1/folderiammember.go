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

// FolderIamMembersGetter has a method to return a FolderIamMemberInterface.
// A group's client should implement this interface.
type FolderIamMembersGetter interface {
	FolderIamMembers(namespace string) FolderIamMemberInterface
}

// FolderIamMemberInterface has methods to work with FolderIamMember resources.
type FolderIamMemberInterface interface {
	Create(*v1alpha1.FolderIamMember) (*v1alpha1.FolderIamMember, error)
	Update(*v1alpha1.FolderIamMember) (*v1alpha1.FolderIamMember, error)
	UpdateStatus(*v1alpha1.FolderIamMember) (*v1alpha1.FolderIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FolderIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.FolderIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FolderIamMember, err error)
	FolderIamMemberExpansion
}

// folderIamMembers implements FolderIamMemberInterface
type folderIamMembers struct {
	client rest.Interface
	ns     string
}

// newFolderIamMembers returns a FolderIamMembers
func newFolderIamMembers(c *GoogleV1alpha1Client, namespace string) *folderIamMembers {
	return &folderIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the folderIamMember, and returns the corresponding folderIamMember object, and an error if there is any.
func (c *folderIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.FolderIamMember, err error) {
	result = &v1alpha1.FolderIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("folderiammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FolderIamMembers that match those selectors.
func (c *folderIamMembers) List(opts v1.ListOptions) (result *v1alpha1.FolderIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FolderIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("folderiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested folderIamMembers.
func (c *folderIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("folderiammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a folderIamMember and creates it.  Returns the server's representation of the folderIamMember, and an error, if there is any.
func (c *folderIamMembers) Create(folderIamMember *v1alpha1.FolderIamMember) (result *v1alpha1.FolderIamMember, err error) {
	result = &v1alpha1.FolderIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("folderiammembers").
		Body(folderIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a folderIamMember and updates it. Returns the server's representation of the folderIamMember, and an error, if there is any.
func (c *folderIamMembers) Update(folderIamMember *v1alpha1.FolderIamMember) (result *v1alpha1.FolderIamMember, err error) {
	result = &v1alpha1.FolderIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("folderiammembers").
		Name(folderIamMember.Name).
		Body(folderIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *folderIamMembers) UpdateStatus(folderIamMember *v1alpha1.FolderIamMember) (result *v1alpha1.FolderIamMember, err error) {
	result = &v1alpha1.FolderIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("folderiammembers").
		Name(folderIamMember.Name).
		SubResource("status").
		Body(folderIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the folderIamMember and deletes it. Returns an error if one occurs.
func (c *folderIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("folderiammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *folderIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("folderiammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched folderIamMember.
func (c *folderIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FolderIamMember, err error) {
	result = &v1alpha1.FolderIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("folderiammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
