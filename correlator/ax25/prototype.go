package ax25

import (
	"encoding/binary"
	"log"
)

type packetPrototype struct {
	header      []byte
	control     uint8
	protocol    uint8
	information []byte
	checksum    uint16
}

func (pp *packetPrototype) clear() {
	pp.header = make([]byte, 0, 70)
	pp.control = 0
	pp.protocol = 0
	pp.information = make([]byte, 0, 256)
	pp.checksum = 0
}

func (pp *packetPrototype) appendHeader(b byte) bool {
	pp.header = append(pp.header, b)
	return indicatesContinuation(b)
}

func indicatesContinuation(b byte) bool {
	return (b & 1) == 0
}

func (pp *packetPrototype) likelyValidHeader() bool {
	return (len(pp.header) >= 14) && (len(pp.header)%7 == 0)
}

func (pp *packetPrototype) appendData(b byte) {
	pp.information = append(pp.information, b)
}

func (pp *packetPrototype) extractFields() {
	if len(pp.information) >= 4 {
		pp.control = pp.information[0]
		pp.protocol = pp.information[1]
		pp.checksum = binary.BigEndian.Uint16(pp.information[len(pp.information)-2:])
		pp.information = pp.information[2 : len(pp.information)-2]
	}
}

func (pp *packetPrototype) isValid() bool {
	var crc CRC16
	crc.clear()
	crc.feedBytes(pp.header)
	crc.feedByte(pp.control)
	crc.feedByte(pp.protocol)
	crc.feedBytes(pp.information)
	return crc.get() == pp.checksum
}

func (pp *packetPrototype) print() {

	h := make([]byte, len(pp.header))
	for i, b := range pp.header {
		h[i] = b >> 1
	}
	log.Println(pp.control, pp.protocol, string(h), string(pp.information))
}
