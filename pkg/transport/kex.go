package transport

import (
	"crypto/rand"
)

type KexInitPacket struct {
	Cookie                              [16]byte
	KexAlgorithms                       []string
	ServerHostKeyAlgorithms             []string
	EncryptionAlgorithmsClientToServer  []string
	EncryptionAlgorithmsServerToClient  []string
	MACAlgorithmsClientToServer         []string
	MACAlgorithmsServerToClient         []string
	CompressionAlgorithmsClientToServer []string
	CompressionAlgorithmsServerToClient []string
	LanguagesClientToServer             []string
	LanguagesServerToClient             []string
	FirstKexPacketFollows               bool
}

func (k *KexInitPacket) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (k *KexInitPacket) CreateCookie() {
	rand.Read(k.Cookie[:])
}
