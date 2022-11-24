package ax25

type state int

const (
	none      state = 0
	header    state = 10
	remainder state = 20
)

const flagByte = 0b01111110

type AX25Correlator struct {
	state             state
	stateBytesCounter int
	packetPrototype   packetPrototype
}

func (cor *AX25Correlator) Initialize() {
	cor.packetPrototype.clear()
}

func (cor *AX25Correlator) RX(b byte) []byte {
	packet := []byte{}

	b = bitrev(b)
	cor.stateBytesCounter++

	switch cor.state {
	case none:
		if b == flagByte {
			cor.packetPrototype.clear()
			cor.state = header
			cor.stateBytesCounter = 0
		}

	case header:
		if b == flagByte {
			cor.stateBytesCounter--
			break
		}
		continuation := cor.packetPrototype.appendHeader(b)
		if !continuation {
			if cor.packetPrototype.likelyValidHeader() {
				cor.state = remainder
			} else {
				cor.state = none
			}
			cor.stateBytesCounter = 0
		}

	case remainder:
		if b == flagByte {
			cor.state = none
			cor.stateBytesCounter = 0
			cor.packetPrototype.extractFields()
			validPacket := cor.packetPrototype.isValid()
			if validPacket {
				cor.packetPrototype.print()
				packet = cor.packetPrototype.packet()
			}
			cor.packetPrototype.clear()
		} else {
			cor.packetPrototype.appendData(b)
		}
	}

	return packet
}

func bitrev(n byte) byte {
	tmp := byte(0x55)
	n = (((n >> 1) & tmp) | ((n & tmp) << 1))
	tmp = byte(0x33)
	n = (((n >> 2) & tmp) | ((n & tmp) << 2))
	return ((n >> 4) | (n << 4))
}
