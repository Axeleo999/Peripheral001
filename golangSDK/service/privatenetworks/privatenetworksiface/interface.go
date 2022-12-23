// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package privatenetworksiface provides an interface to enable mocking the AWS Private 5G service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package privatenetworksiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/privatenetworks"
)

// PrivateNetworksAPI provides an interface to enable mocking the
// privatenetworks.PrivateNetworks service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Private 5G.
//	func myFunc(svc privatenetworksiface.PrivateNetworksAPI) bool {
//	    // Make svc.AcknowledgeOrderReceipt request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := privatenetworks.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPrivateNetworksClient struct {
//	    privatenetworksiface.PrivateNetworksAPI
//	}
//	func (m *mockPrivateNetworksClient) AcknowledgeOrderReceipt(input *privatenetworks.AcknowledgeOrderReceiptInput) (*privatenetworks.AcknowledgeOrderReceiptOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPrivateNetworksClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PrivateNetworksAPI interface {
	AcknowledgeOrderReceipt(*privatenetworks.AcknowledgeOrderReceiptInput) (*privatenetworks.AcknowledgeOrderReceiptOutput, error)
	AcknowledgeOrderReceiptWithContext(aws.Context, *privatenetworks.AcknowledgeOrderReceiptInput, ...request.Option) (*privatenetworks.AcknowledgeOrderReceiptOutput, error)
	AcknowledgeOrderReceiptRequest(*privatenetworks.AcknowledgeOrderReceiptInput) (*request.Request, *privatenetworks.AcknowledgeOrderReceiptOutput)

	ActivateDeviceIdentifier(*privatenetworks.ActivateDeviceIdentifierInput) (*privatenetworks.ActivateDeviceIdentifierOutput, error)
	ActivateDeviceIdentifierWithContext(aws.Context, *privatenetworks.ActivateDeviceIdentifierInput, ...request.Option) (*privatenetworks.ActivateDeviceIdentifierOutput, error)
	ActivateDeviceIdentifierRequest(*privatenetworks.ActivateDeviceIdentifierInput) (*request.Request, *privatenetworks.ActivateDeviceIdentifierOutput)

	ActivateNetworkSite(*privatenetworks.ActivateNetworkSiteInput) (*privatenetworks.ActivateNetworkSiteOutput, error)
	ActivateNetworkSiteWithContext(aws.Context, *privatenetworks.ActivateNetworkSiteInput, ...request.Option) (*privatenetworks.ActivateNetworkSiteOutput, error)
	ActivateNetworkSiteRequest(*privatenetworks.ActivateNetworkSiteInput) (*request.Request, *privatenetworks.ActivateNetworkSiteOutput)

	ConfigureAccessPoint(*privatenetworks.ConfigureAccessPointInput) (*privatenetworks.ConfigureAccessPointOutput, error)
	ConfigureAccessPointWithContext(aws.Context, *privatenetworks.ConfigureAccessPointInput, ...request.Option) (*privatenetworks.ConfigureAccessPointOutput, error)
	ConfigureAccessPointRequest(*privatenetworks.ConfigureAccessPointInput) (*request.Request, *privatenetworks.ConfigureAccessPointOutput)

	CreateNetwork(*privatenetworks.CreateNetworkInput) (*privatenetworks.CreateNetworkOutput, error)
	CreateNetworkWithContext(aws.Context, *privatenetworks.CreateNetworkInput, ...request.Option) (*privatenetworks.CreateNetworkOutput, error)
	CreateNetworkRequest(*privatenetworks.CreateNetworkInput) (*request.Request, *privatenetworks.CreateNetworkOutput)

	CreateNetworkSite(*privatenetworks.CreateNetworkSiteInput) (*privatenetworks.CreateNetworkSiteOutput, error)
	CreateNetworkSiteWithContext(aws.Context, *privatenetworks.CreateNetworkSiteInput, ...request.Option) (*privatenetworks.CreateNetworkSiteOutput, error)
	CreateNetworkSiteRequest(*privatenetworks.CreateNetworkSiteInput) (*request.Request, *privatenetworks.CreateNetworkSiteOutput)

	DeactivateDeviceIdentifier(*privatenetworks.DeactivateDeviceIdentifierInput) (*privatenetworks.DeactivateDeviceIdentifierOutput, error)
	DeactivateDeviceIdentifierWithContext(aws.Context, *privatenetworks.DeactivateDeviceIdentifierInput, ...request.Option) (*privatenetworks.DeactivateDeviceIdentifierOutput, error)
	DeactivateDeviceIdentifierRequest(*privatenetworks.DeactivateDeviceIdentifierInput) (*request.Request, *privatenetworks.DeactivateDeviceIdentifierOutput)

	DeleteNetwork(*privatenetworks.DeleteNetworkInput) (*privatenetworks.DeleteNetworkOutput, error)
	DeleteNetworkWithContext(aws.Context, *privatenetworks.DeleteNetworkInput, ...request.Option) (*privatenetworks.DeleteNetworkOutput, error)
	DeleteNetworkRequest(*privatenetworks.DeleteNetworkInput) (*request.Request, *privatenetworks.DeleteNetworkOutput)

	DeleteNetworkSite(*privatenetworks.DeleteNetworkSiteInput) (*privatenetworks.DeleteNetworkSiteOutput, error)
	DeleteNetworkSiteWithContext(aws.Context, *privatenetworks.DeleteNetworkSiteInput, ...request.Option) (*privatenetworks.DeleteNetworkSiteOutput, error)
	DeleteNetworkSiteRequest(*privatenetworks.DeleteNetworkSiteInput) (*request.Request, *privatenetworks.DeleteNetworkSiteOutput)

	GetDeviceIdentifier(*privatenetworks.GetDeviceIdentifierInput) (*privatenetworks.GetDeviceIdentifierOutput, error)
	GetDeviceIdentifierWithContext(aws.Context, *privatenetworks.GetDeviceIdentifierInput, ...request.Option) (*privatenetworks.GetDeviceIdentifierOutput, error)
	GetDeviceIdentifierRequest(*privatenetworks.GetDeviceIdentifierInput) (*request.Request, *privatenetworks.GetDeviceIdentifierOutput)

	GetNetwork(*privatenetworks.GetNetworkInput) (*privatenetworks.GetNetworkOutput, error)
	GetNetworkWithContext(aws.Context, *privatenetworks.GetNetworkInput, ...request.Option) (*privatenetworks.GetNetworkOutput, error)
	GetNetworkRequest(*privatenetworks.GetNetworkInput) (*request.Request, *privatenetworks.GetNetworkOutput)

	GetNetworkResource(*privatenetworks.GetNetworkResourceInput) (*privatenetworks.GetNetworkResourceOutput, error)
	GetNetworkResourceWithContext(aws.Context, *privatenetworks.GetNetworkResourceInput, ...request.Option) (*privatenetworks.GetNetworkResourceOutput, error)
	GetNetworkResourceRequest(*privatenetworks.GetNetworkResourceInput) (*request.Request, *privatenetworks.GetNetworkResourceOutput)

	GetNetworkSite(*privatenetworks.GetNetworkSiteInput) (*privatenetworks.GetNetworkSiteOutput, error)
	GetNetworkSiteWithContext(aws.Context, *privatenetworks.GetNetworkSiteInput, ...request.Option) (*privatenetworks.GetNetworkSiteOutput, error)
	GetNetworkSiteRequest(*privatenetworks.GetNetworkSiteInput) (*request.Request, *privatenetworks.GetNetworkSiteOutput)

	GetOrder(*privatenetworks.GetOrderInput) (*privatenetworks.GetOrderOutput, error)
	GetOrderWithContext(aws.Context, *privatenetworks.GetOrderInput, ...request.Option) (*privatenetworks.GetOrderOutput, error)
	GetOrderRequest(*privatenetworks.GetOrderInput) (*request.Request, *privatenetworks.GetOrderOutput)

	ListDeviceIdentifiers(*privatenetworks.ListDeviceIdentifiersInput) (*privatenetworks.ListDeviceIdentifiersOutput, error)
	ListDeviceIdentifiersWithContext(aws.Context, *privatenetworks.ListDeviceIdentifiersInput, ...request.Option) (*privatenetworks.ListDeviceIdentifiersOutput, error)
	ListDeviceIdentifiersRequest(*privatenetworks.ListDeviceIdentifiersInput) (*request.Request, *privatenetworks.ListDeviceIdentifiersOutput)

	ListDeviceIdentifiersPages(*privatenetworks.ListDeviceIdentifiersInput, func(*privatenetworks.ListDeviceIdentifiersOutput, bool) bool) error
	ListDeviceIdentifiersPagesWithContext(aws.Context, *privatenetworks.ListDeviceIdentifiersInput, func(*privatenetworks.ListDeviceIdentifiersOutput, bool) bool, ...request.Option) error

	ListNetworkResources(*privatenetworks.ListNetworkResourcesInput) (*privatenetworks.ListNetworkResourcesOutput, error)
	ListNetworkResourcesWithContext(aws.Context, *privatenetworks.ListNetworkResourcesInput, ...request.Option) (*privatenetworks.ListNetworkResourcesOutput, error)
	ListNetworkResourcesRequest(*privatenetworks.ListNetworkResourcesInput) (*request.Request, *privatenetworks.ListNetworkResourcesOutput)

	ListNetworkResourcesPages(*privatenetworks.ListNetworkResourcesInput, func(*privatenetworks.ListNetworkResourcesOutput, bool) bool) error
	ListNetworkResourcesPagesWithContext(aws.Context, *privatenetworks.ListNetworkResourcesInput, func(*privatenetworks.ListNetworkResourcesOutput, bool) bool, ...request.Option) error

	ListNetworkSites(*privatenetworks.ListNetworkSitesInput) (*privatenetworks.ListNetworkSitesOutput, error)
	ListNetworkSitesWithContext(aws.Context, *privatenetworks.ListNetworkSitesInput, ...request.Option) (*privatenetworks.ListNetworkSitesOutput, error)
	ListNetworkSitesRequest(*privatenetworks.ListNetworkSitesInput) (*request.Request, *privatenetworks.ListNetworkSitesOutput)

	ListNetworkSitesPages(*privatenetworks.ListNetworkSitesInput, func(*privatenetworks.ListNetworkSitesOutput, bool) bool) error
	ListNetworkSitesPagesWithContext(aws.Context, *privatenetworks.ListNetworkSitesInput, func(*privatenetworks.ListNetworkSitesOutput, bool) bool, ...request.Option) error

	ListNetworks(*privatenetworks.ListNetworksInput) (*privatenetworks.ListNetworksOutput, error)
	ListNetworksWithContext(aws.Context, *privatenetworks.ListNetworksInput, ...request.Option) (*privatenetworks.ListNetworksOutput, error)
	ListNetworksRequest(*privatenetworks.ListNetworksInput) (*request.Request, *privatenetworks.ListNetworksOutput)

	ListNetworksPages(*privatenetworks.ListNetworksInput, func(*privatenetworks.ListNetworksOutput, bool) bool) error
	ListNetworksPagesWithContext(aws.Context, *privatenetworks.ListNetworksInput, func(*privatenetworks.ListNetworksOutput, bool) bool, ...request.Option) error

	ListOrders(*privatenetworks.ListOrdersInput) (*privatenetworks.ListOrdersOutput, error)
	ListOrdersWithContext(aws.Context, *privatenetworks.ListOrdersInput, ...request.Option) (*privatenetworks.ListOrdersOutput, error)
	ListOrdersRequest(*privatenetworks.ListOrdersInput) (*request.Request, *privatenetworks.ListOrdersOutput)

	ListOrdersPages(*privatenetworks.ListOrdersInput, func(*privatenetworks.ListOrdersOutput, bool) bool) error
	ListOrdersPagesWithContext(aws.Context, *privatenetworks.ListOrdersInput, func(*privatenetworks.ListOrdersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*privatenetworks.ListTagsForResourceInput) (*privatenetworks.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *privatenetworks.ListTagsForResourceInput, ...request.Option) (*privatenetworks.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*privatenetworks.ListTagsForResourceInput) (*request.Request, *privatenetworks.ListTagsForResourceOutput)

	Ping(*privatenetworks.PingInput) (*privatenetworks.PingOutput, error)
	PingWithContext(aws.Context, *privatenetworks.PingInput, ...request.Option) (*privatenetworks.PingOutput, error)
	PingRequest(*privatenetworks.PingInput) (*request.Request, *privatenetworks.PingOutput)

	TagResource(*privatenetworks.TagResourceInput) (*privatenetworks.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *privatenetworks.TagResourceInput, ...request.Option) (*privatenetworks.TagResourceOutput, error)
	TagResourceRequest(*privatenetworks.TagResourceInput) (*request.Request, *privatenetworks.TagResourceOutput)

	UntagResource(*privatenetworks.UntagResourceInput) (*privatenetworks.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *privatenetworks.UntagResourceInput, ...request.Option) (*privatenetworks.UntagResourceOutput, error)
	UntagResourceRequest(*privatenetworks.UntagResourceInput) (*request.Request, *privatenetworks.UntagResourceOutput)

	UpdateNetworkSite(*privatenetworks.UpdateNetworkSiteInput) (*privatenetworks.UpdateNetworkSiteOutput, error)
	UpdateNetworkSiteWithContext(aws.Context, *privatenetworks.UpdateNetworkSiteInput, ...request.Option) (*privatenetworks.UpdateNetworkSiteOutput, error)
	UpdateNetworkSiteRequest(*privatenetworks.UpdateNetworkSiteInput) (*request.Request, *privatenetworks.UpdateNetworkSiteOutput)

	UpdateNetworkSitePlan(*privatenetworks.UpdateNetworkSitePlanInput) (*privatenetworks.UpdateNetworkSitePlanOutput, error)
	UpdateNetworkSitePlanWithContext(aws.Context, *privatenetworks.UpdateNetworkSitePlanInput, ...request.Option) (*privatenetworks.UpdateNetworkSitePlanOutput, error)
	UpdateNetworkSitePlanRequest(*privatenetworks.UpdateNetworkSitePlanInput) (*request.Request, *privatenetworks.UpdateNetworkSitePlanOutput)
}

var _ PrivateNetworksAPI = (*privatenetworks.PrivateNetworks)(nil)