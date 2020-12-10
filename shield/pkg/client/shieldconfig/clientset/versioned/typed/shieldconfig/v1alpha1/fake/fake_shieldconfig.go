//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/IBM/integrity-enforcer/shield/pkg/apis/shieldconfig/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShieldConfigs implements ShieldConfigInterface
type FakeShieldConfigs struct {
	Fake *FakeApisV1alpha1
	ns   string
}

var shieldconfigsResource = schema.GroupVersionResource{Group: "apis.integrityshield.io", Version: "v1alpha1", Resource: "shieldconfigs"}

var shieldconfigsKind = schema.GroupVersionKind{Group: "apis.integrityshield.io", Version: "v1alpha1", Kind: "ShieldConfig"}

// Get takes name of the shieldConfig, and returns the corresponding shieldConfig object, and an error if there is any.
func (c *FakeShieldConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShieldConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shieldconfigsResource, c.ns, name), &v1alpha1.ShieldConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShieldConfig), err
}

// List takes label and field selectors, and returns the list of ShieldConfigs that match those selectors.
func (c *FakeShieldConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShieldConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shieldconfigsResource, shieldconfigsKind, c.ns, opts), &v1alpha1.ShieldConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShieldConfigList{ListMeta: obj.(*v1alpha1.ShieldConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShieldConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shieldConfigs.
func (c *FakeShieldConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shieldconfigsResource, c.ns, opts))

}

// Create takes the representation of a shieldConfig and creates it.  Returns the server's representation of the shieldConfig, and an error, if there is any.
func (c *FakeShieldConfigs) Create(ctx context.Context, shieldConfig *v1alpha1.ShieldConfig, opts v1.CreateOptions) (result *v1alpha1.ShieldConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shieldconfigsResource, c.ns, shieldConfig), &v1alpha1.ShieldConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShieldConfig), err
}

// Update takes the representation of a shieldConfig and updates it. Returns the server's representation of the shieldConfig, and an error, if there is any.
func (c *FakeShieldConfigs) Update(ctx context.Context, shieldConfig *v1alpha1.ShieldConfig, opts v1.UpdateOptions) (result *v1alpha1.ShieldConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shieldconfigsResource, c.ns, shieldConfig), &v1alpha1.ShieldConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShieldConfig), err
}

// Delete takes name of the shieldConfig and deletes it. Returns an error if one occurs.
func (c *FakeShieldConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shieldconfigsResource, c.ns, name), &v1alpha1.ShieldConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShieldConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shieldconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShieldConfigList{})
	return err
}

// Patch applies the patch and returns the patched shieldConfig.
func (c *FakeShieldConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShieldConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shieldconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShieldConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShieldConfig), err
}
