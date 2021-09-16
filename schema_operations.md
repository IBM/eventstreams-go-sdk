# Event Streams Schema Registry REST API
This REST API allows users of the
[IBM Event Streams service](https://cloud.ibm.com/docs/services/EventStreams/index.html)
to manage
[schema](#using-the-rest-api-to-manage-schema)
associated with an instance of the service. You can use this API to perform the following
operations:
  - [Create a schema](#creating-a-schema)
  - [Getting the latest version of schema](#getting-the-latest-version-of-schema)
  - [Listing schemas](#listing-schemas)
  - [Updating a schema](#updating-a-schema)
  - [Deleting schema](#deleting-schema)
  - [Creating a new version of schema](#creating-a-new-version-of-schema)
  - [Getting a specific version of schema](#getting-a-specific-version-of-schema)
  - [List all of the versions of a schema](#list-all-of-the-versions-of-a-schema)
  - [Deleting a version of schema](#deleting-a-version-of-schema)
  - [Getting the global rule](#getting-the-global-rule)
  - [Updating the global rule](#updating-the-global-rule)
  - [Creating schema rule](#creating-schema-rule)
  - [Getting schema rule](#getting-schema-rule)
  - [Updating schema rule](#updating-schema-rule)
  - [Deleting schema rule](#deleting-schema-rule)


## Access control
---

All requests support below authorization methods:
 * Basic authorization with user and password. (user is 'token', password is the API key from `ibmcloud resource service-keys` for the service instance.)
 * Bearer authorization with bearer token. (This token can be either API key or JWT token obtained from IAM upon login to IBM Cloud. Use `ibmcloud iam oauth-tokens` to retrieve the token after `ibmcloud login`)

##  Schema Registry API endpoint
---
Schema registry API endpoint is the `kafka_http_url` property in the service key for the service instance. This command can be used to retrieve this property.
```bash
$ ibmcloud resource service-key "${service_instance_key_name}" --output json | jq -r '.[]|.credentials.kafka_http_url'
```

## Using the REST API to manage schema
---

To run the example in [examples/schema](./examples/schema) :-

Compile the code.
```sh
cd examples/schema && go build -o example
```
Or simply 
```sh
make build
```

Set the required environment variables
```sh
# Set your API KEY (or a bearer token could be used by setting the BEARER_TOKEN environment variable instead, but not both)
export API_KEY="abc123456789"

# Set the Schema Registry Endpoint to point to your cluster.
export KAFKA_HTTP_URL="https://xyzclustername.svc01.region.eventstreams.cloud.ibm.com"

```

Run the example
```sh
./examples/schema/example 
```

## REST API 
---
The following sections explain how the REST API works with examples.

### Code Setup

```golang
// Code Setup
import (
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

// End Code Setup
```

### Authentication
---
Use one of the following methods to authenticate:

- To authenticate using Basic Auth:
  Place these values into the Authorization header of the HTTP request in the form Basic <credentials> 
  where <credentials> is the username and password joined by a single colon `:` base64 encoded. 
  For example:
  ```sh
  echo -n "token:<APIKEY>" | base64
  ```

- To authenticate using a bearer token:
  To obtain your token using the IBM Cloud CLI, first log in to IBM Cloud, then run the following command:
  ```
  ibmcloud iam oauth-tokens
  ```
  Place this token in the Authorization header of the HTTP request in the form Bearer. Both API key or JWT tokens are supported.

- To authenticate directly using the api_key:
  Place the key directly as the value of the X-Auth-Token HTTP header.

#### Example

Here's an example of how to create the authenticator using either an API key or a BEARER_TOKEN

```golang
	// Create Authenticator
	var authenticator core.Authenticator

	if apiKey != "" {
		var err error
		// Create an Basic IAM authenticator.
		authenticator, err = core.NewBasicAuthenticator("token", apiKey)
		if err != nil {
			fmt.Printf("failed to create new basic authenticator: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		var err error
		// Create an IAM Bearer Token authenticator.
		authenticator, err = core.NewBearerTokenAuthenticator(bearerToken)
		if err != nil {
			fmt.Printf("failed to create new bearer token authenticator: %s\n", err.Error())
			os.Exit(1)
		}
	}
	// End Authenticator
```

### Creating a client for the Schema Registry
---
Create a new client

```golang
// Create Service
	esClient, err := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
		Authenticator: authenticator,
		URL:           url,
	})
// End Create Service
```

### Representation of schema
---
`create` schema and `update` schema operations require the schema to be set in request body as a JSON document, for example:

```json
{
  "type": "record",
  "name": "Citizen",
  "fields": [
    {
      "name": "firstName",
      "type": "string"
    }
  ]
}
```
This is represented as `map[string]interface{}` type in Go.


### Creating a schema
---
To create a schema, the schema REST SDK issues a POST request to the `/artifacts` path.

The body of the request contains a JSON document, see [Representation of schema](#representation-of-schema)

Expected HTTP status codes:
- 200: Schema metadata is returned as JSON in the following format:
```json
{
  "id": "schema-id",
  "type": "AVRO",
  "version": 1,
  "createdOn": 1631518689408,
  "modifiedOn": 1631518689408,
  "globalId": 451
}
```

- 400: Invalid request JSON body
- 403: Not authorized to create a schema.
- 409: A schema with the specified schema ID already exists.

If the request to create a schema succeeds then HTTP status code 200 (OK) is returned with information about the newly created schema included in the response body.
If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example

```golang
func createSchema(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Form a schema of type map[string]interface{}
	schema := make(map[string]interface{})
	record := `{"type":"record","name":"Citizen","fields":[{"name":"firstName","type":"string"}]}`
	err := json.Unmarshal([]byte(record), &schema)
	if err != nil {
		return fmt.Errorf("error while unmarshalling schema, %s", err)
	}

	// Construct an instance of the createSchemaOptions
	createSchemaOptions := esClient.NewCreateSchemaOptions()
	createSchemaOptions.SetID("schema-id")
	createSchemaOptions.SetSchema(schema)

	// Create the schema
	schemaMetadata, response, operationErr := esClient.CreateSchema(createSchemaOptions)

	if operationErr != nil {
		return fmt.Errorf("error creating schema: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		operationErr = fmt.Errorf("creating schema failed with response: %v", response)
		return operationErr
	}

	// Print the schema metadata
	fmt.Printf("schema with ID schema-id created, following is the schema metadata:\n")
	fmt.Printf("\tcreated on: %s\n", time.Unix(*schemaMetadata.CreatedOn, 0).Format(time.UnixDate))
	fmt.Printf("\tglobal id: %d\n", *schemaMetadata.GlobalID)
	fmt.Printf("\tschema id: %s\n", *schemaMetadata.ID)
	fmt.Printf("\tmodified on: %s\n", time.Unix(*schemaMetadata.ModifiedOn, 0).Format(time.UnixDate))
	fmt.Printf("\ttype: %s\n", *schemaMetadata.Type)
	fmt.Printf("\tversion: %d\n", *schemaMetadata.Version)

	return nil
}
```

### Getting the latest version of a schema
---
To get the latest version of a schema, the schema REST SDK issues a GET request to the `/artifacts/ID` path, where `ID` is ID of the schema.

Expected return codes:
- 200: Schema of the specified schema ID is returned as JSON in the following format:
```json
{
  "type": "record",
  "name": "Citizen",
  "fields": [
    {
      "name": "firstName",
      "type": "string"
    }
  ]
}
```
- 400: Invalid request.
- 403: Not authorized to get schema.
- 404: Schema with the specified schema ID not found.

If the request to get a schema succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func getLatestSchema(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the GetLatestSchemaOptions model
	getLatestSchemaOptions := esClient.NewGetLatestSchemaOptions("schema-id")

	// Get Latest Schema
	schema, response, operationErr := esClient.GetLatestSchema(getLatestSchemaOptions)
	if operationErr != nil {
		return fmt.Errorf("error getting latest schema: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("getting schema failed with response: %v", response)
	}

	// Print the result
	fmt.Printf("schema with ID schema-id retrieved, following is the schema:\n")
	for k, v := range schema {
		fmt.Println("\t", k, ":", v)
	}

	return nil
}

```

### Listing schemas
---
To list the schemas, the schema REST SDK issues a GET request to the `/artifacts/` path.

Expected return codes:
- 200: List of comma separated schema IDs are returned as JSON in the following format:
```json
["schema-id"]
```
- 400: Invalid request.
- 403: Not authorized to list schema.

If the request to list schema succeeds, then HTTP status code 200 (OK) is returned with list of schema ids in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example
```golang
func listSchemas(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the ListSchemasOptions model
	listSchemasOptions := esClient.NewListSchemasOptions()

	// Get List of Schemas
	schemaList, response, operationErr := esClient.ListSchemas(listSchemasOptions)
	if operationErr != nil {
		return fmt.Errorf("error listing schema: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("listing schemas failed with response: %v", response)
	}

	fmt.Printf("list of schemas retrieved: %v\n", schemaList)

	return nil
}

```

### Updating a schema
---
To update a schema, the admin REST SDK issues a PUT request to the `/artifacts/id` path, where `ID` is the ID of the schema.

The body of the request contains a JSON document, see [Representation of schema](#representation-of-schema)

Expected return codes:
- 200: Schema metadata of specified schema ID is returned as JSON in the following format:
```json
{
  "id": "schema-id",
  "type": "AVRO",
  "version": 2,
  "createdOn": 1631519288967,
  "modifiedOn": 1631519433477,
  "globalId": 454
}
```
- 400: Invalid request.
- 403: Not authorized to update schema.
- 404: Schema with the specified schema ID not found.

If the request to update schema succeeds, then HTTP status code 200 (OK) is returned. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example

```golang
func updateSchema(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Form a schema of type map[string]interface{}
	schema := make(map[string]interface{})
	record := `{"type":"record","name":"Citizen","fields":[{"name":"first_name","type":"int"}]}`
	err := json.Unmarshal([]byte(record), &schema)
	if err != nil {
		return fmt.Errorf("error while unmarshalling schema, %s", err)
	}

	// Construct an instance of the UpdateSchemaOptions model
	updateSchemaOptions := esClient.NewUpdateSchemaOptions("schema-id")
	updateSchemaOptions.SetSchema(schema)

	// Update Schema
	schemaMetadata, response, operationErr := esClient.UpdateSchema(updateSchemaOptions)
	if operationErr != nil {
		return fmt.Errorf("error updating schema: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("updating schema failed with response: %v", response)
	}

	// Print the schema metadata
	fmt.Printf("schema with ID schema-id is updated, following is the schema metadata:\n")
	fmt.Printf("\tcreated on: %s\n", time.Unix(*schemaMetadata.CreatedOn, 0).Format(time.UnixDate))
	fmt.Printf("\tglobal id: %d\n", *schemaMetadata.GlobalID)
	fmt.Printf("\tschema id: %s\n", *schemaMetadata.ID)
	fmt.Printf("\tmodified on: %s\n", time.Unix(*schemaMetadata.ModifiedOn, 0).Format(time.UnixDate))
	fmt.Printf("\ttype: %s\n", *schemaMetadata.Type)
	fmt.Printf("\tversion: %d\n", *schemaMetadata.Version)

	return nil
}

```

### Deleting schema
---
To delete a schema, the schema REST SDK issues a DELETE request to the `/artifacts/ID` path, where `ID` is the ID of the schema.

Expected HTTP status codes:
- 204: Schema deletion of specified schema ID is completed and no content is returned.
- 400: Invalid request.
- 403: Not authorized to delete a schema.
- 404: Schema with the specified schema ID not found.

If the request to delete schema succeeds then HTTP status code 204 (No content) is returned. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example

```golang
func deleteSchema(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the DeleteSchemaOptions
	deleteSchemaOptions := esClient.NewDeleteSchemaOptions("schema-id")

	// Delete schema
	response, operationErr := esClient.DeleteSchema(deleteSchemaOptions)
	if operationErr != nil {
		return fmt.Errorf("error deleting schema: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("deleting schema failed with response: %v", response)
	}

	fmt.Printf("schema with ID schema-id deleted\n")

	return nil
}
```

### Creating a new version of schema
---
To create a new version of a schema, the schema REST SDK issues a POST request to the `/artifacts/ID/versions` path, where `ID` is the ID of schema for which new version is to be created.

The body of the request contains a JSON document, see [Representation of schema](#representation-of-schema)

Expected HTTP status codes:
- 200: Schema metadata of the new version is returned as JSON in the following format:
```json
{
  "id": "schema-id",
  "type": "AVRO",
  "version": 3,
  "createdOn": 1631519288967,
  "modifiedOn": 1631519524248,
  "globalId": 455
}
```
- 400: Invalid request.
- 403: Not authorized to create a new version of schema.
- 409: Schema with the specified schema ID not found.

If the request to create a new version of schema succeeds then HTTP status code 200 (OK) is returned with information about the newly created schema included in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example

```golang
func createVersion(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Form a schema of type map[string]interface{}
	schema := make(map[string]interface{})
	record := `{"type":"record","name":"Citizen","fields":[{"name":"first_name","type":"string"},{"name":"last_name","type":"string"}]}`
	err := json.Unmarshal([]byte(record), &schema)
	if err != nil {
		return fmt.Errorf("error while unmarshalling schema, %s", err)
	}

	// Construct an instance of the CreateVersionOptions
	createVersionOptions := esClient.NewCreateVersionOptions("schema-id")
	createVersionOptions.SetSchema(schema)

	// Create the new version of schema
	schemaMetadata, response, operationErr := esClient.CreateVersion(createVersionOptions)
	if operationErr != nil {
		return fmt.Errorf("error creating new version of schema: %s", operationErr.Error())
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("creating new version of schema with ID schema-id failed with response: %v", response)
	}

	// Print the schema metadata
	fmt.Printf("new version of schema with ID schema-id is created, following is the schema metadata\n")
	fmt.Printf("\tcreated on: %s\n", time.Unix(*schemaMetadata.CreatedOn, 0).Format(time.UnixDate))
	fmt.Printf("\tglobal id: %d\n", *schemaMetadata.GlobalID)
	fmt.Printf("\tschema id: %s\n", *schemaMetadata.ID)
	fmt.Printf("\tmodified on: %s\n", time.Unix(*schemaMetadata.ModifiedOn, 0).Format(time.UnixDate))
	fmt.Printf("\ttype: %s\n", *schemaMetadata.Type)
	fmt.Printf("\tversion: %d\n", *schemaMetadata.Version)

	return nil
}
```

### Getting a specific version of a schema
To get a specific version of a schema, the schema REST SDK issues a GET request to the `/artifacts/ID/versions/VERSION` path, `VERSION` is the specific version number of a schema with ID as `ID`.

Expected return codes:
- 200: Schema of the specified ID is returned as JSON in the following format:
```json
{
  "type": "record",
  "name": "Citizen",
  "fields": [
    {
      "name": "firstName",
      "type": "string"
    },
    {
      "name": "lastName",
      "type": "string"
    }
  ]
}
```
- 400: Invalid request.
- 403: Not authorized to get specified version of schema, or the schema itself.
- 404: Either the schema with specified ID or the version of specified version number of that schema not found.

If the request to get a version of schema succeeds, then HTTP status code 200 (OK) is returned with specified version of schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func getVersion(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the GetVersionOptions model
	getVersionOptions := esClient.NewGetVersionOptions("schema-id", 2)

	// Get specific version of Schema
	schema, response, operationErr := esClient.GetVersion(getVersionOptions)
	if operationErr != nil {
		return fmt.Errorf("error getting schema with ID schema-id of version 2: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("getting schema with ID schema-id of version 2 failed with response: %v", response)

	}

	// Print the schema
	fmt.Printf("version 2 of schema with ID schema-id fetched, following is the schema:\n")
	for k, v := range schema {
		fmt.Println("\t", k, ":", v)
	}

	return nil
}

```

### List all of the versions of a schema
---
To list all of the versions of schema, the schema REST SDK issues a GET request to the `/artifacts/ID/versions` path, where `ID` is the ID of a schema.

Expected return codes:
- 200: A comma separated list of all version numbers of schema with specified schema ID is returned in following format:
```json
[1,2,3]
```
- 400: Invalid request.
- 403: Not authorized to list all versions of schema.
- 404: Schema with specified ID not found.

If the request to list all the versions of scheema succeeds, then HTTP status code 200 (OK) is returned with list of versions in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example
```golang
func listVersions(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the ListVersionsOptions model
	listVersionsOptions := esClient.NewListVersionsOptions("schema-id")

	// Get version list of specific schema
	versionList, response, operationErr := esClient.ListVersions(listVersionsOptions)
	if operationErr != nil {
		return fmt.Errorf("error listing all the versions of schema with ID schema-id: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("listing all the versions of schema with ID schema-id failed with response: %v", response)
	}

	fmt.Printf("list of all the versions of schema with ID schema-id retrieved: %v\n", versionList)
	return nil
}

```

### Deleting a version of a schema:
---
---
To delete a specific version of a schema, the schema REST SDK issues a DELETE request to the `/artifacts/ID/versions/VERSION` path, where the `VERSION` is the version of the schema with `ID`.

Expected HTTP status codes:
- 204: Schema deletion completed and no content is returned.
- 400: Invalid request.
- 403: Not authorized to delete a version of a schema.
- 409: A schema with the specified schema ID already exists.

If the request to delete a specific version of schema succeeds then HTTP status code 204 (No content) is returned. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.

#### Example

```golang
func deleteVersion(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the DeleteSchemaOptions
	deleteVersionOptions := esClient.NewDeleteVersionOptions("schema-id", 2)

	// Delete a version of schema
	response, operationErr := esClient.DeleteVersion(deleteVersionOptions)
	if operationErr != nil {
		return fmt.Errorf("error deleting version 2 of schema with ID schema-id: %s", operationErr.Error())
	}

	// Check the result
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("deleting version 2 of schema with ID schema-id failed with response: %v", response)
	}

	fmt.Printf("version 2 of schema with ID schema-id deleted\n")

	return nil
}

```

### Getting the global rule
---
To get the global rule, the schema REST SDK issues a GET request to the `/rules/RULE` path, where `RULE` is the type of rule to be retrieved. Currently `RULE` can only have value `COMPATIBILITY`.

Expected return codes:
- 200: Global rule is returned as JSON in the following format:
```json
{
  "type": "COMPATIBILITY",
  "config": "NONE"
}
```
- 400: Invalid request.
- 403: Not authorized to get rule.

If the request to get the global rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func getGlobalRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the GetGlobalRuleOptions model
	getGlobalRuleOptions := esClient.NewGetGlobalRuleOptions(
		schemaregistryv1.GetGlobalRuleOptionsRuleCompatibilityConst)

	// Get the global rule
	globalRule, response, operationErr := esClient.GetGlobalRule(getGlobalRuleOptions)
	if operationErr != nil {
		return fmt.Errorf("error getting the global rule: %s", operationErr)
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("getting the global rule failed with response: %v", response)
	}

	fmt.Printf("fetched the global rule: %v\n", *globalRule.Config)

	return nil
}
```

### Updating the global rule
---
To update the global rule, the schema REST SDK issues a PUT request to the `/rules/RULE` path, where `RULE` is the type of rule to be retrieved. Currently `RULE` can only have value `COMPATIBILITY`.

The request body is in JSON, for example:

```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```

Expected return codes:
- 200: Updated global rule is returned in JSON in the following format:
```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```
- 400: Invalid request.
- 403: Not authorized to update global rule.

If the request to update the global rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func updateGlobalRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the UpdateGlobalRuleOptions
	updateGlobalRuleOptions := esClient.NewUpdateGlobalRuleOptions(
		schemaregistryv1.UpdateGlobalRuleOptionsTypeCompatibilityConst,
		schemaregistryv1.UpdateGlobalRuleOptionsRuleCompatibilityConst,
		schemaregistryv1.UpdateGlobalRuleOptionsConfigNoneConst)

	// Update the global rule
	globalRule, response, operationErr := esClient.UpdateGlobalRule(updateGlobalRuleOptions)
	if operationErr != nil {
		return fmt.Errorf("error updating the global rule: %s", operationErr)
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("updating the global rule failed with response: %v", response)
	}

	fmt.Printf("updated the global rule: %s\n", *globalRule.Config)

	return nil
}
```

### Creating a per-schema rule
---
To create a rule per schema, the schema REST SDK issues a POST request to the `/artifacts/ID/rules` path, where `ID` is the ID of the schema.

The request body is in JSON, for example:

```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```

Expected return codes:
- 200: Created rule of the schema with specified schema ID is returned in JSON in the following format:
```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```

- 400: Invalid request.
- 403: Not authorized to update schema rule.
- 404: Schema with specified ID not found.

If the request to create the schema rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func createSchemaRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the CreateSchemaRuleOptions model
	createSchemaRuleOptions := esClient.NewCreateSchemaRuleOptions(
		"schema-id",
		schemaregistryv1.CreateSchemaRuleOptionsTypeCompatibilityConst,
		schemaregistryv1.CreateSchemaRuleOptionsConfigForwardConst)

	// Create schema rule
	schemaRule, response, operationErr := esClient.CreateSchemaRule(createSchemaRuleOptions)
	if operationErr != nil {
		return fmt.Errorf("error creating rule for schema with ID schema-id: %s", operationErr)
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("creating rule for schema with ID schema-id failed with response: %v", response)
	}

	fmt.Printf("created a rule for schema with ID schema-id: %s\n", *schemaRule.Config)

	return nil
}
```

### Getting a per-schema rule
---
To get the rule of a schema, the schema REST SDK issues a GET request to the `/artifacts/ID/rules/RULE` path, where  `RULE` is the type of the rule to retrieve for schema withe ID as `ID`. Currently `RULE` can only have value `COMPATIBILITY`.

Expected return codes:
- 200: The rule of the schema with specified ID is returned as JSON in the following format:
```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```
- 400: Invalid request.
- 403: Not authorized to update schema rule.
- 404: Schema with specified ID not found.

If the request to get the schema rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func getSchemaRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the GetSchemaRuleOptions model
	getSchemaRuleOptions := esClient.NewGetSchemaRuleOptions(
		"schema-id",
		schemaregistryv1.GetSchemaRuleOptionsRuleCompatibilityConst)

	// Get schema rule
	schemaRule, response, operationErr := esClient.GetSchemaRule(getSchemaRuleOptions)
	if operationErr != nil {
		return fmt.Errorf("error getting rule for schema with ID schema-id: %s", operationErr)
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("getting a rule for schema with ID schema-id failed with response: %v", response)
	}

	fmt.Printf("retrieved the rule for schema with ID schema-id: %s\n", *schemaRule.Config)
	return nil
}
```

### Updating a per-schema rule
---
To update the rule of a schema, the schema REST SDK issues a PUT request to the `/artifacts/ID/rules/RULE` path, where `RULE` is the type of rule to update for a schema with ID as `ID`. Currently RULE can only have value `COMPATIBILITY`.

The request body is in JSON, for example:

```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```

Expected return codes:
- 200: Updated rule of the schema with given ID is returned as JSON in following format:
```json
{
  "type": "COMPATIBILITY",
  "config": "BACKWARD"
}
```
- 400: Invalid request.
- 403: Not authorized to update schema rule.
- 404: Schema with specified ID not found.

If the request to update the schema rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func updateSchemaRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the UpdateSchemaRuleOptions model
	updateSchemaRuleOptions := esClient.NewUpdateSchemaRuleOptions("schema-id",
		schemaregistryv1.UpdateSchemaRuleOptionsRuleCompatibilityConst,
		schemaregistryv1.UpdateSchemaRuleOptionsRuleCompatibilityConst,
		schemaregistryv1.UpdateSchemaRuleOptionsConfigBackwardConst)

	// Update schema rule
	schemaRule, response, operationErr := esClient.UpdateSchemaRule(updateSchemaRuleOptions)
	if operationErr != nil {
		return fmt.Errorf("error updating rule for schema with ID schema-id: %s", operationErr)
	}

	// Check the result
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("updating rule for schema with ID schema-id failed with response: %v", response)
	}

	fmt.Printf("updated the rule for schema with ID schema-id: %s\n", *schemaRule.Config)

	return nil
}
```

### Deleting a per-schema rule
---
To delete the rule of a schema, the schema REST SDK issues a DELETE request to the `/artifacts/ID/rules/RULE` path, where `RULE` is the type of rule to retrieve for a schema with ID as `ID`. Currently RULE can only have value `COMPATIBILITY`.

Expected return codes:
- 204: Deleting schema rule completed and no content is returned.
- 400: Invalid request.
- 403: Not authorized to update schema rule.
- 404: Schema with specified ID not found.

If the request to delete the schema rule succeeds, then HTTP status code 200 (OK) is returned with latest schema in the response body. If the operation fails with HTTP status code of 400 (Bad request), additional information about the failure is returned as the body of the response.


#### Example
```golang
func deleteSchemaRule(esClient *schemaregistryv1.SchemaregistryV1) error {
	// Construct an instance of the DeleteSchemaRuleOptions
	deleteSchemaRuleOptions := esClient.NewDeleteSchemaRuleOptions("schema-id",
		schemaregistryv1.DeleteSchemaRuleOptionsRuleCompatibilityConst)

	// Delete schema rule
	response, err := esClient.DeleteSchemaRule(deleteSchemaRuleOptions)
	if err != nil {
		return err
	}

	// Check the result
	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("updating the rule for schema with ID schema-id failed with response: %v", response)
	}

	return nil
}
```