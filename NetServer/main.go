package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

type Data struct {

}

var conf *Config

func init() {
    //conf = NewConfig("conf.json", Data{})
}

func main() {
    log.Println("begin dial...")
    conn, err := net.Dial("tcp", "baidu.com:80")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")


    l, err := net.Listen("tcp", ":7071")
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }

    var i int
    for {
        _, err := l.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            break
        }

        i++
        log.Printf("%d: accept a new connection\n", i)

        //go handleConn(c)
    }
}

func handleConn(c net.Conn)  {
    defer c.Close()

    for {

        t := time.Unix(60, 0)
        c.SetDeadline(t)
    }
}
