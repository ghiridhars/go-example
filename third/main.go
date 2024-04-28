package main

import "fmt"

func main() {

	age := 25.326
	name := "Myocardial infarction"

	// print
	fmt.Print("we learn about print now \n")
	fmt.Println("next line is automatically added")

	// print with parameters
	fmt.Println("my age is", age, "and I can have\n", name)

	// formated strings
	// %v - format specifiers
	fmt.Printf("my age is %v and I can have %v\n", age, name)

	// %q - format specifiers to add quote - has to be string - else error
	fmt.Printf("my age is %v and I can have %q\n", age, name)

	// %T - format specifiers to know what type is the var
	fmt.Printf("my age is of type %T\n", age)

	//formating decimal places of float
	fmt.Printf("I have %f points\n", 152.6)
	fmt.Printf("I have %0.2f points\n", 152.6)

	// SPRNTF - save print f
	var str = fmt.Sprintf("my age is %v and I can have %q\n", age, name)
	fmt.Println("Saved STRING : :", str)

}
