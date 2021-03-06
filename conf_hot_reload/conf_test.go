package conf_hot_reload_test

import (
	"github.com/flxxyz/hello-go/conf_hot_reload"
	"testing"
)

func TestRun(t *testing.T) {
	//json配置文件结构体
	type Content struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Passwd string `json:"passwd"`
		Db     int    `json:"db"`
	}

	data := &Content{
		Host:   "127.0.0.1",
		Port:   6379,
		Passwd: "",
		Db:     0,
	}
	_ = conf_hot_reload.NewConfig("conf.json", data)

	select {}

	//end := make(chan bool, 1)
	//<-end
}
