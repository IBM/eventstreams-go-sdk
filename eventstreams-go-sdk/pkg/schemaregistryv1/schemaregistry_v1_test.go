/**
 * (C) Copyright IBM Corp. 2023.
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

package schemaregistryv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`SchemaregistryV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(schemaregistryService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(schemaregistryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
				URL: "https://schemaregistryv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(schemaregistryService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMAREGISTRY_URL":       "https://schemaregistryv1/api",
				"SCHEMAREGISTRY_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1UsingExternalConfig(&schemaregistryv1.SchemaregistryV1Options{})
				Expect(schemaregistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := schemaregistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != schemaregistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(schemaregistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(schemaregistryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1UsingExternalConfig(&schemaregistryv1.SchemaregistryV1Options{
					URL: "https://testService/api",
				})
				Expect(schemaregistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := schemaregistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != schemaregistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(schemaregistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(schemaregistryService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1UsingExternalConfig(&schemaregistryv1.SchemaregistryV1Options{})
				err := schemaregistryService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := schemaregistryService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != schemaregistryService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(schemaregistryService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(schemaregistryService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMAREGISTRY_URL":       "https://schemaregistryv1/api",
				"SCHEMAREGISTRY_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1UsingExternalConfig(&schemaregistryv1.SchemaregistryV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(schemaregistryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SCHEMAREGISTRY_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1UsingExternalConfig(&schemaregistryv1.SchemaregistryV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(schemaregistryService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = schemaregistryv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetGlobalRule(getGlobalRuleOptions *GetGlobalRuleOptions) - Operation response error`, func() {
		getGlobalRulePath := "/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGlobalRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGlobalRule with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetGlobalRuleOptions model
				getGlobalRuleOptionsModel := new(schemaregistryv1.GetGlobalRuleOptions)
				getGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGlobalRule(getGlobalRuleOptions *GetGlobalRuleOptions)`, func() {
		getGlobalRulePath := "/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGlobalRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke GetGlobalRule successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the GetGlobalRuleOptions model
				getGlobalRuleOptionsModel := new(schemaregistryv1.GetGlobalRuleOptions)
				getGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.GetGlobalRuleWithContext(ctx, getGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.GetGlobalRuleWithContext(ctx, getGlobalRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getGlobalRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke GetGlobalRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.GetGlobalRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGlobalRuleOptions model
				getGlobalRuleOptionsModel := new(schemaregistryv1.GetGlobalRuleOptions)
				getGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGlobalRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetGlobalRuleOptions model
				getGlobalRuleOptionsModel := new(schemaregistryv1.GetGlobalRuleOptions)
				getGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGlobalRuleOptions model with no property values
				getGlobalRuleOptionsModelNew := new(schemaregistryv1.GetGlobalRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModelNew)
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
			It(`Invoke GetGlobalRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetGlobalRuleOptions model
				getGlobalRuleOptionsModel := new(schemaregistryv1.GetGlobalRuleOptions)
				getGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.GetGlobalRule(getGlobalRuleOptionsModel)
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
	Describe(`UpdateGlobalRule(updateGlobalRuleOptions *UpdateGlobalRuleOptions) - Operation response error`, func() {
		updateGlobalRulePath := "/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGlobalRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateGlobalRule with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateGlobalRuleOptions model
				updateGlobalRuleOptionsModel := new(schemaregistryv1.UpdateGlobalRuleOptions)
				updateGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateGlobalRule(updateGlobalRuleOptions *UpdateGlobalRuleOptions)`, func() {
		updateGlobalRulePath := "/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateGlobalRulePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke UpdateGlobalRule successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateGlobalRuleOptions model
				updateGlobalRuleOptionsModel := new(schemaregistryv1.UpdateGlobalRuleOptions)
				updateGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.UpdateGlobalRuleWithContext(ctx, updateGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.UpdateGlobalRuleWithContext(ctx, updateGlobalRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateGlobalRulePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke UpdateGlobalRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.UpdateGlobalRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateGlobalRuleOptions model
				updateGlobalRuleOptionsModel := new(schemaregistryv1.UpdateGlobalRuleOptions)
				updateGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateGlobalRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateGlobalRuleOptions model
				updateGlobalRuleOptionsModel := new(schemaregistryv1.UpdateGlobalRuleOptions)
				updateGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateGlobalRuleOptions model with no property values
				updateGlobalRuleOptionsModelNew := new(schemaregistryv1.UpdateGlobalRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModelNew)
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
			It(`Invoke UpdateGlobalRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateGlobalRuleOptions model
				updateGlobalRuleOptionsModel := new(schemaregistryv1.UpdateGlobalRuleOptions)
				updateGlobalRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateGlobalRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateGlobalRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.UpdateGlobalRule(updateGlobalRuleOptionsModel)
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
	Describe(`CreateSchemaRule(createSchemaRuleOptions *CreateSchemaRuleOptions) - Operation response error`, func() {
		createSchemaRulePath := "/artifacts/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSchemaRule with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaRuleOptions model
				createSchemaRuleOptionsModel := new(schemaregistryv1.CreateSchemaRuleOptions)
				createSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				createSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				createSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				createSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSchemaRule(createSchemaRuleOptions *CreateSchemaRuleOptions)`, func() {
		createSchemaRulePath := "/artifacts/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaRulePath))
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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke CreateSchemaRule successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the CreateSchemaRuleOptions model
				createSchemaRuleOptionsModel := new(schemaregistryv1.CreateSchemaRuleOptions)
				createSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				createSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				createSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				createSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.CreateSchemaRuleWithContext(ctx, createSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.CreateSchemaRuleWithContext(ctx, createSchemaRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaRulePath))
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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke CreateSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.CreateSchemaRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSchemaRuleOptions model
				createSchemaRuleOptionsModel := new(schemaregistryv1.CreateSchemaRuleOptions)
				createSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				createSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				createSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				createSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSchemaRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaRuleOptions model
				createSchemaRuleOptionsModel := new(schemaregistryv1.CreateSchemaRuleOptions)
				createSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				createSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				createSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				createSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSchemaRuleOptions model with no property values
				createSchemaRuleOptionsModelNew := new(schemaregistryv1.CreateSchemaRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModelNew)
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
			It(`Invoke CreateSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaRuleOptions model
				createSchemaRuleOptionsModel := new(schemaregistryv1.CreateSchemaRuleOptions)
				createSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				createSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				createSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				createSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.CreateSchemaRule(createSchemaRuleOptionsModel)
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
	Describe(`GetSchemaRule(getSchemaRuleOptions *GetSchemaRuleOptions) - Operation response error`, func() {
		getSchemaRulePath := "/artifacts/testString/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchemaRulePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSchemaRule with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetSchemaRuleOptions model
				getSchemaRuleOptionsModel := new(schemaregistryv1.GetSchemaRuleOptions)
				getSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				getSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSchemaRule(getSchemaRuleOptions *GetSchemaRuleOptions)`, func() {
		getSchemaRulePath := "/artifacts/testString/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSchemaRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke GetSchemaRule successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the GetSchemaRuleOptions model
				getSchemaRuleOptionsModel := new(schemaregistryv1.GetSchemaRuleOptions)
				getSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				getSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.GetSchemaRuleWithContext(ctx, getSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.GetSchemaRuleWithContext(ctx, getSchemaRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSchemaRulePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke GetSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.GetSchemaRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchemaRuleOptions model
				getSchemaRuleOptionsModel := new(schemaregistryv1.GetSchemaRuleOptions)
				getSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				getSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSchemaRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetSchemaRuleOptions model
				getSchemaRuleOptionsModel := new(schemaregistryv1.GetSchemaRuleOptions)
				getSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				getSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSchemaRuleOptions model with no property values
				getSchemaRuleOptionsModelNew := new(schemaregistryv1.GetSchemaRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModelNew)
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
			It(`Invoke GetSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetSchemaRuleOptions model
				getSchemaRuleOptionsModel := new(schemaregistryv1.GetSchemaRuleOptions)
				getSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				getSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				getSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.GetSchemaRule(getSchemaRuleOptionsModel)
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
	Describe(`UpdateSchemaRule(updateSchemaRuleOptions *UpdateSchemaRuleOptions) - Operation response error`, func() {
		updateSchemaRulePath := "/artifacts/testString/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSchemaRule with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaRuleOptions model
				updateSchemaRuleOptionsModel := new(schemaregistryv1.UpdateSchemaRuleOptions)
				updateSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				updateSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSchemaRule(updateSchemaRuleOptions *UpdateSchemaRuleOptions)`, func() {
		updateSchemaRulePath := "/artifacts/testString/rules/COMPATIBILITY"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaRulePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke UpdateSchemaRule successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSchemaRuleOptions model
				updateSchemaRuleOptionsModel := new(schemaregistryv1.UpdateSchemaRuleOptions)
				updateSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				updateSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.UpdateSchemaRuleWithContext(ctx, updateSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.UpdateSchemaRuleWithContext(ctx, updateSchemaRuleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaRulePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"type": "COMPATIBILITY", "config": "BACKWARD"}`)
				}))
			})
			It(`Invoke UpdateSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.UpdateSchemaRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSchemaRuleOptions model
				updateSchemaRuleOptionsModel := new(schemaregistryv1.UpdateSchemaRuleOptions)
				updateSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				updateSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSchemaRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaRuleOptions model
				updateSchemaRuleOptionsModel := new(schemaregistryv1.UpdateSchemaRuleOptions)
				updateSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				updateSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSchemaRuleOptions model with no property values
				updateSchemaRuleOptionsModelNew := new(schemaregistryv1.UpdateSchemaRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModelNew)
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
			It(`Invoke UpdateSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaRuleOptions model
				updateSchemaRuleOptionsModel := new(schemaregistryv1.UpdateSchemaRuleOptions)
				updateSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				updateSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Type = core.StringPtr("COMPATIBILITY")
				updateSchemaRuleOptionsModel.Config = core.StringPtr("BACKWARD")
				updateSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.UpdateSchemaRule(updateSchemaRuleOptionsModel)
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
	Describe(`DeleteSchemaRule(deleteSchemaRuleOptions *DeleteSchemaRuleOptions)`, func() {
		deleteSchemaRulePath := "/artifacts/testString/rules/COMPATIBILITY"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSchemaRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSchemaRule successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := schemaregistryService.DeleteSchemaRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSchemaRuleOptions model
				deleteSchemaRuleOptionsModel := new(schemaregistryv1.DeleteSchemaRuleOptions)
				deleteSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				deleteSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				deleteSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schemaregistryService.DeleteSchemaRule(deleteSchemaRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSchemaRule with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the DeleteSchemaRuleOptions model
				deleteSchemaRuleOptionsModel := new(schemaregistryv1.DeleteSchemaRuleOptions)
				deleteSchemaRuleOptionsModel.ID = core.StringPtr("testString")
				deleteSchemaRuleOptionsModel.Rule = core.StringPtr("COMPATIBILITY")
				deleteSchemaRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schemaregistryService.DeleteSchemaRule(deleteSchemaRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSchemaRuleOptions model with no property values
				deleteSchemaRuleOptionsModelNew := new(schemaregistryv1.DeleteSchemaRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schemaregistryService.DeleteSchemaRule(deleteSchemaRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetSchemaState(setSchemaStateOptions *SetSchemaStateOptions)`, func() {
		setSchemaStatePath := "/artifacts/testString/state"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSchemaStatePath))
					Expect(req.Method).To(Equal("PUT"))

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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SetSchemaState successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := schemaregistryService.SetSchemaState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetSchemaStateOptions model
				setSchemaStateOptionsModel := new(schemaregistryv1.SetSchemaStateOptions)
				setSchemaStateOptionsModel.ID = core.StringPtr("testString")
				setSchemaStateOptionsModel.State = core.StringPtr("ENABLED")
				setSchemaStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schemaregistryService.SetSchemaState(setSchemaStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SetSchemaState with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the SetSchemaStateOptions model
				setSchemaStateOptionsModel := new(schemaregistryv1.SetSchemaStateOptions)
				setSchemaStateOptionsModel.ID = core.StringPtr("testString")
				setSchemaStateOptionsModel.State = core.StringPtr("ENABLED")
				setSchemaStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schemaregistryService.SetSchemaState(setSchemaStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SetSchemaStateOptions model with no property values
				setSchemaStateOptionsModelNew := new(schemaregistryv1.SetSchemaStateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schemaregistryService.SetSchemaState(setSchemaStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetSchemaVersionState(setSchemaVersionStateOptions *SetSchemaVersionStateOptions)`, func() {
		setSchemaVersionStatePath := "/artifacts/testString/versions/38/state"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(setSchemaVersionStatePath))
					Expect(req.Method).To(Equal("PUT"))

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

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SetSchemaVersionState successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := schemaregistryService.SetSchemaVersionState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SetSchemaVersionStateOptions model
				setSchemaVersionStateOptionsModel := new(schemaregistryv1.SetSchemaVersionStateOptions)
				setSchemaVersionStateOptionsModel.ID = core.StringPtr("testString")
				setSchemaVersionStateOptionsModel.Version = core.Int64Ptr(int64(38))
				setSchemaVersionStateOptionsModel.State = core.StringPtr("ENABLED")
				setSchemaVersionStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schemaregistryService.SetSchemaVersionState(setSchemaVersionStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SetSchemaVersionState with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the SetSchemaVersionStateOptions model
				setSchemaVersionStateOptionsModel := new(schemaregistryv1.SetSchemaVersionStateOptions)
				setSchemaVersionStateOptionsModel.ID = core.StringPtr("testString")
				setSchemaVersionStateOptionsModel.Version = core.Int64Ptr(int64(38))
				setSchemaVersionStateOptionsModel.State = core.StringPtr("ENABLED")
				setSchemaVersionStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schemaregistryService.SetSchemaVersionState(setSchemaVersionStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SetSchemaVersionStateOptions model with no property values
				setSchemaVersionStateOptionsModelNew := new(schemaregistryv1.SetSchemaVersionStateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schemaregistryService.SetSchemaVersionState(setSchemaVersionStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVersions(listVersionsOptions *ListVersionsOptions)`, func() {
		listVersionsPath := "/artifacts/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["jsonformat"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[17]`)
				}))
			})
			It(`Invoke ListVersions successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(schemaregistryv1.ListVersionsOptions)
				listVersionsOptionsModel.ID = core.StringPtr("testString")
				listVersionsOptionsModel.Jsonformat = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.ListVersionsWithContext(ctx, listVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.ListVersionsWithContext(ctx, listVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["jsonformat"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `[17]`)
				}))
			})
			It(`Invoke ListVersions successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.ListVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(schemaregistryv1.ListVersionsOptions)
				listVersionsOptionsModel.ID = core.StringPtr("testString")
				listVersionsOptionsModel.Jsonformat = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVersions with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(schemaregistryv1.ListVersionsOptions)
				listVersionsOptionsModel.ID = core.StringPtr("testString")
				listVersionsOptionsModel.Jsonformat = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.ListVersions(listVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListVersionsOptions model with no property values
				listVersionsOptionsModelNew := new(schemaregistryv1.ListVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.ListVersions(listVersionsOptionsModelNew)
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
			It(`Invoke ListVersions successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the ListVersionsOptions model
				listVersionsOptionsModel := new(schemaregistryv1.ListVersionsOptions)
				listVersionsOptionsModel.ID = core.StringPtr("testString")
				listVersionsOptionsModel.Jsonformat = core.StringPtr("testString")
				listVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.ListVersions(listVersionsOptionsModel)
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
	Describe(`CreateVersion(createVersionOptions *CreateVersionOptions) - Operation response error`, func() {
		createVersionPath := "/artifacts/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVersion with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateVersionOptions model
				createVersionOptionsModel := new(schemaregistryv1.CreateVersionOptions)
				createVersionOptionsModel.ID = core.StringPtr("testString")
				createVersionOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.CreateVersion(createVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.CreateVersion(createVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVersion(createVersionOptions *CreateVersionOptions)`, func() {
		createVersionPath := "/artifacts/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVersionPath))
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
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke CreateVersion successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the CreateVersionOptions model
				createVersionOptionsModel := new(schemaregistryv1.CreateVersionOptions)
				createVersionOptionsModel.ID = core.StringPtr("testString")
				createVersionOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.CreateVersionWithContext(ctx, createVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.CreateVersion(createVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.CreateVersionWithContext(ctx, createVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createVersionPath))
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
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke CreateVersion successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.CreateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateVersionOptions model
				createVersionOptionsModel := new(schemaregistryv1.CreateVersionOptions)
				createVersionOptionsModel.ID = core.StringPtr("testString")
				createVersionOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.CreateVersion(createVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVersion with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateVersionOptions model
				createVersionOptionsModel := new(schemaregistryv1.CreateVersionOptions)
				createVersionOptionsModel.ID = core.StringPtr("testString")
				createVersionOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.CreateVersion(createVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVersionOptions model with no property values
				createVersionOptionsModelNew := new(schemaregistryv1.CreateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.CreateVersion(createVersionOptionsModelNew)
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
			It(`Invoke CreateVersion successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateVersionOptions model
				createVersionOptionsModel := new(schemaregistryv1.CreateVersionOptions)
				createVersionOptionsModel.ID = core.StringPtr("testString")
				createVersionOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.CreateVersion(createVersionOptionsModel)
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
	Describe(`GetVersion(getVersionOptions *GetVersionOptions) - Operation response error`, func() {
		getVersionPath := "/artifacts/testString/versions/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVersion with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(schemaregistryv1.GetVersionOptions)
				getVersionOptionsModel.ID = core.StringPtr("testString")
				getVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVersion(getVersionOptions *GetVersionOptions)`, func() {
		getVersionPath := "/artifacts/testString/versions/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetVersion successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(schemaregistryv1.GetVersionOptions)
				getVersionOptionsModel.ID = core.StringPtr("testString")
				getVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.GetVersionWithContext(ctx, getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.GetVersionWithContext(ctx, getVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetVersion successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.GetVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(schemaregistryv1.GetVersionOptions)
				getVersionOptionsModel.ID = core.StringPtr("testString")
				getVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVersion with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(schemaregistryv1.GetVersionOptions)
				getVersionOptionsModel.ID = core.StringPtr("testString")
				getVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.GetVersion(getVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVersionOptions model with no property values
				getVersionOptionsModelNew := new(schemaregistryv1.GetVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.GetVersion(getVersionOptionsModelNew)
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
			It(`Invoke GetVersion successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetVersionOptions model
				getVersionOptionsModel := new(schemaregistryv1.GetVersionOptions)
				getVersionOptionsModel.ID = core.StringPtr("testString")
				getVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.GetVersion(getVersionOptionsModel)
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
	Describe(`DeleteVersion(deleteVersionOptions *DeleteVersionOptions)`, func() {
		deleteVersionPath := "/artifacts/testString/versions/38"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteVersion successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := schemaregistryService.DeleteVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteVersionOptions model
				deleteVersionOptionsModel := new(schemaregistryv1.DeleteVersionOptions)
				deleteVersionOptionsModel.ID = core.StringPtr("testString")
				deleteVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schemaregistryService.DeleteVersion(deleteVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteVersion with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the DeleteVersionOptions model
				deleteVersionOptionsModel := new(schemaregistryv1.DeleteVersionOptions)
				deleteVersionOptionsModel.ID = core.StringPtr("testString")
				deleteVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schemaregistryService.DeleteVersion(deleteVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteVersionOptions model with no property values
				deleteVersionOptionsModelNew := new(schemaregistryv1.DeleteVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schemaregistryService.DeleteVersion(deleteVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSchemas(listSchemasOptions *ListSchemasOptions)`, func() {
		listSchemasPath := "/artifacts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["jsonformat"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `["OperationResponse"]`)
				}))
			})
			It(`Invoke ListSchemas successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(schemaregistryv1.ListSchemasOptions)
				listSchemasOptionsModel.Jsonformat = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.ListSchemasWithContext(ctx, listSchemasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.ListSchemasWithContext(ctx, listSchemasOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSchemasPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["jsonformat"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `["OperationResponse"]`)
				}))
			})
			It(`Invoke ListSchemas successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.ListSchemas(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(schemaregistryv1.ListSchemasOptions)
				listSchemasOptionsModel.Jsonformat = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.ListSchemas(listSchemasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSchemas with error: Operation request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(schemaregistryv1.ListSchemasOptions)
				listSchemasOptionsModel.Jsonformat = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.ListSchemas(listSchemasOptionsModel)
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
			It(`Invoke ListSchemas successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := new(schemaregistryv1.ListSchemasOptions)
				listSchemasOptionsModel.Jsonformat = core.StringPtr("testString")
				listSchemasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.ListSchemas(listSchemasOptionsModel)
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
	Describe(`CreateSchema(createSchemaOptions *CreateSchemaOptions) - Operation response error`, func() {
		createSchemaPath := "/artifacts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Registry-Artifactid"]).ToNot(BeNil())
					Expect(req.Header["X-Registry-Artifactid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSchema with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(schemaregistryv1.CreateSchemaOptions)
				createSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createSchemaOptionsModel.XRegistryArtifactID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
		createSchemaPath := "/artifacts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
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

					Expect(req.Header["X-Registry-Artifactid"]).ToNot(BeNil())
					Expect(req.Header["X-Registry-Artifactid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke CreateSchema successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(schemaregistryv1.CreateSchemaOptions)
				createSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createSchemaOptionsModel.XRegistryArtifactID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.CreateSchemaWithContext(ctx, createSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.CreateSchemaWithContext(ctx, createSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSchemaPath))
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

					Expect(req.Header["X-Registry-Artifactid"]).ToNot(BeNil())
					Expect(req.Header["X-Registry-Artifactid"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke CreateSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.CreateSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(schemaregistryv1.CreateSchemaOptions)
				createSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createSchemaOptionsModel.XRegistryArtifactID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSchema with error: Operation request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(schemaregistryv1.CreateSchemaOptions)
				createSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createSchemaOptionsModel.XRegistryArtifactID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.CreateSchema(createSchemaOptionsModel)
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
			It(`Invoke CreateSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(schemaregistryv1.CreateSchemaOptions)
				createSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				createSchemaOptionsModel.XRegistryArtifactID = core.StringPtr("testString")
				createSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.CreateSchema(createSchemaOptionsModel)
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
	Describe(`GetLatestSchema(getLatestSchemaOptions *GetLatestSchemaOptions) - Operation response error`, func() {
		getLatestSchemaPath := "/artifacts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestSchemaPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLatestSchema with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetLatestSchemaOptions model
				getLatestSchemaOptionsModel := new(schemaregistryv1.GetLatestSchemaOptions)
				getLatestSchemaOptionsModel.ID = core.StringPtr("testString")
				getLatestSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLatestSchema(getLatestSchemaOptions *GetLatestSchemaOptions)`, func() {
		getLatestSchemaPath := "/artifacts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLatestSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetLatestSchema successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the GetLatestSchemaOptions model
				getLatestSchemaOptionsModel := new(schemaregistryv1.GetLatestSchemaOptions)
				getLatestSchemaOptionsModel.ID = core.StringPtr("testString")
				getLatestSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.GetLatestSchemaWithContext(ctx, getLatestSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.GetLatestSchemaWithContext(ctx, getLatestSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getLatestSchemaPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"schema": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetLatestSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.GetLatestSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLatestSchemaOptions model
				getLatestSchemaOptionsModel := new(schemaregistryv1.GetLatestSchemaOptions)
				getLatestSchemaOptionsModel.ID = core.StringPtr("testString")
				getLatestSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLatestSchema with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetLatestSchemaOptions model
				getLatestSchemaOptionsModel := new(schemaregistryv1.GetLatestSchemaOptions)
				getLatestSchemaOptionsModel.ID = core.StringPtr("testString")
				getLatestSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLatestSchemaOptions model with no property values
				getLatestSchemaOptionsModelNew := new(schemaregistryv1.GetLatestSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModelNew)
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
			It(`Invoke GetLatestSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the GetLatestSchemaOptions model
				getLatestSchemaOptionsModel := new(schemaregistryv1.GetLatestSchemaOptions)
				getLatestSchemaOptionsModel.ID = core.StringPtr("testString")
				getLatestSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.GetLatestSchema(getLatestSchemaOptionsModel)
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
	Describe(`DeleteSchema(deleteSchemaOptions *DeleteSchemaOptions)`, func() {
		deleteSchemaPath := "/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSchemaPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := schemaregistryService.DeleteSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsModel := new(schemaregistryv1.DeleteSchemaOptions)
				deleteSchemaOptionsModel.ID = core.StringPtr("testString")
				deleteSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = schemaregistryService.DeleteSchema(deleteSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSchema with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the DeleteSchemaOptions model
				deleteSchemaOptionsModel := new(schemaregistryv1.DeleteSchemaOptions)
				deleteSchemaOptionsModel.ID = core.StringPtr("testString")
				deleteSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := schemaregistryService.DeleteSchema(deleteSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSchemaOptions model with no property values
				deleteSchemaOptionsModelNew := new(schemaregistryv1.DeleteSchemaOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = schemaregistryService.DeleteSchema(deleteSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSchema(updateSchemaOptions *UpdateSchemaOptions) - Operation response error`, func() {
		updateSchemaPath := "/artifacts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSchema with error: Operation response processing error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaOptions model
				updateSchemaOptionsModel := new(schemaregistryv1.UpdateSchemaOptions)
				updateSchemaOptionsModel.ID = core.StringPtr("testString")
				updateSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				updateSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				schemaregistryService.EnableRetries(0, 0)
				result, response, operationErr = schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSchema(updateSchemaOptions *UpdateSchemaOptions)`, func() {
		updateSchemaPath := "/artifacts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke UpdateSchema successfully with retries`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())
				schemaregistryService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSchemaOptions model
				updateSchemaOptionsModel := new(schemaregistryv1.UpdateSchemaOptions)
				updateSchemaOptionsModel.ID = core.StringPtr("testString")
				updateSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				updateSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := schemaregistryService.UpdateSchemaWithContext(ctx, updateSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				schemaregistryService.DisableRetries()
				result, response, operationErr := schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = schemaregistryService.UpdateSchemaWithContext(ctx, updateSchemaOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSchemaPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"createdOn": 9, "globalId": 8, "id": "ID", "modifiedOn": 10, "type": "Type", "version": 7}`)
				}))
			})
			It(`Invoke UpdateSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := schemaregistryService.UpdateSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSchemaOptions model
				updateSchemaOptionsModel := new(schemaregistryv1.UpdateSchemaOptions)
				updateSchemaOptionsModel.ID = core.StringPtr("testString")
				updateSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				updateSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSchema with error: Operation validation and request error`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaOptions model
				updateSchemaOptionsModel := new(schemaregistryv1.UpdateSchemaOptions)
				updateSchemaOptionsModel.ID = core.StringPtr("testString")
				updateSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				updateSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := schemaregistryService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSchemaOptions model with no property values
				updateSchemaOptionsModelNew := new(schemaregistryv1.UpdateSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = schemaregistryService.UpdateSchema(updateSchemaOptionsModelNew)
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
			It(`Invoke UpdateSchema successfully`, func() {
				schemaregistryService, serviceErr := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(schemaregistryService).ToNot(BeNil())

				// Construct an instance of the UpdateSchemaOptions model
				updateSchemaOptionsModel := new(schemaregistryv1.UpdateSchemaOptions)
				updateSchemaOptionsModel.ID = core.StringPtr("testString")
				updateSchemaOptionsModel.Schema = map[string]interface{}{"anyKey": "anyValue"}
				updateSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := schemaregistryService.UpdateSchema(updateSchemaOptionsModel)
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
			schemaregistryService, _ := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
				URL:           "http://schemaregistryv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateSchemaOptions successfully`, func() {
				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := schemaregistryService.NewCreateSchemaOptions()
				createSchemaOptionsModel.SetSchema(map[string]interface{}{"anyKey": "anyValue"})
				createSchemaOptionsModel.SetXRegistryArtifactID("testString")
				createSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSchemaOptionsModel).ToNot(BeNil())
				Expect(createSchemaOptionsModel.Schema).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createSchemaOptionsModel.XRegistryArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSchemaRuleOptions successfully`, func() {
				// Construct an instance of the CreateSchemaRuleOptions model
				id := "testString"
				createSchemaRuleOptionsType := "COMPATIBILITY"
				createSchemaRuleOptionsConfig := "BACKWARD"
				createSchemaRuleOptionsModel := schemaregistryService.NewCreateSchemaRuleOptions(id, createSchemaRuleOptionsType, createSchemaRuleOptionsConfig)
				createSchemaRuleOptionsModel.SetID("testString")
				createSchemaRuleOptionsModel.SetType("COMPATIBILITY")
				createSchemaRuleOptionsModel.SetConfig("BACKWARD")
				createSchemaRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSchemaRuleOptionsModel).ToNot(BeNil())
				Expect(createSchemaRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createSchemaRuleOptionsModel.Type).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(createSchemaRuleOptionsModel.Config).To(Equal(core.StringPtr("BACKWARD")))
				Expect(createSchemaRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVersionOptions successfully`, func() {
				// Construct an instance of the CreateVersionOptions model
				id := "testString"
				createVersionOptionsModel := schemaregistryService.NewCreateVersionOptions(id)
				createVersionOptionsModel.SetID("testString")
				createVersionOptionsModel.SetSchema(map[string]interface{}{"anyKey": "anyValue"})
				createVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVersionOptionsModel).ToNot(BeNil())
				Expect(createVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createVersionOptionsModel.Schema).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSchemaOptions successfully`, func() {
				// Construct an instance of the DeleteSchemaOptions model
				id := "testString"
				deleteSchemaOptionsModel := schemaregistryService.NewDeleteSchemaOptions(id)
				deleteSchemaOptionsModel.SetID("testString")
				deleteSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSchemaOptionsModel).ToNot(BeNil())
				Expect(deleteSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSchemaRuleOptions successfully`, func() {
				// Construct an instance of the DeleteSchemaRuleOptions model
				id := "testString"
				rule := "COMPATIBILITY"
				deleteSchemaRuleOptionsModel := schemaregistryService.NewDeleteSchemaRuleOptions(id, rule)
				deleteSchemaRuleOptionsModel.SetID("testString")
				deleteSchemaRuleOptionsModel.SetRule("COMPATIBILITY")
				deleteSchemaRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSchemaRuleOptionsModel).ToNot(BeNil())
				Expect(deleteSchemaRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSchemaRuleOptionsModel.Rule).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(deleteSchemaRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVersionOptions successfully`, func() {
				// Construct an instance of the DeleteVersionOptions model
				id := "testString"
				version := int64(38)
				deleteVersionOptionsModel := schemaregistryService.NewDeleteVersionOptions(id, version)
				deleteVersionOptionsModel.SetID("testString")
				deleteVersionOptionsModel.SetVersion(int64(38))
				deleteVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVersionOptionsModel).ToNot(BeNil())
				Expect(deleteVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVersionOptionsModel.Version).To(Equal(core.Int64Ptr(int64(38))))
				Expect(deleteVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGlobalRuleOptions successfully`, func() {
				// Construct an instance of the GetGlobalRuleOptions model
				rule := "COMPATIBILITY"
				getGlobalRuleOptionsModel := schemaregistryService.NewGetGlobalRuleOptions(rule)
				getGlobalRuleOptionsModel.SetRule("COMPATIBILITY")
				getGlobalRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGlobalRuleOptionsModel).ToNot(BeNil())
				Expect(getGlobalRuleOptionsModel.Rule).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(getGlobalRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLatestSchemaOptions successfully`, func() {
				// Construct an instance of the GetLatestSchemaOptions model
				id := "testString"
				getLatestSchemaOptionsModel := schemaregistryService.NewGetLatestSchemaOptions(id)
				getLatestSchemaOptionsModel.SetID("testString")
				getLatestSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLatestSchemaOptionsModel).ToNot(BeNil())
				Expect(getLatestSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getLatestSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSchemaRuleOptions successfully`, func() {
				// Construct an instance of the GetSchemaRuleOptions model
				id := "testString"
				rule := "COMPATIBILITY"
				getSchemaRuleOptionsModel := schemaregistryService.NewGetSchemaRuleOptions(id, rule)
				getSchemaRuleOptionsModel.SetID("testString")
				getSchemaRuleOptionsModel.SetRule("COMPATIBILITY")
				getSchemaRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSchemaRuleOptionsModel).ToNot(BeNil())
				Expect(getSchemaRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getSchemaRuleOptionsModel.Rule).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(getSchemaRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVersionOptions successfully`, func() {
				// Construct an instance of the GetVersionOptions model
				id := "testString"
				version := int64(38)
				getVersionOptionsModel := schemaregistryService.NewGetVersionOptions(id, version)
				getVersionOptionsModel.SetID("testString")
				getVersionOptionsModel.SetVersion(int64(38))
				getVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVersionOptionsModel).ToNot(BeNil())
				Expect(getVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getVersionOptionsModel.Version).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSchemasOptions successfully`, func() {
				// Construct an instance of the ListSchemasOptions model
				listSchemasOptionsModel := schemaregistryService.NewListSchemasOptions()
				listSchemasOptionsModel.SetJsonformat("testString")
				listSchemasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSchemasOptionsModel).ToNot(BeNil())
				Expect(listSchemasOptionsModel.Jsonformat).To(Equal(core.StringPtr("testString")))
				Expect(listSchemasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVersionsOptions successfully`, func() {
				// Construct an instance of the ListVersionsOptions model
				id := "testString"
				listVersionsOptionsModel := schemaregistryService.NewListVersionsOptions(id)
				listVersionsOptionsModel.SetID("testString")
				listVersionsOptionsModel.SetJsonformat("testString")
				listVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVersionsOptionsModel).ToNot(BeNil())
				Expect(listVersionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOptionsModel.Jsonformat).To(Equal(core.StringPtr("testString")))
				Expect(listVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetSchemaStateOptions successfully`, func() {
				// Construct an instance of the SetSchemaStateOptions model
				id := "testString"
				setSchemaStateOptionsState := "ENABLED"
				setSchemaStateOptionsModel := schemaregistryService.NewSetSchemaStateOptions(id, setSchemaStateOptionsState)
				setSchemaStateOptionsModel.SetID("testString")
				setSchemaStateOptionsModel.SetState("ENABLED")
				setSchemaStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setSchemaStateOptionsModel).ToNot(BeNil())
				Expect(setSchemaStateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setSchemaStateOptionsModel.State).To(Equal(core.StringPtr("ENABLED")))
				Expect(setSchemaStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetSchemaVersionStateOptions successfully`, func() {
				// Construct an instance of the SetSchemaVersionStateOptions model
				id := "testString"
				version := int64(38)
				setSchemaVersionStateOptionsState := "ENABLED"
				setSchemaVersionStateOptionsModel := schemaregistryService.NewSetSchemaVersionStateOptions(id, version, setSchemaVersionStateOptionsState)
				setSchemaVersionStateOptionsModel.SetID("testString")
				setSchemaVersionStateOptionsModel.SetVersion(int64(38))
				setSchemaVersionStateOptionsModel.SetState("ENABLED")
				setSchemaVersionStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setSchemaVersionStateOptionsModel).ToNot(BeNil())
				Expect(setSchemaVersionStateOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(setSchemaVersionStateOptionsModel.Version).To(Equal(core.Int64Ptr(int64(38))))
				Expect(setSchemaVersionStateOptionsModel.State).To(Equal(core.StringPtr("ENABLED")))
				Expect(setSchemaVersionStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateGlobalRuleOptions successfully`, func() {
				// Construct an instance of the UpdateGlobalRuleOptions model
				rule := "COMPATIBILITY"
				updateGlobalRuleOptionsType := "COMPATIBILITY"
				updateGlobalRuleOptionsConfig := "BACKWARD"
				updateGlobalRuleOptionsModel := schemaregistryService.NewUpdateGlobalRuleOptions(rule, updateGlobalRuleOptionsType, updateGlobalRuleOptionsConfig)
				updateGlobalRuleOptionsModel.SetRule("COMPATIBILITY")
				updateGlobalRuleOptionsModel.SetType("COMPATIBILITY")
				updateGlobalRuleOptionsModel.SetConfig("BACKWARD")
				updateGlobalRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateGlobalRuleOptionsModel).ToNot(BeNil())
				Expect(updateGlobalRuleOptionsModel.Rule).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(updateGlobalRuleOptionsModel.Type).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(updateGlobalRuleOptionsModel.Config).To(Equal(core.StringPtr("BACKWARD")))
				Expect(updateGlobalRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSchemaOptions successfully`, func() {
				// Construct an instance of the UpdateSchemaOptions model
				id := "testString"
				updateSchemaOptionsModel := schemaregistryService.NewUpdateSchemaOptions(id)
				updateSchemaOptionsModel.SetID("testString")
				updateSchemaOptionsModel.SetSchema(map[string]interface{}{"anyKey": "anyValue"})
				updateSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSchemaOptionsModel).ToNot(BeNil())
				Expect(updateSchemaOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateSchemaOptionsModel.Schema).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSchemaRuleOptions successfully`, func() {
				// Construct an instance of the UpdateSchemaRuleOptions model
				id := "testString"
				rule := "COMPATIBILITY"
				updateSchemaRuleOptionsType := "COMPATIBILITY"
				updateSchemaRuleOptionsConfig := "BACKWARD"
				updateSchemaRuleOptionsModel := schemaregistryService.NewUpdateSchemaRuleOptions(id, rule, updateSchemaRuleOptionsType, updateSchemaRuleOptionsConfig)
				updateSchemaRuleOptionsModel.SetID("testString")
				updateSchemaRuleOptionsModel.SetRule("COMPATIBILITY")
				updateSchemaRuleOptionsModel.SetType("COMPATIBILITY")
				updateSchemaRuleOptionsModel.SetConfig("BACKWARD")
				updateSchemaRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSchemaRuleOptionsModel).ToNot(BeNil())
				Expect(updateSchemaRuleOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateSchemaRuleOptionsModel.Rule).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(updateSchemaRuleOptionsModel.Type).To(Equal(core.StringPtr("COMPATIBILITY")))
				Expect(updateSchemaRuleOptionsModel.Config).To(Equal(core.StringPtr("BACKWARD")))
				Expect(updateSchemaRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRule successfully`, func() {
				typeVar := "COMPATIBILITY"
				config := "BACKWARD"
				_model, err := schemaregistryService.NewRule(typeVar, config)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
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
