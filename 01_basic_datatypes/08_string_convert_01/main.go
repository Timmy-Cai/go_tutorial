package main

import (
	"fmt"
	"strconv"
)

//基本类型转string
func main() {

	//golang 基本类型转成string
	var num1 int = 100
	var num2 float64 = 2323.23
	var b bool = false
	var char byte = 'h'
	var str string

	//第一种方式，fmt.Sprintf
	//golang中文文档。nhttps://studygolang.com/pkgdoc
	str = fmt.Sprintf("%d", num1) //int -> string
	fmt.Printf("str type %T str=%v \n", num1, str)

	str = fmt.Sprintf("%f", num2) //float64 -> string
	fmt.Printf("str type %T str=%v \n", num2, str)

	str = fmt.Sprintf("%t", b) //bool -> string
	fmt.Printf("str type %T str=%v \n", b, str)

	str = fmt.Sprintf("%c", char) //byte -> string
	fmt.Printf("str type %T str=%v \n", char, str)

	//第二种方式，strconv 函数
	str = strconv.FormatInt(int64(num1), 10) //int -> string, base表示进制，10表示10进制
	fmt.Printf("str type %T str=%v \n", num1, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64) //float64 -> string 'f'格式(可参考手册),10表示小数点后10位，64表示float64
	fmt.Printf("str type %T str=%v \n", num2, str)

	str = strconv.FormatBool(b) //bool -> string
	fmt.Printf("str type %T str=%v \n", b, str)

	str = strconv.Itoa(num1) //int -> string
	fmt.Printf("str type %T str=%v \n", num1, str)

}
