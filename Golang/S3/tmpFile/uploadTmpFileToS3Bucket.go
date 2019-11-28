package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// This script creates a tmp file and uploads the file to an AWS S3 bucket
	content := []byte("content of temp file")
	tmpfile, err := ioutil.TempFile("", "tmpFilename.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Delete created File
	defer os.Remove(tmpfile.Name()) 

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	svc := s3.New(session.New(), &aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	})
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(tmpfile.Name())),
		Bucket: aws.String("<S3-Bucket-Name>"),
		Key:    aws.String("<file_name_for_bucket>"),
		// Tagging: aws.String("key1=value1&key2=value2"),
	}

	result, err := svc.PutObject(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
