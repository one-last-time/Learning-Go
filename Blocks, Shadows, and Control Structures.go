package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	z := true
	if x > 5 {
		x, y := 5, 20 //shadowing x
		fmt.Println(true)
		true := 10 //shadowing true which is predefined in universal block
		fmt.Println(true)
		fmt.Println(x, y, z)
	}
	fmt.Println(x)
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	// uncomment the following line will cause compilation err since n is only available in condition and statement block
	//fmt.Println(n)
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)

	for i := 0; i <= 10; i++ {
		if i == 3 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}
