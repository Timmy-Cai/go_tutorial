package main

import (
	"fmt"
	"unsafe"
)

//布尔类型
func main() {

	var b = true
	fmt.Println(b)

	//bool占用一个字节
	fmt.Printf("bool占用%d个字节", unsafe.Sizeof(b))

	//bool只能用true或者false，而非0或1
	//错误示范
	//b = 1
}
