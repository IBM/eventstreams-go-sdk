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

// Package adminrestv1 : Operations and models for the AdminrestV1 service
package adminrestv1

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

// AdminrestV1 : The administration REST API for IBM Event Streams on Cloud.
//
// API Version: 1.3.0
type AdminrestV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "adminrest"

// AdminrestV1Options : Service options
type AdminrestV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAdminrestV1UsingExternalConfig : constructs an instance of AdminrestV1 with passed in options and external configuration.
func NewAdminrestV1UsingExternalConfig(options *AdminrestV1Options) (adminrest *AdminrestV1, err error) {
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

	adminrest, err = NewAdminrestV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = adminrest.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = adminrest.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewAdminrestV1 : constructs an instance of AdminrestV1 with passed in options.
func NewAdminrestV1(options *AdminrestV1Options) (service *AdminrestV1, err error) {
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

	service = &AdminrestV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "adminrest" suitable for processing requests.
func (adminrest *AdminrestV1) Clone() *AdminrestV1 {
	if core.IsNil(adminrest) {
		return nil
	}
	clone := *adminrest
	clone.Service = adminrest.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (adminrest *AdminrestV1) SetServiceURL(url string) error {
	err := adminrest.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (adminrest *AdminrestV1) GetServiceURL() string {
	return adminrest.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (adminrest *AdminrestV1) SetDefaultHeaders(headers http.Header) {
	adminrest.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (adminrest *AdminrestV1) SetEnableGzipCompression(enableGzip bool) {
	adminrest.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (adminrest *AdminrestV1) GetEnableGzipCompression() bool {
	return adminrest.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (adminrest *AdminrestV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	adminrest.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (adminrest *AdminrestV1) DisableRetries() {
	adminrest.Service.DisableRetries()
}

// CreateTopic : Create a new topic
// Create a new topic.
func (adminrest *AdminrestV1) CreateTopic(createTopicOptions *CreateTopicOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.CreateTopicWithContext(context.Background(), createTopicOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateTopicWithContext is an alternate form of the CreateTopic method which supports a Context parameter
func (adminrest *AdminrestV1) CreateTopicWithContext(ctx context.Context, createTopicOptions *CreateTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTopicOptions, "createTopicOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createTopicOptions, "createTopicOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "CreateTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createTopicOptions.Name != nil {
		body["name"] = createTopicOptions.Name
	}
	if createTopicOptions.Partitions != nil {
		body["partitions"] = createTopicOptions.Partitions
	}
	if createTopicOptions.PartitionCount != nil {
		body["partition_count"] = createTopicOptions.PartitionCount
	}
	if createTopicOptions.Configs != nil {
		body["configs"] = createTopicOptions.Configs
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

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "CreateTopic", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// Alive : Basic health check for Admin REST API
func (adminrest *AdminrestV1) Alive(aliveOptions *AliveOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.AliveWithContext(context.Background(), aliveOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// AliveWithContext is an alternate form of the Alive method which supports a Context parameter
func (adminrest *AdminrestV1) AliveWithContext(ctx context.Context, aliveOptions *AliveOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(aliveOptions, "aliveOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/alive`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range aliveOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "Alive")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "alive", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListTopics : Get a list of topics
// Returns a list containing information about all of the Kafka topics that are defined for an instance of the Event
// Streams service. If there are currently no topics defined then an empty list is returned.
func (adminrest *AdminrestV1) ListTopics(listTopicsOptions *ListTopicsOptions) (result []TopicDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.ListTopicsWithContext(context.Background(), listTopicsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListTopicsWithContext is an alternate form of the ListTopics method which supports a Context parameter
func (adminrest *AdminrestV1) ListTopicsWithContext(ctx context.Context, listTopicsOptions *ListTopicsOptions) (result []TopicDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTopicsOptions, "listTopicsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listTopicsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ListTopics")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listTopicsOptions.TopicFilter != nil {
		builder.AddQuery("topic_filter", fmt.Sprint(*listTopicsOptions.TopicFilter))
	}
	if listTopicsOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listTopicsOptions.PerPage))
	}
	if listTopicsOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listTopicsOptions.Page))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse []json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "ListTopics", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTopicDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetTopic : Get detailed information on a topic
// Get detailed information on a topic.
func (adminrest *AdminrestV1) GetTopic(getTopicOptions *GetTopicOptions) (result *TopicDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetTopicWithContext(context.Background(), getTopicOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetTopicWithContext is an alternate form of the GetTopic method which supports a Context parameter
func (adminrest *AdminrestV1) GetTopicWithContext(ctx context.Context, getTopicOptions *GetTopicOptions) (result *TopicDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTopicOptions, "getTopicOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getTopicOptions, "getTopicOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *getTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetTopic")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetTopic", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTopicDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteTopic : Delete a topic
// Delete a topic.
func (adminrest *AdminrestV1) DeleteTopic(deleteTopicOptions *DeleteTopicOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.DeleteTopicWithContext(context.Background(), deleteTopicOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteTopicWithContext is an alternate form of the DeleteTopic method which supports a Context parameter
func (adminrest *AdminrestV1) DeleteTopicWithContext(ctx context.Context, deleteTopicOptions *DeleteTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTopicOptions, "deleteTopicOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteTopicOptions, "deleteTopicOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *deleteTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "DeleteTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "DeleteTopic", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateTopic : Increase the number of partitions and/or update one or more topic configuration parameters
// Increase the number of partitions and/or update one or more topic configuration parameters.
func (adminrest *AdminrestV1) UpdateTopic(updateTopicOptions *UpdateTopicOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.UpdateTopicWithContext(context.Background(), updateTopicOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateTopicWithContext is an alternate form of the UpdateTopic method which supports a Context parameter
func (adminrest *AdminrestV1) UpdateTopicWithContext(ctx context.Context, updateTopicOptions *UpdateTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTopicOptions, "updateTopicOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateTopicOptions, "updateTopicOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *updateTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "UpdateTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateTopicOptions.NewTotalPartitionCount != nil {
		body["new_total_partition_count"] = updateTopicOptions.NewTotalPartitionCount
	}
	if updateTopicOptions.Configs != nil {
		body["configs"] = updateTopicOptions.Configs
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

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "UpdateTopic", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// DeleteTopicRecords : Delete records before the given offset on a topic
// Delete records before the given offset on a topic.
func (adminrest *AdminrestV1) DeleteTopicRecords(deleteTopicRecordsOptions *DeleteTopicRecordsOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.DeleteTopicRecordsWithContext(context.Background(), deleteTopicRecordsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteTopicRecordsWithContext is an alternate form of the DeleteTopicRecords method which supports a Context parameter
func (adminrest *AdminrestV1) DeleteTopicRecordsWithContext(ctx context.Context, deleteTopicRecordsOptions *DeleteTopicRecordsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTopicRecordsOptions, "deleteTopicRecordsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteTopicRecordsOptions, "deleteTopicRecordsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *deleteTopicRecordsOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}/records`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteTopicRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "DeleteTopicRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if deleteTopicRecordsOptions.RecordsToDelete != nil {
		body["records_to_delete"] = deleteTopicRecordsOptions.RecordsToDelete
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

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "DeleteTopicRecords", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// CreateQuota : Create a new quota
// Create a new quota.
func (adminrest *AdminrestV1) CreateQuota(createQuotaOptions *CreateQuotaOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.CreateQuotaWithContext(context.Background(), createQuotaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateQuotaWithContext is an alternate form of the CreateQuota method which supports a Context parameter
func (adminrest *AdminrestV1) CreateQuotaWithContext(ctx context.Context, createQuotaOptions *CreateQuotaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createQuotaOptions, "createQuotaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createQuotaOptions, "createQuotaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"entity_name": *createQuotaOptions.EntityName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/quotas/{entity_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range createQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "CreateQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createQuotaOptions.ProducerByteRate != nil {
		body["producer_byte_rate"] = createQuotaOptions.ProducerByteRate
	}
	if createQuotaOptions.ConsumerByteRate != nil {
		body["consumer_byte_rate"] = createQuotaOptions.ConsumerByteRate
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

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_quota", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateQuota : Update a quota
// Update an entity's quota.
func (adminrest *AdminrestV1) UpdateQuota(updateQuotaOptions *UpdateQuotaOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.UpdateQuotaWithContext(context.Background(), updateQuotaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateQuotaWithContext is an alternate form of the UpdateQuota method which supports a Context parameter
func (adminrest *AdminrestV1) UpdateQuotaWithContext(ctx context.Context, updateQuotaOptions *UpdateQuotaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateQuotaOptions, "updateQuotaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateQuotaOptions, "updateQuotaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"entity_name": *updateQuotaOptions.EntityName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/quotas/{entity_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "UpdateQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateQuotaOptions.ProducerByteRate != nil {
		body["producer_byte_rate"] = updateQuotaOptions.ProducerByteRate
	}
	if updateQuotaOptions.ConsumerByteRate != nil {
		body["consumer_byte_rate"] = updateQuotaOptions.ConsumerByteRate
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

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "update_quota", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// DeleteQuota : Delete a quota
// Delete an entity's quota.
func (adminrest *AdminrestV1) DeleteQuota(deleteQuotaOptions *DeleteQuotaOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.DeleteQuotaWithContext(context.Background(), deleteQuotaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteQuotaWithContext is an alternate form of the DeleteQuota method which supports a Context parameter
func (adminrest *AdminrestV1) DeleteQuotaWithContext(ctx context.Context, deleteQuotaOptions *DeleteQuotaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteQuotaOptions, "deleteQuotaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteQuotaOptions, "deleteQuotaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"entity_name": *deleteQuotaOptions.EntityName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/quotas/{entity_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "DeleteQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_quota", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetQuota : Get quota information for an entity
// Get quota information for an entity.
func (adminrest *AdminrestV1) GetQuota(getQuotaOptions *GetQuotaOptions) (result *QuotaDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetQuotaWithContext(context.Background(), getQuotaOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetQuotaWithContext is an alternate form of the GetQuota method which supports a Context parameter
func (adminrest *AdminrestV1) GetQuotaWithContext(ctx context.Context, getQuotaOptions *GetQuotaOptions) (result *QuotaDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getQuotaOptions, "getQuotaOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getQuotaOptions, "getQuotaOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"entity_name": *getQuotaOptions.EntityName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/quotas/{entity_name}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetQuota")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_quota", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQuotaDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListQuotas : List each entity's quota information
// List each entity's quota information.
func (adminrest *AdminrestV1) ListQuotas(listQuotasOptions *ListQuotasOptions) (result *QuotaList, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.ListQuotasWithContext(context.Background(), listQuotasOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListQuotasWithContext is an alternate form of the ListQuotas method which supports a Context parameter
func (adminrest *AdminrestV1) ListQuotasWithContext(ctx context.Context, listQuotasOptions *ListQuotasOptions) (result *QuotaList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listQuotasOptions, "listQuotasOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/quotas`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listQuotasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ListQuotas")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_quotas", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalQuotaList)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListBrokers : Get a list of brokers in the cluster
// Get a list of brokers in the cluster.
func (adminrest *AdminrestV1) ListBrokers(listBrokersOptions *ListBrokersOptions) (result []BrokerSummary, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.ListBrokersWithContext(context.Background(), listBrokersOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListBrokersWithContext is an alternate form of the ListBrokers method which supports a Context parameter
func (adminrest *AdminrestV1) ListBrokersWithContext(ctx context.Context, listBrokersOptions *ListBrokersOptions) (result []BrokerSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listBrokersOptions, "listBrokersOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/brokers`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listBrokersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ListBrokers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse []json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "ListBrokers", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrokerSummary)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBroker : Get detailed information for a single broker
// Get detailed information for a single broker.
func (adminrest *AdminrestV1) GetBroker(getBrokerOptions *GetBrokerOptions) (result *BrokerDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetBrokerWithContext(context.Background(), getBrokerOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBrokerWithContext is an alternate form of the GetBroker method which supports a Context parameter
func (adminrest *AdminrestV1) GetBrokerWithContext(ctx context.Context, getBrokerOptions *GetBrokerOptions) (result *BrokerDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBrokerOptions, "getBrokerOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getBrokerOptions, "getBrokerOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"broker_id": fmt.Sprint(*getBrokerOptions.BrokerID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/brokers/{broker_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBrokerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetBroker")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetBroker", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrokerDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetBrokerConfig : Get all configuration parameters for a single broker
// Get all configuration parameters for a single broker.
func (adminrest *AdminrestV1) GetBrokerConfig(getBrokerConfigOptions *GetBrokerConfigOptions) (result *BrokerDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetBrokerConfigWithContext(context.Background(), getBrokerConfigOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetBrokerConfigWithContext is an alternate form of the GetBrokerConfig method which supports a Context parameter
func (adminrest *AdminrestV1) GetBrokerConfigWithContext(ctx context.Context, getBrokerConfigOptions *GetBrokerConfigOptions) (result *BrokerDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBrokerConfigOptions, "getBrokerConfigOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getBrokerConfigOptions, "getBrokerConfigOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"broker_id": fmt.Sprint(*getBrokerConfigOptions.BrokerID),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/brokers/{broker_id}/configs`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getBrokerConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetBrokerConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getBrokerConfigOptions.ConfigFilter != nil {
		builder.AddQuery("config_filter", fmt.Sprint(*getBrokerConfigOptions.ConfigFilter))
	}
	if getBrokerConfigOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*getBrokerConfigOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetBrokerConfig", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrokerDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetCluster : Get information about the cluster
// Get information about the cluster.
func (adminrest *AdminrestV1) GetCluster(getClusterOptions *GetClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetClusterWithContext(context.Background(), getClusterOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetClusterWithContext is an alternate form of the GetCluster method which supports a Context parameter
func (adminrest *AdminrestV1) GetClusterWithContext(ctx context.Context, getClusterOptions *GetClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getClusterOptions, "getClusterOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/cluster`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetCluster")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetCluster", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCluster)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ListConsumerGroups : Get a list of consumer group IDs
// Get a list of consumer group IDs.
func (adminrest *AdminrestV1) ListConsumerGroups(listConsumerGroupsOptions *ListConsumerGroupsOptions) (result []string, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.ListConsumerGroupsWithContext(context.Background(), listConsumerGroupsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListConsumerGroupsWithContext is an alternate form of the ListConsumerGroups method which supports a Context parameter
func (adminrest *AdminrestV1) ListConsumerGroupsWithContext(ctx context.Context, listConsumerGroupsOptions *ListConsumerGroupsOptions) (result []string, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listConsumerGroupsOptions, "listConsumerGroupsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/consumergroups`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range listConsumerGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ListConsumerGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listConsumerGroupsOptions.GroupFilter != nil {
		builder.AddQuery("group_filter", fmt.Sprint(*listConsumerGroupsOptions.GroupFilter))
	}
	if listConsumerGroupsOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listConsumerGroupsOptions.PerPage))
	}
	if listConsumerGroupsOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listConsumerGroupsOptions.Page))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = adminrest.Service.Request(request, &result)
	if err != nil {
		core.EnrichHTTPProblem(err, "ListConsumerGroups", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// GetConsumerGroup : Get detailed information on a consumer group
// Get detailed information on a consumer group.
func (adminrest *AdminrestV1) GetConsumerGroup(getConsumerGroupOptions *GetConsumerGroupOptions) (result *GroupDetail, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetConsumerGroupWithContext(context.Background(), getConsumerGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetConsumerGroupWithContext is an alternate form of the GetConsumerGroup method which supports a Context parameter
func (adminrest *AdminrestV1) GetConsumerGroupWithContext(ctx context.Context, getConsumerGroupOptions *GetConsumerGroupOptions) (result *GroupDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConsumerGroupOptions, "getConsumerGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getConsumerGroupOptions, "getConsumerGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"group_id": *getConsumerGroupOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/consumergroups/{group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getConsumerGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetConsumerGroup")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetConsumerGroup", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroupDetail)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteConsumerGroup : Delete a consumer group
// Delete a consumer group.
func (adminrest *AdminrestV1) DeleteConsumerGroup(deleteConsumerGroupOptions *DeleteConsumerGroupOptions) (response *core.DetailedResponse, err error) {
	response, err = adminrest.DeleteConsumerGroupWithContext(context.Background(), deleteConsumerGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteConsumerGroupWithContext is an alternate form of the DeleteConsumerGroup method which supports a Context parameter
func (adminrest *AdminrestV1) DeleteConsumerGroupWithContext(ctx context.Context, deleteConsumerGroupOptions *DeleteConsumerGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConsumerGroupOptions, "deleteConsumerGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteConsumerGroupOptions, "deleteConsumerGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"group_id": *deleteConsumerGroupOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/consumergroups/{group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range deleteConsumerGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "DeleteConsumerGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = adminrest.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "DeleteConsumerGroup", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// UpdateConsumerGroup : Update the offsets of a consumer group
// Update the offsets of a consumer group using various modes, eg. latest, earliest, datetime,etc.
func (adminrest *AdminrestV1) UpdateConsumerGroup(updateConsumerGroupOptions *UpdateConsumerGroupOptions) (result []GroupResetResultsItem, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.UpdateConsumerGroupWithContext(context.Background(), updateConsumerGroupOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// UpdateConsumerGroupWithContext is an alternate form of the UpdateConsumerGroup method which supports a Context parameter
func (adminrest *AdminrestV1) UpdateConsumerGroupWithContext(ctx context.Context, updateConsumerGroupOptions *UpdateConsumerGroupOptions) (result []GroupResetResultsItem, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConsumerGroupOptions, "updateConsumerGroupOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(updateConsumerGroupOptions, "updateConsumerGroupOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"group_id": *updateConsumerGroupOptions.GroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/consumergroups/{group_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range updateConsumerGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "UpdateConsumerGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateConsumerGroupOptions.Topic != nil {
		body["topic"] = updateConsumerGroupOptions.Topic
	}
	if updateConsumerGroupOptions.Mode != nil {
		body["mode"] = updateConsumerGroupOptions.Mode
	}
	if updateConsumerGroupOptions.Value != nil {
		body["value"] = updateConsumerGroupOptions.Value
	}
	if updateConsumerGroupOptions.Execute != nil {
		body["execute"] = updateConsumerGroupOptions.Execute
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

	var rawResponse []json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "UpdateConsumerGroup", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroupResetResultsItem)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMirroringTopicSelection : Get current topic selection for mirroring
// Get current topic selection for mirroring.
func (adminrest *AdminrestV1) GetMirroringTopicSelection(getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetMirroringTopicSelectionWithContext(context.Background(), getMirroringTopicSelectionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMirroringTopicSelectionWithContext is an alternate form of the GetMirroringTopicSelection method which supports a Context parameter
func (adminrest *AdminrestV1) GetMirroringTopicSelectionWithContext(ctx context.Context, getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMirroringTopicSelectionOptions, "getMirroringTopicSelectionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/topic-selection`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMirroringTopicSelectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetMirroringTopicSelection")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetMirroringTopicSelection", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringTopicSelection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceMirroringTopicSelection : Replace topic selection for mirroring
// Replace topic selection for mirroring. This operation replaces the complete set of mirroring topic selections.
func (adminrest *AdminrestV1) ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.ReplaceMirroringTopicSelectionWithContext(context.Background(), replaceMirroringTopicSelectionOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceMirroringTopicSelectionWithContext is an alternate form of the ReplaceMirroringTopicSelection method which supports a Context parameter
func (adminrest *AdminrestV1) ReplaceMirroringTopicSelectionWithContext(ctx context.Context, replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceMirroringTopicSelectionOptions, "replaceMirroringTopicSelectionOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceMirroringTopicSelectionOptions, "replaceMirroringTopicSelectionOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/topic-selection`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range replaceMirroringTopicSelectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ReplaceMirroringTopicSelection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceMirroringTopicSelectionOptions.Includes != nil {
		body["includes"] = replaceMirroringTopicSelectionOptions.Includes
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "ReplaceMirroringTopicSelection", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringTopicSelection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetMirroringActiveTopics : Get topics that are being actively mirrored
// Get topics that are being actively mirrored.
func (adminrest *AdminrestV1) GetMirroringActiveTopics(getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions) (result *MirroringActiveTopics, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetMirroringActiveTopicsWithContext(context.Background(), getMirroringActiveTopicsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetMirroringActiveTopicsWithContext is an alternate form of the GetMirroringActiveTopics method which supports a Context parameter
func (adminrest *AdminrestV1) GetMirroringActiveTopicsWithContext(ctx context.Context, getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions) (result *MirroringActiveTopics, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMirroringActiveTopicsOptions, "getMirroringActiveTopicsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/active-topics`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getMirroringActiveTopicsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetMirroringActiveTopics")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetMirroringActiveTopics", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringActiveTopics)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetStatus : Get the status of the instance
// Get the status of the instance.
func (adminrest *AdminrestV1) GetStatus(getStatusOptions *GetStatusOptions) (result *InstanceStatus, response *core.DetailedResponse, err error) {
	result, response, err = adminrest.GetStatusWithContext(context.Background(), getStatusOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetStatusWithContext is an alternate form of the GetStatus method which supports a Context parameter
func (adminrest *AdminrestV1) GetStatusWithContext(ctx context.Context, getStatusOptions *GetStatusOptions) (result *InstanceStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getStatusOptions, "getStatusOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/status`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	for headerName, headerValue := range getStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetStatus")
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
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "GetStatus", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceStatus)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.3.0")
}

// AliveOptions : The Alive options.
type AliveOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewAliveOptions : Instantiate AliveOptions
func (*AdminrestV1) NewAliveOptions() *AliveOptions {
	return &AliveOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *AliveOptions) SetHeaders(param map[string]string) *AliveOptions {
	options.Headers = param
	return options
}

// BrokerDetailConfigsItem : BrokerDetailConfigsItem struct
type BrokerDetailConfigsItem struct {
	// The name of the config property.
	Name *string `json:"name,omitempty"`

	// The value for a config property.
	Value *string `json:"value,omitempty"`

	// When true, the value cannot be displayed and will be returned with a null value.
	IsSensitive *bool `json:"is_sensitive,omitempty"`
}

// UnmarshalBrokerDetailConfigsItem unmarshals an instance of BrokerDetailConfigsItem from the specified map of raw messages.
func UnmarshalBrokerDetailConfigsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrokerDetailConfigsItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_sensitive", &obj.IsSensitive)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_sensitive-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateQuotaOptions : The CreateQuota options.
type CreateQuotaOptions struct {
	// The entity name of the quotas can be `default` or an IAM Service ID that starts with an `iam-ServiceId` prefix.
	EntityName *string `json:"entity_name" validate:"required,ne="`

	// The producer byte rate quota value.
	ProducerByteRate *int64 `json:"producer_byte_rate,omitempty"`

	// The consumer byte rate quota value.
	ConsumerByteRate *int64 `json:"consumer_byte_rate,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateQuotaOptions : Instantiate CreateQuotaOptions
func (*AdminrestV1) NewCreateQuotaOptions(entityName string) *CreateQuotaOptions {
	return &CreateQuotaOptions{
		EntityName: core.StringPtr(entityName),
	}
}

// SetEntityName : Allow user to set EntityName
func (_options *CreateQuotaOptions) SetEntityName(entityName string) *CreateQuotaOptions {
	_options.EntityName = core.StringPtr(entityName)
	return _options
}

// SetProducerByteRate : Allow user to set ProducerByteRate
func (_options *CreateQuotaOptions) SetProducerByteRate(producerByteRate int64) *CreateQuotaOptions {
	_options.ProducerByteRate = core.Int64Ptr(producerByteRate)
	return _options
}

// SetConsumerByteRate : Allow user to set ConsumerByteRate
func (_options *CreateQuotaOptions) SetConsumerByteRate(consumerByteRate int64) *CreateQuotaOptions {
	_options.ConsumerByteRate = core.Int64Ptr(consumerByteRate)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateQuotaOptions) SetHeaders(param map[string]string) *CreateQuotaOptions {
	options.Headers = param
	return options
}

// CreateTopicOptions : The CreateTopic options.
type CreateTopicOptions struct {
	// The name of topic to be created.
	Name *string `json:"name,omitempty"`

	// The number of partitions.
	Partitions *int64 `json:"partitions,omitempty"`

	// The number of partitions, this field takes precedence over 'partitions'. Default value is 1 if not specified.
	PartitionCount *int64 `json:"partition_count,omitempty"`

	// The config properties to be set for the new topic.
	Configs []TopicCreateRequestConfigsItem `json:"configs,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateTopicOptions : Instantiate CreateTopicOptions
func (*AdminrestV1) NewCreateTopicOptions() *CreateTopicOptions {
	return &CreateTopicOptions{}
}

// SetName : Allow user to set Name
func (_options *CreateTopicOptions) SetName(name string) *CreateTopicOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPartitions : Allow user to set Partitions
func (_options *CreateTopicOptions) SetPartitions(partitions int64) *CreateTopicOptions {
	_options.Partitions = core.Int64Ptr(partitions)
	return _options
}

// SetPartitionCount : Allow user to set PartitionCount
func (_options *CreateTopicOptions) SetPartitionCount(partitionCount int64) *CreateTopicOptions {
	_options.PartitionCount = core.Int64Ptr(partitionCount)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *CreateTopicOptions) SetConfigs(configs []TopicCreateRequestConfigsItem) *CreateTopicOptions {
	_options.Configs = configs
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTopicOptions) SetHeaders(param map[string]string) *CreateTopicOptions {
	options.Headers = param
	return options
}

// DeleteConsumerGroupOptions : The DeleteConsumerGroup options.
type DeleteConsumerGroupOptions struct {
	// The group ID for the consumer group to be deleted.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteConsumerGroupOptions : Instantiate DeleteConsumerGroupOptions
func (*AdminrestV1) NewDeleteConsumerGroupOptions(groupID string) *DeleteConsumerGroupOptions {
	return &DeleteConsumerGroupOptions{
		GroupID: core.StringPtr(groupID),
	}
}

// SetGroupID : Allow user to set GroupID
func (_options *DeleteConsumerGroupOptions) SetGroupID(groupID string) *DeleteConsumerGroupOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConsumerGroupOptions) SetHeaders(param map[string]string) *DeleteConsumerGroupOptions {
	options.Headers = param
	return options
}

// DeleteQuotaOptions : The DeleteQuota options.
type DeleteQuotaOptions struct {
	// The entity name of the quotas can be `default` or an IAM Service ID that starts with an `iam-ServiceId` prefix.
	EntityName *string `json:"entity_name" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteQuotaOptions : Instantiate DeleteQuotaOptions
func (*AdminrestV1) NewDeleteQuotaOptions(entityName string) *DeleteQuotaOptions {
	return &DeleteQuotaOptions{
		EntityName: core.StringPtr(entityName),
	}
}

// SetEntityName : Allow user to set EntityName
func (_options *DeleteQuotaOptions) SetEntityName(entityName string) *DeleteQuotaOptions {
	_options.EntityName = core.StringPtr(entityName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteQuotaOptions) SetHeaders(param map[string]string) *DeleteQuotaOptions {
	options.Headers = param
	return options
}

// DeleteTopicOptions : The DeleteTopic options.
type DeleteTopicOptions struct {
	// The topic name for the topic to be deleted.
	TopicName *string `json:"topic_name" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteTopicOptions : Instantiate DeleteTopicOptions
func (*AdminrestV1) NewDeleteTopicOptions(topicName string) *DeleteTopicOptions {
	return &DeleteTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (_options *DeleteTopicOptions) SetTopicName(topicName string) *DeleteTopicOptions {
	_options.TopicName = core.StringPtr(topicName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTopicOptions) SetHeaders(param map[string]string) *DeleteTopicOptions {
	options.Headers = param
	return options
}

// DeleteTopicRecordsOptions : The DeleteTopicRecords options.
type DeleteTopicRecordsOptions struct {
	// The topic name of the records to be deleted.
	TopicName *string `json:"topic_name" validate:"required,ne="`

	RecordsToDelete []RecordDeleteRequestRecordsToDeleteItem `json:"records_to_delete,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteTopicRecordsOptions : Instantiate DeleteTopicRecordsOptions
func (*AdminrestV1) NewDeleteTopicRecordsOptions(topicName string) *DeleteTopicRecordsOptions {
	return &DeleteTopicRecordsOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (_options *DeleteTopicRecordsOptions) SetTopicName(topicName string) *DeleteTopicRecordsOptions {
	_options.TopicName = core.StringPtr(topicName)
	return _options
}

// SetRecordsToDelete : Allow user to set RecordsToDelete
func (_options *DeleteTopicRecordsOptions) SetRecordsToDelete(recordsToDelete []RecordDeleteRequestRecordsToDeleteItem) *DeleteTopicRecordsOptions {
	_options.RecordsToDelete = recordsToDelete
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTopicRecordsOptions) SetHeaders(param map[string]string) *DeleteTopicRecordsOptions {
	options.Headers = param
	return options
}

// GetBrokerConfigOptions : The GetBrokerConfig options.
type GetBrokerConfigOptions struct {
	// The broker ID of the broker to be described.
	BrokerID *int64 `json:"broker_id" validate:"required"`

	// A filter to be applied to the config names. A simple filter can be specified as a string with asterisk (`*`)
	// wildcards representing 0 or more characters, e.g. `file*` will filter all config names that begin with the string
	// `file` followed by any character sequence. A more complex filter pattern can be used by surrounding a regular
	// expression in forward slash (`/`) delimiters, e.g. `/file.* /`.
	ConfigFilter *string `json:"config_filter,omitempty"`

	// When true, all information about the config properties is returned including the source of the configuration
	// indicating its scope and whether it's dynamic.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBrokerConfigOptions : Instantiate GetBrokerConfigOptions
func (*AdminrestV1) NewGetBrokerConfigOptions(brokerID int64) *GetBrokerConfigOptions {
	return &GetBrokerConfigOptions{
		BrokerID: core.Int64Ptr(brokerID),
	}
}

// SetBrokerID : Allow user to set BrokerID
func (_options *GetBrokerConfigOptions) SetBrokerID(brokerID int64) *GetBrokerConfigOptions {
	_options.BrokerID = core.Int64Ptr(brokerID)
	return _options
}

// SetConfigFilter : Allow user to set ConfigFilter
func (_options *GetBrokerConfigOptions) SetConfigFilter(configFilter string) *GetBrokerConfigOptions {
	_options.ConfigFilter = core.StringPtr(configFilter)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *GetBrokerConfigOptions) SetVerbose(verbose bool) *GetBrokerConfigOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBrokerConfigOptions) SetHeaders(param map[string]string) *GetBrokerConfigOptions {
	options.Headers = param
	return options
}

// GetBrokerOptions : The GetBroker options.
type GetBrokerOptions struct {
	// The broker ID of the broker to be described.
	BrokerID *int64 `json:"broker_id" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetBrokerOptions : Instantiate GetBrokerOptions
func (*AdminrestV1) NewGetBrokerOptions(brokerID int64) *GetBrokerOptions {
	return &GetBrokerOptions{
		BrokerID: core.Int64Ptr(brokerID),
	}
}

// SetBrokerID : Allow user to set BrokerID
func (_options *GetBrokerOptions) SetBrokerID(brokerID int64) *GetBrokerOptions {
	_options.BrokerID = core.Int64Ptr(brokerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBrokerOptions) SetHeaders(param map[string]string) *GetBrokerOptions {
	options.Headers = param
	return options
}

// GetClusterOptions : The GetCluster options.
type GetClusterOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetClusterOptions : Instantiate GetClusterOptions
func (*AdminrestV1) NewGetClusterOptions() *GetClusterOptions {
	return &GetClusterOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetClusterOptions) SetHeaders(param map[string]string) *GetClusterOptions {
	options.Headers = param
	return options
}

// GetConsumerGroupOptions : The GetConsumerGroup options.
type GetConsumerGroupOptions struct {
	// The group ID for the consumer group to be described.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetConsumerGroupOptions : Instantiate GetConsumerGroupOptions
func (*AdminrestV1) NewGetConsumerGroupOptions(groupID string) *GetConsumerGroupOptions {
	return &GetConsumerGroupOptions{
		GroupID: core.StringPtr(groupID),
	}
}

// SetGroupID : Allow user to set GroupID
func (_options *GetConsumerGroupOptions) SetGroupID(groupID string) *GetConsumerGroupOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConsumerGroupOptions) SetHeaders(param map[string]string) *GetConsumerGroupOptions {
	options.Headers = param
	return options
}

// GetMirroringActiveTopicsOptions : The GetMirroringActiveTopics options.
type GetMirroringActiveTopicsOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMirroringActiveTopicsOptions : Instantiate GetMirroringActiveTopicsOptions
func (*AdminrestV1) NewGetMirroringActiveTopicsOptions() *GetMirroringActiveTopicsOptions {
	return &GetMirroringActiveTopicsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMirroringActiveTopicsOptions) SetHeaders(param map[string]string) *GetMirroringActiveTopicsOptions {
	options.Headers = param
	return options
}

// GetMirroringTopicSelectionOptions : The GetMirroringTopicSelection options.
type GetMirroringTopicSelectionOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetMirroringTopicSelectionOptions : Instantiate GetMirroringTopicSelectionOptions
func (*AdminrestV1) NewGetMirroringTopicSelectionOptions() *GetMirroringTopicSelectionOptions {
	return &GetMirroringTopicSelectionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMirroringTopicSelectionOptions) SetHeaders(param map[string]string) *GetMirroringTopicSelectionOptions {
	options.Headers = param
	return options
}

// GetQuotaOptions : The GetQuota options.
type GetQuotaOptions struct {
	// The entity name of the quotas can be `default` or an IAM Service ID that starts with an `iam-ServiceId` prefix.
	EntityName *string `json:"entity_name" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetQuotaOptions : Instantiate GetQuotaOptions
func (*AdminrestV1) NewGetQuotaOptions(entityName string) *GetQuotaOptions {
	return &GetQuotaOptions{
		EntityName: core.StringPtr(entityName),
	}
}

// SetEntityName : Allow user to set EntityName
func (_options *GetQuotaOptions) SetEntityName(entityName string) *GetQuotaOptions {
	_options.EntityName = core.StringPtr(entityName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetQuotaOptions) SetHeaders(param map[string]string) *GetQuotaOptions {
	options.Headers = param
	return options
}

// GetStatusOptions : The GetStatus options.
type GetStatusOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetStatusOptions : Instantiate GetStatusOptions
func (*AdminrestV1) NewGetStatusOptions() *GetStatusOptions {
	return &GetStatusOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetStatusOptions) SetHeaders(param map[string]string) *GetStatusOptions {
	options.Headers = param
	return options
}

// GetTopicOptions : The GetTopic options.
type GetTopicOptions struct {
	// The topic name for the topic to be described.
	TopicName *string `json:"topic_name" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetTopicOptions : Instantiate GetTopicOptions
func (*AdminrestV1) NewGetTopicOptions(topicName string) *GetTopicOptions {
	return &GetTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (_options *GetTopicOptions) SetTopicName(topicName string) *GetTopicOptions {
	_options.TopicName = core.StringPtr(topicName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTopicOptions) SetHeaders(param map[string]string) *GetTopicOptions {
	options.Headers = param
	return options
}

// GroupResetResultsItem : The new offset for one partition of one topic after resetting consumer group's offset.
type GroupResetResultsItem struct {
	Topic *string `json:"topic,omitempty"`

	Partition *int64 `json:"partition,omitempty"`

	Offset *int64 `json:"offset,omitempty"`
}

// UnmarshalGroupResetResultsItem unmarshals an instance of GroupResetResultsItem from the specified map of raw messages.
func UnmarshalGroupResetResultsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupResetResultsItem)
	err = core.UnmarshalPrimitive(m, "topic", &obj.Topic)
	if err != nil {
		err = core.SDKErrorf(err, "", "topic-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "partition", &obj.Partition)
	if err != nil {
		err = core.SDKErrorf(err, "", "partition-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		err = core.SDKErrorf(err, "", "offset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBrokersOptions : The ListBrokers options.
type ListBrokersOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListBrokersOptions : Instantiate ListBrokersOptions
func (*AdminrestV1) NewListBrokersOptions() *ListBrokersOptions {
	return &ListBrokersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListBrokersOptions) SetHeaders(param map[string]string) *ListBrokersOptions {
	options.Headers = param
	return options
}

// ListConsumerGroupsOptions : The ListConsumerGroups options.
type ListConsumerGroupsOptions struct {
	// A filter to be applied to the consumer group IDs. A simple filter can be specified as a string with asterisk (`*`)
	// wildcards representing 0 or more characters, e.g. `group_id*` will filter all group IDs that begin with the string
	// `group_id` followed by any character sequence. A more complex filter pattern can be used by surrounding a regular
	// expression in forward slash (`/`) delimiters, e.g. `/group_id.* /`.
	GroupFilter *string `json:"group_filter,omitempty"`

	// The number of consumer groups to be returned.
	PerPage *int64 `json:"per_page,omitempty"`

	// The page number to be returned. The number 1 represents the first page. The default value is 1.
	Page *int64 `json:"page,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListConsumerGroupsOptions : Instantiate ListConsumerGroupsOptions
func (*AdminrestV1) NewListConsumerGroupsOptions() *ListConsumerGroupsOptions {
	return &ListConsumerGroupsOptions{}
}

// SetGroupFilter : Allow user to set GroupFilter
func (_options *ListConsumerGroupsOptions) SetGroupFilter(groupFilter string) *ListConsumerGroupsOptions {
	_options.GroupFilter = core.StringPtr(groupFilter)
	return _options
}

// SetPerPage : Allow user to set PerPage
func (_options *ListConsumerGroupsOptions) SetPerPage(perPage int64) *ListConsumerGroupsOptions {
	_options.PerPage = core.Int64Ptr(perPage)
	return _options
}

// SetPage : Allow user to set Page
func (_options *ListConsumerGroupsOptions) SetPage(page int64) *ListConsumerGroupsOptions {
	_options.Page = core.Int64Ptr(page)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConsumerGroupsOptions) SetHeaders(param map[string]string) *ListConsumerGroupsOptions {
	options.Headers = param
	return options
}

// ListQuotasOptions : The ListQuotas options.
type ListQuotasOptions struct {

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListQuotasOptions : Instantiate ListQuotasOptions
func (*AdminrestV1) NewListQuotasOptions() *ListQuotasOptions {
	return &ListQuotasOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListQuotasOptions) SetHeaders(param map[string]string) *ListQuotasOptions {
	options.Headers = param
	return options
}

// ListTopicsOptions : The ListTopics options.
type ListTopicsOptions struct {
	// A filter to be applied to the topic names. A simple filter can be specified as a string with asterisk (`*`)
	// wildcards representing 0 or more characters, e.g. `topic-name*` will filter all topic names that begin with the
	// string `topic-name` followed by any character sequence. A more complex filter pattern can be used by surrounding a
	// regular expression in forward slash (`/`) delimiters, e.g. `/topic-name.* /`.
	TopicFilter *string `json:"topic_filter,omitempty"`

	// The number of topic names to be returned.
	PerPage *int64 `json:"per_page,omitempty"`

	// The page number to be returned. The number 1 represents the first page. The default value is 1.
	Page *int64 `json:"page,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListTopicsOptions : Instantiate ListTopicsOptions
func (*AdminrestV1) NewListTopicsOptions() *ListTopicsOptions {
	return &ListTopicsOptions{}
}

// SetTopicFilter : Allow user to set TopicFilter
func (_options *ListTopicsOptions) SetTopicFilter(topicFilter string) *ListTopicsOptions {
	_options.TopicFilter = core.StringPtr(topicFilter)
	return _options
}

// SetPerPage : Allow user to set PerPage
func (_options *ListTopicsOptions) SetPerPage(perPage int64) *ListTopicsOptions {
	_options.PerPage = core.Int64Ptr(perPage)
	return _options
}

// SetPage : Allow user to set Page
func (_options *ListTopicsOptions) SetPage(page int64) *ListTopicsOptions {
	_options.Page = core.Int64Ptr(page)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTopicsOptions) SetHeaders(param map[string]string) *ListTopicsOptions {
	options.Headers = param
	return options
}

// MemberAssignmentsItem : The topic partitions assigned for the consumer group member.
type MemberAssignmentsItem struct {
	// The name of the topic.
	Topic *string `json:"topic,omitempty"`

	// The ID of the partition.
	Partition *int64 `json:"partition,omitempty"`
}

// UnmarshalMemberAssignmentsItem unmarshals an instance of MemberAssignmentsItem from the specified map of raw messages.
func UnmarshalMemberAssignmentsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberAssignmentsItem)
	err = core.UnmarshalPrimitive(m, "topic", &obj.Topic)
	if err != nil {
		err = core.SDKErrorf(err, "", "topic-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "partition", &obj.Partition)
	if err != nil {
		err = core.SDKErrorf(err, "", "partition-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordDeleteRequestRecordsToDeleteItem : RecordDeleteRequestRecordsToDeleteItem struct
type RecordDeleteRequestRecordsToDeleteItem struct {
	// The number of partitions.
	Partition *int64 `json:"partition,omitempty"`

	// The offset number before which records to be deleted.
	BeforeOffset *int64 `json:"before_offset,omitempty"`
}

// UnmarshalRecordDeleteRequestRecordsToDeleteItem unmarshals an instance of RecordDeleteRequestRecordsToDeleteItem from the specified map of raw messages.
func UnmarshalRecordDeleteRequestRecordsToDeleteItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordDeleteRequestRecordsToDeleteItem)
	err = core.UnmarshalPrimitive(m, "partition", &obj.Partition)
	if err != nil {
		err = core.SDKErrorf(err, "", "partition-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "before_offset", &obj.BeforeOffset)
	if err != nil {
		err = core.SDKErrorf(err, "", "before_offset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceMirroringTopicSelectionOptions : The ReplaceMirroringTopicSelection options.
type ReplaceMirroringTopicSelectionOptions struct {
	Includes []string `json:"includes,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceMirroringTopicSelectionOptions : Instantiate ReplaceMirroringTopicSelectionOptions
func (*AdminrestV1) NewReplaceMirroringTopicSelectionOptions() *ReplaceMirroringTopicSelectionOptions {
	return &ReplaceMirroringTopicSelectionOptions{}
}

// SetIncludes : Allow user to set Includes
func (_options *ReplaceMirroringTopicSelectionOptions) SetIncludes(includes []string) *ReplaceMirroringTopicSelectionOptions {
	_options.Includes = includes
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceMirroringTopicSelectionOptions) SetHeaders(param map[string]string) *ReplaceMirroringTopicSelectionOptions {
	options.Headers = param
	return options
}

// TopicCreateRequestConfigsItem : TopicCreateRequestConfigsItem struct
type TopicCreateRequestConfigsItem struct {
	// The name of the config property.
	Name *string `json:"name,omitempty"`

	// The value for a config property.
	Value *string `json:"value,omitempty"`
}

// UnmarshalTopicCreateRequestConfigsItem unmarshals an instance of TopicCreateRequestConfigsItem from the specified map of raw messages.
func UnmarshalTopicCreateRequestConfigsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicCreateRequestConfigsItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicDetailReplicaAssignmentsItem : TopicDetailReplicaAssignmentsItem struct
type TopicDetailReplicaAssignmentsItem struct {
	// The ID of the partition.
	ID *int64 `json:"id,omitempty"`

	Brokers *TopicDetailReplicaAssignmentsItemBrokers `json:"brokers,omitempty"`
}

// UnmarshalTopicDetailReplicaAssignmentsItem unmarshals an instance of TopicDetailReplicaAssignmentsItem from the specified map of raw messages.
func UnmarshalTopicDetailReplicaAssignmentsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicDetailReplicaAssignmentsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "brokers", &obj.Brokers, UnmarshalTopicDetailReplicaAssignmentsItemBrokers)
	if err != nil {
		err = core.SDKErrorf(err, "", "brokers-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicDetailReplicaAssignmentsItemBrokers : TopicDetailReplicaAssignmentsItemBrokers struct
type TopicDetailReplicaAssignmentsItemBrokers struct {
	Replicas []int64 `json:"replicas,omitempty"`
}

// UnmarshalTopicDetailReplicaAssignmentsItemBrokers unmarshals an instance of TopicDetailReplicaAssignmentsItemBrokers from the specified map of raw messages.
func UnmarshalTopicDetailReplicaAssignmentsItemBrokers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicDetailReplicaAssignmentsItemBrokers)
	err = core.UnmarshalPrimitive(m, "replicas", &obj.Replicas)
	if err != nil {
		err = core.SDKErrorf(err, "", "replicas-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicUpdateRequestConfigsItem : TopicUpdateRequestConfigsItem struct
type TopicUpdateRequestConfigsItem struct {
	// The name of the config property.
	Name *string `json:"name,omitempty"`

	// The value of a config property.
	Value *string `json:"value,omitempty"`

	// When true, the value of the config property is reset to its default value.
	ResetToDefault *bool `json:"reset_to_default,omitempty"`
}

// UnmarshalTopicUpdateRequestConfigsItem unmarshals an instance of TopicUpdateRequestConfigsItem from the specified map of raw messages.
func UnmarshalTopicUpdateRequestConfigsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicUpdateRequestConfigsItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		err = core.SDKErrorf(err, "", "value-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "reset_to_default", &obj.ResetToDefault)
	if err != nil {
		err = core.SDKErrorf(err, "", "reset_to_default-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateConsumerGroupOptions : The UpdateConsumerGroup options.
type UpdateConsumerGroupOptions struct {
	// The group ID for the consumer group to be updated.
	GroupID *string `json:"group_id" validate:"required,ne="`

	// The name of the topic to be reset.  If missing or blank, the operation applies to all topics read by the consumer
	// group.
	Topic *string `json:"topic,omitempty"`

	// Mode of shift operation.  Valid values are 'earliest', 'latest', 'datetime'.
	Mode *string `json:"mode,omitempty"`

	// Value for resetting offsets, based on 'mode=datetime', omit for 'earliest' and 'latest'.
	Value *string `json:"value,omitempty"`

	// Whether to execute the operation of resetting the offsets.
	Execute *bool `json:"execute,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateConsumerGroupOptions : Instantiate UpdateConsumerGroupOptions
func (*AdminrestV1) NewUpdateConsumerGroupOptions(groupID string) *UpdateConsumerGroupOptions {
	return &UpdateConsumerGroupOptions{
		GroupID: core.StringPtr(groupID),
	}
}

// SetGroupID : Allow user to set GroupID
func (_options *UpdateConsumerGroupOptions) SetGroupID(groupID string) *UpdateConsumerGroupOptions {
	_options.GroupID = core.StringPtr(groupID)
	return _options
}

// SetTopic : Allow user to set Topic
func (_options *UpdateConsumerGroupOptions) SetTopic(topic string) *UpdateConsumerGroupOptions {
	_options.Topic = core.StringPtr(topic)
	return _options
}

// SetMode : Allow user to set Mode
func (_options *UpdateConsumerGroupOptions) SetMode(mode string) *UpdateConsumerGroupOptions {
	_options.Mode = core.StringPtr(mode)
	return _options
}

// SetValue : Allow user to set Value
func (_options *UpdateConsumerGroupOptions) SetValue(value string) *UpdateConsumerGroupOptions {
	_options.Value = core.StringPtr(value)
	return _options
}

// SetExecute : Allow user to set Execute
func (_options *UpdateConsumerGroupOptions) SetExecute(execute bool) *UpdateConsumerGroupOptions {
	_options.Execute = core.BoolPtr(execute)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConsumerGroupOptions) SetHeaders(param map[string]string) *UpdateConsumerGroupOptions {
	options.Headers = param
	return options
}

// UpdateQuotaOptions : The UpdateQuota options.
type UpdateQuotaOptions struct {
	// The entity name of the quotas can be `default` or an IAM Service ID that starts with an `iam-ServiceId` prefix.
	EntityName *string `json:"entity_name" validate:"required,ne="`

	// The producer byte rate quota value.
	ProducerByteRate *int64 `json:"producer_byte_rate,omitempty"`

	// The consumer byte rate quota value.
	ConsumerByteRate *int64 `json:"consumer_byte_rate,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateQuotaOptions : Instantiate UpdateQuotaOptions
func (*AdminrestV1) NewUpdateQuotaOptions(entityName string) *UpdateQuotaOptions {
	return &UpdateQuotaOptions{
		EntityName: core.StringPtr(entityName),
	}
}

// SetEntityName : Allow user to set EntityName
func (_options *UpdateQuotaOptions) SetEntityName(entityName string) *UpdateQuotaOptions {
	_options.EntityName = core.StringPtr(entityName)
	return _options
}

// SetProducerByteRate : Allow user to set ProducerByteRate
func (_options *UpdateQuotaOptions) SetProducerByteRate(producerByteRate int64) *UpdateQuotaOptions {
	_options.ProducerByteRate = core.Int64Ptr(producerByteRate)
	return _options
}

// SetConsumerByteRate : Allow user to set ConsumerByteRate
func (_options *UpdateQuotaOptions) SetConsumerByteRate(consumerByteRate int64) *UpdateQuotaOptions {
	_options.ConsumerByteRate = core.Int64Ptr(consumerByteRate)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateQuotaOptions) SetHeaders(param map[string]string) *UpdateQuotaOptions {
	options.Headers = param
	return options
}

// UpdateTopicOptions : The UpdateTopic options.
type UpdateTopicOptions struct {
	// The topic name for the topic to be updated.
	TopicName *string `json:"topic_name" validate:"required,ne="`

	// The new partition number to be increased to.
	NewTotalPartitionCount *int64 `json:"new_total_partition_count,omitempty"`

	// The config properties to be updated for the topic. Valid config names are 'cleanup.policy', 'retention.ms',
	// 'retention.bytes', 'segment.bytes', 'segment.ms', 'segment.index.bytes'.
	Configs []TopicUpdateRequestConfigsItem `json:"configs,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewUpdateTopicOptions : Instantiate UpdateTopicOptions
func (*AdminrestV1) NewUpdateTopicOptions(topicName string) *UpdateTopicOptions {
	return &UpdateTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (_options *UpdateTopicOptions) SetTopicName(topicName string) *UpdateTopicOptions {
	_options.TopicName = core.StringPtr(topicName)
	return _options
}

// SetNewTotalPartitionCount : Allow user to set NewTotalPartitionCount
func (_options *UpdateTopicOptions) SetNewTotalPartitionCount(newTotalPartitionCount int64) *UpdateTopicOptions {
	_options.NewTotalPartitionCount = core.Int64Ptr(newTotalPartitionCount)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *UpdateTopicOptions) SetConfigs(configs []TopicUpdateRequestConfigsItem) *UpdateTopicOptions {
	_options.Configs = configs
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTopicOptions) SetHeaders(param map[string]string) *UpdateTopicOptions {
	options.Headers = param
	return options
}

// BrokerDetail : BrokerDetail struct
type BrokerDetail struct {
	// The ID of the broker configured in the 'broker.id' broker config property.
	ID *int64 `json:"id,omitempty"`

	// The hostname that the broker is listening on and which is configured in the 'advertised.listeners' broker config
	// property.
	Host *string `json:"host,omitempty"`

	// The port that the broker is listening on and which is configured in the 'advertised.listeners' broker config
	// property.
	Port *int64 `json:"port,omitempty"`

	// The rack of the broker used in rack aware replication assignment for fault tolerance. It is configure in the
	// 'broker.rack' broker config property.
	Rack *string `json:"rack,omitempty"`

	Configs []BrokerDetailConfigsItem `json:"configs,omitempty"`
}

// UnmarshalBrokerDetail unmarshals an instance of BrokerDetail from the specified map of raw messages.
func UnmarshalBrokerDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrokerDetail)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "rack", &obj.Rack)
	if err != nil {
		err = core.SDKErrorf(err, "", "rack-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalBrokerDetailConfigsItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "configs-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BrokerSummary : BrokerSummary struct
type BrokerSummary struct {
	// The ID of the broker configured in the 'broker.id' broker config property.
	ID *int64 `json:"id,omitempty"`

	// The hostname that the broker is listening on and which is configured in the 'advertised.listeners' broker config
	// property.
	Host *string `json:"host,omitempty"`

	// The port that the broker is listening on and which is configured in the 'advertised.listeners' broker config
	// property.
	Port *int64 `json:"port,omitempty"`

	// The rack of the broker used in rack aware replication assignment for fault tolerance. It is configure in the
	// 'broker.rack' broker config property.
	Rack *string `json:"rack,omitempty"`
}

// UnmarshalBrokerSummary unmarshals an instance of BrokerSummary from the specified map of raw messages.
func UnmarshalBrokerSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrokerSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		err = core.SDKErrorf(err, "", "port-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "rack", &obj.Rack)
	if err != nil {
		err = core.SDKErrorf(err, "", "rack-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Cluster : Cluster struct
type Cluster struct {
	// The ID of the cluster.
	ID *string `json:"id,omitempty"`

	Controller *BrokerSummary `json:"controller,omitempty"`

	// List of brokers in the cluster.
	Brokers []BrokerSummary `json:"brokers,omitempty"`
}

// UnmarshalCluster unmarshals an instance of Cluster from the specified map of raw messages.
func UnmarshalCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Cluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "controller", &obj.Controller, UnmarshalBrokerSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "controller-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "brokers", &obj.Brokers, UnmarshalBrokerSummary)
	if err != nil {
		err = core.SDKErrorf(err, "", "brokers-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntityQuotaDetail : EntityQuotaDetail struct
type EntityQuotaDetail struct {
	// The name of the entity.
	EntityName *string `json:"entity_name" validate:"required"`

	// The producer byte rate quota value.
	ProducerByteRate *int64 `json:"producer_byte_rate,omitempty"`

	// The consumer byte rate quota value.
	ConsumerByteRate *int64 `json:"consumer_byte_rate,omitempty"`
}

// UnmarshalEntityQuotaDetail unmarshals an instance of EntityQuotaDetail from the specified map of raw messages.
func UnmarshalEntityQuotaDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntityQuotaDetail)
	err = core.UnmarshalPrimitive(m, "entity_name", &obj.EntityName)
	if err != nil {
		err = core.SDKErrorf(err, "", "entity_name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "producer_byte_rate", &obj.ProducerByteRate)
	if err != nil {
		err = core.SDKErrorf(err, "", "producer_byte_rate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "consumer_byte_rate", &obj.ConsumerByteRate)
	if err != nil {
		err = core.SDKErrorf(err, "", "consumer_byte_rate-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupDetail : GroupDetail struct
type GroupDetail struct {
	// The ID of the consumer group.
	GroupID *string `json:"group_id,omitempty"`

	// THe state of the consumer group.
	State *string `json:"state,omitempty"`

	// Members in the consumer group.
	Members []Member `json:"members,omitempty"`

	// The offsets of the consumer group.
	Offsets []TopicPartitionOffset `json:"offsets,omitempty"`
}

// UnmarshalGroupDetail unmarshals an instance of GroupDetail from the specified map of raw messages.
func UnmarshalGroupDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupDetail)
	err = core.UnmarshalPrimitive(m, "group_id", &obj.GroupID)
	if err != nil {
		err = core.SDKErrorf(err, "", "group_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalMember)
	if err != nil {
		err = core.SDKErrorf(err, "", "members-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "offsets", &obj.Offsets, UnmarshalTopicPartitionOffset)
	if err != nil {
		err = core.SDKErrorf(err, "", "offsets-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceStatus : Information about the status of the instance.
type InstanceStatus struct {
	// The status of the instance: * `available` - the instance is functioning as expected * `degraded` - the instance is
	// in a degraded state, some operations may not complete successfully * `offline` - the instance is offline, all
	// operations attempted against the instance will fail * `unknown` - the state of the instance is not known at this
	// time.
	Status *string `json:"status,omitempty"`
}

// Constants associated with the InstanceStatus.Status property.
// The status of the instance: * `available` - the instance is functioning as expected * `degraded` - the instance is in
// a degraded state, some operations may not complete successfully * `offline` - the instance is offline, all operations
// attempted against the instance will fail * `unknown` - the state of the instance is not known at this time.
const (
	InstanceStatusStatusAvailableConst = "available"
	InstanceStatusStatusDegradedConst  = "degraded"
	InstanceStatusStatusOfflineConst   = "offline"
	InstanceStatusStatusUnknownConst   = "unknown"
)

// UnmarshalInstanceStatus unmarshals an instance of InstanceStatus from the specified map of raw messages.
func UnmarshalInstanceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceStatus)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Member : Member struct
type Member struct {
	// The consumer ID of the consumer group member.
	ConsumerID *string `json:"consumer_id,omitempty"`

	// The client ID of the consumer group member.
	ClientID *string `json:"client_id,omitempty"`

	// The hostname of the machine where the consumer group member is running.
	Host *string `json:"host,omitempty"`

	// The assignments of the group member.
	Assignments []MemberAssignmentsItem `json:"assignments,omitempty"`
}

// UnmarshalMember unmarshals an instance of Member from the specified map of raw messages.
func UnmarshalMember(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Member)
	err = core.UnmarshalPrimitive(m, "consumer_id", &obj.ConsumerID)
	if err != nil {
		err = core.SDKErrorf(err, "", "consumer_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "client_id", &obj.ClientID)
	if err != nil {
		err = core.SDKErrorf(err, "", "client_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		err = core.SDKErrorf(err, "", "host-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "assignments", &obj.Assignments, UnmarshalMemberAssignmentsItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "assignments-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MirroringActiveTopics : Topics that are being actively mirrored.
type MirroringActiveTopics struct {
	ActiveTopics []string `json:"active_topics,omitempty"`
}

// UnmarshalMirroringActiveTopics unmarshals an instance of MirroringActiveTopics from the specified map of raw messages.
func UnmarshalMirroringActiveTopics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MirroringActiveTopics)
	err = core.UnmarshalPrimitive(m, "active_topics", &obj.ActiveTopics)
	if err != nil {
		err = core.SDKErrorf(err, "", "active_topics-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MirroringTopicSelection : Mirroring topic selection payload.
type MirroringTopicSelection struct {
	Includes []string `json:"includes,omitempty"`
}

// UnmarshalMirroringTopicSelection unmarshals an instance of MirroringTopicSelection from the specified map of raw messages.
func UnmarshalMirroringTopicSelection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MirroringTopicSelection)
	err = core.UnmarshalPrimitive(m, "includes", &obj.Includes)
	if err != nil {
		err = core.SDKErrorf(err, "", "includes-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QuotaDetail : QuotaDetail struct
type QuotaDetail struct {
	// The producer byte rate quota value.
	ProducerByteRate *int64 `json:"producer_byte_rate,omitempty"`

	// The consumer byte rate quota value.
	ConsumerByteRate *int64 `json:"consumer_byte_rate,omitempty"`
}

// UnmarshalQuotaDetail unmarshals an instance of QuotaDetail from the specified map of raw messages.
func UnmarshalQuotaDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QuotaDetail)
	err = core.UnmarshalPrimitive(m, "producer_byte_rate", &obj.ProducerByteRate)
	if err != nil {
		err = core.SDKErrorf(err, "", "producer_byte_rate-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "consumer_byte_rate", &obj.ConsumerByteRate)
	if err != nil {
		err = core.SDKErrorf(err, "", "consumer_byte_rate-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QuotaList : A list of 'quota_detail' is returned.
type QuotaList struct {
	Data []EntityQuotaDetail `json:"data,omitempty"`
}

// UnmarshalQuotaList unmarshals an instance of QuotaList from the specified map of raw messages.
func UnmarshalQuotaList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(QuotaList)
	err = core.UnmarshalModel(m, "data", &obj.Data, UnmarshalEntityQuotaDetail)
	if err != nil {
		err = core.SDKErrorf(err, "", "data-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicConfigs : TopicConfigs struct
type TopicConfigs struct {
	// The value of config property 'retention.bytes'.
	RetentionBytes *string `json:"retention.bytes,omitempty"`

	// The value of config property 'segment.bytes'.
	SegmentBytes *string `json:"segment.bytes,omitempty"`

	// The value of config property 'segment.index.bytes'.
	SegmentIndexBytes *string `json:"segment.index.bytes,omitempty"`

	// The value of config property 'segment.ms'.
	SegmentMs *string `json:"segment.ms,omitempty"`
}

// UnmarshalTopicConfigs unmarshals an instance of TopicConfigs from the specified map of raw messages.
func UnmarshalTopicConfigs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicConfigs)
	err = core.UnmarshalPrimitive(m, "retention.bytes", &obj.RetentionBytes)
	if err != nil {
		err = core.SDKErrorf(err, "", "retention.bytes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.bytes", &obj.SegmentBytes)
	if err != nil {
		err = core.SDKErrorf(err, "", "segment.bytes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.index.bytes", &obj.SegmentIndexBytes)
	if err != nil {
		err = core.SDKErrorf(err, "", "segment.index.bytes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.ms", &obj.SegmentMs)
	if err != nil {
		err = core.SDKErrorf(err, "", "segment.ms-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicDetail : TopicDetail struct
type TopicDetail struct {
	// The name of the topic.
	Name *string `json:"name,omitempty"`

	// The number of partitions.
	Partitions *int64 `json:"partitions,omitempty"`

	// The number of replication factor.
	ReplicationFactor *int64 `json:"replicationFactor,omitempty"`

	// The value of config property 'retention.ms'.
	RetentionMs *int64 `json:"retentionMs,omitempty"`

	// The value of config property 'cleanup.policy'.
	CleanupPolicy *string `json:"cleanupPolicy,omitempty"`

	Configs *TopicConfigs `json:"configs,omitempty"`

	// The replia assignment of the topic.
	ReplicaAssignments []TopicDetailReplicaAssignmentsItem `json:"replicaAssignments,omitempty"`
}

// UnmarshalTopicDetail unmarshals an instance of TopicDetail from the specified map of raw messages.
func UnmarshalTopicDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicDetail)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "partitions", &obj.Partitions)
	if err != nil {
		err = core.SDKErrorf(err, "", "partitions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "replicationFactor", &obj.ReplicationFactor)
	if err != nil {
		err = core.SDKErrorf(err, "", "replicationFactor-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "retentionMs", &obj.RetentionMs)
	if err != nil {
		err = core.SDKErrorf(err, "", "retentionMs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "cleanupPolicy", &obj.CleanupPolicy)
	if err != nil {
		err = core.SDKErrorf(err, "", "cleanupPolicy-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalTopicConfigs)
	if err != nil {
		err = core.SDKErrorf(err, "", "configs-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "replicaAssignments", &obj.ReplicaAssignments, UnmarshalTopicDetailReplicaAssignmentsItem)
	if err != nil {
		err = core.SDKErrorf(err, "", "replicaAssignments-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicPartitionOffset : The offsets of a topic partition.
type TopicPartitionOffset struct {
	// The name of the topic.
	Topic *string `json:"topic,omitempty"`

	// The ID of the partition.
	Partition *int64 `json:"partition,omitempty"`

	// Current offset of the partition.
	CurrentOffset *int64 `json:"current_offset,omitempty"`

	// End offset of the partition.
	EndOffset *int64 `json:"end_offset,omitempty"`
}

// UnmarshalTopicPartitionOffset unmarshals an instance of TopicPartitionOffset from the specified map of raw messages.
func UnmarshalTopicPartitionOffset(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicPartitionOffset)
	err = core.UnmarshalPrimitive(m, "topic", &obj.Topic)
	if err != nil {
		err = core.SDKErrorf(err, "", "topic-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "partition", &obj.Partition)
	if err != nil {
		err = core.SDKErrorf(err, "", "partition-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "current_offset", &obj.CurrentOffset)
	if err != nil {
		err = core.SDKErrorf(err, "", "current_offset-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_offset", &obj.EndOffset)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_offset-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
