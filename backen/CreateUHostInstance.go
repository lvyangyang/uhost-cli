package backen

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/thedevsaddam/gojsonq"

	"github.com/spf13/viper"
)

var imageNumQueryLimit = 1000

//CreateUHostInstance  启动host
func CreateUHostInstance() {

	if APIArgs.LoginMode == "" {
		APIArgs.LoginMode = "Password"
		APIArgs.Password = viper.Get("password").(string)
	}
	if APIArgs.ImageId == "" {

		APIArgs.ImageId = getImageID(viper.Get("imagename").(string))
		if APIArgs.ImageId == "" {
			fmt.Println("cant find image " + viper.Get("imagename").(string))
			os.Exit(0)
		}
	}
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "CreateUHostInstance"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId
	APIArgs.Password = base64.StdEncoding.EncodeToString([]byte(APIArgs.Password))
	//其他参数
	populateParams(APIArgs.actionParams, paraments)
	populateParams(APIArgs.resourceParams, paraments)
	populateParams(APIArgs.arrayParams, paraments)

	MakeAPIRequset(paraments, APIArgs.PrivateKey)
}

func getImageID(imagename string) string {
	paraments := make(map[string]interface{})
	//公共参数
	paraments["Action"] = "DescribeImage"
	paraments["PublicKey"] = APIArgs.PublicKey
	paraments["ProjectId"] = APIArgs.ProjectId
	APIArgs.Password = base64.StdEncoding.EncodeToString([]byte(APIArgs.Password))
	//其他参数
	paraments["Limit"] = imageNumQueryLimit
	populateParams(APIArgs.actionParams, paraments)
	populateParams(APIArgs.resourceParams, paraments)
	populateParams(APIArgs.arrayParams, paraments)

	body := makeAPIRequsetNoPrint(paraments, APIArgs.PrivateKey)

	jq := gojsonq.New().JSONString(string(body)).From("ImageSet").Select("ImageId", "ImageName")
	imageInfoList, ok := jq.Get().([]interface{})
	if !ok {
		fmt.Println("parse imageList  error")
		os.Exit(0)
	}
	for _, imageInfo := range imageInfoList {
		imageInfoMap, ok := imageInfo.(map[string]interface{})
		if !ok {
			fmt.Println("parse imageList  error")
			os.Exit(0)
		}
		if imageInfoMap["ImageName"] == imagename {
			return imageInfoMap["ImageId"].(string)
		}

	}
	return ""
}
