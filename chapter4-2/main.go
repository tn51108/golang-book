package main

import "fmt"

func main()  {
	fmt.Print("Fahrenheit : ")
	var input1 float64
	fmt.Scanf("%f", &input1)
	output1 := (input1 - 32) * 5 / 9
	fmt.Printf("Celsius : %.2f\n",output1)
}