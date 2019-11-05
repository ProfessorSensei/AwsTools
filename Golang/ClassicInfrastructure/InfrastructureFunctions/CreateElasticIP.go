package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// allocate IP address
func AllocateIPA(vpc string) {
	svc := ec2.New(session.New())
	input := &ec2.AllocateAddressInput{
		Domain: aws.String("vpc"),
	}

	result, err := svc.AllocateAddress(input)
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
	// return allocation id

	fmt.Println(result)
}
