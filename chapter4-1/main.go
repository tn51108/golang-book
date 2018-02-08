package main

import "fmt"

func main()  {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	const a string = "Hello, World"
	// cannot assign to a
	//a = "Other string"

	var (
		d = 5
		e = 10
		f = 15
	)
	fmt.Println(d,e,f)

	v1, v2 := "first", "sec"
	fmt.Printf("Type v1: %T\n", v1)
	fmt.Printf("Type v2: %T\n", v2)
	fmt.Println(v1, v2)
	v1, v2 = v2, v1
	fmt.Println(v1, v2)
}

func f()  {
	//undefind x
	//fmt.Println(x)
	
}