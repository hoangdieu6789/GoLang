package main

import (
	"fmt"
)

const Dieu string = "hello"

func main() {
	var a, e string = "hello", "world"
	var b, f = 1, "a"
	var Go bool
	d, g := 100, "hkdh"
	myslice := []int{}

	var cars = [4]int{1, 2, 3}
	fmt.Println(a, e)

	fmt.Println(b, f)
	fmt.Println(Go)
	fmt.Println(d, g)

	fmt.Println(Dieu)
	fmt.Println(cars)

	myslice = append(myslice, 1)
	fmt.Println(myslice)
}
