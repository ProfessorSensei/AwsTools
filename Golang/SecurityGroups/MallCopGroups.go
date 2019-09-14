package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	for _, region := range pawsRegionList() {
		pcheckSecurityGroups(region)
	}

	fmt.Println("About to exit")
}

func pawsRegionList() []string {
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

func pcheckSecurityGroups(s string) {
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeSecurityGroupsInput{}

	result, err := svc.DescribeSecurityGroups(input)
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
	fmt.Println("Region:", s)
	for _, sb := range result.SecurityGroups {
		fmt.Println(aws.StringValue(sb.GroupName))
	}

}
