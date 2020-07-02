package main

import "fmt"

const (
	// cat = iota + 100
	// dog
	// dragon

	// api   = "https://api.google.com"
	// Tiger = iota + 100

	Execute = 1 << iota
	Write
	Read
)

func main() {
	// fmt.Println(cat, dog, dragon, Tiger)
	fmt.Println(Execute, Write, Read)
	i := 0
	i++
	fmt.Println(i)
}
