package ini_config

import (
	"fmt"
	"github.com/go-ini/ini"
	"testing"
)

func TestGetIniData(t *testing.T) {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	// 获取配置项的值
	section := cfg.Section("database")
	username := section.Key("username").String()
	password := section.Key("password").String()

	// 输出配置值
	fmt.Printf("Username: %s, Password: %s\n", username, password)
}

func TestSaveIniData(t *testing.T) {
	source := "./config/config.ini"
	cfg, err := ini.Load(source)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	// 获取配置项的值
	section := cfg.Section("database")
	section.Key("username").SetValue("kolnick")
	section.Key("password").SetValue("88888")
	cfg.SaveTo(source)
}
