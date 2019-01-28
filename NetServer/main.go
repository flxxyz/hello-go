package main

import (
    "github.com/flxxyz/hello/NetServer/tcp"
)

type Data struct {
}

var conf *Config

func init() {
    //conf = NewConfig("conf.json", Data{})
}

func main() {
    //log.Println("begin dial...")
    //conn, err := net.Dial("tcp", "baidu.com:80")
    //if err != nil {
    //    log.Println("dial error:", err)
    //    return
    //}
    //defer conn.Close()
    //log.Println("dial ok")

    //l, err := net.Listen("tcp", ":7071")
    //if err != nil {
    //    fmt.Println("listen error:", err)
    //    return
    //}
    //
    //var i int
    //for {
    //    _, err := l.Accept()
    //    if err != nil {
    //        fmt.Println("accept error:", err)
    //        break
    //    }
    //
    //    i++
    //    log.Printf("%d: accept a new connection\n", i)
    //
    //    //go handleConn(c)
    //}

    protocol := "tcp"
    addr := "127.0.0.1:9999"
    tcp.Run(protocol, addr)
}
