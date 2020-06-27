package main

import "fmt"

//字符串类型(Golang中的字符串由单个字节一个个组成的，使用的是UTF-8)
func main() {

	var address string = "我爱中国 010101 啦啦啦"
	fmt.Println(address)

	//字符串不可变
	//错误示范
	//address[0] = 'h'

	//通过反引号打印出原文。实现防止攻击
	var text string = `
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
			fmt.Printf("bool占用%d个字节",unsafe.Sizeof(b))
		
			//bool只能用true或者false，而非0或1
			//错误示范
			//b = 1
		}
		`
	fmt.Println(text)

	//字符串拼接方式
	var str string = "中号" + "衣服"
	fmt.Println(str)
	//分行行
	var str2 string = "中号" +
		"衣服"
	fmt.Println(str2)
}
