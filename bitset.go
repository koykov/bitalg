package bitset

type Interface interface {
	SetBit(int, bool)
	CheckBit(int) bool
	Reset()
}

type Bitset = Bitset64
