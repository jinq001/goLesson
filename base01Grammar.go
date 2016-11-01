// baseGrammar
package main

import (
	"fmt"
	"reflect"
)

func one() {
	var a int = 1
	var b int8 = 2
	var c int64 = 2312222222222222221
	//	fmt.Printf(reflect.TypeOf(c).Size())
	fmt.Printf("%t", b)
	d := false
	var e = 333.2
	var f, g int
	var h, i string
	var k, j string = "a", "b"
	l := j
	m := &j
	fmt.Println(a, b, c, d, e, f, g, h, i, k, &j, l, &l, m)
	//指针分配
	//	j = &"str"
	fmt.Println(j)
	fmt.Println()
}
func two() {
	const b = 1
	fmt.Println(b)
	const (
		a = iota
		c
	)

	const (
		d = iota
	)
	fmt.Println(a, c, "重置iota:", d)
	var eInt int = 55
	var eIntPoint *int
	eIntPoint = &eInt
	fmt.Println("eInt内存地址：", &eInt, "指针一致：", eIntPoint, "值在么：", *eIntPoint)
	eInt = 66
	fmt.Println("引用的位置没变，但是改变值了：", *eIntPoint)
	fmt.Println(reflect.TypeOf(eInt).Size())
}
func delayFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
func indefiniteParam(arg ...int) {
	for a := range arg {
		fmt.Println(a)
	}

}

/////////////////////////
//一下是接口例子

//func main() {
//	one()
//	two()
//	delayFunc()
//	indefiniteParam(3, 4, 5, 6)

//	//

//	//		somefunc(new InterfaceInstance())
//}
