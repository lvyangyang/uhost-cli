package backen

//StartUHostInstance  启动host
func StartUHostInstance() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "StartUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	paraments["UHostId"] = APIArgs.UHostId
	
	if APIArgs.DiskPassword != "" {
		paraments["DiskPassword"] = APIArgs.DiskPassword
	}
	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
