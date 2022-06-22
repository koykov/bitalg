package bitset

// Bitset8 represents sets of 8 bits.
type Bitset8 uint8

// SetBit sets bool value to bit on position pos [0:7].
func (b *Bitset8) SetBit(pos int, value bool) {
	if pos < 0 || pos > 7 {
		return
	}
	v := Bitset8(1 << pos)
	if value {
		*b = (*b) | v
	} else {
		*b = (*b) & ^v
	}
}

// CheckBit checks bool value of bit on position pos [0:7].
func (b *Bitset8) CheckBit(pos int) bool {
	if pos < 0 || pos > 7 {
		return false
	}
	v := Bitset8(1 << pos)
	return (*b)&v != 0
}

// Reset sets all bits to false.
func (b *Bitset8) Reset() {
	*b = 0
}
