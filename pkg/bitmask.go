package pkg

import "fmt"

type BitMask uint64

func NewBitMask(ids ...uint64) BitMask {
	var b BitMask
	for _, id := range ids {
		b.Set(id)
	}
	return b
}

func (b *BitMask) Set(flag uint64) BitMask {
	*b |= BitMask(1 << flag)
	return *b
}

func (b *BitMask) Clear(flag uint64) BitMask {
	*b = *b & BitMask(1<<flag)
	return *b
}

func (b *BitMask) Toggle(flag uint64) BitMask {
	*b = *b ^ BitMask(1<<flag)
	return *b
}

func (b *BitMask) Checks(flag uint64) bool {
	val := *b & BitMask(1<<flag)
	return val != 0
}

func (b *BitMask) Print() {
	fmt.Printf("%08b\n", *b)
}
