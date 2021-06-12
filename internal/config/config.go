package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	// C 配置文件
	C    = new(CType)
	once sync.Once
)

// InitConfig 获得配置
func InitConfig(mode string) {
	once.Do(func() {

		path := ""
		if mode == "dev" {
			path = "application_dev.yml"
		} else {
			path = "application.yml"
		}

		file, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err.Error())
		}

		err = yaml.Unmarshal(file, &C)
		if err != nil {
			panic(err.Error())
		}
		C.Mode = mode
	})
}
