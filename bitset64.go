package bitset

type Bitset64 uint64

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

func (b *Bitset64) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset64(1 << pos)
	return (*b)&v != 0
}

func (b *Bitset64) Reset() {
	*b = 0
}
