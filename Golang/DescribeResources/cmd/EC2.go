package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// check instances in VPC
func checkInstancesInVPC(s, b string) {
	// Struct for instances
	var instances []InstanceStruct
	// string for instance output
	var INSTDesiredOutput = ""
	// Session for describing Instances
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)
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
	fmt.Println("Instances:\n")
	if len(result.Reservations) > 0 {
		bs, err := json.MarshalIndent(result.Reservations, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(string(bs)), &instances)
		// If the VPCId matches, return the json string
		for _, instance := range instances {
			for _, ec2 := range instance.Instances {
				if b == ec2.VpcID {
					IData, err := json.Marshal(ec2)
					if err != nil {
						fmt.Println(err)
					}
					INSTDesiredOutput = string(IData)
					fmt.Println(INSTDesiredOutput)
				}
			}
		}
	} else {
		fmt.Println("Zero instances in this VPC")
	}

}
