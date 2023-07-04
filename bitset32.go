package bitset

import "io"

// Bitset32 represents sets of 32 bits.
type Bitset32 uint32

// SetBit sets bool value to bit on position pos [0:31].
func (b *Bitset32) SetBit(pos int, value bool) {
	if pos < 0 || pos > 31 {
		return
	}
	v := Bitset32(1 << pos)
	if value {
		*b = (*b) | v
	} else {
		*b = (*b) & ^v
	}
}

// CheckBit checks bool value of bit on position pos [0:31].
func (b *Bitset32) CheckBit(pos int) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	v := Bitset32(1 << pos)
	return (*b)&v != 0
}

// Reset sets all bits to false.
func (b *Bitset32) Reset() {
	*b = 0
}

// Append appends human-readable view of bitset to buf.
func (b *Bitset32) Append(buf []byte) []byte {
	for i := 0; i < 32; i++ {
		c := byte('0')
		if b.CheckBit(i) {
			c = '1'
		}
		buf = append(buf, c)
	}
	return buf
}

// Write writes human-readable view of bitset to w.
func (b *Bitset32) Write(w io.ByteWriter) (n int, err error) {
	for i := 0; i < 32; i++ {
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

// String returns human-readable view of bitset.
func (b *Bitset32) String() string {
	return str(b, 32)
}
