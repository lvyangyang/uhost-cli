package backen

//ModifyUHostInstanceName  启动host
func ModifyUHostInstanceName() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "ModifyUHostInstanceName"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	paraments["UHostId"] = APIArgs.UHostId

	if APIArgs.Name != "" {
		paraments["Name"] = APIArgs.Name
	}
	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
