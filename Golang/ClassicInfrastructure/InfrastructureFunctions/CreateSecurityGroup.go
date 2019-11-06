package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// create security group
func CreateSG(vpcid, egress, igress string) {
	svc := ec2.New(session.New())
	input := &ec2.CreateSecurityGroupInput{
		Description: aws.String("My security group"),
		GroupName:   aws.String("my-security-group"),
		VpcId:       aws.String(vpcid),
	}

	result, err := svc.CreateSecurityGroup(input)
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
	// pass sg id to CreateAuthorizeGroupIngress
}

// create Authorize Group Ingress Input
func CreateAuthorizeGroupIngress(dstsgId, incmsgId string) {
	svc := ec2.New(session.New())
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(incmsgId),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(80),
				IpProtocol: aws.String("tcp"),
				ToPort:     aws.Int64(80),
				UserIdGroupPairs: []*ec2.UserIdGroupPair{
					{
						Description: aws.String("HTTP access from other instances"),
						GroupId:     aws.String(dstsgId),
					},
				},
			},
		},
	}

	result, err := svc.AuthorizeSecurityGroupIngress(input)
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
