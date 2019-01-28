package tcp

import (
    "fmt"
    "net"
)

const MaxRead = 1024

func Run(protocol string, addr string) {
    l, _ := net.Listen(protocol, addr)
    defer l.Close()

    for {
        tcpConn, err := l.Accept()
        if err != nil {
            continue
        }

        fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())

        go handleConnection(tcpConn)
    }
}

func handleConnection(c net.Conn) {
    gusetIp := c.RemoteAddr().String()

    for {
        buf := make([]byte, MaxRead+1)
        len, err := c.Read(buf[0:MaxRead])
        buf[MaxRead] = 0
        fmt.Println("消息内容=", string(buf), ", 消息长度=", len, ", err=", err)
        switch err {
        case nil:
            n, _ := c.Write(handleMsg(buf))
            fmt.Println("write", n)
        default:
            goto DISCONNECT
        }
    }

DISCONNECT:
    defer func() {
        fmt.Println("disconnected :" + gusetIp)
        c.Close()
    }()
}

func handleMsg(content []byte) []byte {
    return []byte("ase key balabala")
}
