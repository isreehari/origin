package fake

import (
	api "github.com/openshift/origin/pkg/template/api"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplateInstances implements TemplateInstanceInterface
type FakeTemplateInstances struct {
	Fake *FakeTemplate
	ns   string
}

var templateinstancesResource = schema.GroupVersionResource{Group: "template.openshift.io", Version: "", Resource: "templateinstances"}

func (c *FakeTemplateInstances) Create(templateInstance *api.TemplateInstance) (result *api.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templateinstancesResource, c.ns, templateInstance), &api.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.TemplateInstance), err
}

func (c *FakeTemplateInstances) Update(templateInstance *api.TemplateInstance) (result *api.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templateinstancesResource, c.ns, templateInstance), &api.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.TemplateInstance), err
}

func (c *FakeTemplateInstances) UpdateStatus(templateInstance *api.TemplateInstance) (*api.TemplateInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(templateinstancesResource, "status", c.ns, templateInstance), &api.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.TemplateInstance), err
}

func (c *FakeTemplateInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(templateinstancesResource, c.ns, name), &api.TemplateInstance{})

	return err
}

func (c *FakeTemplateInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templateinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &api.TemplateInstanceList{})
	return err
}

func (c *FakeTemplateInstances) Get(name string, options v1.GetOptions) (result *api.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templateinstancesResource, c.ns, name), &api.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.TemplateInstance), err
}

func (c *FakeTemplateInstances) List(opts v1.ListOptions) (result *api.TemplateInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templateinstancesResource, c.ns, opts), &api.TemplateInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &api.TemplateInstanceList{}
	for _, item := range obj.(*api.TemplateInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templateInstances.
func (c *FakeTemplateInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templateinstancesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched templateInstance.
func (c *FakeTemplateInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *api.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templateinstancesResource, c.ns, name, data, subresources...), &api.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*api.TemplateInstance), err
}
