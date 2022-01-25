package main

import "fmt"

func main() {
	slice := make([]int, 2, 5)
	slice[0] = 0
	slice[1] = 5

	slice = append(slice, 10, 15, 20, 25)

	fmt.Println(len(slice), "positions, actual length")
	fmt.Println(cap(slice), "capacity, actual capacity of slice")
	fmt.Println("The slice", slice)

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Length:", len(slice))
		fmt.Println("Capacity:", cap(slice))
	}

	sliceTest := slice
	slice = append(slice, 43, 57, 63, 77)
	slice[0] = 99
	fmt.Println(slice)
	fmt.Println(sliceTest)
}
