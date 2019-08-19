package backen

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

var apiAddr = "https://api.ucloud.cn"

const maxArraySize = 20

var timeoutLimit = 15 * time.Second

type publicParams struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	Signature  string `json:"Signature"`
}

type resourceParams struct {
	Action    string `json:"Action"`
	Region    string `json:"Region"`
	Zone      string `json:"Zone"`
	UHostId   string `json:"UHostId"`
	ProjectId string `json:"ProjectId"`
	GroupId   string `json:"GroupId"`
	GroupName string `json:"GroupName"`
	Remark    string `json:"Remark"`
}

type actionParams struct {
	DiskPassword           string `json:"DiskPassword"`
	Destroy                int    `json:"Destroy"`
	CPU                    int    `json:"CPU"`
	Memory                 int    `json:"Memory"`
	DiskSpace              int    `json:"DiskSpace"`
	BootDiskSpace          int    `json:"BootDiskSpace"`
	NetCapValue            int    `json:"NetCapValue"`
	ImageId                string `json:"ImageId"`
	LoginMode              string `json:"LoginMode"`
	Password               string `json:"Password"`
	MachineType            string `json:"MachineType"`
	UHostType              string `json:"UHostType"`
	MinimalCpuPlatform     string `json:"MinimalCpuPlatform"`
	GpuType                string `json:"GpuType"`
	GPU                    int    `json:"GPU"`
	NetCapability          string `json:"NetCapability"`
	VPCId                  string `json:"VPCId"`
	SubnetId               string `json:"SubnetId"`
	SecurityGroupId        string `json:"SecurityGroupId"`
	AlarmTemplateId        int    `json:"AlarmTemplateId"`
	IsolationGroup         string `json:"IsolationGroup"`
	Name                   string `json:"Name"`
	Tag                    string `json:"Tag"`
	ChargeType             string `json:"ChargeType"`
	Quantity               int    `json:"Quantity"`
	MaxCount               int    `json:"MaxCount"`
	CouponId               string `json:"CouponId"`
	Count                  int    `json:"Count"`
	Offset                 int    `json:"Offset"`
	Limit                  int    `json:"Limit"`
	ImageName              string `json:"ImageName"`
	ImageDescription       string `json:"ImageDescription"`
	BackupMode             string `json:"BackupMode"`
	DiskId                 string `json:"DiskId"`
	SourceImageId          string `json:"SourceImageId"`
	TargetRegion           string `json:"TargetRegion"`
	TargetProjectId        string `json:"TargetProjectId"`
	TargetImageName        string `json:"TargetImageName"`
	TargetImageDescription string `json:"TargetImageDescription"`
	ImageType              string `json:"ImageType"`
	OsType                 string `json:"OsType"`
	OsName                 string `json:"OsName"`
	PriceSet               int    `json:"PriceSet"`
	UFileUrl               string `json:"UFileUrl"`
	Format                 string `json:"Format"`
	Auth                   bool   `json:"Auth"`
	ReserveDisk            string `json:"ReserveDisk"`
	ResourceType           string `json:"ResourceType"`
}

