package main

import "fmt"

func main() {
	x := 1
	//获取x的地址
	p := &x
	//p指向的变量*p
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
}
