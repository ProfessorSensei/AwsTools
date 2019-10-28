package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func describeVPCs(s string) {
	var ivpcid string
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeVpcsInput{}
	result, err := svc.DescribeVpcs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	if len(result.Vpcs) == 1 {
		for _, vpcId1 := range result.Vpcs {
			vpcIds = append(vpcIds, aws.StringValue(vpcId1.VpcId))
			ivpcid = aws.StringValue(vpcId1.VpcId)
			fmt.Printf("\nVPC '%v' in Region: '%v' has the following resources:\n", ivpcid, s)
		}
		fmt.Println(result)
		stringOnlyPassed(s, ivpcid)
		// else statement for more than one vpc in a region
	} else if len(result.Vpcs) > 1 {
		//[]string to hold VPC Id's
		fmt.Printf("\nRegion: '%v' has '%v' VPC's", s, len(result.Vpcs))
		var multivpc []string
		for _, vpcId1 := range result.Vpcs {
			multivpc = append(multivpc, aws.StringValue(vpcId1.VpcId))
		}
		for _, vpc := range multivpc {
			fmt.Println(result)
			fmt.Printf("\nVPC '%v' in Region: '%v' has the following resources:\n", vpc, s)
			stringOnlyPassed(s, vpc)
		}
	}

}
