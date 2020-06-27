package main

import "fmt"

//字符
func main() {
	//golang 存储单个字符一般使用byte,如下
	var i byte = 'a'
	//直接打印byte结果，输出的是对应的十进制ASCII码值
	fmt.Println(i) //97

	//打印出字符本身
	fmt.Printf("i=%c \n", i)

	//错误情况
	//var a byte = '中'
	//fmt.Printf("a=%c", a)  //中 UTF-8编码 十进制 UNICODE值是20013， 超过了byte的最大取值范围

	//正确输出
	var b int = '中'
	fmt.Printf("b=%c b对应的码值:%d \n", b, b) //中 int 取之范围比byte大

	//通过ASCII码值找到对应的字符
	var c1 byte = 97
	fmt.Printf("c1=%c \n", c1)
	//通过UNICODE码值找到对应的字符
	var c2 int = 20013
	fmt.Printf("c2=%c \n", c2)

	//字符之间可以运行，是按照码值运算的
	var c3 int = 1
	fmt.Printf("c3=%d \n", c3+'a') //1 + 97 = 98

	//注意：
	//1.如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]直接可以保存到 byte
	//2.如果我们保存的字符对应码值大于 255,这时我们可以考虑使用 int 类型保存
	//3.如果我们需要安装字符的方式输出，这时我们需要格式化输出，即 fmt.Printf(“%c”, c1)
	//4.golang 字符串是由字节组成，而非字符
	//5.UTF-8编码包括了ASCII码值，只是ASCII码值不够使用，UTF-8进行了拓展，英文字母-1个字节 汉字-3个字节
}
