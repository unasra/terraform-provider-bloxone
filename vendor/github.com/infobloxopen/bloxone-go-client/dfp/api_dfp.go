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

type DfpAPI interface {
	/*
			DfpCreateOrUpdateDfp Update DNS Forwarding Proxy resolvers.

			Use this method to update resolvers for the specified DNS Forwarding Proxy.

		For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries. There is a possibility to set default resolver(s) depending on the license class. DNS forwarding proxy would fallback to the default resolvers when BloxOne Cloud is unreachable.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id The DNS Forwarding Proxy object identifier.
			@return ApiDfpCreateOrUpdateDfpRequest
	*/
	DfpCreateOrUpdateDfp(ctx context.Context, id int32) ApiDfpCreateOrUpdateDfpRequest

	// DfpCreateOrUpdateDfpExecute executes the request
	//  @return AtcdfpDfpCreateOrUpdateResponse
	DfpCreateOrUpdateDfpExecute(r ApiDfpCreateOrUpdateDfpRequest) (*AtcdfpDfpCreateOrUpdateResponse, *http.Response, error)
	/*
			DfpListDfp List DNS Forwarding Proxies.

			Use this method to retrieve information on all DNS Forwarding Proxy objects for the account.

		For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.



			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiDfpListDfpRequest
	*/
	DfpListDfp(ctx context.Context) ApiDfpListDfpRequest

	// DfpListDfpExecute executes the request
	//  @return AtcdfpDfpListResponse
	DfpListDfpExecute(r ApiDfpListDfpRequest) (*AtcdfpDfpListResponse, *http.Response, error)
	/*
			DfpReadDfp Read DNS Forwarding Proxy.

			Use this method to retrieve information on the specified DNS Forwarding Proxy object.

		For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.

		Note that DNS Forwarding Proxy cannot be created (all information regarding DFP is synchronized from hostapp service).


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id The DNS Forwarding Proxy object identifier.
			@return ApiDfpReadDfpRequest
	*/
	DfpReadDfp(ctx context.Context, id int32) ApiDfpReadDfpRequest

	// DfpReadDfpExecute executes the request
	//  @return AtcdfpDfpReadResponse
	DfpReadDfpExecute(r ApiDfpReadDfpRequest) (*AtcdfpDfpReadResponse, *http.Response, error)
}

// DfpAPIService DfpAPI service
type DfpAPIService internal.Service

type ApiDfpCreateOrUpdateDfpRequest struct {
	ctx        context.Context
	ApiService DfpAPI
	id         int32
	body       *AtcdfpDfpCreateOrUpdatePayload
}

// The DNS Forwarding Proxy object.
func (r ApiDfpCreateOrUpdateDfpRequest) Body(body AtcdfpDfpCreateOrUpdatePayload) ApiDfpCreateOrUpdateDfpRequest {
	r.body = &body
	return r
}

func (r ApiDfpCreateOrUpdateDfpRequest) Execute() (*AtcdfpDfpCreateOrUpdateResponse, *http.Response, error) {
	return r.ApiService.DfpCreateOrUpdateDfpExecute(r)
}

