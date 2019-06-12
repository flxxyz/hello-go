package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://baidu.com"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "punkbot/0.1 (+https://punkbot.buger.dev)")
	client := http.Client{}
	r, _ := client.Do(req)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
}
