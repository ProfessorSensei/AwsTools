package InfrastructureFunctions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

// Create ALB (load balancer v2) - internet facing
func CreateIFALB(subnid string) {
	svc := elbv2.New(session.New())
	input := &elbv2.CreateLoadBalancerInput{
		Name: aws.String("my-load-balancer"),
		Subnets: []*string{
			aws.String(subnid),
			// aws.String("subnet-8360a9e7"),
		},
	}

	result, err := svc.CreateLoadBalancer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elbv2.ErrCodeDuplicateLoadBalancerNameException:
				fmt.Println(elbv2.ErrCodeDuplicateLoadBalancerNameException, aerr.Error())
			case elbv2.ErrCodeTooManyLoadBalancersException:
				fmt.Println(elbv2.ErrCodeTooManyLoadBalancersException, aerr.Error())
			case elbv2.ErrCodeInvalidConfigurationRequestException:
				fmt.Println(elbv2.ErrCodeInvalidConfigurationRequestException, aerr.Error())
			case elbv2.ErrCodeSubnetNotFoundException:
				fmt.Println(elbv2.ErrCodeSubnetNotFoundException, aerr.Error())
			case elbv2.ErrCodeInvalidSubnetException:
				fmt.Println(elbv2.ErrCodeInvalidSubnetException, aerr.Error())
			case elbv2.ErrCodeInvalidSecurityGroupException:
				fmt.Println(elbv2.ErrCodeInvalidSecurityGroupException, aerr.Error())
			case elbv2.ErrCodeInvalidSchemeException:
				fmt.Println(elbv2.ErrCodeInvalidSchemeException, aerr.Error())
			case elbv2.ErrCodeTooManyTagsException:
				fmt.Println(elbv2.ErrCodeTooManyTagsException, aerr.Error())
			case elbv2.ErrCodeDuplicateTagKeysException:
				fmt.Println(elbv2.ErrCodeDuplicateTagKeysException, aerr.Error())
			case elbv2.ErrCodeResourceInUseException:
				fmt.Println(elbv2.ErrCodeResourceInUseException, aerr.Error())
			case elbv2.ErrCodeAllocationIdNotFoundException:
				fmt.Println(elbv2.ErrCodeAllocationIdNotFoundException, aerr.Error())
			case elbv2.ErrCodeAvailabilityZoneNotSupportedException:
				fmt.Println(elbv2.ErrCodeAvailabilityZoneNotSupportedException, aerr.Error())
			case elbv2.ErrCodeOperationNotPermittedException:
				fmt.Println(elbv2.ErrCodeOperationNotPermittedException, aerr.Error())
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

// create internal facing ALB
func CreateINALB() {
	svc := elbv2.New(session.New())
	input := &elbv2.CreateLoadBalancerInput{
		Name:   aws.String("my-internal-load-balancer"),
		Scheme: aws.String("internal"),
		Subnets: []*string{
			aws.String("subnet-b7d581c0"),
			aws.String("subnet-8360a9e7"),
		},
	}

	result, err := svc.CreateLoadBalancer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elbv2.ErrCodeDuplicateLoadBalancerNameException:
				fmt.Println(elbv2.ErrCodeDuplicateLoadBalancerNameException, aerr.Error())
			case elbv2.ErrCodeTooManyLoadBalancersException:
				fmt.Println(elbv2.ErrCodeTooManyLoadBalancersException, aerr.Error())
			case elbv2.ErrCodeInvalidConfigurationRequestException:
				fmt.Println(elbv2.ErrCodeInvalidConfigurationRequestException, aerr.Error())
			case elbv2.ErrCodeSubnetNotFoundException:
				fmt.Println(elbv2.ErrCodeSubnetNotFoundException, aerr.Error())
			case elbv2.ErrCodeInvalidSubnetException:
				fmt.Println(elbv2.ErrCodeInvalidSubnetException, aerr.Error())
			case elbv2.ErrCodeInvalidSecurityGroupException:
				fmt.Println(elbv2.ErrCodeInvalidSecurityGroupException, aerr.Error())
			case elbv2.ErrCodeInvalidSchemeException:
				fmt.Println(elbv2.ErrCodeInvalidSchemeException, aerr.Error())
			case elbv2.ErrCodeTooManyTagsException:
				fmt.Println(elbv2.ErrCodeTooManyTagsException, aerr.Error())
			case elbv2.ErrCodeDuplicateTagKeysException:
				fmt.Println(elbv2.ErrCodeDuplicateTagKeysException, aerr.Error())
			case elbv2.ErrCodeResourceInUseException:
				fmt.Println(elbv2.ErrCodeResourceInUseException, aerr.Error())
			case elbv2.ErrCodeAllocationIdNotFoundException:
				fmt.Println(elbv2.ErrCodeAllocationIdNotFoundException, aerr.Error())
			case elbv2.ErrCodeAvailabilityZoneNotSupportedException:
				fmt.Println(elbv2.ErrCodeAvailabilityZoneNotSupportedException, aerr.Error())
			case elbv2.ErrCodeOperationNotPermittedException:
				fmt.Println(elbv2.ErrCodeOperationNotPermittedException, aerr.Error())
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
