package main

import (
	"fmt"
)

func main()  {
	weatherCelsius(0)
	//fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
}

func weatherCelsius(incels int)  {
	
	//runes := []rune(cels)
	/*
	// A string.
    value := "Welcome, my friend"

    // Take substring of first word with runes.
    // ... This handles any kind of rune in the string.
    runes := []rune(value)
    // ... Convert back into a string from rune slice.
    safeSubstring := string(runes[0:7])
    fmt.Println(" RUNE SUBSTRING:", safeSubstring)

    // Take substring of first word with direct byte slice.
    // ... This handles only ASCII correctly.
    asciiSubstring := value[0:7]
    fmt.Println("ASCII SUBSTRING:", asciiSubstring)
	*/
	l1, l2, l3 := plotcel(incels)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)

}

func plotcel(cels int) (line1 string, line2 string, line3 string) {
	/*if cels==1 {
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
	}*/
	
	switch cels {
	case 0: {
		line1 := " _ "
		line2 := "| |"
		line3 := "|_|"
	}
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
	return line1, line2, line3
}