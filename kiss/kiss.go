package kiss

const (
	FEND  = 0xC0
	FESC  = 0xDB
	TFEND = 0xDC
	TFESC = 0xDD
)

const (
	DataFrame = 0
	TXDELAY   = 1
	PERSIST   = 2
	SLOTTIME  = 3
	TXTAIL    = 4
	FULLDUP   = 5
	SetH      = 6
	Return    = 255
)

func Encode(port byte, command byte, data []byte) []byte {
	frame := make([]byte, 0, 3+2*len(data))
	frame = append(frame, FEND)

	typeByte := (command & 0xf) | ((port & 0xf) << 4)
	frame = append(frame, typeByte)

	for _, b := range data {
		if b == FEND {
			frame = append(frame, FESC)
			frame = append(frame, TFEND)
		} else if b == FESC {
			frame = append(frame, FESC)
			frame = append(frame, TFESC)
		} else {
			frame = append(frame, b)
		}
	}

	frame = append(frame, FEND)
	return frame
}
