package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Create target group
func CreateTargetGroup(nme, prtcl, vpcid string, prt int64) {
	svc := elbv2.New(session.New())
	input := &elbv2.CreateTargetGroupInput{
		Name:     aws.String(nme),
		Port:     aws.Int64(prt),
		Protocol: aws.String(prtcl),
		VpcId:    aws.String(vpcid),
	}

	result, err := svc.CreateTargetGroup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elbv2.ErrCodeDuplicateTargetGroupNameException:
				fmt.Println(elbv2.ErrCodeDuplicateTargetGroupNameException, aerr.Error())
			case elbv2.ErrCodeTooManyTargetGroupsException:
				fmt.Println(elbv2.ErrCodeTooManyTargetGroupsException, aerr.Error())
			case elbv2.ErrCodeInvalidConfigurationRequestException:
				fmt.Println(elbv2.ErrCodeInvalidConfigurationRequestException, aerr.Error())
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
