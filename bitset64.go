package bitset

import (
	"io"
)

// Bitset64 represents sets of 64 bits.
type Bitset64 uint64

// SetBit sets bool value to bit on position pos [0:63].
func (b *Bitset64) SetBit(pos int, value bool) {
	if pos < 0 || pos > 63 {
		return
	}
	v := Bitset64(1 << pos)
	if value {
		*b = (*b) | v
	} else {
		*b = (*b) & ^v
	}
}

// CheckBit checks bool value of bit on position pos [0:63].
func (b *Bitset64) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset64(1 << pos)
	return (*b)&v != 0
}

// Reset sets all bits to false.
func (b *Bitset64) Reset() {
	*b = 0
}

// Append appends human-readable view of bitset to buf.
func (b *Bitset64) Append(buf []byte) []byte {
	return append_(b, buf, 64)
}

// Write writes human-readable view of bitset to w.
func (b *Bitset64) Write(w io.ByteWriter) (n int, err error) {
	return write(b, w, 64)
}

// String returns human-readable view of bitset.
func (b *Bitset64) String() string {
	return str(b, 64)
}
