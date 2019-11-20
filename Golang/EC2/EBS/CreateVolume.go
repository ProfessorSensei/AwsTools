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

// user inputs
var reg, az, vtype string

func main() {
	// region
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region to create this volume in:")
	reg, _ = reader.ReadString('\n')
	// remove new line
	reg = strings.Replace(reg, "\n", "", -1)
	// availability zone
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter availability zone:")
	az, _ = reader3.ReadString('\n')
	// remove new line
	az = strings.Replace(az, "\n", "", -1)
	// volume type
	reader4 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter volume type:")
	vtype, _ = reader4.ReadString('\n')
	// remove new line
	vtype = strings.Replace(vtype, "\n", "", -1)
	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	svc := ec2.New(sess)
	input := &ec2.CreateVolumeInput{
		AvailabilityZone: aws.String(az),
		Size:             aws.Int64(80),
		VolumeType:       aws.String(vtype),
	}

	result, err := svc.CreateVolume(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
