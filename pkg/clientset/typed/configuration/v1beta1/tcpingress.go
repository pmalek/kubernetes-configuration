/*
Copyright 2021 Kong, Inc.

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
	context "context"

	configurationv1beta1 "github.com/kong/kubernetes-configuration/api/configuration/v1beta1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TCPIngressesGetter has a method to return a TCPIngressInterface.
// A group's client should implement this interface.
type TCPIngressesGetter interface {
	TCPIngresses(namespace string) TCPIngressInterface
}

// TCPIngressInterface has methods to work with TCPIngress resources.
type TCPIngressInterface interface {
	Create(ctx context.Context, tCPIngress *configurationv1beta1.TCPIngress, opts v1.CreateOptions) (*configurationv1beta1.TCPIngress, error)
	Update(ctx context.Context, tCPIngress *configurationv1beta1.TCPIngress, opts v1.UpdateOptions) (*configurationv1beta1.TCPIngress, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, tCPIngress *configurationv1beta1.TCPIngress, opts v1.UpdateOptions) (*configurationv1beta1.TCPIngress, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*configurationv1beta1.TCPIngress, error)
	List(ctx context.Context, opts v1.ListOptions) (*configurationv1beta1.TCPIngressList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configurationv1beta1.TCPIngress, err error)
	TCPIngressExpansion
}

// tCPIngresses implements TCPIngressInterface
type tCPIngresses struct {
	*gentype.ClientWithList[*configurationv1beta1.TCPIngress, *configurationv1beta1.TCPIngressList]
}

// newTCPIngresses returns a TCPIngresses
func newTCPIngresses(c *ConfigurationV1beta1Client, namespace string) *tCPIngresses {
	return &tCPIngresses{
		gentype.NewClientWithList[*configurationv1beta1.TCPIngress, *configurationv1beta1.TCPIngressList](
			"tcpingresses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *configurationv1beta1.TCPIngress { return &configurationv1beta1.TCPIngress{} },
			func() *configurationv1beta1.TCPIngressList { return &configurationv1beta1.TCPIngressList{} },
		),
	}
}
