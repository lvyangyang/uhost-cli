package backen

//StopUHostInstance  启动host
func StopUHostInstance() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "StopUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	paraments["UHostId"] = APIArgs.UHostId

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
