/*
 * Project: AWS Examples
 * Author: Joshua Poland
 * Date: 2022-12-01
 */

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/go-yaml/yaml"
)

func main() {
	// Start AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	},
	)

	// New ResourceStruct Group Tagging API service
	tagsvc := resourcegroupstaggingapi.New(sess)

	// Get a list of all the resources, no parameters set
	result, err := tagsvc.GetResources(&resourcegroupstaggingapi.GetResourcesInput{})

	// Return error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Instantiate variables
	var allResources = make(map[string]interface{})
	var tempTag = make(map[string]string)
	var allTags []interface{}

	// Loop through resources
	for _, res := range result.ResourceTagMappingList {

		// Loop through resource tags
		for _, tag := range res.Tags {
			tempTag[*tag.Key] = *tag.Value
			allTags = append(allTags, tempTag)
		}

		// Save tags under ARN
		allResources[*res.ResourceARN] = allTags

	}

	// Start YAML creation
	resourcesYAML, err := yaml.Marshal(allResources)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err.Error())
	}

	// Output YAML
	fmt.Println(string(resourcesYAML))
}
