package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	gcfg "gopkg.in/gcfg.v1"
)

const (
	ConfigPath = "config/config.ctmpl"

	ProjectName = "tkpd-recom-engine"
)

var cfg *Config

func init() {
	cfg = &Config{}
}

func GetConfig() *Config {
	return cfg
}

func InitConfig() (err error) {
	ok := readConfig(cfg, ConfigPath) || readConfig(cfg, getCurrentDirPath()+ConfigPath)
	if !ok {
		fmt.Errorf("failed to read config file")
	}
	return nil
}

func readConfig(cfg interface{}, filePath string) bool {
	err := gcfg.ReadFileInto(cfg, filePath)
	if err == nil {
		log.Println("read config from ", filePath)
		return true
	}

	return false
}

func getCurrentDirPath() string {
	path := ""
	curDir, _ := os.Getwd()
	dirNames := strings.Split(curDir, "/")

	for _, dirName := range dirNames {
		path += dirName + "/"
		if dirName == ProjectName {
			break
		}
	}
	return path
}
