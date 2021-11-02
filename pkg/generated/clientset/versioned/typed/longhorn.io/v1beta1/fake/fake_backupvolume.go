/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupVolumes implements BackupVolumeInterface
type FakeBackupVolumes struct {
	Fake *FakeLonghornV1beta1
	ns   string
}

var backupvolumesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backupvolumes"}

var backupvolumesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackupVolume"}

// Get takes name of the backupVolume, and returns the corresponding backupVolume object, and an error if there is any.
func (c *FakeBackupVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupvolumesResource, c.ns, name), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

// List takes label and field selectors, and returns the list of BackupVolumes that match those selectors.
func (c *FakeBackupVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupvolumesResource, backupvolumesKind, c.ns, opts), &v1beta1.BackupVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupVolumeList{ListMeta: obj.(*v1beta1.BackupVolumeList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupVolumes.
func (c *FakeBackupVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupvolumesResource, c.ns, opts))

}

// Create takes the representation of a backupVolume and creates it.  Returns the server's representation of the backupVolume, and an error, if there is any.
func (c *FakeBackupVolumes) Create(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.CreateOptions) (result *v1beta1.BackupVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupvolumesResource, c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

// Update takes the representation of a backupVolume and updates it. Returns the server's representation of the backupVolume, and an error, if there is any.
func (c *FakeBackupVolumes) Update(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.UpdateOptions) (result *v1beta1.BackupVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupvolumesResource, c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupVolumes) UpdateStatus(ctx context.Context, backupVolume *v1beta1.BackupVolume, opts v1.UpdateOptions) (*v1beta1.BackupVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupvolumesResource, "status", c.ns, backupVolume), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}

// Delete takes name of the backupVolume and deletes it. Returns an error if one occurs.
func (c *FakeBackupVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupvolumesResource, c.ns, name), &v1beta1.BackupVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupvolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupVolumeList{})
	return err
}

// Patch applies the patch and returns the patched backupVolume.
func (c *FakeBackupVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupvolumesResource, c.ns, name, pt, data, subresources...), &v1beta1.BackupVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupVolume), err
}
