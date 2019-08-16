package backen

//LeaveIsolationGroup  启动host
func LeaveIsolationGroup() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "LeaveIsolationGroup"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}

	paraments["UHostId"] = APIArgs.UHostId
	paraments["GroupId"] = APIArgs.GroupId
	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
