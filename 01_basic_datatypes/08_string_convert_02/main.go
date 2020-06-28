package main

import (
	"fmt"
	"strconv"
)

//string转基本类型
func main() {

	//string->bool
	var str string = "true"
	b, _ := strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n", b, b)

	//string->int
	var str2 string = "123123123"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64) //10表示10进制，64表示int64
	fmt.Printf("n1 type %T b=%v \n", n1, n1)

	//string->float
	var str3 string = "234.23423"
	var f float64
	f, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f type %T b=%v \n", f, f)

	//注意:要确保 String 类型能够转成有效的数据，否则会返回默认值，如下:
	n2, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(n2)
}
