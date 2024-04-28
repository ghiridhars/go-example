package main

import "fmt"

func main() {

	// ARRAYS

	// var ages [3]int = [3]int{1, 2, 3} is same as below
	var ages = [3]int{1, 2, 3}

	fmt.Printf("Arrays %v of length %v", ages, len(ages))

	names := [3]string{"me", "myself", "ghiri"}
	names[1] = "Meself"

	fmt.Printf("\nArrays %v of length %v", names, len(names))

	// SLICES -  uses arrays internally
	var marks = []int32{87, 88, 87}
	marks = append(marks, 76)
	fmt.Printf("\nArrays %v of length %v", marks, len(marks))

	// slice ranges

	marks1 := marks[1:2]
	names1 := append(names[:], "And", "asa")
	fmt.Println("\n", marks1, "::", names1)

}
