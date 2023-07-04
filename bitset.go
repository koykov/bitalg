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

func str(b Interface, size int) string {
	var buf bytes.Buffer
	buf.Grow(size)
	buf.Reset()
	_, _ = b.Write(&buf)
	return buf.String()
}
