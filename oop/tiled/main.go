package main

import "fmt"

/**
平铺式的模块设计
那么作为interface数据类型，他存在的意义在哪呢？
实际上是为了满足一些面向对象的编程思想。
我们知道，软件设计的最高目标就是高内聚，低耦合。
那么其中有一个设计原则叫开闭原则。什么是开闭原则呢，接下来我们看一个例子：
*/

// Banker 我们要写一个类,Banker银行业务员
type Banker struct {
}

// Save 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// Transfer 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

// Pay 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

func main() {
	banker := &Banker{}
	
	banker.Save()
	banker.Transfer()
	banker.Pay()
}
