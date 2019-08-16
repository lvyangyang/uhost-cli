package backen

import "encoding/base64"

//ResetUHostInstancePassword  启动host
func ResetUHostInstancePassword() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "ResetUHostInstancePassword"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId

	//其他参数
	paraments["Region"] = APIArgs.Region
	paraments["Zone"] = APIArgs.Zone
	paraments["UHostId"] = APIArgs.UHostId
	paraments["Password"] = base64.StdEncoding.EncodeToString([]byte(APIArgs.Password))

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
