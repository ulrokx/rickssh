package transport

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

func HandleConnection(conn net.Conn, i int) {
	log.Infof("handing connection # %d from %s", i, conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			log.Infof("connection %d ended: %v", i, err)
			break
		}
		log.Debugf("read %d bytes from connection # %d", len, i)
		log.Tracef("connection # %d: read %s", i, buf[:len])
		conn.Write([]byte(fmt.Sprintf("%d %s", i, buf[:len])))
	}
}
