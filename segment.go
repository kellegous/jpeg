package jpeg

type Segment []byte

type SegmentType byte

func (s Segment) Type() SegmentType {
	return SegmentType(s[1])
}

func (s Segment) Len() int {
	return len(s)
}

func (s Segment) Data() []byte {
	return s
}

// TODO(kellegous): DAC, SIZ?
const (
	SOI   SegmentType = 0xd8
	SOF0  SegmentType = 0xc0
	SOF1  SegmentType = 0xc1
	SOF2  SegmentType = 0xc2
	DHT   SegmentType = 0xc4
	DQT   SegmentType = 0xdb
	DRI   SegmentType = 0xdd
	SOS   SegmentType = 0xda
	RST0  SegmentType = 0xd0
	RST1  SegmentType = 0xd1
	RST2  SegmentType = 0xd2
	RST3  SegmentType = 0xd3
	RST4  SegmentType = 0xd4
	RST5  SegmentType = 0xd5
	RST6  SegmentType = 0xd6
	RST7  SegmentType = 0xd7
	APP0  SegmentType = 0xe0
	APP1  SegmentType = 0xe1
	APP2  SegmentType = 0xe2
	APP3  SegmentType = 0xe3
	APP4  SegmentType = 0xe4
	APP5  SegmentType = 0xe5
	APP6  SegmentType = 0xe6
	APP7  SegmentType = 0xe7
	APP8  SegmentType = 0xe8
	APP9  SegmentType = 0xe9
	APP10 SegmentType = 0xea
	APP11 SegmentType = 0xeb
	APP12 SegmentType = 0xec
	APP13 SegmentType = 0xed
	APP14 SegmentType = 0xee
	APP15 SegmentType = 0xef
	COM   SegmentType = 0xfe
	EOI   SegmentType = 0xd9
)

func (t SegmentType) Name() string {
	switch t {
	case SOI:
		return "SOI"
	case SOF0:
		return "SOF0"
	case SOF1:
		return "SOF1"
	case SOF2:
		return "SOF2"
	case DHT:
		return "DHT"
	case DQT:
		return "DQT"
	case DRI:
		return "DRI"
	case SOS:
		return "SOS"
	case RST0:
		return "RST0"
	case RST1:
		return "RST1"
	case RST2:
		return "RST2"
	case RST3:
		return "RST3"
	case RST4:
		return "RST4"
	case RST5:
		return "RST5"
	case RST6:
		return "RST6"
	case RST7:
		return "RST7"
	case APP0:
		return "APP0"
	case APP1:
		return "APP1"
	case APP2:
		return "APP2"
	case APP3:
		return "APP3"
	case APP4:
		return "APP4"
	case APP5:
		return "APP5"
	case APP6:
		return "APP6"
	case APP7:
		return "APP7"
	case APP8:
		return "APP8"
	case APP9:
		return "APP9"
	case APP10:
		return "APP10"
	case APP11:
		return "APP11"
	case APP12:
		return "APP12"
	case APP13:
		return "APP13"
	case APP14:
		return "APP14"
	case APP15:
		return "APP15"
	case COM:
		return "COM"
	case EOI:
		return "EOI"
	}
	return ""
}
