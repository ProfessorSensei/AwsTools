package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func NWIGather(s, b string) {
	// Struct for Network Interfaces
	var NIFace []NetWorkInterFaces
	// []string for returned value
	var NIFDesiredOutput []string
	// EC2 session to gather NetworkInterfaces
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeNetworkInterfacesInput{}
	result, err := svc.DescribeNetworkInterfaces(input)
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
	fmt.Println("NetworkInterfaces:\n")
	if len(result.NetworkInterfaces) > 0 {
		bs, err := json.MarshalIndent(result.NetworkInterfaces, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(bs), &NIFace)
		// if the VPC Id matches, return json string
		for _, NI := range NIFace {
			if b == NI.VpcID {
				ts, err := json.Marshal(NI)
				if err != nil {
					fmt.Println(err)
				}
				NIFDesiredOutput = append(NIFDesiredOutput, string(ts))

			}
		}
		fmt.Println(NIFDesiredOutput)
	} else {
		fmt.Println("There are 0 NetworkInterfaces in this region")
	}
}
