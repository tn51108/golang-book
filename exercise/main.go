package main

import (
	"fmt"
)

func main()  {
	weatherCelsius(0)
}

func weatherCelsius(incels int)  {
	
	//runes := []rune(cels)
	writecels(incels)
}

func writecels(cels int)  {
	if cels==1 {
		fmt.Println("  ")
		fmt.Println(" |")
		fmt.Println(" |")
	} else if cels==2 {
		fmt.Println(" _")
		fmt.Println(" _|")
		fmt.Println("|_")
	} else if cels==3 {
		fmt.Println("_")
		fmt.Println("_|")
		fmt.Println("_|")
	} else if cels==4 {
		fmt.Println(" ")
		fmt.Println("|_|")
		fmt.Println("  |")
	} else if cels==5 {
		fmt.Println(" _")
		fmt.Println("|_")
		fmt.Println(" _|")
	} else if cels==6 {
		fmt.Println(" _")
		fmt.Println("|_")
		fmt.Println("|_|")
	} else if cels==7 {
		fmt.Println("__")
		fmt.Println(" /")
		fmt.Println("/")
	} else if cels==8 {
		fmt.Println(" _")
		fmt.Println("|_|")
		fmt.Println("|_|")
	} else if cels==9 {
		fmt.Println(" _")
		fmt.Println("|_|")
		fmt.Println(" _|")
	}  else if cels==0 {
		fmt.Println(" _")
		fmt.Println("| |")
		fmt.Println("|_|")
	}
}