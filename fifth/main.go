package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// STRINGS package examples
	message := "Hola amigos! como estas?"

	fmt.Println(strings.Contains(message, "comos"))
	fmt.Println(strings.ReplaceAll(message, "como", "bueno"))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.Index(message, "ol"))
	fmt.Println(strings.Split(message, " "))

	fmt.Println("original value ::: ", message)

	// SORT package

	age := []int{23, 4, 123, 56, 7, 53, 908}

	fmt.Println("BEFORE	:", age)
	sort.Ints(age)
	fmt.Println("AFTER	:", age)

	// WILL ONLY WORK on sorted array
	index := sort.SearchInts(age, -32)
	fmt.Println("index::", index)

	names := []string{"I", "zebra", "no", "love", "apple"}

	fmt.Println("BEFORE	:", names)
	sort.Strings(names)
	fmt.Println("AFTER	:", names)

	// WILL ONLY WORK on sorted array
	sindex := sort.SearchStrings(names, "apple")
	fmt.Println("index::", sindex)

}
