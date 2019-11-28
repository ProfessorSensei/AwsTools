package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./S3Funcs"
)

// region for bucket
var reg, buckName, s3Opt string

func main() {
	// ask user for bucket option they would like
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`Please enter option: cb(create bucket) | cbv(create bucket with versioning 
		dlb(delete bucket) | dsb(describe bucket`)
	s3Opt, _ = reader.ReadString('\n')
	// remove new line
	s3Opt = strings.Replace(s3Opt, "\n", "", -1)
	if strings.ToLower(s3Opt) == "cb" {
		fmt.Println("this worked")
		buckinfo()

	} else if strings.ToLower(s3Opt) == "dlb" {
		delbuck()
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
