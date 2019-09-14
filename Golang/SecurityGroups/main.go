package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan []string)

	// get channel with values from aws regions
	go populate(c1)

	// use region from c1 and perform describe funcs
	go fanOutIn(c1, c2)

	// pull value created from above func calls and print
	for v := range c2 {
		fmt.Println(v)
	}

}

func populate(c chan string) {
	for _, region := range awsRegionList() {
		c <- region
	}
	close(c)
}

func fanOutIn(c1 chan string, c2 chan []string) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 string) {
					c2 <- checkSecurityGroups(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func awsRegionList() []string {
	var awsRegions []string
	// create ec2 session
	ec2session := ec2.New(session.New(), &aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	})
	// describe regions
	regions, err := ec2session.DescribeRegions(&ec2.DescribeRegionsInput{})

	if err != nil {
		panic(err)
	}

	for _, region := range regions.Regions {
		awsRegions = append(awsRegions, *region.RegionName)
	}

	return awsRegions
}

func checkSecurityGroups(s string) []string {
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(s))
	input := &ec2.DescribeSecurityGroupsInput{}

	result, err := svc.DescribeSecurityGroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	var a []string
	if len(result.SecurityGroups) > 0 {
		fmt.Printf("'%v' Internet Gateways:\n", s)
		for _, sb := range result.SecurityGroups {
			a = append(a, aws.StringValue(sb.GroupName))
		}

	}
	return a

}
