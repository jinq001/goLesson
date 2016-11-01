// base03Pointer
package main

import (
	"fmt"
)

var p *int

func pointExample1() {
	var i int
	//指向i的地址
	p = &i
	fmt.Println(p)

	*p = 8
	fmt.Println("修改指针指向的值*p=", *p)
	fmt.Println("i=", i)

	*p++
	fmt.Println("没有指针运算*p++=", *p)
}

//func main3() {
//	fmt.Println(p)
//	pointExample1()
//}

//内存回收
