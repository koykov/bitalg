package bitset

type Bitset uint64

func (b *Bitset) SetBit(pos int, value bool) {
	if pos < 0 || pos > 63 {
		return
	}
	v := uint64(1 << pos)
	if value {
		*b = Bitset((uint64)(*b) | v)
	} else {
		*b = Bitset((uint64)(*b) & ^v)
	}
}

func (b *Bitset) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := uint64(1 << pos)
	return (uint64)(*b)&v != 0
}

func (b *Bitset) Reset() {
	*b = 0
}
