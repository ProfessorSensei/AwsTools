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

// add user input for volume id
var volId, reg string

func main() {
	// user input volume ID
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the volume ID to be deleted:")
	volId, _ = reader.ReadString('\n')
	// remove new line
	volId = strings.Replace(volId, "\n", "", -1)
	// user input region
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region of the to be deleted:")
	reg, _ = reader2.ReadString('\n')
	// remove new line
	reg = strings.Replace(reg, "\n", "", -1)
	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	svc := ec2.New(sess)
	input := &ec2.DeleteVolumeInput{
		VolumeId: aws.String(volId),
	}

	result, err := svc.DeleteVolume(input)
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
