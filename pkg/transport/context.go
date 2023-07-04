package transport

import (
	"net"
)

type ConnContext struct {
	Conn                 net.Conn
	SoftwareVersion      string
	SoftwareComment      string
	MACAlgorithm         string
	EncryptionAlgorithm  string
	CompressionAlgorithm string
	ConnectionID         int
}
