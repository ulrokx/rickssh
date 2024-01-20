package packet

import (
	"github.com/ulrokx/rickssh/pkg/transport"
)

type Packet struct {
	PacketLength  uint32
	PaddingLength uint8
	Payload       []byte
	Padding       []byte
	MAC           []byte
}

func (p *Packet) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (p *Packet) UnmarshalBinary([]byte) error {
	return nil
}

func CreatePacketFromPayload(ctx transport.ConnContext, payload []byte) *Packet {
	if ctx.MACAlgorithm != "" {

	}
	return &Packet{}
}
