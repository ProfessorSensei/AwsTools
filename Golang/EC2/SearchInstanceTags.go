package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// a quick way to retrieve instance data filtering with tags input by user

var tNme, tVal, awsReg string

func main() {
	// User input tag name
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter tag name:")
	tNme, _ = reader.ReadString('\n')
	// remove new line
	tNme = strings.Replace(tNme, "\n", "", -1)
	// User input tag value
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter tag value:")
	tVal, _ = reader2.ReadString('\n')
	// remove new line
	tVal = strings.Replace(tVal, "\n", "", -1)
	// User input Region
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region the instance is located:")
	awsReg, _ = reader3.ReadString('\n')
	// remove new line
	awsReg = strings.Replace(awsReg, "\n", "", -1)

	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsReg)},
	)
	// Tag to filter for
	filTagName := fmt.Sprintf("tag:%v", tNme)
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(filTagName),
				Values: []*string{
					aws.String(tVal),
				},
			},
		},
	}
	// describe filtered instance
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
	// print filtered instance output
	fmt.Println(result)
}
