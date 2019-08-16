package backen

import (
	"strconv"
	"strings"
)

//UpgradeToArkUHostInstance  启动host
func UpgradeToArkUHostInstance() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "UpgradeToArkUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他
	paraments["Region"] = APIArgs.Region
	paraments["Zone"] = APIArgs.Zone

	tagString := "UHostIds.N"
	if len(APIArgs.UHostIdsN) > 0 {
		length := len(APIArgs.UHostIdsN)
		for j := 0; j < length; j++ {
			paraments[strings.Replace(tagString, ".N", "."+strconv.Itoa(j), 1)] = APIArgs.UHostIdsN[j]
		}
	}
	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
