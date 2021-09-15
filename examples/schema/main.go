package main

// Code Setup
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

// End Code Setup

func main() {
	fmt.Println("Schema Registry Go SDK")
	url := os.Getenv("KAFKA_ADMIN_URL")
	apiKey := os.Getenv("API_KEY")
	bearerToken := os.Getenv("BEARER_TOKEN")

	if url == "" {
		fmt.Println("please set env KAFKA_ADMIN_URL")
		os.Exit(1)
	}

	if apiKey == "" && bearerToken == "" {
		fmt.Println("please set either an API_KEY or a BEARER_TOKEN")
		os.Exit(1)
	}

	if apiKey != "" && bearerToken != "" {
		fmt.Println("please set either an API_KEY or a BEARER_TOKEN not both")
		os.Exit(1)
	}

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

	// Create Service
	esClient, err := schemaregistryv1.NewSchemaregistryV1(&schemaregistryv1.SchemaregistryV1Options{
		Authenticator: authenticator,
		URL:           url,
	})
	// End Create Service

	if err != nil {
		log.Printf("error occurred while configuring event streams schema registry: %q", err)
		os.Exit(1)
	}

	// Try to delete Schema before creating
	deleteSchema(esClient)

	// Create Schema
	err = createSchema(esClient)
	if err != nil {
		log.Printf("error occurred while creating the schema: %q", err)
		os.Exit(1)
	}
	log.Printf("Creating schema successful")

	// Get Schema
	if err = getLatestSchema(esClient); err != nil {
		log.Printf("error occurred while getting schema: %q", err)
		os.Exit(1)
	}
	log.Printf("fetching schema successful")

	// List Schema
	if err = listSchemas(esClient); err != nil {
		log.Printf("error occurred while listing schema: %q", err)
		os.Exit(1)
	}
	log.Printf("listing schema successful")

	// Update Schema
	if err = updateSchema(esClient); err != nil {
		log.Printf("error occurred while updating schema: %q", err)
		os.Exit(1)
	}
	log.Printf("updating schema successful")

	// Create Version
	err = createVersion(esClient)
	if err != nil {
		log.Printf("error occurred while creating version")
	}
	log.Printf("creating version successful")

	// Get Version
	err = getVersion(esClient)
	if err != nil {
		log.Printf("error occurred while getting version 2")
	}
	log.Printf("getting version successful")

	// List Versions
	err = listVersions(esClient)
	if err != nil {
		log.Printf("error occurred while listing versions: %q", err)
		os.Exit(1)
	}
	log.Printf("listing versions successful")

	// Delete Version
	if err = deleteVersion(esClient); err != nil {
		log.Printf("error occurred while deleting the version: %q", err)
		os.Exit(1)
	}
	log.Printf("deleting version number 2 successful")

	// Update Global Rule
	if err = updateGlobalRule(esClient); err != nil {
		log.Printf("error occurred while updating global rule: %q", err)
		os.Exit(1)
	}
	log.Printf("updating global rule successful")

	// Get Global Rule
	if err = getGlobalRule(esClient); err != nil {
		log.Printf("Error occurred while getting global config: %q", err)
		os.Exit(1)
	}
	log.Printf("fetching global rule successful")

	// Create Schema Rule
	if err = createSchemaRule(esClient); err != nil {
		log.Printf("error occurred while creating schema rule: %q", err)
		os.Exit(1)
	}
	log.Printf("creating schema rule successful")

	// Get Schema Rule
	if err = getSchemaRule(esClient); err != nil {
		log.Printf("error occurred while getting schema rule: %q", err)
		os.Exit(1)
	}
	log.Printf("fetching schema rule successful")

	// Update Schema Rule
	if err = updateSchemaRule(esClient); err != nil {
		log.Printf("error occurred while updating schema rule: %q", err)
		os.Exit(1)
	}
	log.Printf("updating schema rule successful")

	// Delete Schema Rule
	if err = deleteSchemaRule(esClient); err != nil {
		log.Printf("error occurred while deleting schema rule: %q", err)
		os.Exit(1)
	}
	log.Printf("deleting schema rule successful")

	// Delete Schema
	if err = deleteSchema(esClient); err != nil {
		log.Printf("error occurred while deleting the schema: %q", err)
		os.Exit(1)
	}
	log.Printf("deleting schema successful")
}

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end

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
} // func.end
