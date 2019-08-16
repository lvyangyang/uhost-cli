package backen

//ResizeUHostInstance  启动host
func ResizeUHostInstance() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "ResizeUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region

	if APIArgs.Zone != "" {
		paraments["Zone"] = APIArgs.Zone
	}
	paraments["UHostId"] = APIArgs.UHostId
	if APIArgs.CPU > 0 {
		paraments["CPU"] = APIArgs.CPU
	}
	if APIArgs.Memory > 0 {
		paraments["Memory"] = APIArgs.Memory
	}
	if APIArgs.DiskSpace > 0 {
		paraments["DiskSpace"] = APIArgs.DiskSpace
	}
	if APIArgs.BootDiskSpace > 0 {
		paraments["BootDiskSpace"] = APIArgs.BootDiskSpace
	}
	paraments["NetCapValue"] = APIArgs.NetCapValue

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
