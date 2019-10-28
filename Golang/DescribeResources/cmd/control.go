/*
Copyright © 2019 Peter Menage petermenage1@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// string region flag variable
var selectedRegion string

// string VPCID flag variable
var selectedVPCId string

// []string to hold VPC Id's and pass to other functions
var vpcIds []string

// helloWorldCmd represents the helloWorld command
var helloWorldCmd = &cobra.Command{
	Use:   "helloWorld",
	Short: "Map resources to VPC ID",
	Long: `VPCResources goes through an AWS account (assuming you have programmatic access to an account)
and returns a list of Resources mapped to the VPC that is either specified by the user, or discovered 
in the AWS account. If you do not pass either a region, or VPCId, each region will be searched for VPCs and
resources that map back to the VPC/VPCs on a per region basis`,
	Run: func(cmd *cobra.Command, args []string) {
		testFunc(vpcIdList(awsRegionList()))
	},
}

func init() {
	rootCmd.AddCommand(helloWorldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloWorldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// string region flag
	helloWorldCmd.Flags().StringVarP(&selectedRegion, "selectedRegion", "r", "", "This is not a test")
	// string VPCId flag
	helloWorldCmd.Flags().StringVarP(&selectedVPCId, "selectedVPCId", "v", "", "This is not a test")
}

func testFunc(s []string) {
	if len(selectedRegion) > 0 {
		fmt.Printf("Checking for Region: '%v'\n........", selectedRegion)
		// use AccountContains to check if provided region is valid
		if AccountContains(awsRegionList(), selectedRegion) == true {
			describeVPCs(selectedRegion)
		} else {
			fmt.Printf("'%v' is not a valid region. Please try again\n", selectedRegion)
		}
	} else if len(selectedVPCId) > 0 {
		fmt.Printf("Searching for VPC: '%v'\n", selectedVPCId)
		if AccountContains(vpcIdList(awsRegionList()), selectedVPCId) == true {
			var veepz string
			var regs string
			for _, az := range s {
				veepz = az
			}
			for _, region := range awsRegionList() {
				regs = region
			}
			// regs is the region of the vpcId
			// veepz is the VPC Id
			fmt.Println("Gathering VPC Resources.....")
			flagVPCs(regs, veepz)

		} else if AccountContains(vpcIdList(awsRegionList()), selectedVPCId) != true {
			fmt.Printf("Could not locate VPC: '%v'.\nPlease enter make sure you entered a valid VPC Id and try again.\n", selectedVPCId)
		}
	} else {
		fmt.Println("Gathering VPC Resources.....")
		for _, region := range awsRegionList() {
			describeVPCs(region)
		}
	}
}
