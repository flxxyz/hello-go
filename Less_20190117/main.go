package main

import (
    "log"
    "net"
)

func init() {

}

func main() {
    log.Println("begin dial...")
    conn, err := net.Dial("tcp", ":7071")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")
}
