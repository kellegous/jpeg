package jpeg

type Chunk interface {
	Data() []byte
	isChunk()
}

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
