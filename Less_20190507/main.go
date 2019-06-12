package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rl ResponseList

func main() {
	url := "http://baidu.com"

	go httpGet(url, 1)
	go httpGet(url, 2)
	go httpGet(url, 3)

	time.Sleep(time.Duration(1) * time.Second)

	rl.echo()
}

func httpGet(url string, index int) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get() err=", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll() err=", err)
	}

	r := &Response{
		Index:   index,
		Content: body,
		Time:    time.Now().Unix(),
	}

	rl.append(r)
}
