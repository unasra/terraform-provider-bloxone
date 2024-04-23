/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"github.com/infobloxopen/bloxone-go-client/internal"
)

var ServiceBasePath = "/api/anycast/v1"

// APIClient manages communication with the BloxOne Anycast API API vv1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*internal.APIClient

	// API Services
	OnPremAnycastManagerAPI OnPremAnycastManagerAPI
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *internal.Configuration) *APIClient {
	c := &APIClient{}
	c.APIClient = internal.NewAPIClient(cfg)

	// API Services
	c.OnPremAnycastManagerAPI = (*OnPremAnycastManagerAPIService)(&c.Common)

	return c
}
