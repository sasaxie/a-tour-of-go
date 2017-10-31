package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	first, second := 0, 1

	return func() int {
		temp := first

		first, second = second, first+second

		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/*
0
1
1
2
3
5
8
13
21
34
*/
