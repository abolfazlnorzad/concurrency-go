package main

import (
	"io"
	"net"
	"os"
)

func main() {
	c, _ := net.Dial("tcp", "localhost:7777")
	io.Copy(os.Stdout, c)
}
