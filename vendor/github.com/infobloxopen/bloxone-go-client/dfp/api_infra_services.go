/*
DFP API

BloxOne Cloud is a SaaS offering designed to provide protection to devices on and off-premises, including roaming, remote, and branch offices. It provides visibility into infected and compromised devices, prevents DNS-based data exfiltration, and automatically stops device communications with command-and-control servers (C&Cs) and botnets, in addition to providing recursive DNS services in the cloud. You can access the services by deploying the BloxOne Endpoint agent or the DNS forwarding proxy.  For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.  By implementing the DNS forwarding proxy, you can rest assured that BloxOne Cloud effectively enforces DNS client-based security policies at your remote sites. On-premises devices that send DNS queries reveal their actual client IP addresses (instead of their NAT IP address), which allows BloxOne Cloud to apply the security policies applicable to the respective endpoints and identify infected clients.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfp

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type InfraServicesAPI interface {
	/*
		CreateOrUpdateDfpService Update DNS Forwarding Proxy services.

		Use this method to update resolvers for the specified DNS Forwarding Proxy Service.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param payloadServiceId The DNS Forwarding Proxy Service ID object identifier.
		@return InfraServicesAPICreateOrUpdateDfpServiceRequest
	*/
	CreateOrUpdateDfpService(ctx context.Context, payloadServiceId string) InfraServicesAPICreateOrUpdateDfpServiceRequest

	// CreateOrUpdateDfpServiceExecute executes the request
	//  @return DfpCreateOrUpdateResponse
	CreateOrUpdateDfpServiceExecute(r InfraServicesAPICreateOrUpdateDfpServiceRequest) (*DfpCreateOrUpdateResponse, *http.Response, error)
	/*
		ListDfpServices List DNS Forwarding Proxy services.

		Use this method to retrieve information on all DNS Forwarding Proxy services.



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return InfraServicesAPIListDfpServicesRequest
	*/
	ListDfpServices(ctx context.Context) InfraServicesAPIListDfpServicesRequest

	// ListDfpServicesExecute executes the request
	//  @return DfpListResponse
	ListDfpServicesExecute(r InfraServicesAPIListDfpServicesRequest) (*DfpListResponse, *http.Response, error)
	/*
		ReadDfpService Read DNS Forwarding Proxy services.

		Use this method to retrieve information on the specified DNS Forwarding Proxy service.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param serviceId The On-Prem Application Service identifier. For internal Use only
		@return InfraServicesAPIReadDfpServiceRequest
	*/
	ReadDfpService(ctx context.Context, serviceId string) InfraServicesAPIReadDfpServiceRequest

	// ReadDfpServiceExecute executes the request
	//  @return DfpReadResponse
	ReadDfpServiceExecute(r InfraServicesAPIReadDfpServiceRequest) (*DfpReadResponse, *http.Response, error)
}

// InfraServicesAPIService InfraServicesAPI service
type InfraServicesAPIService internal.Service

type InfraServicesAPICreateOrUpdateDfpServiceRequest struct {
	ctx              context.Context
	ApiService       InfraServicesAPI
	payloadServiceId string
	body             *DfpCreateOrUpdatePayload
}

// The DNS Forwarding Proxy object.
func (r InfraServicesAPICreateOrUpdateDfpServiceRequest) Body(body DfpCreateOrUpdatePayload) InfraServicesAPICreateOrUpdateDfpServiceRequest {
	r.body = &body
	return r
}

func (r InfraServicesAPICreateOrUpdateDfpServiceRequest) Execute() (*DfpCreateOrUpdateResponse, *http.Response, error) {
	return r.ApiService.CreateOrUpdateDfpServiceExecute(r)
}

/*
CreateOrUpdateDfpService Update DNS Forwarding Proxy services.

Use this method to update resolvers for the specified DNS Forwarding Proxy Service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param payloadServiceId The DNS Forwarding Proxy Service ID object identifier.
	@return InfraServicesAPICreateOrUpdateDfpServiceRequest
*/
func (a *InfraServicesAPIService) CreateOrUpdateDfpService(ctx context.Context, payloadServiceId string) InfraServicesAPICreateOrUpdateDfpServiceRequest {
	return InfraServicesAPICreateOrUpdateDfpServiceRequest{
		ApiService:       a,
		ctx:              ctx,
		payloadServiceId: payloadServiceId,
	}
}

