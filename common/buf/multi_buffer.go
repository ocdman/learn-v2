package buf

type MultiBuffer []*Buffer

// ReleaseMulti release all content of the MultiBuffer, and returns an empty MultiBuffer.
func ReleaseMulti(mb MultiBuffer) MultiBuffer {
	for i := range mb {
		mb[i].Release()
		mb[i] = nil
	}
	return mb[:0]
}

// SplitBytes splits the given amount of bytes from the beginning of the MultiBuffer.
// It returns the new address of MultiBuffer leftover, and number of bytes written into the input byte slice.
func SplitBytes(mb MultiBuffer, b []byte) (MultiBuffer, int) {
	totalBytes := 0
	endIndex := -1
	for i := range mb {
		pBuffer := mb[i]
		nBytes, _ := pBuffer.Read(b)
		totalBytes += nBytes
		b = b[nBytes:]
		if !pBuffer.IsEmpty() {
			endIndex = i
			break
		}
		pBuffer.Release()
		mb[i] = nil
	}

	if endIndex == -1 {
		mb = mb[:0]
	} else {
		mb = mb[endIndex:]
	}

	return mb, totalBytes
}

// IsEmpty return true if the MultiBuffer has no content.
func (mb MultiBuffer) IsEmpty() bool {
	for _, b := range mb {
		if !b.IsEmpty() {
			return false
		}
	}
	return true
}

// MultiBufferContainer is a ReadWriteCloser wrapper over MultiBuffer.
type MultiBufferContainer struct {
	MultiBuffer
}

// ReadMultiBuffer implements Reader.
func (c *MultiBufferContainer) ReadMultiBuffer() (MultiBuffer, error) {
	mb := c.MultiBuffer
	c.MultiBuffer = nil
	return mb, nil
}
