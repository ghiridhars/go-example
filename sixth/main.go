package main

import "fmt"

func main() {

	// variation 1
	x := 0
	for x < 3 {
		fmt.Println("value of x: ", x)
		x++
	}

	// variation 2
	for i := 0; i < 3; i++ {
		fmt.Println("value of i: ", i)
	}

	names := []string{"i", "am", "death", "destroyer", "of", "the", "worlds"}
	// loop array variation 1
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i], " ")
	}

	// loop array variation 2
	for index, value := range names {
		fmt.Printf("index: %v has value: %v\n", index, value)
	}

	// loop array variation 3 - value and index - local copy - wont change actual array
	for index, value := range names {
		fmt.Print(value, index)
		value = "new str"
	}
	fmt.Print(names)

}
