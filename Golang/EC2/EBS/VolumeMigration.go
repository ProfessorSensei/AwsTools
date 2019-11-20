// VolumeMigration takes a snapshot of an existing volume and copies
// the snapshot to another region to create new volume. User decides
// if they would like to encrypt the new volume
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var boocryp bool
var regf, reg2, encryp, snapId, descrip string

func main() {
	fmt.Println("here goes")
	// user input region from
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region this volume is located in:")
	regf, _ = reader2.ReadString('\n')
	// remove new line
	regf = strings.Replace(regf, "\n", "", -1)
	// user input region to
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the region to create the new volume in:")
	reg2, _ = reader.ReadString('\n')
	// remove new line
	reg2 = strings.Replace(reg2, "\n", "", -1)
	// encryption boolean
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("Encrypt volume: (true|false)")
	encryp, _ = reader3.ReadString('\n')
	// remove new line
	encryp = strings.Replace(encryp, "\n", "", -1)
	// convert to boolean
	if boocryp, err := strconv.ParseBool(encryp); err == nil {
		fmt.Printf("%T, %v\n", boocryp, boocryp)
	}
	// snapshot ID
	reader4 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the snapshot ID to create volume:")
	snapId, _ = reader4.ReadString('\n')
	// remove new line
	snapId = strings.Replace(snapId, "\n", "", -1)
	// user input description
	reader5 := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the description for the snapshot:")
	descrip, _ = reader5.ReadString('\n')
	// remove new line
	descrip = strings.Replace(descrip, "\n", "", -1)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regf)},
	)
	svc2 := ec2.New(sess)
	resultsnap, err := svc2.CopySnapshot(createSnap(boocryp, reg2, regf, snapId, descrip))

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
	fmt.Println(resultsnap)
}

// Create CopySnapshotInput from user provided information
func createSnap(encryp bool, reg2, regf, snapId, descrip string) *ec2.CopySnapshotInput {
	csinput := &ec2.CopySnapshotInput{
		Description:       aws.String(descrip),
		DestinationRegion: aws.String(reg2),
		Encrypted:         &encryp,
		SourceRegion:      aws.String(regf),
		SourceSnapshotId:  aws.String(snapId),
	}
	return csinput
}
