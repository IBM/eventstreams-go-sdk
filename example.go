// Golang Template

//

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/eventstreams-go-sdk/pkg/adminrestv1"
)

func main() {
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

	basicAuthenticator, _ := core.NewBasicAuthenticator("token", apiKey)

	serviceAPI, serviceErr := adminrestv1.NewAdminrestV1(&adminrestv1.AdminrestV1Options{

		URL:           URL,
		Authenticator: basicAuthenticator,
	})
	if serviceErr != nil {
		fmt.Printf("Error Creating Service")
		os.Exit(1)
	}

	// Always try to delete test-topic
	fmt.Printf("Delete Topic\n")
	_ = deleteTopic(serviceAPI)

	fmt.Printf("List Topics\n")
	err := listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Create Topic\n")
	err = createTopic(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Print Topic Details\n")
	err = topicDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("List Topics\n")
	err = listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Update Topic Details\n")
	err = updateDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Print Topic Details\n")
	err = topicDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	// Uncomment these examples if you are running against a Event Streams Mirrored Target Cluster.
	// fmt.Printf("List Active Mirroring Topics\n")
	// err = getMirroringActiveTopics(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Printf("Replace Mirroring Topics\n")
	// err = replaceMirroringTopicSelection(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Printf("List Mirroring Topic Selection\n")
	// err = ListMirroringTopics(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	fmt.Printf("Delete Topic\n")
	err = deleteTopic(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("List Topics\n")
	err = listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func listTopics(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the ListTopicsOptions model
	listTopicsOptionsModel := new(adminrestv1.ListTopicsOptions)

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
	return nil
}

func topicDetails(serviceAPI *adminrestv1.AdminrestV1) error {

	// Construct an instance of the GetTopicOptions model
	getTopicOptionsModel := new(adminrestv1.GetTopicOptions)
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

	return nil

}

func createTopic(serviceAPI *adminrestv1.AdminrestV1) error {
	// Set the retries policy.
	serviceAPI.EnableRetries(0, 0)

	// Construct an instance of the createTopicOptionsModel.
	createTopicOptionsModel := new(adminrestv1.CreateTopicOptions)
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

	return nil
}

func deleteTopic(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the DeleteTopicOptions model
	deleteTopicOptionsModel := new(adminrestv1.DeleteTopicOptions)
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
	return nil
}

func updateDetails(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the UpdateTopicOptions model
	updateTopicOptionsModel := new(adminrestv1.UpdateTopicOptions)
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

	return nil
}

//func replaceMirroringTopicSelection(serviceAPI *adminrestv1.AdminrestV1) error {
// 	// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
// 	replaceMirroringTopicSelectionOptionsModel := new(adminrestv1.ReplaceMirroringTopicSelectionOptions)
// 	replaceMirroringTopicSelectionOptionsModel.Includes = []string{"test-topic"}
// 	replaceMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

// 	// Enable retries.
// 	serviceAPI.EnableRetries(0, 0)

// 	// Invoke operation with valid options model (positive test)
// 	result, response, operationErr := serviceAPI.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
// 	if operationErr != nil {
// 		return fmt.Errorf("Error Replacing Mirroring Topics: %s\n", operationErr.Error())
// 	}

// 	// Check the result.
// 	if response.StatusCode != http.StatusAccepted {
// 		return fmt.Errorf("Error Replacing Mirroring Topics: status %d\n", response.StatusCode)
// 	}

// 	// Loop and print mirroring topics.
// 	for _, topicName := range result.Includes {
// 		fmt.Printf("\ttopic added: %s\n", topicName)
// 	}

// 	return nil
// }

// func ListMirroringTopics(serviceAPI *adminrestv1.AdminrestV1) error {
// 	// Construct an instance of the GetMirroringTopicSelectionOptions model
// 	getMirroringTopicSelectionOptionsModel := new(adminrestv1.GetMirroringTopicSelectionOptions)
// 	getMirroringTopicSelectionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

// 	// Enable retries.
// 	serviceAPI.EnableRetries(0, 0)

// 	// Call GetMirroringTopicSelection.
// 	result, response, operationErr := serviceAPI.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
// 	if operationErr != nil {
// 		return fmt.Errorf("Error Listing Mirroring Topics: %s\n", operationErr.Error())
// 	}

// 	// Check the result.
// 	if response.StatusCode != http.StatusAccepted {
// 		return fmt.Errorf("Error Listing Mirroring Topics: status %d\n", response.StatusCode)
// 	}

// 	// Loop and print mirroring topics.
// 	for _, topicName := range result.Includes {
// 		fmt.Printf("\tname: %s\n", topicName)
// 	}

// 	return nil
// }

// func getMirroringActiveTopics(serviceAPI *adminrestv1.AdminrestV1) error {
// 	// Construct an instance of the GetMirroringActiveTopicsOptions model
// 	getMirroringActiveTopicsOptionsModel := new(adminrestv1.GetMirroringActiveTopicsOptions)
// 	getMirroringActiveTopicsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

// 	// Call GetMirroringActiveTopics.
// 	result, response, operationErr := serviceAPI.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
// 	if operationErr != nil {
// 		return fmt.Errorf("Error Listing Active Mirroring Topics: %s\n", operationErr.Error())
// 	}

// 	// Check the result.
// 	if response.StatusCode != http.StatusAccepted {
// 		return fmt.Errorf("Error Listing Active Mirroring Topics: status %d\n", response.StatusCode)
// 	}

// 	// Loop and print mirroring topics.
// 	for _, topicName := range result.ActiveTopics {
// 		fmt.Printf("\tname: %s\n", topicName)
// 	}

// 	return nil
// }