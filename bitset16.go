package bitset

type Bitset16 uint16

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

func (b *Bitset16) CheckBit(pos int) bool {
	if pos < 0 || pos > 63 {
		return false
	}
	v := Bitset16(1 << pos)
	return (*b)&v != 0
}

func (b *Bitset16) Reset() {
	*b = 0
}
