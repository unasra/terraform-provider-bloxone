/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type LeasesCommandAPI interface {

	/*
		LeasesCommandCreate Perform actions like clearing DHCP lease(s).

		Use this method to create a __LeasesCommand__ object.
	The __LeasesCommand__ object (_dhcp/leases_command_) is used for performing an action like clearing DHCP lease(s).

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiLeasesCommandCreateRequest
	*/
	LeasesCommandCreate(ctx context.Context) ApiLeasesCommandCreateRequest

	// LeasesCommandCreateExecute executes the request
	//  @return IpamsvcCreateLeasesCommandResponse
	LeasesCommandCreateExecute(r ApiLeasesCommandCreateRequest) (*IpamsvcCreateLeasesCommandResponse, *http.Response, error)
}

// LeasesCommandAPIService LeasesCommandAPI service
type LeasesCommandAPIService internal.Service

type ApiLeasesCommandCreateRequest struct {
	ctx        context.Context
	ApiService LeasesCommandAPI
	body       *IpamsvcLeasesCommand
}

func (r ApiLeasesCommandCreateRequest) Body(body IpamsvcLeasesCommand) ApiLeasesCommandCreateRequest {
	r.body = &body
	return r
}

func (r ApiLeasesCommandCreateRequest) Execute() (*IpamsvcCreateLeasesCommandResponse, *http.Response, error) {
	return r.ApiService.LeasesCommandCreateExecute(r)
}

/*
LeasesCommandCreate Perform actions like clearing DHCP lease(s).

Use this method to create a __LeasesCommand__ object.
The __LeasesCommand__ object (_dhcp/leases_command_) is used for performing an action like clearing DHCP lease(s).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiLeasesCommandCreateRequest
*/
func (a *LeasesCommandAPIService) LeasesCommandCreate(ctx context.Context) ApiLeasesCommandCreateRequest {
	return ApiLeasesCommandCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcCreateLeasesCommandResponse
func (a *LeasesCommandAPIService) LeasesCommandCreateExecute(r ApiLeasesCommandCreateRequest) (*IpamsvcCreateLeasesCommandResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcCreateLeasesCommandResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "LeasesCommandAPIService.LeasesCommandCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/leases_command"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
