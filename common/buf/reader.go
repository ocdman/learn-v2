package buf

import (
	"io"
)

func readOneUDP(r io.Reader) (*Buffer, error) {
	b := New()
	for i := 0; i < 64; i++ {
		_, err := b.ReadFrom(r)
		if !b.IsEmpty() {
			return b, nil
		}
		if err != nil {
			b.Release()
			return nil, err
		}
	}

	b.Release()
	return nil, newError("Reader returns too many empty payloads.")
}

// ReadBuffer reads a Buffer from the given reader.
func ReadBuffer(r io.Reader) (*Buffer, error) {
	b := New()
	n, err := b.ReadFrom(r)
	if n > 0 {
		return b, err
	}
	b.Release()
	return nil, err
}

// BufferedReader is a Reader that keeps its internal buffer.
type BufferedReader struct {
	// Reader is the underlying reader to be read from
	Reader Reader
	// Buffer is the internal buffer to be read from first
	Buffer MultiBuffer
	// Spliter is a function to read bytes from MultiBuffer
	Spliter func(MultiBuffer, []byte) (MultiBuffer, int)
}

// ReadByte implements io.ByteReader.
func (r *BufferedReader) ReadByte() (byte, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err
}

func (r *BufferedReader) Read(b []byte) (int, error) {
	spliter := r.Spliter
	if spliter == nil {
		spliter = SplitBytes
	}

	if !r.Buffer.IsEmpty() {
		buffer, nBytes := spliter(r.Buffer, b)
		r.Buffer = buffer
		if r.Buffer.IsEmpty() {
			r.Buffer = nil
		}
		return nBytes, nil
	}

	mb, err := r.Reader.ReadMultiBuffer()
	if err != nil {
		return 0, err
	}

	mb, nBytes := spliter(mb, b)
	if !mb.IsEmpty() {
		r.Buffer = mb
	}
	return nBytes, nil
}

// SingleReader is a Reader that read one Buffer every time.
type SingleReader struct {
	io.Reader
}

// ReadMultiBuffer implements Reader.
func (r *SingleReader) ReadMultiBuffer() (MultiBuffer, error) {
	b, err := ReadBuffer(r.Reader)
	return MultiBuffer{b}, err
}

// PacketReader is a Reader that read one Buffer every time.
type PacketReader struct {
	io.Reader
}

// ReadMultiBuffer implements Reader.
func (r *PacketReader) ReadMultiBuffer() (MultiBuffer, error) {
	b, err := readOneUDP(r.Reader)
	if err != nil {
		return nil, err
	}
	return MultiBuffer{b}, nil
}
