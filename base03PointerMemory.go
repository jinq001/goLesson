// base03Pointer
package main

import (
	"fmt"
)

var p *int

//go的指针操作，go没有指针运算，指针是为了提高函数穿参的效率
func pointExample1() {
	var i int
	//指向i的地址
	p = &i
	fmt.Println("p=&i 地址=", p)
	fmt.Println("i现在=", i)
	*p = 8
	fmt.Println("修改指针指向的值*p=", *p)
	fmt.Println("i=", i)

	*p++
	fmt.Println("没有指针运算*p++实际等于*p(++) ", *p)
}

//内存分配
// 1.new分配  2、make分配

func mem() {

}
