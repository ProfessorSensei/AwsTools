package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Create route table
func CreateRT() {
	svc := ec2.New(session.New())
	input := &ec2.CreateRouteTableInput{
		VpcId: aws.String("vpc-a01106c2"),
	}

	result, err := svc.CreateRouteTable(input)
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
	// return route table id

	fmt.Println(result)
}
