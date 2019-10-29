package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkNatGatewayVPC(s, b string) {
	// Struct for NateGateways
	var nateGWay []NGateways
	var NGDesiredOutput []string
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeNatGatewaysInput{
		Filter: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String(b),
				},
			},
		},
	}

	result, err := svc.DescribeNatGateways(input)
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
	fmt.Println("NatGateways:\n")
	if len(result.NatGateways) > 0 {
		bs, err := json.MarshalIndent(result.NatGateways, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(bs), &nateGWay)
		for _, ngway := range nateGWay {
			NGData, err := json.Marshal(ngway)
			if err != nil {
				fmt.Println(err)
			}
			// resources per VPC
			NGDesiredOutput = append(NGDesiredOutput, string(NGData))
		}
		fmt.Println(NGDesiredOutput)

	} else {
		fmt.Println("0 NatGateways\n")
	}
}
