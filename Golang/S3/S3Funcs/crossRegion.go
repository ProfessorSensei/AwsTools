package S3Funcs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func crossRegRep(reg, buck, dstbuckarn, strclass, repcnfrole string) {
	svc := s3.New(session.New())
	input := &s3.PutBucketReplicationInput{
		Bucket: aws.String("examplebucket"),
		ReplicationConfiguration: &s3.ReplicationConfiguration{
			Role: aws.String("arn:aws:iam::123456789012:role/examplerole"),
			Rules: []*s3.ReplicationRule{
				{
					Destination: &s3.Destination{
						Bucket:       aws.String("arn:aws:s3:::destinationbucket"),
						StorageClass: aws.String("STANDARD"),
					},
					Prefix: aws.String(""),
					Status: aws.String("Enabled"),
				},
			},
		},
	}

	result, err := svc.PutBucketReplication(input)
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
