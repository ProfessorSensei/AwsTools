package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// this function is the function that calls all of the search functions
// Call the functions and unmarshal the data
func stringOnlyPassed(s, r string) {
	checkInstancesInVPC(s, r)
	checkNaclInVPC(s, r)
	checkRTInVPC(s, r)
	checkSGInVPC(s, r)
	checkSubNetInVPC(s, r)
	checkNatGatewayVPC(s, r)
	checkInternetGatewayVPC(s, r)
	NWIGather(s, r)
}

// function for when the VPC ID argument is passed and approved
func flagVPCs(s, b string) {
	var ivpcid string
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeVpcsInput{}
	result, err := svc.DescribeVpcs(input)
	if err != nil {
		fmt.Println(err)
	}
	if len(result.Vpcs) == 1 {
		for _, vpcId1 := range result.Vpcs {
			ivpcid = aws.StringValue(vpcId1.VpcId)
		}
		if ivpcid == b {
			fmt.Printf("Found VPC Id:'%v' in Region:'%v'\n", b, s)
			fmt.Printf("\nVPC '%v' has the following resources:\n", b)
			stringOnlyPassed(s, ivpcid)
		}
		// statement for if there are more than one vpc in a region
	} else if len(result.Vpcs) > 1 {
		var multivpc []string
		for _, vpcId1 := range result.Vpcs {
			multivpc = append(multivpc, aws.StringValue(vpcId1.VpcId))
		}
		for _, vpc := range multivpc {
			if vpc == b {
				fmt.Printf("Found VPC Id:'%v' in Region:'%v'\n", b, s)
				fmt.Printf("\nVPC '%v' has the following resources:\n", b)
				stringOnlyPassed(s, vpc)
			}
		}
	}

}
