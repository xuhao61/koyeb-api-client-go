/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type LogsApi interface {

	/*
	 * TailLogs Tails logs
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiTailLogsRequest
	 */
	TailLogs(ctx _context.Context) ApiTailLogsRequest

	/*
	 * TailLogsExecute executes the request
	 * @return StreamResultOfLogEntry
	 */
	TailLogsExecute(r ApiTailLogsRequest) (StreamResultOfLogEntry, *_nethttp.Response, error)
}

// LogsApiService LogsApi service
type LogsApiService service

type ApiTailLogsRequest struct {
	ctx _context.Context
	ApiService LogsApi
	type_ *string
	appId *string
	serviceId *string
	deploymentId *string
	regionalDeploymentId *string
	instanceId *string
	stream *string
	start *string
	limit *string
}

func (r ApiTailLogsRequest) Type_(type_ string) ApiTailLogsRequest {
	r.type_ = &type_
	return r
}
func (r ApiTailLogsRequest) AppId(appId string) ApiTailLogsRequest {
	r.appId = &appId
	return r
}
func (r ApiTailLogsRequest) ServiceId(serviceId string) ApiTailLogsRequest {
	r.serviceId = &serviceId
	return r
}
func (r ApiTailLogsRequest) DeploymentId(deploymentId string) ApiTailLogsRequest {
	r.deploymentId = &deploymentId
	return r
}
func (r ApiTailLogsRequest) RegionalDeploymentId(regionalDeploymentId string) ApiTailLogsRequest {
	r.regionalDeploymentId = &regionalDeploymentId
	return r
}
func (r ApiTailLogsRequest) InstanceId(instanceId string) ApiTailLogsRequest {
	r.instanceId = &instanceId
	return r
}
func (r ApiTailLogsRequest) Stream(stream string) ApiTailLogsRequest {
	r.stream = &stream
	return r
}
func (r ApiTailLogsRequest) Start(start string) ApiTailLogsRequest {
	r.start = &start
	return r
}
func (r ApiTailLogsRequest) Limit(limit string) ApiTailLogsRequest {
	r.limit = &limit
	return r
}

func (r ApiTailLogsRequest) Execute() (StreamResultOfLogEntry, *_nethttp.Response, error) {
	return r.ApiService.TailLogsExecute(r)
}

/*
 * TailLogs Tails logs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTailLogsRequest
 */
func (a *LogsApiService) TailLogs(ctx _context.Context) ApiTailLogsRequest {
	return ApiTailLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return StreamResultOfLogEntry
 */
func (a *LogsApiService) TailLogsExecute(r ApiTailLogsRequest) (StreamResultOfLogEntry, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StreamResultOfLogEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsApiService.TailLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/streams/logs/tail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.appId != nil {
		localVarQueryParams.Add("app_id", parameterToString(*r.appId, ""))
	}
	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.deploymentId != nil {
		localVarQueryParams.Add("deployment_id", parameterToString(*r.deploymentId, ""))
	}
	if r.regionalDeploymentId != nil {
		localVarQueryParams.Add("regional_deployment_id", parameterToString(*r.regionalDeploymentId, ""))
	}
	if r.instanceId != nil {
		localVarQueryParams.Add("instance_id", parameterToString(*r.instanceId, ""))
	}
	if r.stream != nil {
		localVarQueryParams.Add("stream", parameterToString(*r.stream, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorWithFields
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v GoogleRpcStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
