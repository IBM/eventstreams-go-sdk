/**
 * (C) Copyright IBM Corp. 2024.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.93.0-c40121e6-20240729-182103
 */

// Package schemaregistryv1 : Operations and models for the SchemaregistryV1 service
package schemaregistryv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/eventstreams-go-sdk/pkg/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// SchemaregistryV1 : IBM Event Streams schema registry management
//
// API Version: 1.4.1
type SchemaregistryV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "schemaregistry"

// SchemaregistryV1Options : Service options
type SchemaregistryV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSchemaregistryV1UsingExternalConfig : constructs an instance of SchemaregistryV1 with passed in options and external configuration.
func NewSchemaregistryV1UsingExternalConfig(options *SchemaregistryV1Options) (schemaregistry *SchemaregistryV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	schemaregistry, err = NewSchemaregistryV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = schemaregistry.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = schemaregistry.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewSchemaregistryV1 : constructs an instance of SchemaregistryV1 with passed in options.
func NewSchemaregistryV1(options *SchemaregistryV1Options) (service *SchemaregistryV1, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &SchemaregistryV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "schemaregistry" suitable for processing requests.
func (schemaregistry *SchemaregistryV1) Clone() *SchemaregistryV1 {
	if core.IsNil(schemaregistry) {
		return nil
	}
	clone := *schemaregistry
	clone.Service = schemaregistry.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (schemaregistry *SchemaregistryV1) SetServiceURL(url string) error {
	err := schemaregistry.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (schemaregistry *SchemaregistryV1) GetServiceURL() string {
	return schemaregistry.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (schemaregistry *SchemaregistryV1) SetDefaultHeaders(headers http.Header) {
	schemaregistry.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (schemaregistry *SchemaregistryV1) SetEnableGzipCompression(enableGzip bool) {
	schemaregistry.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (schemaregistry *SchemaregistryV1) GetEnableGzipCompression() bool {
	return schemaregistry.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (schemaregistry *SchemaregistryV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	schemaregistry.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (schemaregistry *SchemaregistryV1) DisableRetries() {
	schemaregistry.Service.DisableRetries()
}

// GetGlobalRule : Retrieve the configuration for a global rule
// Retrieves the configuration for the specified global rule. The value of the global rule is used as the _default_ when
// a schema does not have a corresponding schema compatibility rule defined.
func (schemaregistry *SchemaregistryV1) GetGlobalRule(getGlobalRuleOptions *GetGlobalRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.GetGlobalRuleWithContext(context.Background(), getGlobalRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetGlobalRuleWithContext is an alternate form of the GetGlobalRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) GetGlobalRuleWithContext(ctx context.Context, getGlobalRuleOptions *GetGlobalRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getGlobalRuleOptions, "getGlobalRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getGlobalRuleOptions, "getGlobalRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"rule": *getGlobalRuleOptions.Rule,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/rules/{rule}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getGlobalRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "GetGlobalRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getGlobalRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateGlobalRule : Update the configuration for a global rule
// Update the configuration for the specified global rule. The value of the global rule is used as the _default_ when a
// schema does not have a corresponding schema compatibility rule defined.
func (schemaregistry *SchemaregistryV1) UpdateGlobalRule(updateGlobalRuleOptions *UpdateGlobalRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.UpdateGlobalRuleWithContext(context.Background(), updateGlobalRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateGlobalRuleWithContext is an alternate form of the UpdateGlobalRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) UpdateGlobalRuleWithContext(ctx context.Context, updateGlobalRuleOptions *UpdateGlobalRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateGlobalRuleOptions, "updateGlobalRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateGlobalRuleOptions, "updateGlobalRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"rule": *updateGlobalRuleOptions.Rule,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/rules/{rule}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateGlobalRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "UpdateGlobalRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateGlobalRuleOptions.Type != nil {
		body["type"] = updateGlobalRuleOptions.Type
	}
	if updateGlobalRuleOptions.Config != nil {
		body["config"] = updateGlobalRuleOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "updateGlobalRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateSchemaRule : Create a schema rule
// Create a new rule that controls compatibility checks for a particular schema. Schema rules override the registries
// global compatibility rule setting.
func (schemaregistry *SchemaregistryV1) CreateSchemaRule(createSchemaRuleOptions *CreateSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.CreateSchemaRuleWithContext(context.Background(), createSchemaRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSchemaRuleWithContext is an alternate form of the CreateSchemaRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) CreateSchemaRuleWithContext(ctx context.Context, createSchemaRuleOptions *CreateSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaRuleOptions, "createSchemaRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSchemaRuleOptions, "createSchemaRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *createSchemaRuleOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/rules`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSchemaRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "CreateSchemaRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createSchemaRuleOptions.Type != nil {
		body["type"] = createSchemaRuleOptions.Type
	}
	if createSchemaRuleOptions.Config != nil {
		body["config"] = createSchemaRuleOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "createSchemaRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetSchemaRule : Get a schema rule configuration
// Retrieves the current configuration for a schema rule. If a schema rule exists then it overrides the corresponding
// global rule that would otherwise be applied.
func (schemaregistry *SchemaregistryV1) GetSchemaRule(getSchemaRuleOptions *GetSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.GetSchemaRuleWithContext(context.Background(), getSchemaRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetSchemaRuleWithContext is an alternate form of the GetSchemaRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) GetSchemaRuleWithContext(ctx context.Context, getSchemaRuleOptions *GetSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSchemaRuleOptions, "getSchemaRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getSchemaRuleOptions, "getSchemaRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":   *getSchemaRuleOptions.ID,
		"rule": *getSchemaRuleOptions.Rule,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/rules/{rule}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getSchemaRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "GetSchemaRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getSchemaRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// UpdateSchemaRule : Update the configuration of a schema rule
// Updates the configuration of an existing schema rule. The updated rule will be applied to the specified schema,
// overriding the value set for the corresponding global rule.
func (schemaregistry *SchemaregistryV1) UpdateSchemaRule(updateSchemaRuleOptions *UpdateSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.UpdateSchemaRuleWithContext(context.Background(), updateSchemaRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSchemaRuleWithContext is an alternate form of the UpdateSchemaRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) UpdateSchemaRuleWithContext(ctx context.Context, updateSchemaRuleOptions *UpdateSchemaRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSchemaRuleOptions, "updateSchemaRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateSchemaRuleOptions, "updateSchemaRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":   *updateSchemaRuleOptions.ID,
		"rule": *updateSchemaRuleOptions.Rule,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/rules/{rule}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateSchemaRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "UpdateSchemaRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSchemaRuleOptions.Type != nil {
		body["type"] = updateSchemaRuleOptions.Type
	}
	if updateSchemaRuleOptions.Config != nil {
		body["config"] = updateSchemaRuleOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "updateSchemaRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSchemaRule : Delete a schema rule
// Delete a rule that controls compatibility checks for a particular schema. After this operation completes the schema
// will be subject to compatibility checking defined by the global compatibility rule setting for the registry.
func (schemaregistry *SchemaregistryV1) DeleteSchemaRule(deleteSchemaRuleOptions *DeleteSchemaRuleOptions) (response *core.DetailedResponse, err error) {
	response, err = schemaregistry.DeleteSchemaRuleWithContext(context.Background(), deleteSchemaRuleOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSchemaRuleWithContext is an alternate form of the DeleteSchemaRule method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) DeleteSchemaRuleWithContext(ctx context.Context, deleteSchemaRuleOptions *DeleteSchemaRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSchemaRuleOptions, "deleteSchemaRuleOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSchemaRuleOptions, "deleteSchemaRuleOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":   *deleteSchemaRuleOptions.ID,
		"rule": *deleteSchemaRuleOptions.Rule,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/rules/{rule}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSchemaRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "DeleteSchemaRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "deleteSchemaRule", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// SetSchemaState : Set schema state
// Sets schema state.
func (schemaregistry *SchemaregistryV1) SetSchemaState(setSchemaStateOptions *SetSchemaStateOptions) (response *core.DetailedResponse, err error) {
	response, err = schemaregistry.SetSchemaStateWithContext(context.Background(), setSchemaStateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// SetSchemaStateWithContext is an alternate form of the SetSchemaState method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) SetSchemaStateWithContext(ctx context.Context, setSchemaStateOptions *SetSchemaStateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setSchemaStateOptions, "setSchemaStateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(setSchemaStateOptions, "setSchemaStateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *setSchemaStateOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/state`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range setSchemaStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "SetSchemaState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setSchemaStateOptions.State != nil {
		body["state"] = setSchemaStateOptions.State
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "setSchemaState", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// SetSchemaVersionState : Set schema version state
// Sets schema version state.
func (schemaregistry *SchemaregistryV1) SetSchemaVersionState(setSchemaVersionStateOptions *SetSchemaVersionStateOptions) (response *core.DetailedResponse, err error) {
	response, err = schemaregistry.SetSchemaVersionStateWithContext(context.Background(), setSchemaVersionStateOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// SetSchemaVersionStateWithContext is an alternate form of the SetSchemaVersionState method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) SetSchemaVersionStateWithContext(ctx context.Context, setSchemaVersionStateOptions *SetSchemaVersionStateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setSchemaVersionStateOptions, "setSchemaVersionStateOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(setSchemaVersionStateOptions, "setSchemaVersionStateOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":      *setSchemaVersionStateOptions.ID,
		"version": fmt.Sprint(*setSchemaVersionStateOptions.Version),
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/versions/{version}/state`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range setSchemaVersionStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "SetSchemaVersionState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setSchemaVersionStateOptions.State != nil {
		body["state"] = setSchemaVersionStateOptions.State
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "setSchemaVersionState", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListVersions : List the versions of a schema
// Returns an array containing the version numbers of all of the versions of the specified schema.
func (schemaregistry *SchemaregistryV1) ListVersions(listVersionsOptions *ListVersionsOptions) (result []int64, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.ListVersionsWithContext(context.Background(), listVersionsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListVersionsWithContext is an alternate form of the ListVersions method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) ListVersionsWithContext(ctx context.Context, listVersionsOptions *ListVersionsOptions) (result []int64, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listVersionsOptions, "listVersionsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listVersionsOptions, "listVersionsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *listVersionsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/versions`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "ListVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listVersionsOptions.Jsonformat != nil {
		builder.AddQuery("jsonformat", fmt.Sprint(*listVersionsOptions.Jsonformat))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "listVersions", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateVersion : Create a new schema version
// Creates a new version of a schema using the AVRO schema supplied in the request body.
func (schemaregistry *SchemaregistryV1) CreateVersion(createVersionOptions *CreateVersionOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.CreateVersionWithContext(context.Background(), createVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateVersionWithContext is an alternate form of the CreateVersion method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) CreateVersionWithContext(ctx context.Context, createVersionOptions *CreateVersionOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVersionOptions, "createVersionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createVersionOptions, "createVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *createVersionOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/versions`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "CreateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createVersionOptions.Schema != nil {
		body["schema"] = createVersionOptions.Schema
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "createVersion", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSchemaMetadata)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetVersion : Get a version of the schema
// Retrieve a particular version of the schema.
func (schemaregistry *SchemaregistryV1) GetVersion(getVersionOptions *GetVersionOptions) (result *AvroSchema, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.GetVersionWithContext(context.Background(), getVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetVersionWithContext is an alternate form of the GetVersion method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) GetVersionWithContext(ctx context.Context, getVersionOptions *GetVersionOptions) (result *AvroSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionOptions, "getVersionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getVersionOptions, "getVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":      *getVersionOptions.ID,
		"version": fmt.Sprint(*getVersionOptions.Version),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/versions/{version}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "GetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getVersion", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAvroSchema)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteVersion : Delete a version of the schema
// Delete a version of the schema. If this was the only version of the schema then the whole schema will be deleted.
func (schemaregistry *SchemaregistryV1) DeleteVersion(deleteVersionOptions *DeleteVersionOptions) (response *core.DetailedResponse, err error) {
	response, err = schemaregistry.DeleteVersionWithContext(context.Background(), deleteVersionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteVersionWithContext is an alternate form of the DeleteVersion method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) DeleteVersionWithContext(ctx context.Context, deleteVersionOptions *DeleteVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVersionOptions, "deleteVersionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteVersionOptions, "deleteVersionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id":      *deleteVersionOptions.ID,
		"version": fmt.Sprint(*deleteVersionOptions.Version),
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}/versions/{version}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "DeleteVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "deleteVersion", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListSchemas : List schema IDs
// Returns an array containing the schema IDs of all of the schemas that are stored in the registry.
func (schemaregistry *SchemaregistryV1) ListSchemas(listSchemasOptions *ListSchemasOptions) (result []string, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.ListSchemasWithContext(context.Background(), listSchemasOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListSchemasWithContext is an alternate form of the ListSchemas method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) ListSchemasWithContext(ctx context.Context, listSchemasOptions *ListSchemasOptions) (result []string, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSchemasOptions, "listSchemasOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listSchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "ListSchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listSchemasOptions.Jsonformat != nil {
		builder.AddQuery("jsonformat", fmt.Sprint(*listSchemasOptions.Jsonformat))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "listSchemas", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateSchema : Create a new schema
// Create a new schema and populate it with an initial schema version containing the AVRO document in the request body.
func (schemaregistry *SchemaregistryV1) CreateSchema(createSchemaOptions *CreateSchemaOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.CreateSchemaWithContext(context.Background(), createSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateSchemaWithContext is an alternate form of the CreateSchema method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) CreateSchemaWithContext(ctx context.Context, createSchemaOptions *CreateSchemaOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaOptions, "createSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createSchemaOptions, "createSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "CreateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createSchemaOptions.XRegistryArtifactID != nil {
		builder.AddHeader("X-Registry-ArtifactId", fmt.Sprint(*createSchemaOptions.XRegistryArtifactID))
	}

	body := make(map[string]interface{})
	if createSchemaOptions.Schema != nil {
		body["schema"] = createSchemaOptions.Schema
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "createSchema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSchemaMetadata)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetLatestSchema : Get the latest version of a schema
// Retrieves the lastest version of the specified schema.
func (schemaregistry *SchemaregistryV1) GetLatestSchema(getLatestSchemaOptions *GetLatestSchemaOptions) (result *AvroSchema, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.GetLatestSchemaWithContext(context.Background(), getLatestSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetLatestSchemaWithContext is an alternate form of the GetLatestSchema method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) GetLatestSchemaWithContext(ctx context.Context, getLatestSchemaOptions *GetLatestSchemaOptions) (result *AvroSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLatestSchemaOptions, "getLatestSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getLatestSchemaOptions, "getLatestSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *getLatestSchemaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getLatestSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "GetLatestSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getLatestSchema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAvroSchema)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteSchema : Delete a schema
// Deletes a schema and all of its versions from the schema registry.
func (schemaregistry *SchemaregistryV1) DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	response, err = schemaregistry.DeleteSchemaWithContext(context.Background(), deleteSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteSchemaWithContext is an alternate form of the DeleteSchema method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) DeleteSchemaWithContext(ctx context.Context, deleteSchemaOptions *DeleteSchemaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSchemaOptions, "deleteSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteSchemaOptions, "deleteSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteSchemaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "DeleteSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = schemaregistry.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "deleteSchema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateSchema : Update a schema
// Updates a schema.
func (schemaregistry *SchemaregistryV1) UpdateSchema(updateSchemaOptions *UpdateSchemaOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	result, response, err = schemaregistry.UpdateSchemaWithContext(context.Background(), updateSchemaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateSchemaWithContext is an alternate form of the UpdateSchema method which supports a Context parameter
func (schemaregistry *SchemaregistryV1) UpdateSchemaWithContext(ctx context.Context, updateSchemaOptions *UpdateSchemaOptions) (result *SchemaMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSchemaOptions, "updateSchemaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateSchemaOptions, "updateSchemaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateSchemaOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = schemaregistry.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(schemaregistry.Service.Options.URL, `/artifacts/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("schemaregistry", "V1", "UpdateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSchemaOptions.Schema != nil {
		body["schema"] = updateSchemaOptions.Schema
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = schemaregistry.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "updateSchema", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSchemaMetadata)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.4.1")
}

// CreateSchemaOptions : The CreateSchema options.
type CreateSchemaOptions struct {
	// The AVRO schema.
	Schema map[string]interface{} `json:"schema,omitempty"`

	// The name to assign to the new schema. This must be unique. If this value is not specified then a UUID is used.
	XRegistryArtifactID *string `json:"X-Registry-ArtifactId,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateSchemaOptions : Instantiate CreateSchemaOptions
func (*SchemaregistryV1) NewCreateSchemaOptions() *CreateSchemaOptions {
	return &CreateSchemaOptions{}
}

// SetSchema : Allow user to set Schema
func (_options *CreateSchemaOptions) SetSchema(schema map[string]interface{}) *CreateSchemaOptions {
	_options.Schema = schema
	return _options
}

// SetXRegistryArtifactID : Allow user to set XRegistryArtifactID
func (_options *CreateSchemaOptions) SetXRegistryArtifactID(xRegistryArtifactID string) *CreateSchemaOptions {
	_options.XRegistryArtifactID = core.StringPtr(xRegistryArtifactID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSchemaOptions) SetHeaders(param map[string]string) *CreateSchemaOptions {
	options.Headers = param
	return options
}

// CreateSchemaRuleOptions : The CreateSchemaRule options.
type CreateSchemaRuleOptions struct {
	// The ID of the schema that the rule is to be associated with.
	ID *string `json:"id" validate:"required,ne="`

	// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
	Type *string `json:"type" validate:"required"`

	// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
	Config *string `json:"config" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the CreateSchemaRuleOptions.Type property.
// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
const (
	CreateSchemaRuleOptionsTypeCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the CreateSchemaRuleOptions.Config property.
// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
const (
	CreateSchemaRuleOptionsConfigBackwardConst           = "BACKWARD"
	CreateSchemaRuleOptionsConfigBackwardTransitiveConst = "BACKWARD_TRANSITIVE"
	CreateSchemaRuleOptionsConfigForwardConst            = "FORWARD"
	CreateSchemaRuleOptionsConfigForwardTransitiveConst  = "FORWARD_TRANSITIVE"
	CreateSchemaRuleOptionsConfigFullConst               = "FULL"
	CreateSchemaRuleOptionsConfigFullTransitiveConst     = "FULL_TRANSITIVE"
	CreateSchemaRuleOptionsConfigNoneConst               = "NONE"
)

// NewCreateSchemaRuleOptions : Instantiate CreateSchemaRuleOptions
func (*SchemaregistryV1) NewCreateSchemaRuleOptions(id string, typeVar string, config string) *CreateSchemaRuleOptions {
	return &CreateSchemaRuleOptions{
		ID:     core.StringPtr(id),
		Type:   core.StringPtr(typeVar),
		Config: core.StringPtr(config),
	}
}

// SetID : Allow user to set ID
func (_options *CreateSchemaRuleOptions) SetID(id string) *CreateSchemaRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateSchemaRuleOptions) SetType(typeVar string) *CreateSchemaRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetConfig : Allow user to set Config
func (_options *CreateSchemaRuleOptions) SetConfig(config string) *CreateSchemaRuleOptions {
	_options.Config = core.StringPtr(config)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSchemaRuleOptions) SetHeaders(param map[string]string) *CreateSchemaRuleOptions {
	options.Headers = param
	return options
}

// CreateVersionOptions : The CreateVersion options.
type CreateVersionOptions struct {
	// A schema ID. This identifies the schema for which a new version will be created.
	ID *string `json:"id" validate:"required,ne="`

	// The AVRO schema.
	Schema map[string]interface{} `json:"schema,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateVersionOptions : Instantiate CreateVersionOptions
func (*SchemaregistryV1) NewCreateVersionOptions(id string) *CreateVersionOptions {
	return &CreateVersionOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *CreateVersionOptions) SetID(id string) *CreateVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetSchema : Allow user to set Schema
func (_options *CreateVersionOptions) SetSchema(schema map[string]interface{}) *CreateVersionOptions {
	_options.Schema = schema
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVersionOptions) SetHeaders(param map[string]string) *CreateVersionOptions {
	options.Headers = param
	return options
}

// DeleteSchemaOptions : The DeleteSchema options.
type DeleteSchemaOptions struct {
	// The ID of the schema to delete.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteSchemaOptions : Instantiate DeleteSchemaOptions
func (*SchemaregistryV1) NewDeleteSchemaOptions(id string) *DeleteSchemaOptions {
	return &DeleteSchemaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteSchemaOptions) SetID(id string) *DeleteSchemaOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSchemaOptions) SetHeaders(param map[string]string) *DeleteSchemaOptions {
	options.Headers = param
	return options
}

// DeleteSchemaRuleOptions : The DeleteSchemaRule options.
type DeleteSchemaRuleOptions struct {
	// The ID of the schema that the rule is to be deleted from.
	ID *string `json:"id" validate:"required,ne="`

	// The type of rule to delete. Currently only the value that can be specified is `COMPATIBILITY`.
	Rule *string `json:"rule" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the DeleteSchemaRuleOptions.Rule property.
// The type of rule to delete. Currently only the value that can be specified is `COMPATIBILITY`.
const (
	DeleteSchemaRuleOptionsRuleCompatibilityConst = "COMPATIBILITY"
)

// NewDeleteSchemaRuleOptions : Instantiate DeleteSchemaRuleOptions
func (*SchemaregistryV1) NewDeleteSchemaRuleOptions(id string, rule string) *DeleteSchemaRuleOptions {
	return &DeleteSchemaRuleOptions{
		ID:   core.StringPtr(id),
		Rule: core.StringPtr(rule),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteSchemaRuleOptions) SetID(id string) *DeleteSchemaRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetRule : Allow user to set Rule
func (_options *DeleteSchemaRuleOptions) SetRule(rule string) *DeleteSchemaRuleOptions {
	_options.Rule = core.StringPtr(rule)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSchemaRuleOptions) SetHeaders(param map[string]string) *DeleteSchemaRuleOptions {
	options.Headers = param
	return options
}

// DeleteVersionOptions : The DeleteVersion options.
type DeleteVersionOptions struct {
	// A schema ID that identifies the schema to delete a version from.
	ID *string `json:"id" validate:"required,ne="`

	// The schema version number to delete.
	Version *int64 `json:"version" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteVersionOptions : Instantiate DeleteVersionOptions
func (*SchemaregistryV1) NewDeleteVersionOptions(id string, version int64) *DeleteVersionOptions {
	return &DeleteVersionOptions{
		ID:      core.StringPtr(id),
		Version: core.Int64Ptr(version),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteVersionOptions) SetID(id string) *DeleteVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *DeleteVersionOptions) SetVersion(version int64) *DeleteVersionOptions {
	_options.Version = core.Int64Ptr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVersionOptions) SetHeaders(param map[string]string) *DeleteVersionOptions {
	options.Headers = param
	return options
}

// GetGlobalRuleOptions : The GetGlobalRule options.
type GetGlobalRuleOptions struct {
	// The type of the global rule to retrieve. Currently only `COMPATIBILITY` is supported.
	Rule *string `json:"rule" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the GetGlobalRuleOptions.Rule property.
// The type of the global rule to retrieve. Currently only `COMPATIBILITY` is supported.
const (
	GetGlobalRuleOptionsRuleCompatibilityConst = "COMPATIBILITY"
)

// NewGetGlobalRuleOptions : Instantiate GetGlobalRuleOptions
func (*SchemaregistryV1) NewGetGlobalRuleOptions(rule string) *GetGlobalRuleOptions {
	return &GetGlobalRuleOptions{
		Rule: core.StringPtr(rule),
	}
}

// SetRule : Allow user to set Rule
func (_options *GetGlobalRuleOptions) SetRule(rule string) *GetGlobalRuleOptions {
	_options.Rule = core.StringPtr(rule)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetGlobalRuleOptions) SetHeaders(param map[string]string) *GetGlobalRuleOptions {
	options.Headers = param
	return options
}

// GetLatestSchemaOptions : The GetLatestSchema options.
type GetLatestSchemaOptions struct {
	// The ID of a schema.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetLatestSchemaOptions : Instantiate GetLatestSchemaOptions
func (*SchemaregistryV1) NewGetLatestSchemaOptions(id string) *GetLatestSchemaOptions {
	return &GetLatestSchemaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetLatestSchemaOptions) SetID(id string) *GetLatestSchemaOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLatestSchemaOptions) SetHeaders(param map[string]string) *GetLatestSchemaOptions {
	options.Headers = param
	return options
}

// GetSchemaRuleOptions : The GetSchemaRule options.
type GetSchemaRuleOptions struct {
	// The ID of the schema to retrieve the rule for.
	ID *string `json:"id" validate:"required,ne="`

	// The type of rule to retrieve. Currently only the value that can be specified is `COMPATIBILITY`.
	Rule *string `json:"rule" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the GetSchemaRuleOptions.Rule property.
// The type of rule to retrieve. Currently only the value that can be specified is `COMPATIBILITY`.
const (
	GetSchemaRuleOptionsRuleCompatibilityConst = "COMPATIBILITY"
)

// NewGetSchemaRuleOptions : Instantiate GetSchemaRuleOptions
func (*SchemaregistryV1) NewGetSchemaRuleOptions(id string, rule string) *GetSchemaRuleOptions {
	return &GetSchemaRuleOptions{
		ID:   core.StringPtr(id),
		Rule: core.StringPtr(rule),
	}
}

// SetID : Allow user to set ID
func (_options *GetSchemaRuleOptions) SetID(id string) *GetSchemaRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetRule : Allow user to set Rule
func (_options *GetSchemaRuleOptions) SetRule(rule string) *GetSchemaRuleOptions {
	_options.Rule = core.StringPtr(rule)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSchemaRuleOptions) SetHeaders(param map[string]string) *GetSchemaRuleOptions {
	options.Headers = param
	return options
}

// GetVersionOptions : The GetVersion options.
type GetVersionOptions struct {
	// The schema ID identifying which schema to return a version from.
	ID *string `json:"id" validate:"required,ne="`

	// The version number that identifies the particular schema version to return.
	Version *int64 `json:"version" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetVersionOptions : Instantiate GetVersionOptions
func (*SchemaregistryV1) NewGetVersionOptions(id string, version int64) *GetVersionOptions {
	return &GetVersionOptions{
		ID:      core.StringPtr(id),
		Version: core.Int64Ptr(version),
	}
}

// SetID : Allow user to set ID
func (_options *GetVersionOptions) SetID(id string) *GetVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *GetVersionOptions) SetVersion(version int64) *GetVersionOptions {
	_options.Version = core.Int64Ptr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionOptions) SetHeaders(param map[string]string) *GetVersionOptions {
	options.Headers = param
	return options
}

// ListSchemasOptions : The ListSchemas options.
type ListSchemasOptions struct {
	// format of the response to be returned, allowed values are 'string' and 'object'.
	Jsonformat *string `json:"jsonformat,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListSchemasOptions : Instantiate ListSchemasOptions
func (*SchemaregistryV1) NewListSchemasOptions() *ListSchemasOptions {
	return &ListSchemasOptions{}
}

// SetJsonformat : Allow user to set Jsonformat
func (_options *ListSchemasOptions) SetJsonformat(jsonformat string) *ListSchemasOptions {
	_options.Jsonformat = core.StringPtr(jsonformat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListSchemasOptions) SetHeaders(param map[string]string) *ListSchemasOptions {
	options.Headers = param
	return options
}

// ListVersionsOptions : The ListVersions options.
type ListVersionsOptions struct {
	// The schema ID for which the list of versions will be returned.
	ID *string `json:"id" validate:"required,ne="`

	// format of the response to be returned, allowed values are 'number' and 'object'.
	Jsonformat *string `json:"jsonformat,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListVersionsOptions : Instantiate ListVersionsOptions
func (*SchemaregistryV1) NewListVersionsOptions(id string) *ListVersionsOptions {
	return &ListVersionsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListVersionsOptions) SetID(id string) *ListVersionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetJsonformat : Allow user to set Jsonformat
func (_options *ListVersionsOptions) SetJsonformat(jsonformat string) *ListVersionsOptions {
	_options.Jsonformat = core.StringPtr(jsonformat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVersionsOptions) SetHeaders(param map[string]string) *ListVersionsOptions {
	options.Headers = param
	return options
}

// SetSchemaStateOptions : The SetSchemaState options.
type SetSchemaStateOptions struct {
	// The ID of a schema.
	ID *string `json:"id" validate:"required,ne="`

	// The state of the schema or schema version.
	State *string `json:"state" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the SetSchemaStateOptions.State property.
// The state of the schema or schema version.
const (
	SetSchemaStateOptionsStateDisabledConst = "DISABLED"
	SetSchemaStateOptionsStateEnabledConst  = "ENABLED"
)

// NewSetSchemaStateOptions : Instantiate SetSchemaStateOptions
func (*SchemaregistryV1) NewSetSchemaStateOptions(id string, state string) *SetSchemaStateOptions {
	return &SetSchemaStateOptions{
		ID:    core.StringPtr(id),
		State: core.StringPtr(state),
	}
}

// SetID : Allow user to set ID
func (_options *SetSchemaStateOptions) SetID(id string) *SetSchemaStateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetState : Allow user to set State
func (_options *SetSchemaStateOptions) SetState(state string) *SetSchemaStateOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetSchemaStateOptions) SetHeaders(param map[string]string) *SetSchemaStateOptions {
	options.Headers = param
	return options
}

// SetSchemaVersionStateOptions : The SetSchemaVersionState options.
type SetSchemaVersionStateOptions struct {
	// The ID of a schema.
	ID *string `json:"id" validate:"required,ne="`

	// The version number that identifies the particular schema version to return.
	Version *int64 `json:"version" validate:"required"`

	// The state of the schema or schema version.
	State *string `json:"state" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the SetSchemaVersionStateOptions.State property.
// The state of the schema or schema version.
const (
	SetSchemaVersionStateOptionsStateDisabledConst = "DISABLED"
	SetSchemaVersionStateOptionsStateEnabledConst  = "ENABLED"
)

// NewSetSchemaVersionStateOptions : Instantiate SetSchemaVersionStateOptions
func (*SchemaregistryV1) NewSetSchemaVersionStateOptions(id string, version int64, state string) *SetSchemaVersionStateOptions {
	return &SetSchemaVersionStateOptions{
		ID:      core.StringPtr(id),
		Version: core.Int64Ptr(version),
		State:   core.StringPtr(state),
	}
}

// SetID : Allow user to set ID
func (_options *SetSchemaVersionStateOptions) SetID(id string) *SetSchemaVersionStateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *SetSchemaVersionStateOptions) SetVersion(version int64) *SetSchemaVersionStateOptions {
	_options.Version = core.Int64Ptr(version)
	return _options
}

// SetState : Allow user to set State
func (_options *SetSchemaVersionStateOptions) SetState(state string) *SetSchemaVersionStateOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SetSchemaVersionStateOptions) SetHeaders(param map[string]string) *SetSchemaVersionStateOptions {
	options.Headers = param
	return options
}

// UpdateGlobalRuleOptions : The UpdateGlobalRule options.
type UpdateGlobalRuleOptions struct {
	// The type of the global rule to update. Currently only `COMPATIBILITY` is supported.
	Rule *string `json:"rule" validate:"required,ne="`

	// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
	Type *string `json:"type" validate:"required"`

	// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
	Config *string `json:"config" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateGlobalRuleOptions.Rule property.
// The type of the global rule to update. Currently only `COMPATIBILITY` is supported.
const (
	UpdateGlobalRuleOptionsRuleCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the UpdateGlobalRuleOptions.Type property.
// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
const (
	UpdateGlobalRuleOptionsTypeCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the UpdateGlobalRuleOptions.Config property.
// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
const (
	UpdateGlobalRuleOptionsConfigBackwardConst           = "BACKWARD"
	UpdateGlobalRuleOptionsConfigBackwardTransitiveConst = "BACKWARD_TRANSITIVE"
	UpdateGlobalRuleOptionsConfigForwardConst            = "FORWARD"
	UpdateGlobalRuleOptionsConfigForwardTransitiveConst  = "FORWARD_TRANSITIVE"
	UpdateGlobalRuleOptionsConfigFullConst               = "FULL"
	UpdateGlobalRuleOptionsConfigFullTransitiveConst     = "FULL_TRANSITIVE"
	UpdateGlobalRuleOptionsConfigNoneConst               = "NONE"
)

// NewUpdateGlobalRuleOptions : Instantiate UpdateGlobalRuleOptions
func (*SchemaregistryV1) NewUpdateGlobalRuleOptions(rule string, typeVar string, config string) *UpdateGlobalRuleOptions {
	return &UpdateGlobalRuleOptions{
		Rule:   core.StringPtr(rule),
		Type:   core.StringPtr(typeVar),
		Config: core.StringPtr(config),
	}
}

// SetRule : Allow user to set Rule
func (_options *UpdateGlobalRuleOptions) SetRule(rule string) *UpdateGlobalRuleOptions {
	_options.Rule = core.StringPtr(rule)
	return _options
}

// SetType : Allow user to set Type
func (_options *UpdateGlobalRuleOptions) SetType(typeVar string) *UpdateGlobalRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetConfig : Allow user to set Config
func (_options *UpdateGlobalRuleOptions) SetConfig(config string) *UpdateGlobalRuleOptions {
	_options.Config = core.StringPtr(config)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateGlobalRuleOptions) SetHeaders(param map[string]string) *UpdateGlobalRuleOptions {
	options.Headers = param
	return options
}

// UpdateSchemaOptions : The UpdateSchema options.
type UpdateSchemaOptions struct {
	// The ID of the schema to update.
	ID *string `json:"id" validate:"required,ne="`

	// The AVRO schema.
	Schema map[string]interface{} `json:"schema,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateSchemaOptions : Instantiate UpdateSchemaOptions
func (*SchemaregistryV1) NewUpdateSchemaOptions(id string) *UpdateSchemaOptions {
	return &UpdateSchemaOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateSchemaOptions) SetID(id string) *UpdateSchemaOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetSchema : Allow user to set Schema
func (_options *UpdateSchemaOptions) SetSchema(schema map[string]interface{}) *UpdateSchemaOptions {
	_options.Schema = schema
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSchemaOptions) SetHeaders(param map[string]string) *UpdateSchemaOptions {
	options.Headers = param
	return options
}

// UpdateSchemaRuleOptions : The UpdateSchemaRule options.
type UpdateSchemaRuleOptions struct {
	// The ID of the schema for which to update the rule configuration.
	ID *string `json:"id" validate:"required,ne="`

	// The type of rule to update. Currently only the value that can be specified is `COMPATIBILITY`.
	Rule *string `json:"rule" validate:"required,ne="`

	// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
	Type *string `json:"type" validate:"required"`

	// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
	Config *string `json:"config" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// Constants associated with the UpdateSchemaRuleOptions.Rule property.
// The type of rule to update. Currently only the value that can be specified is `COMPATIBILITY`.
const (
	UpdateSchemaRuleOptionsRuleCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the UpdateSchemaRuleOptions.Type property.
// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
const (
	UpdateSchemaRuleOptionsTypeCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the UpdateSchemaRuleOptions.Config property.
// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
const (
	UpdateSchemaRuleOptionsConfigBackwardConst           = "BACKWARD"
	UpdateSchemaRuleOptionsConfigBackwardTransitiveConst = "BACKWARD_TRANSITIVE"
	UpdateSchemaRuleOptionsConfigForwardConst            = "FORWARD"
	UpdateSchemaRuleOptionsConfigForwardTransitiveConst  = "FORWARD_TRANSITIVE"
	UpdateSchemaRuleOptionsConfigFullConst               = "FULL"
	UpdateSchemaRuleOptionsConfigFullTransitiveConst     = "FULL_TRANSITIVE"
	UpdateSchemaRuleOptionsConfigNoneConst               = "NONE"
)

// NewUpdateSchemaRuleOptions : Instantiate UpdateSchemaRuleOptions
func (*SchemaregistryV1) NewUpdateSchemaRuleOptions(id string, rule string, typeVar string, config string) *UpdateSchemaRuleOptions {
	return &UpdateSchemaRuleOptions{
		ID:     core.StringPtr(id),
		Rule:   core.StringPtr(rule),
		Type:   core.StringPtr(typeVar),
		Config: core.StringPtr(config),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateSchemaRuleOptions) SetID(id string) *UpdateSchemaRuleOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetRule : Allow user to set Rule
func (_options *UpdateSchemaRuleOptions) SetRule(rule string) *UpdateSchemaRuleOptions {
	_options.Rule = core.StringPtr(rule)
	return _options
}

// SetType : Allow user to set Type
func (_options *UpdateSchemaRuleOptions) SetType(typeVar string) *UpdateSchemaRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetConfig : Allow user to set Config
func (_options *UpdateSchemaRuleOptions) SetConfig(config string) *UpdateSchemaRuleOptions {
	_options.Config = core.StringPtr(config)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSchemaRuleOptions) SetHeaders(param map[string]string) *UpdateSchemaRuleOptions {
	options.Headers = param
	return options
}

// AvroSchema : AvroSchema struct
type AvroSchema struct {
	// The AVRO schema.
	Schema map[string]interface{} `json:"schema,omitempty"`
}

// UnmarshalAvroSchema unmarshals an instance of AvroSchema from the specified map of raw messages.
func UnmarshalAvroSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AvroSchema)
	err = core.UnmarshalPrimitive(m, "schema", &obj.Schema)
	if err != nil {
		err = core.SDKErrorf(err, "", "schema-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : Rules define constraints on whether the schema registry will accept a new version of a schema.
type Rule struct {
	// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
	Type *string `json:"type" validate:"required"`

	// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
	Config *string `json:"config" validate:"required"`
}

// Constants associated with the Rule.Type property.
// The type of the rule. Currently only one type is supported (`COMPATIBILITY`).
const (
	RuleTypeCompatibilityConst = "COMPATIBILITY"
)

// Constants associated with the Rule.Config property.
// The configuration value for the rule. Which values are valid depends on the value of this object's `type` property.
const (
	RuleConfigBackwardConst           = "BACKWARD"
	RuleConfigBackwardTransitiveConst = "BACKWARD_TRANSITIVE"
	RuleConfigForwardConst            = "FORWARD"
	RuleConfigForwardTransitiveConst  = "FORWARD_TRANSITIVE"
	RuleConfigFullConst               = "FULL"
	RuleConfigFullTransitiveConst     = "FULL_TRANSITIVE"
	RuleConfigNoneConst               = "NONE"
)

// NewRule : Instantiate Rule (Generic Model Constructor)
func (*SchemaregistryV1) NewRule(typeVar string, config string) (_model *Rule, err error) {
	_model = &Rule{
		Type:   core.StringPtr(typeVar),
		Config: core.StringPtr(config),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "config", &obj.Config)
	if err != nil {
		err = core.SDKErrorf(err, "", "config-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchemaMetadata : Information about a schema version.
type SchemaMetadata struct {
	// Creation timestamp of the schema in UNIX epoc format.
	CreatedOn *int64 `json:"createdOn" validate:"required"`

	// Globally unique ID assigned to the initial version of the schema.
	GlobalID *int64 `json:"globalId" validate:"required"`

	// The ID of the schema. This is either taken from the `X-Registry-ArtifactId` header when the request is made to
	// create the schema or is an automatically assigned UUID value.
	ID *string `json:"id" validate:"required"`

	// Last modification timestamp of the schema in UNIX epoc format.
	ModifiedOn *int64 `json:"modifiedOn" validate:"required"`

	// Type of the schema. Always the string `AVRO`.
	Type *string `json:"type" validate:"required"`

	// Version number assigned to this version of the schema.
	Version *int64 `json:"version" validate:"required"`
}

// UnmarshalSchemaMetadata unmarshals an instance of SchemaMetadata from the specified map of raw messages.
func UnmarshalSchemaMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchemaMetadata)
	err = core.UnmarshalPrimitive(m, "createdOn", &obj.CreatedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "createdOn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "globalId", &obj.GlobalID)
	if err != nil {
		err = core.SDKErrorf(err, "", "globalId-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "modifiedOn", &obj.ModifiedOn)
	if err != nil {
		err = core.SDKErrorf(err, "", "modifiedOn-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		err = core.SDKErrorf(err, "", "version-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
