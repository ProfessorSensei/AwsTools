package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func regionalVPCIds() []string {
	return vpcIdList(awsRegionList())
}

// function to gather Region names
func awsRegionList() []string {
	var awsRegions []string
	// create ec2 session
	ec2session := ec2.New(session.New(), &aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	})
	// describe regions
	regions, err := ec2session.DescribeRegions(&ec2.DescribeRegionsInput{})

	if err != nil {
		panic(err)
	}

	for _, region := range regions.Regions {
		awsRegions = append(awsRegions, *region.RegionName)
	}

	return awsRegions
}

// separate return []string from awsRegionList and call this function
func vpcIdList(s []string) []string {
	var vpcIds []string
	for _, noS := range s {
		svc := ec2.New(session.New(), aws.NewConfig().WithRegion(noS))
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
		for _, vpcId := range result.Vpcs {
			vpcIds = append(vpcIds, aws.StringValue(vpcId.VpcId))
		}
	}
	return vpcIds
}

func AccountContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
