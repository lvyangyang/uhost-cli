package backen

//TerminateUHostInstance  启动host
func TerminateUHostInstance() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "TerminateUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}
	paraments["UHostId"] = APIArgs.UHostId
	paraments["Destroy"] = APIArgs.Destroy

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
