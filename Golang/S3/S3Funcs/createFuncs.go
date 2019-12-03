package S3Funcs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// create bucket function
func CreateBuck(reg, buckName string) (*s3.CreateBucketOutput, error) {
	// session - change this to an automatically set region since S3 is global
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	svc := s3.New(sess)
	input := &s3.CreateBucketInput{
		Bucket: aws.String(buckName),
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
	return result, err
}

// bucket versioning
func BuckVer(MFA_Delete, buckName, reg string) (*s3.PutBucketVersioningOutput, error) {
	// session
	// MFA_bool, err := strconv.ParseBool(MFA)
	// if err != nil {
	// 	fmt.Println("error not bool value")
	// }
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	svc := s3.New(sess)
	input := &s3.PutBucketVersioningInput{
		Bucket: aws.String(buckName),
		VersioningConfiguration: &s3.VersioningConfiguration{
			MFADelete: aws.String(MFA_Delete),
			Status:    aws.String("Enabled"),
		},
	}

	result, err := svc.PutBucketVersioning(input)
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
	}

	return result, err
}
