package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkNaclInVPC(s, b string) {
	// Struct for NACL
	var nacls []NetworkAcls
	// []string for nacl output
	var NACLDesiredOutput []string
	// session to describe nacl
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeNetworkAclsInput{}
	result, err := svc.DescribeNetworkAcls(input)
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
	fmt.Println("NetworkAcls:\n")
	if len(result.NetworkAcls) > 0 {
		bs, err := json.MarshalIndent(result.NetworkAcls, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(string(bs)), &nacls)
		for _, nl := range nacls {
			if b == nl.VpcID {
				NData, err := json.Marshal(nl)
				if err != nil {
					fmt.Println(err)
				}
				NACLDesiredOutput = append(NACLDesiredOutput, string(NData))
			}

		}

	}
	for _, nacsac := range NACLDesiredOutput {
		fmt.Println(nacsac)
	}
}
