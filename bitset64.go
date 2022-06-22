package bitset

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
