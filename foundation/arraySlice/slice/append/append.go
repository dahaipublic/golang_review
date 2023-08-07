package main

import (
	"fmt"
)

func main() {
	//切片追加, make初始化均为0
	s := make([]int, 10)

	s = append(s, 1, 2, 3)

	fmt.Println(s)

	// 两个slice在append的时候，记住需要进行将第二个slice进行...打散再拼接。
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
