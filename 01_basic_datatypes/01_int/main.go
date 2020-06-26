package main

import (
	"fmt"
	"unsafe"
)

//整数类型
func main() {

	//有符号整数
	var i0 int = 1 //大小跟系统有关。如果在32位的系统上，则表示4个字节。范围同int32;如果在64位的系统上，则表示8个字节。范围同int64
	fmt.Println(i0)
	var i1 int8 = 1 //8位 占用1个字节 范围 -2的7次幂 到 2的7次幂-1
	fmt.Println(i1)
	var i2 int16 = 1 //16位 占用2个字节 范围 -2的15次幂 到 2的15次幂-1
	fmt.Println(i2)
	var i3 int32 = 1 //32位 占用4个字节 范围 -2的31次幂 到 2的31次幂-1
	fmt.Println(i3)
	var i4 int64 = 1 //64位 占用8个字节 范围 -2的63次幂 到 2的63次幂-1
	fmt.Println(i4)

	//错误示范
	//var i5 int8 = -129 // -129超过了 int8范围
	//fmt.Println(i5)

	//无符号整数
	var ui0 uint = 1 //大小跟系统有关。如果在32位的系统上，则表示4个字节。范围同uint32;如果在64位的系统上，则表示8个字节。范围同uint64
	fmt.Println(ui0)
	var ui1 uint8 = 1 //8位 占用1个字节 范围 0 到 2的8次幂-1
	fmt.Println(ui1)
	var ui2 uint16 = 1 //16位 占用2个字节 范围 0 到 2的16次幂-1
	fmt.Println(ui2)
	var ui3 uint32 = 1 //32位 占用4个字节 范围 0 到 2的32次幂-1
	fmt.Println(ui3)
	var ui4 uint64 = 1 //64位 占用8个字节 范围 0 到 2的64次幂-1
	fmt.Println(ui4)

	//错误示范
	//var ui5 uint8 = -1 // -1 超过了 uint8范围
	//fmt.Println(ui5)
	//var ui6 uint8 = 256 // -1 超过了 uint8范围 2的8次幂-1
	//fmt.Println(ui6)

	//其他
	//rune为有符号整数。占用4个字节，占用的空间和数字范围同int32一致，表示一个unicode码。用于存储中文
	//byte为无符号整数。占用1个字节，占用的空间和数字范围同uint8一致，当存储单个字母字符字符时选用byte
	//如果不给类型，默认为int类型
	var i6 = 1
	fmt.Printf("i6 的类型是:%T \n", i6)
	//查看变量类型以及占用的字节数
	var i7 int16 = 1
	fmt.Printf("i7 的类型是:%T i7占用的字节数:%d \n", i7, unsafe.Sizeof(i7))

	//拓展
	//bit是计算机中最小的存储单位。byte是计算机中基本的存储单元。 1byte=8bit
	//尽量选用合适的类型，在数据确定的情况下，保小不保大，占用空间小

}
