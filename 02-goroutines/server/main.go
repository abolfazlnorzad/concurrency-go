package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	l, _ := net.Listen("tcp", "localhost:7777")
	for {
		c, _ := l.Accept()
		fmt.Println("Accepted !")
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	for {
		io.WriteString(c, time.Now().Format(time.ANSIC)+"\n")
		time.Sleep(time.Second)
	}
}
