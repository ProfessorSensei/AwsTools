package cmd

// This is where all the structs will go. replace when there is an improvement

type InstanceStruct struct {
	Instances []struct {
		ImageID           string `json:"ImageId"`
		InstanceID        string `json:"InstanceId"`
		InstanceType      string `json:"InstanceType"`
		KeyName           string `json:"KeyName"`
		LaunchTime        string `json:"LaunchTime"`
		NetworkInterfaces []struct {
			NetworkInterfaceID string `json:"NetworkInterfaceId"`
		} `json:"NetworkInterfaces"`
		PrivateIPAddress string `json:"PrivateIpAddress"`
		PublicIPAddress  string `json:"PublicIpAddress"`
		SecurityGroups   []struct {
			GroupName string `json:"GroupName"`
		} `json:"SecurityGroups"`
		SubnetID string `json:"SubnetId"`
		Tags     []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"Tags"`
		VpcID string `json:"VpcId"`
	} `json:"Instances"`
}

// had to take out associations
// multiple associations printed the same line multiple times
type NetworkAcls struct {
	Entries []struct {
		CidrBlock string `json:"CidrBlock"`
	} `json:"Entries"`
	IsDefault    bool   `json:"IsDefault"`
	NetworkACLID string `json:"NetworkAclId"`
	VpcID        string `json:"VpcId"`
}

type NetWorkInterFaces struct {
	Attachment struct {
		AttachmentID string `json:"AttachmentId"`
	} `json:"Attachment,omitempty"`
	AvailabilityZone   string `json:"AvailabilityZone"`
	Description        string `json:"Description"`
	InterfaceType      string `json:"InterfaceType"`
	MacAddress         string `json:"MacAddress"`
	NetworkInterfaceID string `json:"NetworkInterfaceId"`
	PrivateDNSName     string `json:"PrivateDnsName"`
	PrivateIPAddress   string `json:"PrivateIpAddress"`
	VpcID              string `json:"VpcId"`
	Association        struct {
		PublicDNSName string `json:"PublicDnsName"`
		PublicIP      string `json:"PublicIp"`
	} `json:"Association,omitempty"`
}

type NGateways struct {
	CreateTime   string `json:"CreateTime"`
	NatGatewayID string `json:"NatGatewayId"`
	State        string `json:"State"`
	VpcID        string `json:"VpcId"`
}

// turn this into a variable
type VPCRouteTables struct {
	RouteTableID string `json:"RouteTableId"`
	Routes       []struct {
		DestinationCidrBlock string `json:"DestinationCidrBlock"`
		GatewayID            string `json:"GatewayId"`
		Origin               string `json:"Origin"`
		State                string `json:"State"`
	} `json:"Routes"`
	VpcID string `json:"VpcId"`
}
type SecurityGroups struct {
	Description   string `json:"Description"`
	GroupID       string `json:"GroupId"`
	GroupName     string `json:"GroupName"`
	IPPermissions []struct {
		FromPort   int    `json:"FromPort"`
		IPProtocol string `json:"IpProtocol"`
		IPRanges   []struct {
			CidrIP string `json:"CidrIp"`
		} `json:"IpRanges"`
		ToPort int `json:"ToPort"`
	} `json:"IpPermissions"`
	IPPermissionsEgress []struct {
		IPProtocol string `json:"IpProtocol"`
		IPRanges   []struct {
			CidrIP string `json:"CidrIp"`
		} `json:"IpRanges"`
	} `json:"IpPermissionsEgress"`
	VpcID string `json:"VpcId"`
}

type Subnet struct {
	CidrBlock    string `json:"CidrBlock"`
	DefaultForAz bool   `json:"DefaultForAz"`
	State        string `json:"State"`
	SubnetID     string `json:"SubnetId"`
	VpcID        string `json:"VpcId"`
}
type InterNetGateways struct {
	Attachments []struct {
		State string `json:"State"`
		VpcID string `json:"VpcId"`
	} `json:"Attachments"`
	InternetGatewayID string `json:"InternetGatewayId"`
}
