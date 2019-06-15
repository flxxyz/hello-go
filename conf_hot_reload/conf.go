package conf_hot_reload

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

//配置结构体
type Config struct {
	Filename       string
	LastModifyTime int64
	Lock           *sync.RWMutex
	Data           interface{}
}

func NewConfig(filename string, data interface{}) *Config {
	conf := &Config{
		Filename: filename,
		Data:     data,
		Lock:     &sync.RWMutex{},
	}

	conf.parse()

	go conf.reload()

	return conf
}

//解析函数
func (c *Config) parse() bool {
	//记录最后修改时间
	fileInfo, _ := os.Stat(c.Filename)
	c.LastModifyTime = fileInfo.ModTime().Unix()

	//读取文件内容
	file, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		log.Println("读取配置文件失败:", err)
		//直接退出程序
		os.Exit(1)
	}

	//解json
	temp := c.Data
	if err = json.Unmarshal(file, &temp); err != nil {
		log.Println("解析json出错:", err)
		return false
	}

	c.Lock.Lock()
	c.Data = temp
	c.Lock.Unlock()
	log.Printf("文件内容 = %+v\n", c.Data)

	return true
}

//重载函数
func (c *Config) reload() {
	ticker := time.NewTicker(time.Second * 5)
	for _ = range ticker.C {
		func() {
			fileInfo, _ := os.Stat(c.Filename)
			currModifyTime := fileInfo.ModTime().Unix()
			if currModifyTime > c.LastModifyTime {
				if c.parse() {
					log.Println("重新加载配置文件conf.json")
				}
			}
		}()
	}
}
