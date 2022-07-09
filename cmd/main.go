package main

import (
	"fmt"

	"github.com/b5710546232/bitmask/pkg"
)

func main() {
	b := pkg.NewBitMask(1, 2, 63)
	fmt.Printf("%08b\n", b)
	fmt.Println("b.Checks(1)", b.Checks(63))
}
