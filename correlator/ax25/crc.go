package ax25

var lookup []uint16 = []uint16{
	0x0000, 0x1081, 0x2102, 0x3183,
	0x4204, 0x5285, 0x6306, 0x7387,
	0x8408, 0x9489, 0xa50a, 0xb58b,
	0xc60c, 0xd68d, 0xe70e, 0xf78f,
}

type CRC16 struct {
	crc uint16
}

func (obj *CRC16) clear() {
	obj.crc = 0xffff
}

func (obj *CRC16) feedByte(b byte) {
	obj.crc = (obj.crc >> 4) ^ lookup[int((obj.crc&0xf)^uint16(b&0xf))]
	obj.crc = (obj.crc >> 4) ^ lookup[int((obj.crc&0xf)^uint16(b>>4))]
}

func (obj *CRC16) feedBytes(bytes []byte) {
	for _, b := range bytes {
		obj.feedByte(b)
	}
}

func (obj *CRC16) get() uint16 {
	return ((obj.crc << 8) | ((obj.crc >> 8) & 0xff)) ^ 0xffff
}
