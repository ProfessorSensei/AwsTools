package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// finish functions needed then return information needed
// to pass to other functions
// create VPC √
// Create Internet gateway √
// create subnets √
// create natgateway √
// allocation id √
// create route table√
// create route table association ? may be taken care of in route
// create route √

// create target group √
// create ALB and security group open port 80
// target group only allows traffic from ALB security group
// Create access logs bucket and enable elb (ALB in this case) access logs to bucket
// Create ASG

func main() {
	// User input ImageId

	// User input region

	// User input instance type

	// User input security group

	// create launch config to be used by autoscaler
	fmt.Println("Hello placeholder")
	createASGInput()
}

func createASGInput() *autoscaling.CreateLaunchConfigurationInput {
	svc := autoscaling.New(session.New())
	input := &autoscaling.CreateLaunchConfigurationInput{
		IamInstanceProfile:      aws.String("my-iam-role"),
		ImageId:                 aws.String("ami-12345678"),
		InstanceType:            aws.String("m3.medium"),
		LaunchConfigurationName: aws.String("my-launch-config"),
		SecurityGroups: []*string{
			aws.String("sg-eb2af88e"),
		},
	}

	result, err := svc.CreateLaunchConfiguration(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeAlreadyExistsFault:
				fmt.Println(autoscaling.ErrCodeAlreadyExistsFault, aerr.Error())
			case autoscaling.ErrCodeLimitExceededFault:
				fmt.Println(autoscaling.ErrCodeLimitExceededFault, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	return input

	fmt.Println(result)
}
