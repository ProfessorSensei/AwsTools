package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// region and bucket name
var reg, buckName string

// create bucket function
func main() {
	// gather region and bucket name from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter region for bucket: \n")
	reg, _ = reader.ReadString('\n')
	// remove new line
	reg = strings.Replace(reg, "\n", "", -1)
	// bucket name
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter bucket name: \n")
	buckName, _ = reader2.ReadString('\n')
	// remove new line
	buckName = strings.Replace(buckName, "\n", "", -1)
	// session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	fmt.Println("create bucket function goes here")
	svc := s3.New(sess)
	input := &s3.CreateBucketInput{
		Bucket: aws.String(buckName),
		// CreateBucketConfiguration: &s3.CreateBucketConfiguration{
		// 	LocationConstraint: aws.String(reg),
		// },
	}

	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	fmt.Println(result)
}
