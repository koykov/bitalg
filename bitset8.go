package bitset

type Bitset8 uint8

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

func (b *Bitset8) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset8(1 << pos)
	return (*b)&v != 0
}

func (b *Bitset8) Reset() {
	*b = 0
}
