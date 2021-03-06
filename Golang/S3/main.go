package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./S3Funcs"
)

// region for bucket
var reg, buckName, s3Opt, MFADelete string

func main() {
	// ask user for bucket option they would like
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`Please enter option: cb(create bucket) | cbv(create bucket with versioning) | 
		dlb(delete bucket) | lstb(list buckets)`)
	s3Opt, _ = reader.ReadString('\n')
	// remove new line
	s3Opt = strings.Replace(s3Opt, "\n", "", -1)
	if strings.ToLower(s3Opt) == "cb" {
		fmt.Println("this worked")
		buckinfo()

	} else if strings.ToLower(s3Opt) == "dlb" {
		delbuck()
	} else if strings.ToLower(s3Opt) == "lstb" {
		lsbct()
	} else if strings.ToLower(s3Opt) == "cbv" {
		buckvinfo()
	}

}

// gather bucket info and pass to create function
func buckinfo() {
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
	cb, err := S3Funcs.CreateBuck(reg, buckName)
	if err != nil {
		fmt.Printf("There was an error creating bucket: '%v'. Please check name", buckName)
	}
	// fmt.Println(cb)
	fmt.Printf("Bucket '%v' created \n '%v'", buckName, cb)
}

func delbuck() {
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
	S3Funcs.DeleteBuck(reg, buckName)
}

// list buckets
func lsbct() {
	// gather region and bucket name from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter region for bucket: \n")
	reg, _ = reader.ReadString('\n')
	// remove new line
	reg = strings.Replace(reg, "\n", "", -1)
	lst, err := S3Funcs.ListBuck(reg)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(lst)
}

// create bucket with versioning
func buckvinfo() {
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
	// MFA enable
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Println("MFA enabled? (Disabled|Enabled): \n")
	MFADelete, _ = reader3.ReadString('\n')
	// remove new line
	MFADelete = strings.Replace(MFADelete, "\n", "", -1)
	cb, err := S3Funcs.CreateBuck(reg, buckName)
	if err != nil {
		fmt.Printf("There was an error creating bucket: '%v'. Please check name", buckName)
	}
	// fmt.Println(cb)
	fmt.Printf("Bucket '%v' created \n '%v'", buckName, cb)
	// add waiter here then create versioning for bucket using MFA and bucket name
	// try using output from create bucket for location that is needed to create versioning
	_, err2 := S3Funcs.BuckVer(MFADelete, buckName, reg)
	if err2 != nil {
		fmt.Println("That didn't work, son")
	}
}
