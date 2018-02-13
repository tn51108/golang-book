package main

import "fmt"

func main() {
	var input1 int
	x := 9
	for i := 0; i < 5; i++ {
		fmt.Print("Enter Number:")
		fmt.Scanf("%d\n", &input1)
		if input1 == x {
			fmt.Println("Match")
		} else if input1 > x {
			fmt.Println("Must Less than", input1)
		} else {
			fmt.Println("Must More than", input1)
		}
	}
}
