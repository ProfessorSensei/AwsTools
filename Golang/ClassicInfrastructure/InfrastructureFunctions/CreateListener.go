package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Create HTTP Listener
func CreateHTTPListener(tgrouparn, ldblncrarn string) {
	svc := elbv2.New(session.New())
	input := &elbv2.CreateListenerInput{
		DefaultActions: []*elbv2.Action{
			{
				TargetGroupArn: aws.String(tgrouparn),
				Type:           aws.String("forward"),
			},
		},
		LoadBalancerArn: aws.String(ldblncrarn),
		Port:            aws.Int64(80),
		Protocol:        aws.String("HTTP"),
	}

	result, err := svc.CreateListener(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elbv2.ErrCodeDuplicateListenerException:
				fmt.Println(elbv2.ErrCodeDuplicateListenerException, aerr.Error())
			case elbv2.ErrCodeTooManyListenersException:
				fmt.Println(elbv2.ErrCodeTooManyListenersException, aerr.Error())
			case elbv2.ErrCodeTooManyCertificatesException:
				fmt.Println(elbv2.ErrCodeTooManyCertificatesException, aerr.Error())
			case elbv2.ErrCodeLoadBalancerNotFoundException:
				fmt.Println(elbv2.ErrCodeLoadBalancerNotFoundException, aerr.Error())
			case elbv2.ErrCodeTargetGroupNotFoundException:
				fmt.Println(elbv2.ErrCodeTargetGroupNotFoundException, aerr.Error())
			case elbv2.ErrCodeTargetGroupAssociationLimitException:
				fmt.Println(elbv2.ErrCodeTargetGroupAssociationLimitException, aerr.Error())
			case elbv2.ErrCodeInvalidConfigurationRequestException:
				fmt.Println(elbv2.ErrCodeInvalidConfigurationRequestException, aerr.Error())
			case elbv2.ErrCodeIncompatibleProtocolsException:
				fmt.Println(elbv2.ErrCodeIncompatibleProtocolsException, aerr.Error())
			case elbv2.ErrCodeSSLPolicyNotFoundException:
				fmt.Println(elbv2.ErrCodeSSLPolicyNotFoundException, aerr.Error())
			case elbv2.ErrCodeCertificateNotFoundException:
				fmt.Println(elbv2.ErrCodeCertificateNotFoundException, aerr.Error())
			case elbv2.ErrCodeUnsupportedProtocolException:
				fmt.Println(elbv2.ErrCodeUnsupportedProtocolException, aerr.Error())
			case elbv2.ErrCodeTooManyRegistrationsForTargetIdException:
				fmt.Println(elbv2.ErrCodeTooManyRegistrationsForTargetIdException, aerr.Error())
			case elbv2.ErrCodeTooManyTargetsException:
				fmt.Println(elbv2.ErrCodeTooManyTargetsException, aerr.Error())
			case elbv2.ErrCodeTooManyActionsException:
				fmt.Println(elbv2.ErrCodeTooManyActionsException, aerr.Error())
			case elbv2.ErrCodeInvalidLoadBalancerActionException:
				fmt.Println(elbv2.ErrCodeInvalidLoadBalancerActionException, aerr.Error())
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
