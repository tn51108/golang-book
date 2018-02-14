package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomString(10))
	/*
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
	*/
}

func ranInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i<l; i++ {
		bytes[i] = byte(ranInt(65, 90))
	}
	return string(bytes)
}