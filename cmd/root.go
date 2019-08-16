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
	rootCmd := &cobra.Command{
		Use:   "ucloudcli",
		Short: "ucloudcli is an ucloudapi cli tool",
		Long:  "ucloud api control tool",
		PreRun: func(cmd *cobra.Command, args []string) {
			flagConfigSet(cmd, args)
		},
		//	Run: func(cmd *cobra.Command, args []string) {
		//		fmt.Println("get into root cmd/n")
		//	},
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.json)")
	rootCmd.PersistentFlags().StringVar(&backen.APIArgs.PublicKey, "publickey", "", "user PublicKey")
	rootCmd.PersistentFlags().StringVar(&backen.APIArgs.PrivateKey, "privatekey", "", "user PrivateKey")
	rootCmd.PersistentFlags().StringVar(&backen.APIArgs.ProjectId, "projectid", projectid, "user projectid")

	rootCmd.AddCommand(initUhostCmd())
	return rootCmd
}

func flagConfigSet(cmd *cobra.Command, args []string) {
	if !cmd.Flags().Lookup("publickey").Changed {
		if viper.Get("privatekey") != nil {
			backen.APIArgs.PublicKey = viper.Get("publickey").(string)
		}
	}

	if !cmd.Flags().Lookup("privatekey").Changed {

		if viper.Get("publickey") != nil {
			backen.APIArgs.PrivateKey = viper.Get("privatekey").(string)
		}
	}

	if !cmd.Flags().Lookup("projectid").Changed {

		if viper.Get("projectid") != nil {
			backen.APIArgs.ProjectId = viper.Get("projectid").(string)
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
