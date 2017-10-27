package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		// fallthrough // 继续执行下面的
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s,", os)
	}
}