type arrayParams struct {
	UHostIdsN                             []string `json:"UHostIds.N"`
	DisksNIsBoot                          []string `json:"Disks.N.IsBoot"`
	DisksNType                            []string `json:"Disks.N.Type"`
	DisksNSize                            []int    `json:"Disks.N.Size"`
	DisksNBackupType                      []string `json:"Disks.N.BackupType"`
	DisksNEncrypted                       []bool   `json:"Disks.N.Encrypted"`
	DisksNKmsKeyId                        []string `json:"Disks.N.KmsKeyId"`
	DisksNCouponId                        []string `json:"Disks.N.CouponId"`
	PrivateIpN                            []string `json:"PrivateIp.N"`
	NetworkInterfaceNEIPBandwidth         []int    `json:"NetworkInterface.N.EIP.Bandwidth"`
	NetworkInterfaceNEIPPayMode           []string `json:"NetworkInterface.N.EIP.PayMode"`
	NetworkInterfaceNEIPShareBandwidthId  []string `json:"NetworkInterface.N.EIP.ShareBandwidthId"`
	NetworkInterfaceNEIPGlobalSSHArea     []string `json:"NetworkInterface.N.EIP.GlobalSSH.Area"`
	NetworkInterfaceNEIPOperatorName      []string `json:"NetworkInterface.N.EIP.OperatorName"`
	NetworkInterfaceNEIPGlobalSSHPort     []int    `json:"NetworkInterface.N.EIP.GlobalSSH.Port"`
	NetworkInterfaceNEIPCouponId          []string `json:"NetworkInterface.N.EIP.CouponId"`
	NetworkInterfaceNEIPGlobalSSHAreaCode []string `json:"NetworkInterface.N.EIP.GlobalSSH.AreaCode	"`
	DNSServersN                           []string `json:"DNSServers.N"`
}

type reqParams struct {
	publicParams
	resourceParams
	actionParams
	arrayParams
}

//APIArgs 请求参数
var APIArgs reqParams

func makeAPIRequsetNoPrint(paraments map[string]interface{}, privateKey string) []byte {
	signature := ucloudAPISignGen(paraments, privateKey)
	paraments["Signature"] = signature

	requestParams, err := json.Marshal(paraments)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		os.Exit(-1)
	}

	client := &http.Client{Timeout: timeoutLimit}
	req, err := http.NewRequest("POST", apiAddr, bytes.NewReader(requestParams))
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		os.Exit(-1)
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		os.Exit(-1)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}

//MakeAPIRequset 通过post发起api请求
func MakeAPIRequset(paraments map[string]interface{}, privateKey string) {
	body := makeAPIRequsetNoPrint(paraments, privateKey)

	JSONPrettyPrint(body)
}

func ucloudAPISignGen(sourceMap map[string]interface{}, privateKey string) string {
	var paramsData string

	var keys []string
	for k := range sourceMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		paramsData = paramsData + key + fmt.Sprintf("%v", sourceMap[key])
	}
	paramsData += privateKey

	t := sha1.New()
	io.WriteString(t, paramsData)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func populateParams(APIArgs interface{}, paraments map[string]interface{}) {
	typ := reflect.TypeOf(APIArgs)
	val := reflect.ValueOf(APIArgs)
	num := val.NumField()
	for i := 0; i < num; i++ {

		tagString := typ.Field(i).Tag.Get("json")
		if val.Field(i).Kind() == reflect.String {
			if val.Field(i).String() != "" {
				paraments[tagString] = val.Field(i).String()
			}
			continue
		}
		if val.Field(i).Kind() == reflect.Int {
			if val.Field(i).Int() != 0 {
				paraments[tagString] = val.Field(i).Int()
			}
			continue
		}
		if val.Field(i).Kind() == reflect.Slice {
			if val.Field(i).Len() > 0 {
				length := val.Field(i).Len()
				if val.Field(i).Index(0).Kind() == reflect.String {
					for j := 0; j < length; j++ {
						paraments[strings.Replace(tagString, ".N", "."+strconv.Itoa(j), 1)] = val.Field(i).Index(j).String()
					}
				}
				if val.Field(i).Index(0).Kind() == reflect.Int {
					for j := 0; j < length; j++ {
						paraments[strings.Replace(tagString, ".N", "."+strconv.Itoa(j), 1)] = val.Field(i).Index(j).Int()
					}
				}
				if val.Field(i).Index(0).Kind() == reflect.Bool {
					for j := 0; j < length; j++ {
						paraments[strings.Replace(tagString, ".N", "."+strconv.Itoa(j), 1)] = val.Field(i).Index(j).Bool()
					}
				}

			}
			continue
		}
	}
}
