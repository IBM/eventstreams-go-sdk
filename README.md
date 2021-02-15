# Introduction

IBM Event Streams for IBM Cloud™ is a high-throughput message bus built with Apache Kafka. It is optimized for event ingestion into IBM Cloud and event stream distribution between your services and applications.

Event Streams provides a REST API to help connect your existing systems to your Event Streams Kafka cluster. Using the API, you can integrate Event Streams with any system that supports RESTful APIs.

# Event Streams Administration REST API

This REST API allows users of the
[IBM Event Streams service](https://cloud.ibm.com/docs/services/EventStreams/index.html)
to administer
[Kafka topics](#using-the-rest-api-to-administer-kafka-topics)
associated with an instance of the service. You can use this API to perform the following
operations:
  - [Create a Kafka topic](#creating-a-kafka-topic)
  - [List Kafka topics](#listing-kafka-topics)
  - [Get a Kafka topic](#getting-a-kafka-topic)
  - [Delete a Kafka topic](#deleting-a-kafka-topic)
  - [Update a Kafka topic configuration](#updating-kafka-topics-configuration)
  - [List which topics are mirrored](#list-current-mirroring-topic-selection)
  - [Replace selection of topics which are mirrored](#replace-mirroring-topic-selection)
  - [List active mirroring topics](#list-active-mirroring-topics)
  
The Admin REST API is also [documented using swagger](./admin-rest-api.yaml).

Note: This Admin REST API works with both Enterprise plan and Standard plan. The Admin REST API that works with Classic plan is located in [here](../admin-rest-api-classic-plan-only). They are compatible in topic management capabilities.

## Access control
All requests support below authorization methods:
 * Basic authorization with user and password. (
  For both standard and enterprise plan, user is 'token', password is the API key from `ibmcloud resource service-keys` for the service instance.)
 * Bearer authorization with bearer token. (This token can be either API key or JWT token obtained from IAM upon login to IBM Cloud. Use `ibmcloud iam oauth-tokens` to retrieve the token after `ibmcloud login`)
 * `X-Auth-Token` header to be set to the API key. This header is deprecated.

##  Administration API endpoint
Administration API endpoint is the `kafka_admin_url` property in the service key for the service instance. This command can be used to retrieve this property.
```bash
$ibmcloud resource service-key "${service_instance_key_name}" --output json > jq -r '.[]|.credentials.kafka_admin_url'
```

In addition, the `Content-type` header has to be set to `application/json`.

Common HTTP status codes:
- 200: Request succeeded.
- 202: Request was accepted.
- 400: Invalid request JSON.
- 401: The authentication header is not set or provided information is not valid.
- 403: Not authorized to perform the operation. Usually it means the API key used is missing a certain role. More details on what role can perform what operation refers to this [document](https://cloud.ibm.com/docs/services/EventStreams?topic=eventstreams-security).
- 404: Unable to find the topic with topic name given by user.
- 422: Semantically invalid request.
- 503: An error occurred handling the request.

Error responses carry a JSON body like the following:
```json
{"error_code":50301,"message":"Unknown Kafka Error", "incident_id": "17afe715-0ff5-4c49-9acc-a4204244a331"}
```
Error codes are of the format `HHHKK` where `HHH` is the HTTP Status Code and `KK` is the Kafka protocol error.  

For E2E debugging purposes, the transaction ID of every request is returned in the HTTP header `X-Global-Transaction-Id`.
If the header is set on the request, it will be honored. If not, it will be generated.
In the event of a non-200 error return code, the transaction ID is also returned in the JSON error response as `incident_id`.


## Using the REST API to administer Kafka topics

The remainder of this document describes the Go-lang implementation of the Admin Rest SDK 
and we also provide example `example.go` to show all SDK calls in action. 

To run the example :-

Compile the code.
```sh
	go build -o eventstreams-sdk-example
```
Set the required environment variables
```sh
# Set you API KEY.
export API_KEY="abc123456789"

# Set the Admin Endpoint to point to your cluster.
export ADMIN_ENDPOINT="https://xyzclustername.svc01.region.eventstreams.test.cloud.ibm.com"

```

Run the example
```sh
./eventstreams-example 
```

## REST API 

The following sections explain how the REST API works with examples.

### Code Setup

You will need to import the IBM SDK core functions and the rest API client itself.
```golang
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/eventstreams-go-sdk/pkg/adminrestv1"
```

### Authentication

Use one of the following methods to authenticate:

- To authenticate using Basic Auth:
  Use the user and api_key properties of the above objects as the username and password. Place these values into the Authorization header of the HTTP request in the form Basic <base64 encoding of username and password joined by a single colon (:)>.

- To authenticate using a bearer token:
  To obtain your token using the IBM Cloud CLI, first log in to IBM Cloud, then run the following command:
  ```
    ibmcloud iam oauth-tokens
  ```
    Place this token in the Authorization header of the HTTP request in the form Bearer. Both API key or JWT tokens are supported.

- To authenticate directly using the api_key:
  Place the key directly as the value of the X-Auth-Token HTTP header.

```golang
    URL := os.Getenv("ADMIN_ENDPOINT")
	apiKey := os.Getenv("API_KEY")

	if URL == "" {
		fmt.Println("Please set env ADMIN_ENDPOINT")
		os.Exit(1)
	}

	if apiKey == "" {
		fmt.Println("Please set env API_KEY")
		os.Exit(1)
	}

	basicAuthenticator, err := core.NewBasicAuthenticator("token", apiKey)
	if err != nil {
		fmt.Printf("Failed to Authenticate with error : %s\n", err.Error())
		os.Exit(1)
	}
```


### Create the service

Create a new service object.

```golang
    // Create a new REST service.
	ibmEventStreamsAdminRestApiService, serviceErr := ibmeventstreamsadminrestapiv1.NewIbmEventStreamsAdminRestApiV1(&ibmeventstreamsadminrestapiv1.IbmEventStreamsAdminRestApiV1Options{
		URL:           URL,
		Authenticator: basicAuthenticator,
	})
	if serviceErr != nil {
		fmt.Println("Error Creating Service")
		os.Exit(1)
	}
```

### Creating a Kafka topic

To get a Kafka topic detail information, the admin REST SDK issues a GET request to the `/admin/topics/TOPICNAME`
path (where `TOPICNAME` is the name of the Kafka topic that you want to get).  

Expected status codes
  - 200: Retrieve topic details successfully in following format:


#### Example

```golang
	// Set the retries policy.
	serviceAPI.EnableRetries(0, 0)

	// Construct an instance of the createTopicOptionsModel.
	createTopicOptionsModel := new(ibmeventstreamsadminrestapiv1.CreateTopicOptions)
	createTopicOptionsModel.Name = core.StringPtr("test-topic")
	createTopicOptionsModel.Partitions = core.Int64Ptr(int64(26))
	createTopicOptionsModel.PartitionCount = core.Int64Ptr(int64(1))
	createTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Create the Topic.
	response, operationErr := serviceAPI.CreateTopic(createTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Creating Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Creating Topic: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s created\n", *createTopicOptionsModel.Name) 
```

### Deleting a Kafka topic

To delete a Kafka topic, the admin REST SDK issues a DELETE request to the `/admin/topics/TOPICNAME`
path (where `TOPICNAME` is the name of the Kafka topic that you want to delete).

Expected return codes:
- 202: Topic deletion request was accepted.
- 403: Not authorized to delete topic.
- 404: Topic does not exist.
  
A 202 (Accepted) status code is returned if the REST API accepts the delete
request or status code 422 (Un-processable Entity) if the delete request is
rejected. If a delete request is rejected then the body of the HTTP response
will contain a [JSON object](#information-returned-when-a-request-fails) which
provides additional information about why the request was rejected.

Kafka deletes topics asynchronously. Deleted topics may still appear in the
response to a [list topics request](#listing-kafka-topics) for a short period
of time after the completion of a REST request to delete the topic.

#### Example

First create a delete topic options model which gives details of the topic.
Then call the DeleteTopic function on the ibmEventStreamsAdminRestApiService service.

```golang
	// Construct an instance of the DeleteTopicOptions model
	deleteTopicOptionsModel := new(ibmeventstreamsadminrestapiv1.DeleteTopicOptions)
	deleteTopicOptionsModel.TopicName = core.StringPtr("test-topic")
	deleteTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Delete Topic
	response, operationErr := serviceAPI.DeleteTopic(deleteTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Deleting Topic: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Deleting Topic: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s deleted\n", *deleteTopicOptionsModel.TopicName)
```

### Listing Kafka topics

You can list all of your Kafka topics by issuing a GET request to the
`/admin/topics` path. 

Expected status codes:
  - 200: the topic list is returned as JSON in the following format:
```json
[
  {
    "name": "topic1",
    "partitions": 1,
    "retentionMs": 86400000,
    "cleanupPolicy": "delete"
  },
  { "name": "topic2",
    "partitions": 2,
    "retentionMs": 86400000,
    "cleanupPolicy": "delete"
  }
]
```

A successful response will have HTTP status code 200 (OK) and contain an
array of JSON objects, where each object represents a Kafka topic and has the
following properties:

| Property name     | Description                                             |
|-------------------|---------------------------------------------------------|
| name              | The name of the Kafka topic.                            |
| partitions        | The number of partitions assigned to the Kafka topic.   |
| retentionsMs      | The retention period for messages on the topic (in ms). |
| cleanupPolicy     | The cleanup policy of the Kafka topic.                  |

#### Example

```golang
	// Construct an instance of the ListTopicsOptions model
	listTopicsOptionsModel := new(ibmeventstreamsadminrestapiv1.ListTopicsOptions)

	// Call ListTopics.
	result, response, operationErr := serviceAPI.ListTopics(listTopicsOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Topics" + operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error Listing Topics: status %d\n", response.StatusCode)
	}

	// Loop and print topics.
	for _, topicDetail := range result {
		fmt.Printf("\tname: %s\n", *topicDetail.Name)
	}
```

### Getting a Kafka topic

To get a Kafka topic detail information, issue a GET request to the `/admin/topics/TOPICNAME`
path (where `TOPICNAME` is the name of the Kafka topic that you want to get).  

Expected status codes
  - 200: Retrieve topic details successfully in following format:
```json
{
  "name": "MYTOPIC",
  "partitions": 1,
  "replicationFactor": 3,
  "retentionMs": 86400000,
  "cleanupPolicy": "delete",
  "configs": {
    "cleanup.policy": "delete",
    "min.insync.replicas": "2",
    "retention.bytes": "1073741824",
    "retention.ms": "86400000",
    "segment.bytes": "536870912"
  },
  "replicaAssignments": [
    {
      "id": 0,
      "brokers": {
        "replicas": [
          3,
          2,
          4
        ]
      }
    }
  ]
}
```

#### Example
```golang
	// Construct an instance of the GetTopicOptions model
	getTopicOptionsModel := new(ibmeventstreamsadminrestapiv1.GetTopicOptions)
	getTopicOptionsModel.TopicName = core.StringPtr("test-topic")
	getTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Call List Topic Details.
	result, response, operationErr := serviceAPI.GetTopic(getTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Topic Details" + operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error Listing Topic Details: status %d\n", response.StatusCode)
	}

	// Print topics details.
	fmt.Printf("\tname: \t\t\t%s\n", *result.Name)

	// The number of partitions.
	fmt.Printf("\tno of partitions: \t%d\n", *result.Partitions)

	// The number of replication factor.
	fmt.Printf("\treplication factor: \t%d\n", *result.ReplicationFactor)

	// // The value of config property 'retention.ms'.
	fmt.Printf("\tretention (ms): \t%d\n", *result.RetentionMs)

	// // The value of config property 'cleanup.policy'.
	fmt.Printf("\tcleanup policy: \t%s\n", *result.CleanupPolicy)

	// Configs *TopicConfigs
	fmt.Printf("\ttopic configs: \t\t%+v\n", *result.Configs)

	// The replia assignment of the topic.
	// ReplicaAssignments []ReplicaAssignment
	for _, assignment := range result.ReplicaAssignments {
		fmt.Printf("\tassignment:  \t\tid:%d,  \tbrokers: %+v\n", assignment.ID, assignment.Brokers)
	}
```

### Updating Kafka topic's configuration

To increase a topic's partition number or to update a topic's configuration, the admin REST SDK issues an
`PATCH` request to `/admin/topics/{topic}` with the following body:
```json
{
  "new_total_partition_count": 4,
  "configs": [
    {
      "name": "cleanup.policy",
      "value": "compact"
    }
  ]
}
```
Supported configuration keys are 'cleanup.policy', 'retention.ms', 'retention.bytes', 'segment.bytes', 'segment.ms', 'segment.index.bytes'.
And partition number can only be increased, not decreased.

Expected status codes
  - 202: Update topic request was accepted.
  - 400: Invalid request JSON/number of partitions is invalid.
  - 404: Topic specified does not exist.
  - 422: Semantically invalid request.

#### Example

```golang
	// Construct an instance of the UpdateTopicOptions model
	updateTopicOptionsModel := new(ibmeventstreamsadminrestapiv1.UpdateTopicOptions)
	updateTopicOptionsModel.TopicName = core.StringPtr("test-topic")
	updateTopicOptionsModel.NewTotalPartitionCount = core.Int64Ptr(int64(6))
	updateTopicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Invoke operation with valid options model (positive test)
	response, operationErr := serviceAPI.UpdateTopic(updateTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Updating Topic: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Updating Topics: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s updated\n", *updateTopicOptionsModel.TopicName)

```

### List current mirroring topic selection

Mirroring user controls are available on the target cluster in a mirroring environment.

To get the current topic selection, issue an GET request to /admin/mirroring/topic-selection


Expected status codes
    - 200: Retrieved topic selection successfully in following format:
	{
	"includes": [
		"^prefix1_.*",
		"^prefix2_.*"
	]
	}
    - 403: Unauthorized to use mirroring user controls.
    - 404: Mirroring not enabled. The mirroring user control apis are only available on a target in a - pair of clusters with mirroring enabled between them.
    - 503: An error occurred handling the request.

#### Example

```golang
	// Construct an instance of the GetMirroringTopicSelectionOptions model
	getMirroringTopicSelectionOptionsModel := new(ibmeventstreamsadminrestapiv1.GetMirroringTopicSelectionOptions)
	getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Enable retries.
	serviceAPI.EnableRetries(0, 0)

	// Call GetMirroringTopicSelection.
	result, response, operationErr := serviceAPI.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Listing Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Loop and print mirroring topics.
	for _, topicName := range result.Includes {
		fmt.Printf("\tname: %s\n", topicName)
	}
```

### Replace selection of topics which are mirrored

Replace mirroring topic selection

Mirroring user controls are available on the target cluster in a mirroring environment.

To replace the current topic selection, issue a POST request to /admin/mirroring/topic-selection

Expected status codes

    200: Replaced topic selection successfully. The new selection is returned in following format:

	{
	"includes": [
		"^prefix1_.*",
		"^prefix2_.*"
	]
	}

    400: Invalid request. The request data cannot be parsed and used to replace the topic selection.
    403: Unauthorized to use mirroring user controls.
    404: Mirroring not enabled. The mirroring user control apis are only available on a target in a pair of clusters with mirroring enabled between them.
    415: Unsupported media type. Content-Type header with application/json is required.
    503: An error occurred handling the request.

#### Example

```golang
	replaceMirroringTopicSelectionOptionsModel := new(ibmeventstreamsadminrestapiv1.ReplaceMirroringTopicSelectionOptions)
	replaceMirroringTopicSelectionOptionsModel.Includes = []string{"test-topic"}
	replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Enable retries.
	serviceAPI.EnableRetries(0, 0)

	// Invoke operation with valid options model (positive test)
	result, response, operationErr := serviceAPI.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Replacing Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Replacing Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Loop and print mirroring topics.
	for _, topicName := range result.Includes {
		fmt.Printf("\ttopic added: %s\n", topicName)
	}
	
```

### List active mirroring topics

Mirroring user controls are available on the target cluster in a mirroring environment.

To get the list of currently mirrored topics, issue an GET request to /admin/mirroring/active-topics

Expected status codes

    200: Retrieved active topics successfully in following format:

	{
	"active_topics": [
		"topic1",
		"topic2"
	]
	}

    403: Unauthorized to use mirroring user controls.
    404: Mirroring not enabled. The mirroring user control apis are only available on a target in a pair of clusters with mirroring enabled between them.
    503: An error occurred handling the request.

#### Example

```golang
	// Construct an instance of the GetMirroringActiveTopicsOptions model
	getMirroringActiveTopicsOptionsModel := new(serviceName.GetMirroringActiveTopicsOptions)
	getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

	// Call GetMirroringActiveTopics.
	result, response, operationErr := serviceAPI.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Active Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Listing Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Lopp and print mirroring topics.
	for _, topicName := range result.ActiveTopics {
		fmt.Printf("\tname: %s\n", topicName)
	}
```