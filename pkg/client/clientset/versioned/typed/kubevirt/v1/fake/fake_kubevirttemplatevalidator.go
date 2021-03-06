/*
Copyright The Kubernetes Authors.

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

package fake

import (
	kubevirtv1 "github.com/MarSik/kubevirt-ssp-operator/pkg/apis/kubevirt/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubevirtTemplateValidators implements KubevirtTemplateValidatorInterface
type FakeKubevirtTemplateValidators struct {
	Fake *FakeKubevirtV1
	ns   string
}

var kubevirttemplatevalidatorsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "kubevirttemplatevalidators"}

var kubevirttemplatevalidatorsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "KubevirtTemplateValidator"}

// Get takes name of the kubevirtTemplateValidator, and returns the corresponding kubevirtTemplateValidator object, and an error if there is any.
func (c *FakeKubevirtTemplateValidators) Get(name string, options v1.GetOptions) (result *kubevirtv1.KubevirtTemplateValidator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubevirttemplatevalidatorsResource, c.ns, name), &kubevirtv1.KubevirtTemplateValidator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtTemplateValidator), err
}

// List takes label and field selectors, and returns the list of KubevirtTemplateValidators that match those selectors.
func (c *FakeKubevirtTemplateValidators) List(opts v1.ListOptions) (result *kubevirtv1.KubevirtTemplateValidatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubevirttemplatevalidatorsResource, kubevirttemplatevalidatorsKind, c.ns, opts), &kubevirtv1.KubevirtTemplateValidatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubevirtv1.KubevirtTemplateValidatorList{ListMeta: obj.(*kubevirtv1.KubevirtTemplateValidatorList).ListMeta}
	for _, item := range obj.(*kubevirtv1.KubevirtTemplateValidatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubevirtTemplateValidators.
func (c *FakeKubevirtTemplateValidators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubevirttemplatevalidatorsResource, c.ns, opts))

}

// Create takes the representation of a kubevirtTemplateValidator and creates it.  Returns the server's representation of the kubevirtTemplateValidator, and an error, if there is any.
func (c *FakeKubevirtTemplateValidators) Create(kubevirtTemplateValidator *kubevirtv1.KubevirtTemplateValidator) (result *kubevirtv1.KubevirtTemplateValidator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubevirttemplatevalidatorsResource, c.ns, kubevirtTemplateValidator), &kubevirtv1.KubevirtTemplateValidator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtTemplateValidator), err
}

// Update takes the representation of a kubevirtTemplateValidator and updates it. Returns the server's representation of the kubevirtTemplateValidator, and an error, if there is any.
func (c *FakeKubevirtTemplateValidators) Update(kubevirtTemplateValidator *kubevirtv1.KubevirtTemplateValidator) (result *kubevirtv1.KubevirtTemplateValidator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubevirttemplatevalidatorsResource, c.ns, kubevirtTemplateValidator), &kubevirtv1.KubevirtTemplateValidator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtTemplateValidator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubevirtTemplateValidators) UpdateStatus(kubevirtTemplateValidator *kubevirtv1.KubevirtTemplateValidator) (*kubevirtv1.KubevirtTemplateValidator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubevirttemplatevalidatorsResource, "status", c.ns, kubevirtTemplateValidator), &kubevirtv1.KubevirtTemplateValidator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtTemplateValidator), err
}

// Delete takes name of the kubevirtTemplateValidator and deletes it. Returns an error if one occurs.
func (c *FakeKubevirtTemplateValidators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubevirttemplatevalidatorsResource, c.ns, name), &kubevirtv1.KubevirtTemplateValidator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubevirtTemplateValidators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubevirttemplatevalidatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kubevirtv1.KubevirtTemplateValidatorList{})
	return err
}

// Patch applies the patch and returns the patched kubevirtTemplateValidator.
func (c *FakeKubevirtTemplateValidators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubevirtv1.KubevirtTemplateValidator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubevirttemplatevalidatorsResource, c.ns, name, data, subresources...), &kubevirtv1.KubevirtTemplateValidator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubevirtv1.KubevirtTemplateValidator), err
}
