package main

import (
	"github.com/Luoxin/Eutamias/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig(path string) error {
	if path == "" {
		// 可能存在的目录
		viper.AddConfigPath("./")
		viper.AddConfigPath("../")
		viper.AddConfigPath("./conf/")
		viper.AddConfigPath("../conf/")

		viper.AddConfigPath(utils.GetExecPath())
		viper.AddConfigPath(utils.GetPwd())

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	} else {
		viper.SetConfigFile(path)
	}

	// 配置一些默认值
	viper.SetDefault("dns_nameserver", append(dnsClientList, dnsJsonApiList...))
	viper.SetDefault("domain_list", githubList)
	viper.SetDefault("use_default_domain", true)

	err := viper.ReadInConfig()
	if err != nil {
		switch e := err.(type) {
		case viper.ConfigFileNotFoundError:
			pterm.Warning.Printfln("not found conf file, use default")
		case *os.PathError:
			pterm.Warning.Printfln("not find conf file in %s", e.Path)
		default:
			pterm.Error.Printfln("load config fail:%v", err)
			return err
		}
	}

	return nil
}
