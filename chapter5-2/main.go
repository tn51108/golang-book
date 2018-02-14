package main

import (
	"strconv"
	"fmt"
)

func main()  {
	for number := 1; number <= 100; number++ {
		/*if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println(number, "Fizz")
		} else if number%5 == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
		*/
		fmt.Println(fizzbuzz(number))
	}
}

func fizzbuzz(number int) string {
	
	if number%15 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}