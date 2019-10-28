package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkSGInVPC(s, b string) {
	// Struct for Security Groups
	var SecG []SecurityGroups
	// string for SG output
	var SGDesiredOutPut []string
	// session for describing Security Groups
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
		return
	}
	fmt.Println("Security Groups:\n")
	if len(result.SecurityGroups) > 0 {
		// Marshal output
		bs, err := json.MarshalIndent(result.SecurityGroups, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(string(bs)), &SecG)
		for _, SG := range SecG {
			if b == SG.VpcID {
				// marshall output
				SGData, err := json.Marshal(SG)
				if err != nil {
					fmt.Println(err)
				}
				// This will be the returned string
				SGDesiredOutPut = append(SGDesiredOutPut, string(SGData))
			}
		}

	}
	fmt.Println(SGDesiredOutPut)
}
