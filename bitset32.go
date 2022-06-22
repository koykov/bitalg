package bitset

type Bitset32 uint32

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

func (b *Bitset32) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset32(1 << pos)
	return (*b)&v != 0
}

func (b *Bitset32) Reset() {
	*b = 0
}
