package main

import "fmt"

//浮点型
func main() {

	//单精度，float32 4个字节，类似于Java中的float,取值范围 -3.403E38 ~ 3.403E38
	var price1 float32 = 0.01
	fmt.Printf("price1:%v \n", price1)

	//双精度(默认)，float64 8个字节，类似于Java中的double,取值范围 -1.798E308 ~ 1.798E308
	var price2 float64 = 0.02
	fmt.Printf("price2:%v \n", price2)

	//双精度范围比单精度大，精度更准确一些。float32会丢失精度，如下
	var price3 float32 = 111.000000000101
	fmt.Printf("price3:%v \n", price3)
	var price4 float64 = 111.000000000101
	fmt.Printf("price4:%v \n", price4)

	//解释：符号位+指数位+尾数位 在存储过程中，精度会有丢失，如果我们要保存一个精度高的数，则应该选用 float64
}
