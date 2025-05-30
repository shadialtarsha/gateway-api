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

package v1alpha2

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	apisv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	applyconfigurationapisv1alpha2 "sigs.k8s.io/gateway-api/applyconfiguration/apis/v1alpha2"
	scheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

// UDPRoutesGetter has a method to return a UDPRouteInterface.
// A group's client should implement this interface.
type UDPRoutesGetter interface {
	UDPRoutes(namespace string) UDPRouteInterface
}

// UDPRouteInterface has methods to work with UDPRoute resources.
type UDPRouteInterface interface {
	Create(ctx context.Context, uDPRoute *apisv1alpha2.UDPRoute, opts v1.CreateOptions) (*apisv1alpha2.UDPRoute, error)
	Update(ctx context.Context, uDPRoute *apisv1alpha2.UDPRoute, opts v1.UpdateOptions) (*apisv1alpha2.UDPRoute, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, uDPRoute *apisv1alpha2.UDPRoute, opts v1.UpdateOptions) (*apisv1alpha2.UDPRoute, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*apisv1alpha2.UDPRoute, error)
	List(ctx context.Context, opts v1.ListOptions) (*apisv1alpha2.UDPRouteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apisv1alpha2.UDPRoute, err error)
	Apply(ctx context.Context, uDPRoute *applyconfigurationapisv1alpha2.UDPRouteApplyConfiguration, opts v1.ApplyOptions) (result *apisv1alpha2.UDPRoute, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, uDPRoute *applyconfigurationapisv1alpha2.UDPRouteApplyConfiguration, opts v1.ApplyOptions) (result *apisv1alpha2.UDPRoute, err error)
	UDPRouteExpansion
}

// uDPRoutes implements UDPRouteInterface
type uDPRoutes struct {
	*gentype.ClientWithListAndApply[*apisv1alpha2.UDPRoute, *apisv1alpha2.UDPRouteList, *applyconfigurationapisv1alpha2.UDPRouteApplyConfiguration]
}

// newUDPRoutes returns a UDPRoutes
func newUDPRoutes(c *GatewayV1alpha2Client, namespace string) *uDPRoutes {
	return &uDPRoutes{
		gentype.NewClientWithListAndApply[*apisv1alpha2.UDPRoute, *apisv1alpha2.UDPRouteList, *applyconfigurationapisv1alpha2.UDPRouteApplyConfiguration](
			"udproutes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *apisv1alpha2.UDPRoute { return &apisv1alpha2.UDPRoute{} },
			func() *apisv1alpha2.UDPRouteList { return &apisv1alpha2.UDPRouteList{} },
		),
	}
}
