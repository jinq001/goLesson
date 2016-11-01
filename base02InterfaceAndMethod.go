// interfaceAndMethod
package main

import (
	"fmt"
)

/////////////////////////////////////////////////////////////
//1完整的函数定义
func (p *Struct1) fullFunc(param int) (r int, s string) {
	return 1, "a"
}
func (p *Struct1) fullFunc2(param int) string {
	return "b"
}
func fullFunc3(param int) {

}

/////////////////////////////////////////////////////////////
//2局部变量全局变量
var cycle = 3

func a() {
	cycle := 4
	fmt.Println("局部变量定义cycle := 4 ", cycle)
}
func b() {
	cycle = 10
	fmt.Println("全局变量赋值：", cycle)
}

//func main() {
//	fmt.Println("全局变量：", cycle)
//	a()
//	fmt.Println("局部变量定义不改变全局同名变量  ", cycle)
//	b()
//}

/////////////////////////////////////////////////////////////
//3    多返回值
func returnMany() (v int, b string) {
	return
}

//func main() {
//	var va, bb = returnMany()
//	fmt.Println("定义了名称的返回值可直接操作和返回", va, bb)
//}

/////////////////////////////////////////////////////////////
//4 延迟列表  类似于finally
func delay1(i int) bool {
	defer printDelay()
	if i == 0 {
		return false
	} else if i < 5 {
		return false
	}
	return true

}
func printDelay() {
	fmt.Println("delay func")
}

//func main() {
//	delay1(1)
//}

/////////////////////////////////////////////////////////////
//5 变参
func myfunc(arg ...int) {
}

/////////////////////////////////////////////////////////////
//6 匿名函数
func anonymousFunc(b int) {
	a := func() {
		fmt.Println("匿名了", b)
	}
	fmt.Printf("a的类型 %T\n ", a)
	a()
}

func main() {
	anonymousFunc(1)
	f := anonymousFunc
	callBack(10, f)
}
func callBack(someValue int, f anonymousFunc) {
	f(someValue)
}

//接口定义
type InterfaceOne interface {
	Get() int
	Put(int)
}

//定义结构和结构的方法
type Struct1 struct {
	i int
}

func (p *Struct1) Get() int {
	return p.i
}

func (p *Struct1) Put(v int) {
	p.i = v
}

//接口值的调用方式
func somefunc(p InterfaceOne) {
	fmt.Println(p.Get())
	p.Put(1)
}
func callSomefuncUseStruct1() {
	//StructOne的指针上实现了InterfaceOne，所以能直接调用somefunc方法
	var s Struct1
	somefunc(&s)
	//这里表面无需明确一个类型是否实现一个接口（duck typing模式）
}

type Struct2 struct {
	i int
}

func (p *Struct2) Get() int {
	return p.i
}
func (p *Struct2) Put(v int) {
	p.i = v
}
func callStruct1AndStuct2(p InterfaceOne) {
	//可以接受interfaceOne的多个实现
	//type只能在switch中使用
	switch t := p.(type) {
	case *Struct1:
		fmt.Println(t.Get())
	case *Struct2:
		//	case Struct1:
		//		fmt.Println(t.Get())
		//	case Struct2:
	default:
	}
}

//func main() {
//	s := Struct1()(4)
//	s.Put(5)
//	callStruct1AndStuct2(s)
//}
