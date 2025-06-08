package jpeg

type Chunk interface {
	Data() []byte
	isChunk()
}

// func (c Chunk) AsSegment() Segment {
// 	if c[0] != 0xff {
// 		return nil
// 	}
// 	return Segment(c)
// }

// func (c Chunk) AsEntropyCodedData() []byte {
// 	if c[0] == 0xff {
// 		return nil
// 	}
// 	return c
// }

type EntropyCodedData struct {
	data []byte
}

func (e *EntropyCodedData) Data() []byte {
	return e.data
}

func (e *EntropyCodedData) isChunk() {}

type Detritus struct {
	data []byte
}

func (d *Detritus) Data() []byte {
	return d.data
}

func (d *Detritus) isChunk() {}
