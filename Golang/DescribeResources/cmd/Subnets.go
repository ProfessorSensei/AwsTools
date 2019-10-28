package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkSubNetInVPC(s, b string) {
	// Struct for Subnets
	var subN []Subnet
	// string for subnet output
	var SBDesiredOutput []string
	// session for Describing Subnets
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeSubnetsInput{}
	result, err := svc.DescribeSubnets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println("Subnets:\n")
	if len(result.Subnets) > 0 {
		bs, err := json.MarshalIndent(result.Subnets, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(string(bs)), &subN)
		// Now range over and check VPC ID
		for _, sb := range subN {
			if b == sb.VpcID {
				// marshal output
				SBData, err := json.Marshal(sb)
				if err != nil {
					fmt.Println(err)
				}
				SBDesiredOutput = append(SBDesiredOutput, string(SBData))
			}
		}
		fmt.Println(SBDesiredOutput)

	}
}
