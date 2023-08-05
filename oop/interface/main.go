package main

import "fmt"

/**
interface 是GO语言的基础特性之一。可以理解为一种类型的规范或者约定。
它跟java，C# 不太一样，不需要显示说明实现了某个接口，它没有继承或子类或“implements”关键字，
只是通过约定的形式，隐式的实现interface 中的方法即可。因此，Golang 中的 interface 让编码更灵活、易扩展。

 如何理解go 语言中的interface ？ 只需记住以下三点即可：
1. interface 是方法声明的集合
2. 任何类型的对象实现了在interface 接口中声明的全部方法，则表明该类型实现了该接口。
3. interface 可以作为一种数据类型，实现了该接口的任何对象都可以给对应的接口类型变量赋值。

注意：
　　a. interface 可以被任意对象实现，一个类型/对象也可以实现多个 interface
　　b. 方法不能重载，如 eat(), eat(s string) 不能同时存在
*/

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {
}

func (iPhone ApplePhone) call() {
	fmt.Println("I am Apple Phone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	
	phone = new(ApplePhone)
	phone.call()
}
