package backen

import (
	"encoding/base64"
)

//DescribeUHostInstance  启动host
func DescribeUHostInstance() {
	var setName = "UHostSet"
	var tabs = []string{"UHostId", "Name", "Tag", "OsName", "CPU", "Memory", "State"}
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "DescribeUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId
	APIArgs.Password = base64.StdEncoding.EncodeToString([]byte(APIArgs.Password))
	//其他参数
	populateParams(APIArgs.actionParams, paraments)
	populateParams(APIArgs.resourceParams, paraments)
	populateParams(APIArgs.arrayParams, paraments)

	//MakeAPIRequset(paraments, APIArgs.PrivateKey)
	body := makeAPIRequsetNoPrint(paraments, APIArgs.PrivateKey)
	SimpleResultPrint(setName, tabs, body)
}
