package channel

import (
	"ForeverMS/core/consts"
	"ForeverMS/core/maplepacket"
	"ForeverMS/core/netio"
	"ForeverMS/core/packet79"

	"github.com/juju/errors"
	"github.com/ngaut/log"
)

func (s *ChannelServer) handleConnEvent(evt *netio.ConnEvent) error {
	switch evt.Type {
	case netio.CEConnected:
		{
			return errors.Trace(s.handleEvtConnected(evt))
		}
	case netio.CEDisconnected:
		{
			return errors.Trace(s.handleEvtDisconnected(evt))
		}
	case netio.CERecv:
		{
			return errors.Trace(s.handleEvtRecv(evt))
		}
	}
	return nil
}

func (s *ChannelServer) handleEvtConnected(evt *netio.ConnEvent) error {
	// Once connected, send handshake response
	evt.Conn.Send(packet79.NewHello(consts.ServerVersion,
		evt.Conn.GetSendCipher().GetKey(),
		evt.Conn.GetRecvCipher().GetKey(),
		s.config.TestServer))
	return nil
}

func (s *ChannelServer) handleEvtDisconnected(evt *netio.ConnEvent) error {
	return nil
}

func (s *ChannelServer) handleEvtRecv(evt *netio.ConnEvent) error {
	reader := maplepacket.NewReader(&evt.Packet)
	opcode := reader.ReadInt16()
	// Dispatch packet with opcode
	handler, ok := s.packetDispatchMap[opcode]
	if ok && nil != handler {
		log.Debugf("Handle opcode 0x%04x", opcode)
		return errors.Trace(handler(evt.Conn, &reader))
	}
	log.Warningf("Opcode %d do not setup packet handler", opcode)
	return nil
}
