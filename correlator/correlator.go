package correlator

type Correlator interface {
	Initialize()
	Feed(receivedByte byte) (optionalPacket []byte)
}
