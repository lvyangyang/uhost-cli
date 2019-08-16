package backen

//ModifyUHostInstanceTag  启动host
func ModifyUHostInstanceTag() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "ModifyUHostInstanceTag"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	paraments["UHostId"] = APIArgs.UHostId

	if APIArgs.Tag != "" {
		paraments["Tag"] = APIArgs.Tag
	}
	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
