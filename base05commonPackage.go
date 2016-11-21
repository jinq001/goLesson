// base05io
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func http1() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println("err=", err)
	}

}

//func main() {
//	http1()
//}
