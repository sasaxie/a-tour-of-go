package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"          // %T 打印变量的类型
	fmt.Printf(f, ToBe, ToBe)     // %T ：bool
	fmt.Printf(f, MaxInt, MaxInt) // %T ：uint64
	fmt.Printf(f, z, z)           // %T ：complex128
}
