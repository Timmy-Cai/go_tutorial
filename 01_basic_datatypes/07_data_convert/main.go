package main

import "fmt"

//基本数据类型转换
func main() {
	//golang 数据类型必须强制转换,不能自动转换

	var a int32 = 100
	//int32->float32
	var b float32 = float32(a)
	fmt.Println(b)
	//float32->float64
	var c float64 = float64(b)
	fmt.Println(c)
	//int32->int8
	var d int8 = int8(a)
	fmt.Println(d)

	//以下特殊情况，编译不会报错，结果是按溢出处理，所以类型转换一定要考虑数据范围
	var e uint16 = 1000
	var f uint8 = uint8(e)
	fmt.Println(f)

}
