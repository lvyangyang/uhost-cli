package cmd

import (
	"fmt"
	"os"

	"github.com/lvyangyang/uhost_cli/backen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var publicKey = "default publickey"
var privateKey = "default privatekey"
var projectid = "default projectid"
var cfgFile string
var defaultCfgFile = "config.json"

func initCmdRoot() *cobra.Command {
	cobra.OnInitialize(initConfig)
	cmd := &cobra.Command{
		Use:   "uhost",
		Short: "uhost cli util use ucloud api",
		Long:  "full api function about uhost",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		//	Run: func(cmd *cobra.Command, args []string) {
		//		fmt.Println("get into root cmd/n")
		//	},
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.json)")
	cmd.PersistentFlags().StringVar(&backen.APIArgs.PublicKey, "publickey", "", "user PublicKey")
	cmd.PersistentFlags().StringVar(&backen.APIArgs.PrivateKey, "privatekey", "", "user PrivateKey")
	cmd.PersistentFlags().StringVar(&backen.APIArgs.ProjectId, "projectid", projectid, "user projectid")

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
	cmd.AddCommand(initGetUHostInstancePriceCmd())
	cmd.AddCommand(initDescribeUHostInstanceCmd())
	cmd.AddCommand(initCreateCustomImageCmd())
	cmd.AddCommand(initGetAttachedDiskUpgradePriceCmd())
	cmd.AddCommand(initCopyCustomImageCmd())
	cmd.AddCommand(initDescribeImageCmd())
	cmd.AddCommand(initDeleteIsolationGroupCmd())
	cmd.AddCommand(initModifyUHostInstanceRemarkCmd())
	cmd.AddCommand(initResizeAttachedDiskCmd())
	cmd.AddCommand(initImportCustomImageCmd())
	cmd.AddCommand(initPoweroffUHostInstanceCmd())
	cmd.AddCommand(initReinstallUHostInstanceCmd())
	cmd.AddCommand(initGetUHostUpgradePriceCmd())
	cmd.AddCommand(initTerminateCustomImageCmd())
	cmd.AddCommand(initDescribeIsolationGroupCmd())
	return cmd
}

func flagConfigSet(cmd *cobra.Command, args []string) {
	if !cmd.Flags().Lookup("publickey").Changed {
		if viper.Get("publickey") != nil {
			backen.APIArgs.PublicKey = viper.Get("publickey").(string)
		}
	}

	if !cmd.Flags().Lookup("privatekey").Changed {

		if viper.Get("privatekey") != nil {
			backen.APIArgs.PrivateKey = viper.Get("privatekey").(string)
		}
	}

	if !cmd.Flags().Lookup("projectid").Changed {

		if viper.Get("projectid") != nil {
			backen.APIArgs.ProjectId = viper.Get("projectid").(string)
		}
	}
	if !cmd.Flags().Lookup("Region").Changed {

		if viper.Get("region") != nil {
			backen.APIArgs.Region = viper.Get("region").(string)
		}
	}
	if !cmd.Flags().Lookup("Zone").Changed {

		if viper.Get("zone") != nil {
			backen.APIArgs.Zone = viper.Get("zone").(string)
		}
	}
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = defaultCfgFile
	}
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//Execute 执行入口
func Execute() {
	rootCmd := initCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
