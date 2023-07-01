package main

import (
	"flag"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/ulrokx/rickssh/pkg/transport"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	debug := flag.Bool("d", false, "enable debug logging")
	trace := flag.Bool("t", false, "enable trace logging")
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	}
	if *trace {
		log.SetLevel(log.TraceLevel)
	}

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
