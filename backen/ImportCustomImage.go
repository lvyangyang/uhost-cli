package backen

import (
	"encoding/base64"
)

//ImportCustomImage  启动host
func ImportCustomImage() {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "ImportCustomImage"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId
	APIArgs.Password = base64.StdEncoding.EncodeToString([]byte(APIArgs.Password))
	//其他参数
	populateParams(APIArgs.actionParams, paraments)
	populateParams(APIArgs.resourceParams, paraments)
	populateParams(APIArgs.arrayParams, paraments)

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}