/*
DfpCreateOrUpdateDfp Update DNS Forwarding Proxy resolvers.

Use this method to update resolvers for the specified DNS Forwarding Proxy.

For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries. There is a possibility to set default resolver(s) depending on the license class. DNS forwarding proxy would fallback to the default resolvers when BloxOne Cloud is unreachable.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The DNS Forwarding Proxy object identifier.
	@return ApiDfpCreateOrUpdateDfpRequest
*/
func (a *DfpAPIService) DfpCreateOrUpdateDfp(ctx context.Context, id int32) ApiDfpCreateOrUpdateDfpRequest {
	return ApiDfpCreateOrUpdateDfpRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AtcdfpDfpCreateOrUpdateResponse
func (a *DfpAPIService) DfpCreateOrUpdateDfpExecute(r ApiDfpCreateOrUpdateDfpRequest) (*AtcdfpDfpCreateOrUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *AtcdfpDfpCreateOrUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DfpAPIService.DfpCreateOrUpdateDfp")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dfps/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

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
			var v DfpCreateOrUpdateDfp400Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v DfpReadDfp404Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v DfpListDfp500Response
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

type ApiDfpListDfpRequest struct {
	ctx        context.Context
	ApiService DfpAPI
	filter     *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;.  You can filter by following fields:  | Name                    | type   | Supported Op                | | ----------------------- | ------ | --------------------------- | | name                    | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | site_id                 | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | ophid                   | string | !&#x3D;, &#x3D;&#x3D;, ~, !~, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D; | | policy_id               | int32  | !&#x3D;, &#x3D;&#x3D;, &gt;, &lt;, &lt;&#x3D;, &gt;&#x3D;        | | default_security_policy | bool   | !&#x3D;, &#x3D;&#x3D;                      |  In addition groupping operators are supported:  | Op  | Description          | | --- | -------------------- | | and | Logical AND          | | or  | Logical OR           | | not | Logical NOT          | | ()  | Groupping Operators  |  Example: &#x60;&#x60;&#x60; ?_filter&#x3D;\&quot;((name&#x3D;&#x3D;&#39;dfp1&#39;)or(ophid~&#39;oph&#39;))and(default_security_policy!&#x3D;&#39;true&#39;)\&quot; &#x60;&#x60;&#x60;
func (r ApiDfpListDfpRequest) Filter(filter string) ApiDfpListDfpRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiDfpListDfpRequest) Fields(fields string) ApiDfpListDfpRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiDfpListDfpRequest) Offset(offset int32) ApiDfpListDfpRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiDfpListDfpRequest) Limit(limit int32) ApiDfpListDfpRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiDfpListDfpRequest) PageToken(pageToken string) ApiDfpListDfpRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiDfpListDfpRequest) Execute() (*AtcdfpDfpListResponse, *http.Response, error) {
	return r.ApiService.DfpListDfpExecute(r)
}

/*
DfpListDfp List DNS Forwarding Proxies.

Use this method to retrieve information on all DNS Forwarding Proxy objects for the account.

For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDfpListDfpRequest
*/
func (a *DfpAPIService) DfpListDfp(ctx context.Context) ApiDfpListDfpRequest {
	return ApiDfpListDfpRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AtcdfpDfpListResponse
func (a *DfpAPIService) DfpListDfpExecute(r ApiDfpListDfpRequest) (*AtcdfpDfpListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *AtcdfpDfpListResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DfpAPIService.DfpListDfp")
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
			var v DfpListDfp500Response
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

type ApiDfpReadDfpRequest struct {
	ctx        context.Context
	ApiService DfpAPI
	id         int32
	fields     *string
	name       *string
	serviceId  *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiDfpReadDfpRequest) Fields(fields string) ApiDfpReadDfpRequest {
	r.fields = &fields
	return r
}

// The name of the DNS Forwarding Proxy. Used only if the &#39;id&#39; field is empty.
func (r ApiDfpReadDfpRequest) Name(name string) ApiDfpReadDfpRequest {
	r.name = &name
	return r
}

// The On-Prem Application Service identifier. For internal Use only.
func (r ApiDfpReadDfpRequest) ServiceId(serviceId string) ApiDfpReadDfpRequest {
	r.serviceId = &serviceId
	return r
}

func (r ApiDfpReadDfpRequest) Execute() (*AtcdfpDfpReadResponse, *http.Response, error) {
	return r.ApiService.DfpReadDfpExecute(r)
}

/*
DfpReadDfp Read DNS Forwarding Proxy.

Use this method to retrieve information on the specified DNS Forwarding Proxy object.

For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.

Note that DNS Forwarding Proxy cannot be created (all information regarding DFP is synchronized from hostapp service).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The DNS Forwarding Proxy object identifier.
	@return ApiDfpReadDfpRequest
*/
func (a *DfpAPIService) DfpReadDfp(ctx context.Context, id int32) ApiDfpReadDfpRequest {
	return ApiDfpReadDfpRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AtcdfpDfpReadResponse
func (a *DfpAPIService) DfpReadDfpExecute(r ApiDfpReadDfpRequest) (*AtcdfpDfpReadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *AtcdfpDfpReadResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "DfpAPIService.DfpReadDfp")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dfps/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.name != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.serviceId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "service_id", r.serviceId, "")
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
			var v DfpReadDfp404Response
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr = internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr = internal.NewGenericOpenAPIErrorWithBodyAndModel(localVarHTTPResponse.Status, localVarBody, v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v DfpListDfp500Response
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
