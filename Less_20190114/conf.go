package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type redisStruct struct {
	Addr   string `json:"addr"`
	Passwd string `json:"passwd"`
	Db     int    `json:"db"`
}

type config struct {
	Redis redisStruct `json:"redis"`
}

func NewConf() *config {
	return &config{
		Redis: redisStruct{
			Addr: "localhost:6379",
			Passwd: "",
			Db: 0,
		},
	}
}

func (c *config) Load(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("config.json读取错误")
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatal("json解析失败")
	}
}
