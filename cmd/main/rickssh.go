package main

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/ulrokx/rickssh/pkg/transport"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
	})
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
