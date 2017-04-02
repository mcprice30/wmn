package data

// BufferSize indicates how many elements to store at once.
const BufferSize = 64

// Buffer is responsible for storing up to BufferSize elements at once, while
// maintaining their original order.
// It is implemented using a circular array.
type Buffer struct {
	buff []SensorData
	idx  int
}

// CreateBuffer will create and return a new Buffer.
func CreateBuffer() *Buffer {
	return &Buffer{
		buff: make([]SensorData, BufferSize),
		idx:  0,
	}
}

// Add will add the given data element to the buffer.
func (b *Buffer) Add(d SensorData) {
	b.buff[b.idx%BufferSize] = d
	b.idx++
}

// GetData will return all elements in the buffer, in the order that they
// were added to the buffer.
func (b *Buffer) GetData() []SensorData {
	outSize := BufferSize
	startIdx := (b.idx - BufferSize) % BufferSize
	if b.idx < BufferSize {
		outSize = b.idx
		startIdx = 0
	}
	out := make([]SensorData, outSize)
	for i := range out {
		out[i] = b.buff[(i+startIdx)%BufferSize]
	}
	return out
}
