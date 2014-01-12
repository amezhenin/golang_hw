package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var a [10]int
	fmt.Sprintf("asdf")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i, j := 0, 0
	for i = 0; i < len(a); i++ {
		fmt.Printf(strconv.Itoa(a[i]) + " ")
	}
	fmt.Println("\nWrap array!")
	for i, j = 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	for i = 0; i < len(a); i++ {
		fmt.Printf(strconv.Itoa(a[i]) + " ")
	}

	fmt.Println("Finish!")
}
