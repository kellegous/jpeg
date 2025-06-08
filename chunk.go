package jpeg

type Chunk []byte

func (c Chunk) AsSegment() Segment {
	if c[0] != 0xff {
		return nil
	}
	return Segment(c)
}

func (c Chunk) AsEntropyCodedData() []byte {
	if c[0] == 0xff {
		return nil
	}
	return c
}
