package transport

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

const PROTOCOL_VERSION = "SSH-2.0-rickssh_0.0.1"

func HandleConnection(ctx *ConnContext) {
	log.Infof("handing connection # %d from %s", ctx.ConnectionID, ctx.Conn.RemoteAddr().String())
	defer func() {
		if r := recover(); r != nil {
			ctx.Conn.Close()
			log.Infof("connection # %d ended: %v", ctx.ConnectionID, r)
		}
	}()
	sendProtocolVersionExchange(ctx)
	receieveProtocolVersionExchange(ctx)
	buf := make([]byte, 32565)
	len, _ := ctx.Conn.Read(buf)
	log.Tracef("connection # %d: read %08b", ctx.ConnectionID, buf[:len])
	log.Tracef("context: %#v", ctx)
	for {

	}
}

func sendProtocolVersionExchange(ctx *ConnContext) {
	log.Debugf("connection # %d: sending protocol version exchange: %s\n", ctx.ConnectionID, PROTOCOL_VERSION)
	_, err := ctx.Conn.Write([]byte(PROTOCOL_VERSION + "\r\n"))
	if err != nil {
		log.Infof("connection ended: %v", err)
		panic(err)
	}
}

func receieveProtocolVersionExchange(ctx *ConnContext) {
	buf := make([]byte, 255)
	length, err := ctx.Conn.Read(buf)
	log.Tracef("connection # %d: read %s", ctx.ConnectionID, buf[:length])
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
	comments := strings.Split(versionSpaceComments[1], "\r\n")[0]
	ctx.SoftwareVersion = softwareVersion
	ctx.SoftwareComment = comments
	log.Debugf("connection # %d: software version: %s comments: %s", ctx.ConnectionID, softwareVersion, comments)
}
