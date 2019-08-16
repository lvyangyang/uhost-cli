package cmd

import (
	"fmt"
	"strings"

	"github.com/lvyangyang/uhost_cli/backen"

	"github.com/spf13/cobra"
)

func initUhostCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uhost",
		Short: "uhost util",
		Long:  "create host,get host infor and config host",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into host cmd" + strings.Join(args, " "))
		},
	}
	cmd.AddCommand(initGetUHostInstanceVncInfoCmd())
	cmd.AddCommand(initStartUHostInstanceCmd())
	cmd.AddCommand(initRebootUHostInstanceCmd())
	cmd.AddCommand(initStopUHostInstanceCmd())
	cmd.AddCommand(initDescribeUHostTagsCmd())
	cmd.AddCommand(initTerminateUHostInstanceCmd())
	cmd.AddCommand(initResizeUHostInstanceCmd())
	cmd.AddCommand(initCreateUHostInstanceCmd())
	cmd.AddCommand(initModifyUHostInstanceTagCmd())
	cmd.AddCommand(initModifyUHostInstanceNameCmd())
	cmd.AddCommand(initLeaveIsolationGroupCmd())
	cmd.AddCommand(initCreateIsolationGroupCmd())
	cmd.AddCommand(initResetUHostInstancePasswordCmd())
	cmd.AddCommand(initUpgradeToArkUHostInstanceCmd())
	return cmd
}

func initGetUHostInstanceVncInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "GetUHostInstanceVncInfo",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into GetUHostInstanceVncInfocmd")
			backen.GetUHostInstanceVncInfo()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	return cmd
}

func initStartUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "StartUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into StartUHostInstanceCmd")
			backen.StartUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVarP(&backen.APIArgs.DiskPassword, "DiskPassword", "P", "", "diskpassword")

	return cmd
}

func initRebootUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "RebootUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into RebootUHostInstanceCmd")
			backen.RebootUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVarP(&backen.APIArgs.DiskPassword, "DiskPassword", "P", "", "diskpassword")

	return cmd
}

func initStopUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "StopUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into StopUHostInstanceCmd")
			backen.StopUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")

	return cmd
}

func initDescribeUHostTagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "DescribeUHostTags",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into DescribeUHostTagsCmd")
			backen.DescribeUHostTags()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")

	return cmd
}

func initTerminateUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "TerminateUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into TerminateUHostInstanceCmd")
			backen.TerminateUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().IntVarP(&backen.APIArgs.Destroy, "Destroy", "D", 0, "0 default,may be recycled, 1 delete directly")

	return cmd
}

func initResizeUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ResizeUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into ResizeUHostInstanceCmd")
			backen.ResizeUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().IntVar(&backen.APIArgs.CPU, "CPU", 0, "virtual cpu nums,range:[1,16],must be a power of 2 except 1,default is current cpu nums")
	cmd.Flags().IntVar(&backen.APIArgs.Memory, "Memory", 0, "memory size in MB,range:[2048,65536],default is current memory size")
	cmd.Flags().IntVar(&backen.APIArgs.DiskSpace, "DiskSpace", 0, "Data disk size, unit: GB, range [10,1000]; SSD model, unit: GB, range [100,500]; step size: 10, the default value is the data disk size of the current instance, the data disk does not support shrinkage, Therefore, it is not allowed to enter a value larger than the current instance data disk size.")
	cmd.Flags().IntVar(&backen.APIArgs.BootDiskSpace, "BootDiskSpace", 0, "System disk size, unit: GB, range [20,100], step size: 10, the system disk does not support shrinkage, so it is not allowed to input a smaller value than the current instance system disk")
	cmd.Flags().IntVar(&backen.APIArgs.NetCapValue, "NetCapValue", 0, "NIC lifting level (1, indicating upgrade, 2 means downgrade, 0 means unchanged)")
	return cmd
}

func initCreateUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "CreateUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into CreateUHostInstanceCmd")
			backen.CreateUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")

	cmd.Flags().StringVar(&backen.APIArgs.ImageId, "ImageId", "", "Mirror ID. Please get it through DescribeImage")
	cmd.MarkFlagRequired("ImageId")

	cmd.Flags().StringVar(&backen.APIArgs.LoginMode, "LoginMode", "Password", "Host login mode. (default option): Password.")

	cmd.Flags().StringVarP(&backen.APIArgs.Password, "Password", "P", "", "UHost password.")
	cmd.MarkFlagRequired("Password")

	cmd.Flags().StringVar(&backen.APIArgs.MachineType, "MachineType", "N", "Cloud host model (V2.0), enumeration values ['N', 'C', 'G', 'O']. Refer to the cloud host model description.")

	cmd.Flags().StringVar(&backen.APIArgs.MinimalCpuPlatform, "MinimalCpuPlatform", "Intel/Auto", "Minimum cpu platform, enumeration values [Intel/Auto, Intel/IvyBridge, Intel/Haswell, Intel/Broadwell, Intel/Skylake, Intel/Cascadelake")

	cmd.Flags().IntVar(&backen.APIArgs.CPU, "CPU", 0, "virtual cpu nums,range:[1,16],must be a power of 2 except 1,default is current cpu nums")
	cmd.Flags().IntVar(&backen.APIArgs.Memory, "Memory", 0, "memory size in MB,range:[2048,65536],default is current memory size")
	cmd.Flags().StringVar(&backen.APIArgs.GpuType, "GpuType", "", "GPU type, enumeration value [K80, P40, V100]")
	cmd.Flags().IntVar(&backen.APIArgs.GPU, "GPU", 0, "The number of GPU card cores. This field is only supported on GPU models (optional range is related to UHostType)")

	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNIsBoot, "Disks.N.IsBoot", []string{"True"}, "Whether it is a system disk. Enumeration value:> True, is the system disk> False, is the data disk (default). There is only one disk in the Disks array and it is the system disk.")
	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNType, "Disks.N.Type", []string{"LOCAL_NORMAL"}, "Disk type. Please refer to the disk type.")
	cmd.Flags().IntSliceVar(&backen.APIArgs.DisksNSize, "Disks.N.Size", []int{20}, "Disk size in GB. Please refer to the disk type.")
	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNBackupType, "Disks.N.BackupType", []string{"NONE"}, "Disk backup solution. Enumeration value:> NONE, no backup> DATAARK, data ark Backup mode reference supported by current disk")
	cmd.Flags().BoolSliceVar(&backen.APIArgs.DisksNEncrypted, "Disks.N.Encrypted", []bool{}, "[The function is only partially available, open for technical support.] Is the disk encrypted? Encryption: true, no encryption: false. Encryption must be passed to the corresponding KmsKeyId")
	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNKmsKeyId, "Disks.N.KmsKeyId", []string{}, "[The function is only partially available, open for technical support] kms key id. Required when selecting an encryption disk.")
	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNCouponId, "Disks.N.CouponId", []string{}, "Cloud disk voucher id. Not applicable to system disk/local disk. Please check through the DescribeCoupon interface or log in to the User Center.")
	cmd.Flags().StringVar(&backen.APIArgs.NetCapability, "NetCapability", "Normal", "Network enhancement. Enumeration value: Normal (default), not enabled; Super, open network enhancement 1.0; Ultra, enable network enhancement 2.0 (only support some available areas, please refer to the console)")
	cmd.Flags().StringVar(&backen.APIArgs.VPCId, "VPCId", "", "VPC ID. The default is the default VPC for the current region.")
	cmd.Flags().StringVar(&backen.APIArgs.SubnetId, "SubnetId", "", "Subnet ID. The default is the default subnet of the current region.")
	cmd.Flags().StringSliceVar(&backen.APIArgs.PrivateIpN, "PrivateIp.N", []string{}, "[Array] Specify the intranet IP when creating a cloud host. If no value is passed, the IP under the current subnet is randomly assigned. Example of calling method: PrivateIp.0=x.x.x.x. Currently only one intranet IP is supported.")
	cmd.Flags().StringVar(&backen.APIArgs.SecurityGroupId, "SecurityGroupId", "", "Firewall Id, default: Web recommended firewall. For how to query the SecurityGroupId, see DescribeSecurityGroup.")
	cmd.Flags().IntVar(&backen.APIArgs.AlarmTemplateId, "AlarmTemplateId", 0, "If the id of the alarm template is advertised and the id of the alarm template is correct, the alarm template is bound. Failure to bind an alarm template will only result in a log in the background, which will not affect the process of creating a host, nor will it report an error on the front end.")
	cmd.Flags().IntSliceVar(&backen.APIArgs.NetworkInterfaceNEIPBandwidth, "NetworkInterface.N.EIP.Bandwidth", []int{}, "[If the EIP binding parameter is required] The external network bandwidth of the elastic IP is in Mbps. The shared bandwidth mode must specify 0M bandwidth, and the non-shared bandwidth mode must specify non-0Mbps bandwidth. The bandwidth of the non-shared bandwidth of each region is as follows: Billing [1-300], bandwidth billing [1-800]")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPPayMode, "NetworkInterface.N.EIP.PayMode", []string{}, "Elastic IP billing mode. Enumeration value: Traffic, traffic billing; Bandwidth, bandwidth billing; ShareBandwidth, shared bandwidth mode. Free: free bandwidth mode. The default is Bandwidth")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPShareBandwidthId, "NetworkInterface.N.EIP.ShareBandwidthId", []string{}, "Bind shared bandwidth Id, valid only when PayMode is ShareBandwidth")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPGlobalSSHArea, "NetworkInterface.N.EIP.GlobalSSH.Area", []string{}, "Fill in the names of the regions that support SSH access to IP, such as Los Angeles, Singapore, Hong Kong, Tokyo, Washington, Frankfurt. Area and AreaCode must fill in one")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPOperatorName, "NetworkInterface.N.EIP.OperatorName", []string{}, "[If the EIP binding parameter is required] The flexible IP line is as follows: International: International BGP: Bgp The allowed line parameters for each area are as follows: cn-sh1: Bgp cn-sh2: Bgp cn-gd: Bgp cn-bj1: Bgp Cn-bj2: Bgp hk: International us-ca: International th-bkk: International kr-seoul: International us-ws: International ge-fra: International sg: International tw-kh: International. Other overseas routes are International")
	cmd.Flags().IntSliceVar(&backen.APIArgs.NetworkInterfaceNEIPGlobalSSHPort, "NetworkInterface.N.EIP.GlobalSSH.Port", []int{}, "SSH port, 1-65535 and cannot use port 80,443")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPCouponId, "NetworkInterface.N.EIP.CouponId", []string{}, "Current EIP voucher id. Please check through the DescribeCoupon interface or log in to the User Center.")
	cmd.Flags().StringSliceVar(&backen.APIArgs.NetworkInterfaceNEIPGlobalSSHAreaCode, "NetworkInterface.N.EIP.GlobalSSH.AreaCode", []string{}, "GlobalSSH. AreaCode, Regional Air Port International General Code. Area and AreaCode must fill in one")
	cmd.Flags().StringVar(&backen.APIArgs.IsolationGroup, "IsolationGroup", "", "Hardware isolation group id. Available through the DescribeIsolationGroup.")
	cmd.Flags().StringVar(&backen.APIArgs.Name, "Name", "", "UHost instance name. Default: UHost. Please set the instance name according to the field specification.")
	cmd.Flags().StringVar(&backen.APIArgs.Tag, "Tag", "", "Business group. Default: Default (Default is ungrouped). Please set the business group according to the field specification.")
	cmd.Flags().StringVar(&backen.APIArgs.ChargeType, "ChargeType", "", "Billing mode. The enumeration value is:> Year, pay annually;> Month, paying monthly;> Dynamic, pay by the hour The default is monthly payment")
	cmd.Flags().IntVar(&backen.APIArgs.Quantity, "Quantity", 0, "The length of purchase. Default: value 1. This parameter is not required when purchasing by the hour (Dynamic). When paying monthly, this parameter is 0, which means purchase until the end of the month.")
	cmd.Flags().IntVar(&backen.APIArgs.MaxCount, "MaxCount", 0, "[Required when creating a host in batches] Maximum number of hosts created, in the range of [1,100];")
	cmd.Flags().StringVar(&backen.APIArgs.CouponId, "CouponId", "", "Host voucher ID. Please check through the DescribeCoupon interface or log in to the User Center.")
	return cmd
}

func initModifyUHostInstanceTagCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ModifyUHostInstanceTag",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into ModifyUHostInstanceTagcmd")
			backen.ModifyUHostInstanceTag()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVar(&backen.APIArgs.Tag, "Tag", "", "Business group name")
	return cmd
}
func initModifyUHostInstanceNameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ModifyUHostInstanceName",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into ModifyUHostInstanceNamecmd")
			backen.ModifyUHostInstanceName()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVar(&backen.APIArgs.Name, "Name", "", "UHost instance name")
	return cmd
}

func initLeaveIsolationGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "LeaveIsolationGroup",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into LeaveIsolationGroupcmd")
			backen.LeaveIsolationGroup()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")
	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVar(&backen.APIArgs.GroupId, "GroupId", "", "Hardware isolation group id")
	cmd.MarkFlagRequired("GroupId")
	return cmd
}

func initCreateIsolationGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "CreateIsolationGroup",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into CreateIsolationGroupCmd")
			backen.CreateIsolationGroup()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")

	cmd.Flags().StringVar(&backen.APIArgs.GroupName, "GroupName", "", "Hardware isolation group name. Please set the isolation group name according to the field specification")
	cmd.MarkFlagRequired("GroupName")
	cmd.Flags().StringVar(&backen.APIArgs.Remark, "Remark", "", "Remarks. Please follow the field specification to set the isolation group comment.")
	return cmd
}

func initResetUHostInstancePasswordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ResetUHostInstancePassword",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into ResetUHostInstancePasswordCmd")
			backen.ResetUHostInstancePassword()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")

	cmd.Flags().StringVar(&backen.APIArgs.UHostId, "UHostId", "", "uhost id as a resource")
	cmd.MarkFlagRequired("UHostId")
	cmd.Flags().StringVar(&backen.APIArgs.Password, "Password", "", "UHost new password")
	cmd.MarkFlagRequired("Password")

	return cmd
}

func initUpgradeToArkUHostInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "UpgradeToArkUHostInstance",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into UpgradeToArkUHostInstance")
			backen.UpgradeToArkUHostInstance()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")

	cmd.Flags().StringSliceVar(&backen.APIArgs.UHostIdsN, "UHostIds.n", []string{}, "The resource ID of the UHost host, for example, UHostIds.0 represents the host 1 that you want to upgrade, and UHostIds.1 represents the host 2.")
	cmd.MarkFlagRequired("UHostIds.n")
	cmd.Flags().StringVar(&backen.APIArgs.CouponId, "CouponId", "", "Voucher ID Please refer to the DescribeCoupon interface")

	return cmd
}

func initGetUHostInstancePriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "GetUHostInstancePrice",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("get into GetUHostInstancePriceCmd")
			backen.GetUHostInstancePrice()
		},
	}
	cmd.Flags().StringVarP(&backen.APIArgs.Region, "Region", "R", "", "region where host locate")
	cmd.MarkFlagRequired("Region")
	cmd.Flags().StringVarP(&backen.APIArgs.Zone, "Zone", "Z", "", "zone is a subdistrict of a region")

	cmd.Flags().StringVar(&backen.APIArgs.ImageId, "ImageId", "", "Mirror ID. Please get it through DescribeImage")
	cmd.MarkFlagRequired("ImageId")

	cmd.Flags().StringVar(&backen.APIArgs.MachineType, "MachineType", "N", "Cloud host model (V2.0), enumeration values ['N', 'C', 'G', 'O']. Refer to the cloud host model description.")

	cmd.Flags().IntVar(&backen.APIArgs.CPU, "CPU", 4, "CPU core count. Optional parameters: 1-64. The optional range is referenced to the console. Default: 4")
	cmd.MarkFlagRequired("CPU")

	cmd.Flags().IntVar(&backen.APIArgs.Memory, "Memory", 8192, "memory size. Unit: MB. Range: [1024, 262144], the value is a multiple of 1024 (optional range refers to the good console). Default: 8192")
	cmd.MarkFlagRequired("Memory")

	cmd.Flags().IntVar(&backen.APIArgs.Count, "Count", 1, "Number of purchases, range [1, 5]")
	cmd.MarkFlagRequired("Count")

	cmd.Flags().StringVar(&backen.APIArgs.GpuType, "GpuType", "", "GPU type, enumeration value [K80, P40, V100]")
	cmd.Flags().IntVar(&backen.APIArgs.GPU, "GPU", 0, "The number of GPU card cores. This field is only supported on GPU models (optional range is related to UHostType)")

	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNIsBoot, "Disks.N.IsBoot", []string{"True"}, "Whether it is a system disk. Enumeration value:> True, is the system disk> False, is the data disk (default). There is only one disk in the Disks array and it is the system disk.")
	cmd.MarkFlagRequired("Disks.N.IsBoot")

	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNType, "Disks.N.Type", []string{"LOCAL_NORMAL"}, "Disk type. Please refer to the disk type.")
	cmd.MarkFlagRequired("Disks.N.Type")

	cmd.Flags().IntSliceVar(&backen.APIArgs.DisksNSize, "Disks.N.Size", []int{20}, "Disk size in GB. Please refer to the disk type.")
	cmd.MarkFlagRequired("Disks.N.Size")

	cmd.Flags().StringSliceVar(&backen.APIArgs.DisksNBackupType, "Disks.N.BackupType", []string{"NONE"}, "Disk backup solution. Enumeration value:> NONE, no backup> DATAARK, data ark Backup mode reference supported by current disk")
	cmd.Flags().StringVar(&backen.APIArgs.NetCapability, "NetCapability", "Normal", "Network enhancement. Enumeration value: Normal (default), not enabled; Super, open network enhancement 1.0; Ultra, enable network enhancement 2.0 (only support some available areas, please refer to the console)")
	cmd.Flags().StringVar(&backen.APIArgs.ChargeType, "ChargeType", "", "Billing mode. The enumeration value is:> Year, pay annually;> Month, paying monthly;> Dynamic, pay by the hour The default is monthly payment")
	cmd.Flags().IntVar(&backen.APIArgs.Quantity, "Quantity", 0, "The length of purchase. Default: value 1. This parameter is not required when purchasing by the hour (Dynamic). When paying monthly, this parameter is 0, which means purchase until the end of the month.")

	return cmd
}
