package transport

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

const PROTOCOL_VERSION = "SSH-2.0-rickssh_0.0.1"

func HandleConnection(conn net.Conn, i int) {
	log.Infof("handing connection # %d from %s", i, conn.RemoteAddr().String())
	defer func() {
		if r := recover(); r != nil {
			conn.Close()
			log.Infof("connection # %d ended: %v", i, r)
		}
	}()
	sendProtocolVersionExchange(conn, i)
	receieveProtocolVersionExchange(conn, i)
	buf := make([]byte, 32565)
	conn.Read(buf)
	log.Tracef("connection # %d: read %s", i, buf)
	for {

	}
}

func sendProtocolVersionExchange(conn net.Conn, i int) {
	log.Debugf("connection # %d: sending protocol version exchange: %s\n", i, PROTOCOL_VERSION)
	_, err := conn.Write([]byte(PROTOCOL_VERSION + "\r\n"))
	if err != nil {
		log.Infof("connection ended: %v", err)
		panic(err)
	}
}

func receieveProtocolVersionExchange(conn net.Conn, i int) {
	buf := make([]byte, 255)
	length, err := conn.Read(buf)
	log.Tracef("connection # %d: read %s", i, buf[:length])
	if err != nil {
		log.Debugf("connection ended: %v", err)
		panic(err)
	}
	id := string(buf[:length])
	versionSpaceComments := strings.Split(id, " ")
	if len(versionSpaceComments) != 2 {
		log.Debugf("error splitting on space: %s", id)
		panic("invalid protocol version exchange")
	}
	protocolVersionSoftwareVersion := strings.Split(versionSpaceComments[0], "-")
	if len(protocolVersionSoftwareVersion) != 3 {
		log.Debugf("error splitting on dash: %s", id)
		panic("invalid protocol version exchange")
	}
	if protocolVersionSoftwareVersion[0] != "SSH" {
		log.Debugf("error ssh not found: %s", id)
		panic("invalid protocol version exchange")
	}
	if protocolVersionSoftwareVersion[1] != "2.0" {
		log.Debugf("error version incorrect: %s", id)
		panic("invalid protocol version exchange")
	}
	softwareVersion := protocolVersionSoftwareVersion[2]
	comments := versionSpaceComments[1]
	log.Debugf("connection # %d: software version: %s comments: %s", i, softwareVersion, comments)
}
