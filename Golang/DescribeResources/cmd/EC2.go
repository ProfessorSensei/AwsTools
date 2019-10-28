package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// to store ec2 data from the instances
// var TestSlice []*ec2.Reservation

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
		// TestSlice = result.Reservations
		bs, err := json.MarshalIndent(result.Reservations, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		// stringOutput := string(bs)
		// fmt.Println(stringOutput)
		// probabaly best to send the marshalled string
		// and unmarshal into the master struct from there
		json.Unmarshal([]byte(string(bs)), &instances)
		// If the VPCId matches, return the json string
		for _, instance := range instances {
			for _, ec2 := range instance.Instances {
				if b == ec2.VpcID {
					// fmt.Println(ec2)
					IData, err := json.Marshal(ec2)
					if err != nil {
						fmt.Println(err)
					}
					// fmt.Println(string(IData))
					INSTDesiredOutput = string(IData)
					fmt.Println(INSTDesiredOutput)
				} // 		for _, networkInfo := range ec2.NetworkInterfaces {
				// 			if b == networkInfo.VpcID {
				// 				fmt.Printf("Image ID: '%v'\t Instance ID: '%v'\t Key Pair: '%v'\t Instance Type: '%v'\n", ec2.ImageID, ec2.InstanceID, ec2.KeyName, ec2.InstanceType)
				// 			}
				// 		}
				// 		// fmt.Printf("Image ID: '%v'\n", ec2.ImageID)
				// 		// fmt.Printf("Instance ID: '%v'\n", ec2.InstanceID)
				// 		// fmt.Printf("Key Name: '%v'\n", ec2.KeyName)
				// 		// fmt.Printf("'%T'\n", ec2.KeyName)
			}
		}
		// fmt.Println(instances)
		// for _, instance := range result.Reservations {
		// 	for _, machine := range instance.Instances {
		// 		if aws.StringValue(machine.VpcId) == b {
		// 			fmt.Println(aws.StringValue(machine.InstanceId))
		// 		}
		// 	}
		// }
	} else {
		fmt.Println("Zero instances in this VPC")
	}

}
