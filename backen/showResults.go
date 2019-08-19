package backen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/thedevsaddam/gojsonq"
)

//JSONPrettyPrint 带换行输出json
func JSONPrettyPrint(body []byte) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println("JSON parse error: ", error)
		return
	}
	fmt.Println(string(prettyJSON.Bytes()))
}

//SimpleResultPrint 简化返回列表内容
func SimpleResultPrint(setName string, tabs []string, body []byte) {
	jq := gojsonq.New().JSONString(string(body)).From(setName).Select(tabs...)
	imageInfoList, ok := jq.Get().([]interface{})
	if !ok {
		fmt.Println("parse imageList  error")
		os.Exit(0)
	}
	for _, tab := range tabs {
		fmt.Printf("%s\t", tab)
	}
	fmt.Printf("\n")
	for _, imageInfo := range imageInfoList {
		imageInfoMap, ok := imageInfo.(map[string]interface{})
		if !ok {
			fmt.Println("parse imageList  error")
			os.Exit(0)
		}
		for _, tab := range tabs {
			fmt.Printf("%v\t", imageInfoMap[tab])
		}
		fmt.Printf("\n")
	}
}
