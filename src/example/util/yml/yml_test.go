package yml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)

// 使用yaml 需要安装 go get gopkg.in/yaml.v2
type Config struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database struct {
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func TestLoadYml(t *testing.T) {
	// 读取 YAML 文件
	yamlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 解析 YAML 文件
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	// 打印配置信息
	fmt.Printf("Server: %s\n", config.Server)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Name: %s\n", config.Database.Name)
	fmt.Printf("Database User: %s\n", config.Database.User)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
}
