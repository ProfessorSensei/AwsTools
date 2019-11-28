package S3Funcs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// delete bucket function
func DeleteBuck(reg, buckName string) (*s3.DeleteBucketOutput, error) {
	// will need region after all
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	svc := s3.New(sess)
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(buckName),
	}

	result, err := svc.DeleteBucket(input)
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
