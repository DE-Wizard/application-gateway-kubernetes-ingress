// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package appgw

import (
	networkingv1alpha3 "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

type istioMatchIdentifier struct {
	Namespace      string
	VirtualService *v1alpha3.VirtualService
	Rule           *networkingv1alpha3.HTTPRoute
	Match          *networkingv1alpha3.HTTPMatchRequest
	Destinations   []*networkingv1alpha3.Destination
	Gateways       []string
}

type istioVirtualServiceIdentifier struct {
	Namespace string
	Name      string
}

type istioDestinationIdentifier struct {
	serviceIdentifier
	istioVirtualServiceIdentifier

	DestinationHost   string
	DestinationSubset string
	DestinationPort   uint32
}

const (
	ProtocolHTTP  string = "HTTP"
	ProtocolHTTPS string = "HTTPS"
	ProtocolGRPC  string = "GRPC"
	ProtocolHTTP2 string = "HTTP2"
	ProtocolMongo string = "Mongo"
	ProtocolTCP   string = "TCP"
)
