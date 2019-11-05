package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// create subnet needs VPC ID
func CreateSubnet(vpcid, cdr string) string {
	svc := ec2.New(session.New())
	input := &ec2.CreateSubnetInput{
		CidrBlock: aws.String(cdr),
		VpcId:     aws.String(vpcid),
	}

	result, err := svc.CreateSubnet(input)
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

	// return subnet id
	fmt.Println(result)
}
