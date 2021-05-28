package bitset

type Bitset uint64

func (b *Bitset) SetBit(pos int, value bool) {
	if pos < 0 || pos > 63 {
		return
	}
	v := Bitset(1 << pos)
	if value {
		*b = (*b) | v
	} else {
		*b = (*b) & ^v
	}
}

func (b *Bitset) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset(1 << pos)
	return (*b)&v != 0
}

func (b *Bitset) Reset() {
	*b = 0
}
