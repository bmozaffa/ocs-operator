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

package v1beta1

import (
	"time"

	v1beta1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1beta1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NFSsGetter has a method to return a NFSInterface.
// A group's client should implement this interface.
type NFSsGetter interface {
	NFSs(namespace string) NFSInterface
}

// NFSInterface has methods to work with NFS resources.
type NFSInterface interface {
	Create(*v1beta1.NFS) (*v1beta1.NFS, error)
	Update(*v1beta1.NFS) (*v1beta1.NFS, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.NFS, error)
	List(opts v1.ListOptions) (*v1beta1.NFSList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.NFS, err error)
	NFSExpansion
}

// nFSs implements NFSInterface
type nFSs struct {
	client rest.Interface
	ns     string
}

// newNFSs returns a NFSs
func newNFSs(c *EdgefsV1beta1Client, namespace string) *nFSs {
	return &nFSs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nFS, and returns the corresponding nFS object, and an error if there is any.
func (c *nFSs) Get(name string, options v1.GetOptions) (result *v1beta1.NFS, err error) {
	result = &v1beta1.NFS{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nfss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NFSs that match those selectors.
func (c *nFSs) List(opts v1.ListOptions) (result *v1beta1.NFSList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NFSList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nfss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nFSs.
func (c *nFSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nfss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nFS and creates it.  Returns the server's representation of the nFS, and an error, if there is any.
func (c *nFSs) Create(nFS *v1beta1.NFS) (result *v1beta1.NFS, err error) {
	result = &v1beta1.NFS{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nfss").
		Body(nFS).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nFS and updates it. Returns the server's representation of the nFS, and an error, if there is any.
func (c *nFSs) Update(nFS *v1beta1.NFS) (result *v1beta1.NFS, err error) {
	result = &v1beta1.NFS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nfss").
		Name(nFS.Name).
		Body(nFS).
		Do().
		Into(result)
	return
}

// Delete takes name of the nFS and deletes it. Returns an error if one occurs.
func (c *nFSs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nfss").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nFSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nfss").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nFS.
func (c *nFSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.NFS, err error) {
	result = &v1beta1.NFS{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nfss").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