// Execute executes the request
//
//	@return DfpCreateOrUpdateResponse
func (a *InfraServicesAPIService) CreateOrUpdateDfpServiceExecute(r InfraServicesAPICreateOrUpdateDfpServiceRequest) (*DfpCreateOrUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *DfpCreateOrUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "InfraServicesAPIService.CreateOrUpdateDfpService")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dfp_services/{payload.service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payload.service_id"+"}", url.PathEscape(internal.ParameterValueToString(r.payloadServiceId, "payloadServiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 400 {
			var v InfraServicesCreateOrUpdateDfpService400Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InfraServicesCreateOrUpdateDfpService404Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InfraServicesListDfpServices500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type InfraServicesAPIListDfpServicesRequest struct {
	ctx        context.Context
	ApiService InfraServicesAPI
	filter     *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type     | Supported Op                | | ----------------------- | -------- | --------------------------- | | service_name            | string   | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | internal_domain_lists   | [int32]  | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | policy_id               | int32    | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | default_security_policy | bool     | !&#x3D;, &#x3D;&#x3D;                      |   In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((service_name&#x3D;&#x3D;&#39;dfp1&#39;)or(policy_id~&#39;oph&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;
func (r InfraServicesAPIListDfpServicesRequest) Filter(filter string) InfraServicesAPIListDfpServicesRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r InfraServicesAPIListDfpServicesRequest) Fields(fields string) InfraServicesAPIListDfpServicesRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r InfraServicesAPIListDfpServicesRequest) Offset(offset int32) InfraServicesAPIListDfpServicesRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r InfraServicesAPIListDfpServicesRequest) Limit(limit int32) InfraServicesAPIListDfpServicesRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r InfraServicesAPIListDfpServicesRequest) PageToken(pageToken string) InfraServicesAPIListDfpServicesRequest {
	r.pageToken = &pageToken
	return r
}

func (r InfraServicesAPIListDfpServicesRequest) Execute() (*DfpListResponse, *http.Response, error) {
	return r.ApiService.ListDfpServicesExecute(r)
}

/*
ListDfpServices List DNS Forwarding Proxy services.

Use this method to retrieve information on all DNS Forwarding Proxy services.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InfraServicesAPIListDfpServicesRequest
*/
func (a *InfraServicesAPIService) ListDfpServices(ctx context.Context) InfraServicesAPIListDfpServicesRequest {
	return InfraServicesAPIListDfpServicesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DfpListResponse
func (a *InfraServicesAPIService) ListDfpServicesExecute(r InfraServicesAPIListDfpServicesRequest) (*DfpListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *DfpListResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "InfraServicesAPIService.ListDfpServices")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dfp_services"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 500 {
			var v InfraServicesListDfpServices500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type InfraServicesAPIReadDfpServiceRequest struct {
	ctx        context.Context
	ApiService InfraServicesAPI
	serviceId  string
	id         *int32
	fields     *string
	name       *string
}

// The DNS Forwarding Proxy object identifier.
func (r InfraServicesAPIReadDfpServiceRequest) Id(id int32) InfraServicesAPIReadDfpServiceRequest {
	r.id = &id
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r InfraServicesAPIReadDfpServiceRequest) Fields(fields string) InfraServicesAPIReadDfpServiceRequest {
	r.fields = &fields
	return r
}

// The name of the DNS Forwarding Proxy. Used only if the &#39;id&#39; field is empty.
func (r InfraServicesAPIReadDfpServiceRequest) Name(name string) InfraServicesAPIReadDfpServiceRequest {
	r.name = &name
	return r
}

func (r InfraServicesAPIReadDfpServiceRequest) Execute() (*DfpReadResponse, *http.Response, error) {
	return r.ApiService.ReadDfpServiceExecute(r)
}

/*
ReadDfpService Read DNS Forwarding Proxy services.

Use this method to retrieve information on the specified DNS Forwarding Proxy service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId The On-Prem Application Service identifier. For internal Use only
	@return InfraServicesAPIReadDfpServiceRequest
*/
func (a *InfraServicesAPIService) ReadDfpService(ctx context.Context, serviceId string) InfraServicesAPIReadDfpServiceRequest {
	return InfraServicesAPIReadDfpServiceRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return DfpReadResponse
func (a *InfraServicesAPIService) ReadDfpServiceExecute(r InfraServicesAPIReadDfpServiceRequest) (*DfpReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *DfpReadResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "InfraServicesAPIService.ReadDfpService")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dfp_services/{service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(internal.ParameterValueToString(r.serviceId, "serviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.name != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 404 {
			var v InfraServicesCreateOrUpdateDfpService404Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InfraServicesListDfpServices500Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}