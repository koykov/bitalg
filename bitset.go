package bitset

type Interface interface {
	SetBit(int, bool)
	CheckBit(int) bool
	Reset()
}

// Bitset is an alias of Bitset64.
type Bitset = Bitset64
