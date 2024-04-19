package main

import "fmt"

var outsideName = "declaring string outside function"

// var outsideName := "declaring string outside function"
// above gives error in compile time

func main() {

	//----------------------------
	//	strings
	//----------------------------
	var nameFirst string = "String variable declaration"
	var nameSecond = "String variable declaration variation 2"
	var nameThird string

	fmt.Println(nameFirst, nameSecond, nameThird)

	nameFirst = "String change of values"
	nameThird = " setting value"

	fmt.Println(nameFirst, nameSecond, nameThird)

	// cant use := out side of a function
	nameFourth := "shorthand for declare and initialize at same time. Only used for initializing"

	fmt.Println(nameFourth)

	//----------------------------
	//	integers - int and float
	//----------------------------

	var ageFirst int = 20
	var ageSecond = 154
	ageThird := 120

	fmt.Println(ageFirst, ageSecond, ageThird)

	// bits and memory variations inside int

	// uint8        (0 to 255)
	// uint16       (0 to 65535)
	// uint32       (0 to 4294967295)
	// uint64       (0 to 18446744073709551615)

	// int8         (-128 to 127)
	// int16        (-32768 to 32767)
	// int32        (-2147483648 to 2147483647)
	// int64        (-9223372036854775808 to 9223372036854775807)

	var numFirst int8 = 125
	var numSecond uint8 = 254

	fmt.Println(numFirst, numSecond)

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	var floatFirst float32 = 121.165329                                                   // prints upto 5 decimal -2^5 32 bit
	var floatSecond float64 = 164623164643131316416631216541.1653294611284946416316613145 // used mostly for its higher precision - outputed as -  1.6462316464313132e+29

	fmt.Println(floatFirst, floatSecond)

}
