/**
 * (C) Copyright IBM Corp. 2021.
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

package adminrestv1

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AdminrestV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(adminrestService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(adminrestService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
				URL: "https://adminrestv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(adminrestService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMINREST_URL":       "https://adminrestv1/api",
				"ADMINREST_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminrestService, serviceErr := NewAdminrestV1UsingExternalConfig(&AdminrestV1Options{})
				Expect(adminrestService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := adminrestService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminrestService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminrestService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminrestService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminrestService, serviceErr := NewAdminrestV1UsingExternalConfig(&AdminrestV1Options{
					URL: "https://testService/api",
				})
				Expect(adminrestService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminrestService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminrestService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminrestService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminrestService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				adminrestService, serviceErr := NewAdminrestV1UsingExternalConfig(&AdminrestV1Options{})
				err := adminrestService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := adminrestService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != adminrestService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(adminrestService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(adminrestService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMINREST_URL":       "https://adminrestv1/api",
				"ADMINREST_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminrestService, serviceErr := NewAdminrestV1UsingExternalConfig(&AdminrestV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(adminrestService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ADMINREST_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			adminrestService, serviceErr := NewAdminrestV1UsingExternalConfig(&AdminrestV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(adminrestService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTopic(createTopicOptions *CreateTopicOptions)`, func() {
		createTopicPath := "/admin/topics"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTopicPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateTopic successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := adminrestService.CreateTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ConfigCreate model
				configCreateModel := new(ConfigCreate)
				configCreateModel.Name = core.StringPtr("testString")
				configCreateModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(CreateTopicOptions)
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Partitions = core.Int64Ptr(int64(26))
				createTopicOptionsModel.PartitionCount = core.Int64Ptr(int64(1))
				createTopicOptionsModel.Configs = []ConfigCreate{*configCreateModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = adminrestService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateTopic with error: Operation request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ConfigCreate model
				configCreateModel := new(ConfigCreate)
				configCreateModel.Name = core.StringPtr("testString")
				configCreateModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := new(CreateTopicOptions)
				createTopicOptionsModel.Name = core.StringPtr("testString")
				createTopicOptionsModel.Partitions = core.Int64Ptr(int64(26))
				createTopicOptionsModel.PartitionCount = core.Int64Ptr(int64(1))
				createTopicOptionsModel.Configs = []ConfigCreate{*configCreateModel}
				createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := adminrestService.CreateTopic(createTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTopics(listTopicsOptions *ListTopicsOptions) - Operation response error`, func() {
		listTopicsPath := "/admin/topics"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["topic_filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTopics with error: Operation response processing error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(ListTopicsOptions)
				listTopicsOptionsModel.TopicFilter = core.StringPtr("testString")
				listTopicsOptionsModel.PerPage = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Page = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminrestService.EnableRetries(0, 0)
				result, response, operationErr = adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTopics(listTopicsOptions *ListTopicsOptions)`, func() {
		listTopicsPath := "/admin/topics"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["topic_filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"name": "Name", "partitions": 10, "replicationFactor": 17, "retentionMs": 11, "cleanupPolicy": "CleanupPolicy", "configs": {"cleanup.policy": "CleanupPolicy", "min.insync.replicas": "MinInsyncReplicas", "retention.bytes": "RetentionBytes", "retention.ms": "RetentionMs", "segment.bytes": "SegmentBytes", "segment.index.bytes": "SegmentIndexBytes", "segment.ms": "SegmentMs"}, "replicaAssignments": [{"id": 2, "brokers": {"replicas": [8]}}]}]`)
				}))
			})
			It(`Invoke ListTopics successfully with retries`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				adminrestService.EnableRetries(0, 0)

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(ListTopicsOptions)
				listTopicsOptionsModel.TopicFilter = core.StringPtr("testString")
				listTopicsOptionsModel.PerPage = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Page = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminrestService.ListTopicsWithContext(ctx, listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminrestService.DisableRetries()
				result, response, operationErr := adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminrestService.ListTopicsWithContext(ctx, listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["topic_filter"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[{"name": "Name", "partitions": 10, "replicationFactor": 17, "retentionMs": 11, "cleanupPolicy": "CleanupPolicy", "configs": {"cleanup.policy": "CleanupPolicy", "min.insync.replicas": "MinInsyncReplicas", "retention.bytes": "RetentionBytes", "retention.ms": "RetentionMs", "segment.bytes": "SegmentBytes", "segment.index.bytes": "SegmentIndexBytes", "segment.ms": "SegmentMs"}, "replicaAssignments": [{"id": 2, "brokers": {"replicas": [8]}}]}]`)
				}))
			})
			It(`Invoke ListTopics successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminrestService.ListTopics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(ListTopicsOptions)
				listTopicsOptionsModel.TopicFilter = core.StringPtr("testString")
				listTopicsOptionsModel.PerPage = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Page = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTopics with error: Operation request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(ListTopicsOptions)
				listTopicsOptionsModel.TopicFilter = core.StringPtr("testString")
				listTopicsOptionsModel.PerPage = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Page = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTopics successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := new(ListTopicsOptions)
				listTopicsOptionsModel.TopicFilter = core.StringPtr("testString")
				listTopicsOptionsModel.PerPage = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Page = core.Int64Ptr(int64(38))
				listTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminrestService.ListTopics(listTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTopic(getTopicOptions *GetTopicOptions) - Operation response error`, func() {
		getTopicPath := "/admin/topics/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTopic with error: Operation response processing error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(GetTopicOptions)
				getTopicOptionsModel.TopicName = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminrestService.EnableRetries(0, 0)
				result, response, operationErr = adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTopic(getTopicOptions *GetTopicOptions)`, func() {
		getTopicPath := "/admin/topics/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "partitions": 10, "replicationFactor": 17, "retentionMs": 11, "cleanupPolicy": "CleanupPolicy", "configs": {"cleanup.policy": "CleanupPolicy", "min.insync.replicas": "MinInsyncReplicas", "retention.bytes": "RetentionBytes", "retention.ms": "RetentionMs", "segment.bytes": "SegmentBytes", "segment.index.bytes": "SegmentIndexBytes", "segment.ms": "SegmentMs"}, "replicaAssignments": [{"id": 2, "brokers": {"replicas": [8]}}]}`)
				}))
			})
			It(`Invoke GetTopic successfully with retries`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				adminrestService.EnableRetries(0, 0)

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(GetTopicOptions)
				getTopicOptionsModel.TopicName = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminrestService.GetTopicWithContext(ctx, getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminrestService.DisableRetries()
				result, response, operationErr := adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminrestService.GetTopicWithContext(ctx, getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTopicPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "partitions": 10, "replicationFactor": 17, "retentionMs": 11, "cleanupPolicy": "CleanupPolicy", "configs": {"cleanup.policy": "CleanupPolicy", "min.insync.replicas": "MinInsyncReplicas", "retention.bytes": "RetentionBytes", "retention.ms": "RetentionMs", "segment.bytes": "SegmentBytes", "segment.index.bytes": "SegmentIndexBytes", "segment.ms": "SegmentMs"}, "replicaAssignments": [{"id": 2, "brokers": {"replicas": [8]}}]}`)
				}))
			})
			It(`Invoke GetTopic successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminrestService.GetTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(GetTopicOptions)
				getTopicOptionsModel.TopicName = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTopic with error: Operation validation and request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(GetTopicOptions)
				getTopicOptionsModel.TopicName = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTopicOptions model with no property values
				getTopicOptionsModelNew := new(GetTopicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = adminrestService.GetTopic(getTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTopic successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetTopicOptions model
				getTopicOptionsModel := new(GetTopicOptions)
				getTopicOptionsModel.TopicName = core.StringPtr("testString")
				getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminrestService.GetTopic(getTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTopic(deleteTopicOptions *DeleteTopicOptions)`, func() {
		deleteTopicPath := "/admin/topics/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTopicPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteTopic successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := adminrestService.DeleteTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTopicOptions model
				deleteTopicOptionsModel := new(DeleteTopicOptions)
				deleteTopicOptionsModel.TopicName = core.StringPtr("testString")
				deleteTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = adminrestService.DeleteTopic(deleteTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTopic with error: Operation validation and request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the DeleteTopicOptions model
				deleteTopicOptionsModel := new(DeleteTopicOptions)
				deleteTopicOptionsModel.TopicName = core.StringPtr("testString")
				deleteTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := adminrestService.DeleteTopic(deleteTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTopicOptions model with no property values
				deleteTopicOptionsModelNew := new(DeleteTopicOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = adminrestService.DeleteTopic(deleteTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTopic(updateTopicOptions *UpdateTopicOptions)`, func() {
		updateTopicPath := "/admin/topics/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTopicPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateTopic successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := adminrestService.UpdateTopic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ConfigUpdate model
				configUpdateModel := new(ConfigUpdate)
				configUpdateModel.Name = core.StringPtr("testString")
				configUpdateModel.Value = core.StringPtr("testString")
				configUpdateModel.ResetToDefault = core.BoolPtr(true)

				// Construct an instance of the UpdateTopicOptions model
				updateTopicOptionsModel := new(UpdateTopicOptions)
				updateTopicOptionsModel.TopicName = core.StringPtr("testString")
				updateTopicOptionsModel.NewTotalPartitionCount = core.Int64Ptr(int64(38))
				updateTopicOptionsModel.Configs = []ConfigUpdate{*configUpdateModel}
				updateTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = adminrestService.UpdateTopic(updateTopicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateTopic with error: Operation validation and request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ConfigUpdate model
				configUpdateModel := new(ConfigUpdate)
				configUpdateModel.Name = core.StringPtr("testString")
				configUpdateModel.Value = core.StringPtr("testString")
				configUpdateModel.ResetToDefault = core.BoolPtr(true)

				// Construct an instance of the UpdateTopicOptions model
				updateTopicOptionsModel := new(UpdateTopicOptions)
				updateTopicOptionsModel.TopicName = core.StringPtr("testString")
				updateTopicOptionsModel.NewTotalPartitionCount = core.Int64Ptr(int64(38))
				updateTopicOptionsModel.Configs = []ConfigUpdate{*configUpdateModel}
				updateTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := adminrestService.UpdateTopic(updateTopicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateTopicOptions model with no property values
				updateTopicOptionsModelNew := new(UpdateTopicOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = adminrestService.UpdateTopic(updateTopicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMirroringTopicSelection(getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions) - Operation response error`, func() {
		getMirroringTopicSelectionPath := "/admin/mirroring/topic-selection"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMirroringTopicSelection with error: Operation response processing error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := new(GetMirroringTopicSelectionOptions)
				getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminrestService.EnableRetries(0, 0)
				result, response, operationErr = adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMirroringTopicSelection(getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions)`, func() {
		getMirroringTopicSelectionPath := "/admin/mirroring/topic-selection"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"includes": ["Includes"]}`)
				}))
			})
			It(`Invoke GetMirroringTopicSelection successfully with retries`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				adminrestService.EnableRetries(0, 0)

				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := new(GetMirroringTopicSelectionOptions)
				getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminrestService.GetMirroringTopicSelectionWithContext(ctx, getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminrestService.DisableRetries()
				result, response, operationErr := adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminrestService.GetMirroringTopicSelectionWithContext(ctx, getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"includes": ["Includes"]}`)
				}))
			})
			It(`Invoke GetMirroringTopicSelection successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminrestService.GetMirroringTopicSelection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := new(GetMirroringTopicSelectionOptions)
				getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMirroringTopicSelection with error: Operation request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := new(GetMirroringTopicSelectionOptions)
				getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetMirroringTopicSelection successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := new(GetMirroringTopicSelectionOptions)
				getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminrestService.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions) - Operation response error`, func() {
		replaceMirroringTopicSelectionPath := "/admin/mirroring/topic-selection"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceMirroringTopicSelection with error: Operation response processing error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := new(ReplaceMirroringTopicSelectionOptions)
				replaceMirroringTopicSelectionOptionsModel.Includes = []string{"testString"}
				replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminrestService.EnableRetries(0, 0)
				result, response, operationErr = adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions)`, func() {
		replaceMirroringTopicSelectionPath := "/admin/mirroring/topic-selection"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"includes": ["Includes"]}`)
				}))
			})
			It(`Invoke ReplaceMirroringTopicSelection successfully with retries`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				adminrestService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := new(ReplaceMirroringTopicSelectionOptions)
				replaceMirroringTopicSelectionOptionsModel.Includes = []string{"testString"}
				replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminrestService.ReplaceMirroringTopicSelectionWithContext(ctx, replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminrestService.DisableRetries()
				result, response, operationErr := adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminrestService.ReplaceMirroringTopicSelectionWithContext(ctx, replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceMirroringTopicSelectionPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"includes": ["Includes"]}`)
				}))
			})
			It(`Invoke ReplaceMirroringTopicSelection successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminrestService.ReplaceMirroringTopicSelection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := new(ReplaceMirroringTopicSelectionOptions)
				replaceMirroringTopicSelectionOptionsModel.Includes = []string{"testString"}
				replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceMirroringTopicSelection with error: Operation request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := new(ReplaceMirroringTopicSelectionOptions)
				replaceMirroringTopicSelectionOptionsModel.Includes = []string{"testString"}
				replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceMirroringTopicSelection successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := new(ReplaceMirroringTopicSelectionOptions)
				replaceMirroringTopicSelectionOptionsModel.Includes = []string{"testString"}
				replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminrestService.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMirroringActiveTopics(getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions) - Operation response error`, func() {
		getMirroringActiveTopicsPath := "/admin/mirroring/active-topics"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringActiveTopicsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetMirroringActiveTopics with error: Operation response processing error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := new(GetMirroringActiveTopicsOptions)
				getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				adminrestService.EnableRetries(0, 0)
				result, response, operationErr = adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetMirroringActiveTopics(getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions)`, func() {
		getMirroringActiveTopicsPath := "/admin/mirroring/active-topics"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringActiveTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_topics": ["ActiveTopics"]}`)
				}))
			})
			It(`Invoke GetMirroringActiveTopics successfully with retries`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())
				adminrestService.EnableRetries(0, 0)

				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := new(GetMirroringActiveTopicsOptions)
				getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := adminrestService.GetMirroringActiveTopicsWithContext(ctx, getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				adminrestService.DisableRetries()
				result, response, operationErr := adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = adminrestService.GetMirroringActiveTopicsWithContext(ctx, getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getMirroringActiveTopicsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"active_topics": ["ActiveTopics"]}`)
				}))
			})
			It(`Invoke GetMirroringActiveTopics successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := adminrestService.GetMirroringActiveTopics(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := new(GetMirroringActiveTopicsOptions)
				getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetMirroringActiveTopics with error: Operation request error`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := new(GetMirroringActiveTopicsOptions)
				getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := adminrestService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetMirroringActiveTopics successfully`, func() {
				adminrestService, serviceErr := NewAdminrestV1(&AdminrestV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(adminrestService).ToNot(BeNil())

				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := new(GetMirroringActiveTopicsOptions)
				getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := adminrestService.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			adminrestService, _ := NewAdminrestV1(&AdminrestV1Options{
				URL:           "http://adminrestv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateTopicOptions successfully`, func() {
				// Construct an instance of the ConfigCreate model
				configCreateModel := new(ConfigCreate)
				Expect(configCreateModel).ToNot(BeNil())
				configCreateModel.Name = core.StringPtr("testString")
				configCreateModel.Value = core.StringPtr("testString")
				Expect(configCreateModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configCreateModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateTopicOptions model
				createTopicOptionsModel := adminrestService.NewCreateTopicOptions()
				createTopicOptionsModel.SetName("testString")
				createTopicOptionsModel.SetPartitions(int64(26))
				createTopicOptionsModel.SetPartitionCount(int64(1))
				createTopicOptionsModel.SetConfigs([]ConfigCreate{*configCreateModel})
				createTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTopicOptionsModel).ToNot(BeNil())
				Expect(createTopicOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createTopicOptionsModel.Partitions).To(Equal(core.Int64Ptr(int64(26))))
				Expect(createTopicOptionsModel.PartitionCount).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createTopicOptionsModel.Configs).To(Equal([]ConfigCreate{*configCreateModel}))
				Expect(createTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTopicOptions successfully`, func() {
				// Construct an instance of the DeleteTopicOptions model
				topicName := "testString"
				deleteTopicOptionsModel := adminrestService.NewDeleteTopicOptions(topicName)
				deleteTopicOptionsModel.SetTopicName("testString")
				deleteTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTopicOptionsModel).ToNot(BeNil())
				Expect(deleteTopicOptionsModel.TopicName).To(Equal(core.StringPtr("testString")))
				Expect(deleteTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMirroringActiveTopicsOptions successfully`, func() {
				// Construct an instance of the GetMirroringActiveTopicsOptions model
				getMirroringActiveTopicsOptionsModel := adminrestService.NewGetMirroringActiveTopicsOptions()
				getMirroringActiveTopicsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMirroringActiveTopicsOptionsModel).ToNot(BeNil())
				Expect(getMirroringActiveTopicsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetMirroringTopicSelectionOptions successfully`, func() {
				// Construct an instance of the GetMirroringTopicSelectionOptions model
				getMirroringTopicSelectionOptionsModel := adminrestService.NewGetMirroringTopicSelectionOptions()
				getMirroringTopicSelectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getMirroringTopicSelectionOptionsModel).ToNot(BeNil())
				Expect(getMirroringTopicSelectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTopicOptions successfully`, func() {
				// Construct an instance of the GetTopicOptions model
				topicName := "testString"
				getTopicOptionsModel := adminrestService.NewGetTopicOptions(topicName)
				getTopicOptionsModel.SetTopicName("testString")
				getTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTopicOptionsModel).ToNot(BeNil())
				Expect(getTopicOptionsModel.TopicName).To(Equal(core.StringPtr("testString")))
				Expect(getTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTopicsOptions successfully`, func() {
				// Construct an instance of the ListTopicsOptions model
				listTopicsOptionsModel := adminrestService.NewListTopicsOptions()
				listTopicsOptionsModel.SetTopicFilter("testString")
				listTopicsOptionsModel.SetPerPage(int64(38))
				listTopicsOptionsModel.SetPage(int64(38))
				listTopicsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTopicsOptionsModel).ToNot(BeNil())
				Expect(listTopicsOptionsModel.TopicFilter).To(Equal(core.StringPtr("testString")))
				Expect(listTopicsOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listTopicsOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listTopicsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceMirroringTopicSelectionOptions successfully`, func() {
				// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
				replaceMirroringTopicSelectionOptionsModel := adminrestService.NewReplaceMirroringTopicSelectionOptions()
				replaceMirroringTopicSelectionOptionsModel.SetIncludes([]string{"testString"})
				replaceMirroringTopicSelectionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceMirroringTopicSelectionOptionsModel).ToNot(BeNil())
				Expect(replaceMirroringTopicSelectionOptionsModel.Includes).To(Equal([]string{"testString"}))
				Expect(replaceMirroringTopicSelectionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTopicOptions successfully`, func() {
				// Construct an instance of the ConfigUpdate model
				configUpdateModel := new(ConfigUpdate)
				Expect(configUpdateModel).ToNot(BeNil())
				configUpdateModel.Name = core.StringPtr("testString")
				configUpdateModel.Value = core.StringPtr("testString")
				configUpdateModel.ResetToDefault = core.BoolPtr(true)
				Expect(configUpdateModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(configUpdateModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(configUpdateModel.ResetToDefault).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateTopicOptions model
				topicName := "testString"
				updateTopicOptionsModel := adminrestService.NewUpdateTopicOptions(topicName)
				updateTopicOptionsModel.SetTopicName("testString")
				updateTopicOptionsModel.SetNewTotalPartitionCount(int64(38))
				updateTopicOptionsModel.SetConfigs([]ConfigUpdate{*configUpdateModel})
				updateTopicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTopicOptionsModel).ToNot(BeNil())
				Expect(updateTopicOptionsModel.TopicName).To(Equal(core.StringPtr("testString")))
				Expect(updateTopicOptionsModel.NewTotalPartitionCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(updateTopicOptionsModel.Configs).To(Equal([]ConfigUpdate{*configUpdateModel}))
				Expect(updateTopicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
