package bitset

import (
	"bytes"
	"testing"
)

const (
	flagFoo = 0
	flagBar = 1
)

func TestBitset(t *testing.T) {
	t.Run("io", func(t *testing.T) {
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
	})
	t.Run("print", func(t *testing.T) {
		var b Bitset
		b.SetBit(4, true)
		b.SetBit(12, true)
		b.SetBit(19, true)
		b.SetBit(27, true)
		b.SetBit(35, true)
		b.SetBit(42, true)
		b.SetBit(51, true)
		if b.String() != "0000100000001000000100000001000000010000001000000001000000000000" {
			t.FailNow()
		}
	})
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
	b.Run("print", func(b *testing.B) {
		b.ReportAllocs()
		var (
			bs  Bitset
			buf bytes.Buffer
		)
		for i := 0; i < b.N; i++ {
			bs.SetBit(4, true)
			bs.SetBit(12, true)
			bs.SetBit(19, true)
			bs.SetBit(27, true)
			bs.SetBit(35, true)
			bs.SetBit(42, true)
			bs.SetBit(51, true)
			buf.Reset()
			_, _ = bs.Write(&buf)
			_ = buf.Bytes()
		}
	})
}
