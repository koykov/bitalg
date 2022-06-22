package bitset

import "testing"

const (
	flagFoo = 0
	flagBar = 1
)

func TestBitset(t *testing.T) {
	var b Bitset
	b.SetBit(flagFoo, true)
	if !b.CheckBit(flagFoo) {
		t.Error("bit mismatch flagFoo")
	}
	if b.CheckBit(flagBar) {
		t.Error("bit mismatch flagBar")
	}
	b.Reset()
	if b.CheckBit(flagFoo) {
		t.Error("bit mismatch flagFoo")
	}
}

func BenchmarkBitset(b *testing.B) {
	b.Run("set", func(b *testing.B) {
		var s Bitset
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.SetBit(flagFoo, true)
		}
	})
	b.Run("check", func(b *testing.B) {
		var s Bitset
		s.SetBit(flagFoo, true)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.CheckBit(flagFoo)
		}
	})
}
