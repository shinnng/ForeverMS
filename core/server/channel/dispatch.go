package channel

import (
	"ForeverMS/core/consts/opcode"
	"ForeverMS/core/maplepacket"
	"ForeverMS/core/netio"
)

type packetHandler func(netio.IConn, *maplepacket.Reader) error

func (s *ChannelServer) initPacketDispatchMap() {
	s.packetDispatchMap = map[int16]packetHandler{
		opcode.RecvOps.PONG:            s.handlePong,
		opcode.RecvOps.PLAYER_LOGGEDIN: s.handlePlayerLoggedIn,
	}
}
