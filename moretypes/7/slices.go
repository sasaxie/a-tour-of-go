package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

/*
s == [2 3 5 7 11 13]
s[0] == 2
s[1] == 3
s[2] == 5
s[3] == 7
s[4] == 11
s[5] == 13
*/
