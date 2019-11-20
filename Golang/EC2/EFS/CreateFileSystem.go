package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

// EFS Name, AWS region
var efsName, reg string

func main() {
	// user input efs name
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter efs name:")
	efsName, _ = reader.ReadString('\n')
	// remove new line
	efsName = strings.Replace(efsName, "\n", "", -1)
	// user input aws region
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region to create this volume in:")
	reg, _ = reader2.ReadString('\n')
	// remove new line
	reg = strings.Replace(reg, "\n", "", -1)
	svc := efs.New(session.New())
	input := &efs.CreateFileSystemInput{
		CreationToken:   aws.String("tokenstring"),
		PerformanceMode: aws.String("generalPurpose"),
		Tags: []*efs.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String(efsName),
			},
		},
	}

	result, err := svc.CreateFileSystem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case efs.ErrCodeBadRequest:
				fmt.Println(efs.ErrCodeBadRequest, aerr.Error())
			case efs.ErrCodeInternalServerError:
				fmt.Println(efs.ErrCodeInternalServerError, aerr.Error())
			case efs.ErrCodeFileSystemAlreadyExists:
				fmt.Println(efs.ErrCodeFileSystemAlreadyExists, aerr.Error())
			case efs.ErrCodeFileSystemLimitExceeded:
				fmt.Println(efs.ErrCodeFileSystemLimitExceeded, aerr.Error())
			case efs.ErrCodeInsufficientThroughputCapacity:
				fmt.Println(efs.ErrCodeInsufficientThroughputCapacity, aerr.Error())
			case efs.ErrCodeThroughputLimitExceeded:
				fmt.Println(efs.ErrCodeThroughputLimitExceeded, aerr.Error())
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
