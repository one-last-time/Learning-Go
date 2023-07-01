package main

import (
	"fmt"
)

var x [3]int
var x1 = [3]int{1, 2, 3}
var x2 = [5]int{1, 3: 10, 4}
var x3 = [...]int{1, 2, 10}

var slc []int

func main() {
	fmt.Println(x)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(slc == nil)
	slc = append(slc, 9)
	fmt.Println(len(slc))
	//make
	slc1 := make([]int, 5, 6)
	fmt.Println(slc1)
	fmt.Println(cap(slc1))
	fmt.Println(len(slc1))
	slc1 = append(slc1, 1)
	fmt.Println(slc1)
	fmt.Println(cap(slc1))
	fmt.Println(len(slc1))
	//capacity increase
	slc1 = append(slc1, 1)
	fmt.Println(slc1)
	fmt.Println(cap(slc1))
	fmt.Println(len(slc1))
	//slicing slice
	slc2 := make([]int, 0, 5)
	slc2 = append(slc2, 1, 2, 3, 4)
	slc3 := slc2[:2]
	slc4 := slc2[2:]
	fmt.Println(cap(slc2), cap(slc3), cap(slc4))
	slc3 = append(slc3, 30, 40, 50)
	slc2 = append(slc2, 60)
	slc4 = append(slc4, 70)
	fmt.Println("slc2:", slc2)
	fmt.Println("slc3:", slc3)
	fmt.Println("slc4:", slc4)

	// appending to modified slc3 modified slc2 and sl4 cause they shared unused capacity of parent slice
	//thats why when slicing we should use full expressing the third part indecate the last index of parents capacity

	slc5 := make([]int, 0, 5)
	slc5 = append(slc5, 1, 2, 3, 4)
	slc6 := slc5[:2:2]
	slc7 := slc5[2:4:4]
	fmt.Println(cap(slc5), cap(slc5), cap(slc7))
	fmt.Println("slc5:", slc5)
	fmt.Println("slc6:", slc6)
	fmt.Println("slc7:", slc7)

	slc6 = append(slc6, 30, 40, 50)
	slc5 = append(slc5, 60)
	slc7 = append(slc7, 70)
	fmt.Println("slc5:", slc5)
	fmt.Println("slc6:", slc6)
	fmt.Println("slc7:", slc7)
	//now its alright.
	cnt := copy(slc5, slc6)
	fmt.Println(slc5, cnt)

	str := "hello 😁"
	bs := []byte(str)
	rs := []rune(str)
	fmt.Println(bs)
	fmt.Println(rs)
	fmt.Println(string(bs[4:7])) // it takes 4 bytes to represent smile emoji thats why
	fmt.Println(string(bs[4:10]))

	//map
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	v, ok := totalWins["hello"]
	fmt.Println(v, ok)
	v, ok = totalWins["Lions"]
	fmt.Println(v, ok)
	delete(totalWins, "Lions")
	fmt.Println("delete Lions")
	v, ok = totalWins["Lions"]
	fmt.Println(v, ok)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}
