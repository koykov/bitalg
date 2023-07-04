package bitset

import "io"

// Bitset16 represents sets of 16 bits.
type Bitset16 uint16

// SetBit sets bool value to bit on position pos [0:15].
func (b *Bitset16) SetBit(pos int, value bool) {
	if pos < 0 || pos > 15 {
		return
	}
	v := Bitset16(1 << pos)
	if value {
		*b = (*b) | v
	} else {
		*b = (*b) & ^v
	}
}

// CheckBit checks bool value of bit on position pos [0:15].
func (b *Bitset16) CheckBit(pos int) bool {
	if pos < 0 || pos > 15 {
		return false
	}
	v := Bitset16(1 << pos)
	return (*b)&v != 0
}

// Reset sets all bits to false.
func (b *Bitset16) Reset() {
	*b = 0
}

// Append appends human-readable view of bitset to buf.
func (b *Bitset16) Append(buf []byte) []byte {
	for i := 0; i < 16; i++ {
		c := byte('0')
		if b.CheckBit(i) {
			c = '1'
		}
		buf = append(buf, c)
	}
	return buf
}

// Write writes human-readable view of bitset to w.
func (b *Bitset16) Write(w io.ByteWriter) (n int, err error) {
	for i := 0; i < 16; i++ {
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
func (b *Bitset16) String() string {
	return str(b, 16)
}
