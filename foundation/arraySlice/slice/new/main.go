package main

import "fmt"

/**
切片指针的解引用。
可以使用list:=make([]int,0) list类型为切片
或使用*list = append(*list, 1) list类型为指针

new和make的区别：
二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；
而new用于类型的内存分配，并且内存置为零。所以在我们编写程序的时候，就可以根据自己的需要很好的选择了。

make返回的还是这三个引用类型本身；而new返回的是指向类型的指针。
*/
func main() {

	list := new([]int)

	//list = append(list, 1)
	//list := make([]int, 0)
	*list = append(*list, 1)
	fmt.Println(list)
}
