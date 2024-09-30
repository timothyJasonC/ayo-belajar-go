package main

import (
	"fmt"
)

func main() {
	// var txt = "Hello"

	// fmt.Printf("%s\n", txt)
	// fmt.Printf("%q\n", txt)
	// fmt.Printf("%24s\n", txt)
	// fmt.Printf("%-24s\n", txt)
	// fmt.Printf("%x\n", txt)
	// fmt.Printf("% x\n", txt)

	// var i = true
	// var j = false

	// fmt.Printf("%t\n", i)
	// fmt.Printf("%t\n", j)

	// var arr1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// arr2 := [5]int{4, 5, 6, 7, 8}

	// fmt.Println(arr1)
	// fmt.Println(arr2[2])

	// arr1 := [...]int{0: 12, 1: 10, 2: 40}

	// fmt.Println(arr1)

	// myslice1 := []int{1, 2, 3, 4}
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))
	// fmt.Println(myslice1)

	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))
	// fmt.Println(myslice2)

	// arr1 := [6]int{10, 11, 12, 13, 14, 15}
	// myslice := arr1[1 : len(arr1)-2]
	// fmt.Println(myslice)
	// fmt.Printf("myslice = %v\n", myslice)
	// fmt.Printf("length = %d\n", len(myslice))
	// fmt.Printf("capacity = %d\n", cap(myslice))

	// arr1 := [6]int{10, 11, 12, 13, 14, 15}
	// fmt.Printf("myslice = %v\n", arr1)
	// myslice := append(arr1[:], 2, 3)
	// fmt.Printf("myslice = %v\n", myslice)

	// myslice1 := []int{1, 2, 3}
	// myslice2 := []int{4, 5, 6}
	// myslice3 := append(myslice1, myslice2...)
	// myslice3 = append(myslice3, 1, 2)

	// fmt.Printf("myslice3=%v\n", myslice3)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
