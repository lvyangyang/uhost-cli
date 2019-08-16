package backen

//CreateIsolationGroup  启动host
func CreateIsolationGroup() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "CreateIsolationGroup"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region
	paraments["GroupName"] = APIArgs.GroupName
	paraments["Remark"] = APIArgs.Remark

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
