package main

import "fmt"

// tinyint(4) 1个字节 8bit 2^8 = 256 0~255 -128~127
// smallint 2个字节
// mediumint 3
// int 4
// bigint  8

type myInt int64 //类型再定义 不支持隐式转换

type myInt8 = int8 //类型别名 除了名字不一样 其他都一样

var s2 string

func main() {
	var s2 string
	s2 = "aaa"
	fmt.Println(s2)
	{
		var s2 string
		s2 = "bbb"
		fmt.Println(s2)
		{
			var s2 string
			s2 = "ccc"
			fmt.Println(s2)
		}
	}
	s1 := "abc"
	fmt.Println(s1)
	var i1 myInt
	i1 = 10

	var i2 int64
	i2 = 10

	fmt.Println(i1 == myInt(i2))
}
