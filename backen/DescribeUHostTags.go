package backen

//DescribeUHostTags  启动host
func DescribeUHostTags() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "DescribeUHostTags"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
