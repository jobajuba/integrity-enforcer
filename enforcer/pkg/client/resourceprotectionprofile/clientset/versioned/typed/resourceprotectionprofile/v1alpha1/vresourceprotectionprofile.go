/*
Copyright 2020 IBM Corporation

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

	v1alpha1 "github.com/IBM/integrity-enforcer/enforcer/pkg/apis/resourceprotectionprofile/v1alpha1"
	scheme "github.com/IBM/integrity-enforcer/enforcer/pkg/client/resourceprotectionprofile/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceProtectionProfilesGetter has a method to return a ResourceProtectionProfileInterface.
// A group's client should implement this interface.
type ResourceProtectionProfilesGetter interface {
	ResourceProtectionProfiles(namespace string) ResourceProtectionProfileInterface
}

// ResourceProtectionProfileInterface has methods to work with ResourceProtectionProfile resources.
type ResourceProtectionProfileInterface interface {
	Create(*v1alpha1.ResourceProtectionProfile) (*v1alpha1.ResourceProtectionProfile, error)
	Update(*v1alpha1.ResourceProtectionProfile) (*v1alpha1.ResourceProtectionProfile, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ResourceProtectionProfile, error)
	List(opts v1.ListOptions) (*v1alpha1.ResourceProtectionProfileList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceProtectionProfile, err error)
	ResourceProtectionProfileExpansion
}

// resourceProtectionProfiles implements ResourceProtectionProfileInterface
type resourceProtectionProfiles struct {
	client rest.Interface
	ns     string
}

// newResourceProtectionProfiles returns a ResourceProtectionProfiles
func newResourceProtectionProfiles(c *ResearchV1alpha1Client, namespace string) *resourceProtectionProfiles {
	return &resourceProtectionProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceProtectionProfile, and returns the corresponding resourceProtectionProfile object, and an error if there is any.
func (c *resourceProtectionProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceProtectionProfile, err error) {
	result = &v1alpha1.ResourceProtectionProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceProtectionProfiles that match those selectors.
func (c *resourceProtectionProfiles) List(opts v1.ListOptions) (result *v1alpha1.ResourceProtectionProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceProtectionProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceProtectionProfiles.
func (c *resourceProtectionProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a resourceProtectionProfile and creates it.  Returns the server's representation of the resourceProtectionProfile, and an error, if there is any.
func (c *resourceProtectionProfiles) Create(resourceProtectionProfile *v1alpha1.ResourceProtectionProfile) (result *v1alpha1.ResourceProtectionProfile, err error) {
	result = &v1alpha1.ResourceProtectionProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		Body(resourceProtectionProfile).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resourceProtectionProfile and updates it. Returns the server's representation of the resourceProtectionProfile, and an error, if there is any.
func (c *resourceProtectionProfiles) Update(resourceProtectionProfile *v1alpha1.ResourceProtectionProfile) (result *v1alpha1.ResourceProtectionProfile, err error) {
	result = &v1alpha1.ResourceProtectionProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		Name(resourceProtectionProfile.Name).
		Body(resourceProtectionProfile).
		Do().
		Into(result)
	return
}

// Delete takes name of the resourceProtectionProfile and deletes it. Returns an error if one occurs.
func (c *resourceProtectionProfiles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceProtectionProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resourceProtectionProfile.
func (c *resourceProtectionProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceProtectionProfile, err error) {
	result = &v1alpha1.ResourceProtectionProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceprotectionprofiles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
