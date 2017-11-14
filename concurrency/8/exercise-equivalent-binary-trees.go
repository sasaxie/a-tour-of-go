package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch
func Walk(t *tree.Tree, ch chan int) {
	getValue(t, ch)
	close(ch)
}

func getValue(t *tree.Tree, ch chan int) {
	if t != nil {
		getValue(t.Left, ch)
		ch <- t.Value
		getValue(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
