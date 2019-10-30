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

var PNme, pType, awsRegion string

func main() {
	// User input placement groupname
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Placement Group name:")
	PNme, _ = reader.ReadString('\n')
	// remove new line
	PNme = strings.Replace(PNme, "\n", "", -1)
	// User input cluster type
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter Placement Group strategy 'cluster|spread':")
	pType, _ = reader2.ReadString('\n')
	// remove new line
	pType = strings.Replace(pType, "\n", "", -1)
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
	svc := ec2.New(sess)
	input := &ec2.CreatePlacementGroupInput{
		GroupName: aws.String(PNme),
		Strategy:  aws.String(pType),
	}

	// Create placement group name and type provided by user
	result, err := svc.CreatePlacementGroup(input)
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
