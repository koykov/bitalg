package bitset

import (
	"bytes"
	"io"
)

type Interface interface {
	SetBit(int, bool)
	CheckBit(int) bool
	Reset()
	Write(writer io.ByteWriter) (int, error)
	Append([]byte) []byte
	String() string
}

// Bitset is an alias of Bitset64.
type Bitset = Bitset64

func write(b Interface, w io.ByteWriter, size int) (n int, err error) {
	for i := 0; i < size; i++ {
		c := byte('0')
		if b.CheckBit(i) {
			c = '1'
		}
		if err = w.WriteByte(c); err != nil {
			return
		}
		n++
	}
	return
}

func append_(b Interface, buf []byte, size int) []byte {
	for i := 0; i < size; i++ {
		c := byte('0')
		if b.CheckBit(i) {
			c = '1'
		}
		buf = append(buf, c)
	}
	return buf
}

func str(b Interface, size int) string {
	var buf bytes.Buffer
	buf.Grow(size)
	buf.Reset()
	_, _ = b.Write(&buf)
	return buf.String()
}
