package bitset

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
