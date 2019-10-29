package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkInternetGatewayVPC(s, b string) {
	// Struct for InternetGateways
	var IGateway []InterNetGateways
	// string for returned value
	var IGODesiredOutput []string
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeInternetGatewaysInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("attachment.vpc-id"),
				Values: []*string{
					aws.String(b),
				},
			},
		},
	}

	result, err := svc.DescribeInternetGateways(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		// return
	}
	fmt.Println("Internet Gateways:\n")
	if len(result.InternetGateways) > 0 {
		// marshall output
		bs, err := json.MarshalIndent(result.InternetGateways, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		// unmarshall data to filter
		json.Unmarshal([]byte(string(bs)), &IGateway)
		// if the vpc id matches, return json string
		for _, IG := range IGateway {
			// range to get VPC Id
			ts, err := json.Marshal(IG)
			if err != nil {
				fmt.Println(err)
			}
			IGODesiredOutput = append(IGODesiredOutput, string(ts))
		}
		fmt.Println(IGODesiredOutput)
	} else {
		fmt.Println("There are 0 Internet Gateways in this VPC")
		IGODesiredOutput = append(IGODesiredOutput, "")
	}
}
