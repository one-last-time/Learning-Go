package main

import "fmt"

var flag bool
var flag1 bool = false
var (
	x int = 10
	y     = 20.2
	z     = "hello world"
)

const a = 5

func main() {
	fmt.Println(flag, flag1)
	x, i := z, float64(x)+y
	fmt.Println(x, i, a, z[0], string(z[0]), z)

}
