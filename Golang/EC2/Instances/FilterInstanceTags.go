package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// user input values, tag name, value, and region of instance
var tagName, tagValue, awsRegion string

// Ask user for tag name, value, and instance region, then describe the instance
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter tag name:")
	tagName, _ = reader.ReadString('\n')
	// remove new line
	tagName = strings.Replace(tagName, "\n", "", -1)
	// UserName
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter tag value:")
	tagValue, _ = reader2.ReadString('\n')
	// remove new line
	tagValue = strings.Replace(tagValue, "\n", "", -1)
	// Region
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region the instance is located:")
	awsRegion, _ = reader3.ReadString('\n')
	// remove new line
	awsRegion = strings.Replace(awsRegion, "\n", "", -1)

	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)

	filterTagName := fmt.Sprintf("tag:%v", tagName)
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(filterTagName),
				Values: []*string{
					aws.String(tagValue),
				},
			},
		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}