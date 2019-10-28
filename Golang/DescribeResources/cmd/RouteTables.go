package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func checkRTInVPC(s, b string) {
	// Struct for RouteTables
	var RT []VPCRouteTables
	// string for instance output
	var RTDesiredOutput = ""
	// Session for describing RouteTables
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeRouteTablesInput{}
	result, err := svc.DescribeRouteTables(input)
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
	fmt.Println("RouteTables:\n")
	if len(result.RouteTables) > 0 {
		bs, err := json.MarshalIndent(result.RouteTables, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal([]byte(string(bs)), &RT)
		// // verify VPC Id
		for _, routeTable := range RT {
			if b == routeTable.VpcID {
				RTData, err := json.Marshal(routeTable)
				if err != nil {
					fmt.Println(err)
				}
				RTDesiredOutput = string(RTData)
				fmt.Println(RTDesiredOutput)
			}
		}
	}

}
