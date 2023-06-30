package main

import (
	"fmt"
	"net"

	"github.com/ulrokx/rickssh/pkg/transport"
)

func main() {
	listener, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	for i := 0; ; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go transport.HandleConnection(conn, i)
	}
}
