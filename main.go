package main

import (
	"dsa/arrays"
	"fmt"
)

func main() {
	arr := []int{1, 8, 20, 25, 3}
	el, exist := arrays.LinearSearch(arr, 1)
	if exist {
		fmt.Println(el)
	}
	fmt.Println(exist)
}
