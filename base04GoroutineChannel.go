// base04GoroutineChannel
package main

import (
	"fmt"
	"time"
)

func printSomething(str string, waitSecond int) {
	time.Sleep(time.Duration(waitSecond) * time.Second)
	fmt.Println(str)
}
func exampleRun1() {
	//go 作为关键字运行
	go printSomething("sleep 2", 2)
	go printSomething("sleep 1", 1)

	fmt.Println("print now")
	//主线程不休眠会立刻终止
	time.Sleep(4 * time.Second)
	fmt.Println("end pringsomething****")
}
func main2() {
	//上面是没在channel的实例，主线程退出会退出goroutine##################
	exampleRun1()

	//下面是有channel的实例
	exampleRun2()

}
func exampleRun2() {
	//常用的集中chan make方式
	//chan chanStr:=make(chan string)
	//chan chanInterface:make(chan interface{})
	c = make(chan int)
	go printSomething2("sleep 2", 2)
	go printSomething2("sleep 1", 1)
	fmt.Println("print2 now")
	//	<-c
	//	<-c
	//	<-c  //再取出一次试试

	//另外一种从channl中取出方法
	var i int
L:
	for {
		select {
		case <-c:
			fmt.Println(i)
			i++
			if i > 1 {
				break L
			}
		}
	}
}

var c chan int

func printSomething2(str string, waitSecond int) {
	time.Sleep(time.Duration(waitSecond) * time.Second)
	fmt.Println(str)

	//发送到channel
	c <- 1
}
