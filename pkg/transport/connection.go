package transport

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn, i int) {
	fmt.Printf("connection %d start\n", i)
	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println(fmt.Sprintf("connection %d ended: %v", i, err))
			break
		}
		conn.Write([]byte(fmt.Sprintf("%d %s", i, buf[:len])))
	}
}
