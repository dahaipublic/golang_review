package main

import "fmt"

/*
面向抽象层依赖倒转
如果我们在设计一个系统的时候，将模块分为3个层次，抽象层、实现层、业务逻辑层。
那么，我们首先将抽象层的模块和接口定义出来，这里就需要了interface接口的设计，
然后我们依照抽象层，依次实现每个实现层的模块，在我们写实现层代码的时候，
实际上我们只需要参考对应的抽象层实现就好了，实现每个模块，也和其他的实现的模块没有关系，
这样也符合了上面介绍的开闭原则。这样实现起来每个模块只依赖对象的接口，而和其他模块没关系，
依赖关系单一。系统容易扩展和维护。

我们在指定业务逻辑也是一样，只需要参考抽象层的接口来业务就好了，
抽象层暴露出来的接口就是我们业务层可以使用的方法，然后可以通过多态的线下，
接口指针指向哪个实现模块，调用了就是具体的实现方法，这样我们业务逻辑层也是依赖抽象成编程。

我们就将这种的设计原则叫做依赖倒转原则。
*/

// Car ===== >   抽象层  < ========
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// BenZ ===== >   实现层  < ========
type BenZ struct {
	//...
}

func (benz *BenZ) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
	//...
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type ZhangThree struct {
	//...
}

func (zhang3 *ZhangThree) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

type LiFour struct {
	//...
}

func (li4 *LiFour) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}

// ===== >   业务逻辑层  < ========
func main() {
	//张3 开 宝马
	var bmw Car
	bmw = &Bmw{}
	
	var zhang3 Driver
	zhang3 = &ZhangThree{}
	
	zhang3.Drive(bmw)
	
	//李4 开 奔驰
	var benz Car
	benz = &BenZ{}
	
	var li4 Driver
	li4 = &LiFour{}
	
	li4.Drive(benz)
}
