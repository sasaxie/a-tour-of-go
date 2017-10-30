package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		ss := make([]uint8, dx)

		for j := 0; j < dx; j++ {
			ss[j] = uint8(i * j)
		}

		s[i] = ss
	}

	return s
}

func main() {
	s := Pic(3, 3)

	for _, v := range s {
		for _, vv := range v {
			fmt.Print(vv)
		}

		fmt.Println()
	}
}

/*
000
012
024
*/
