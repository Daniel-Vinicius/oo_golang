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

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Test",
	}

	fmt.Println(sliceString)

	// Take all values until position two in this case it will take position 0 and 1 [Hello World]
	fmt.Println(sliceString[:2])

	// Take all values from position 2 until position 4 (it don't take the fourth position) in this case it will take position 2 and 3 [Much Better]
	fmt.Println(sliceString[2:4])

	// Take all values from position 2 until last position in this case it will take position 2, 3 and 4 [Much Better Test]
	fmt.Println(sliceString[2:])
}
