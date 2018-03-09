package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int, 3)
	go func() { ch <- 1 }()
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("cap:", cap(ch))
	fmt.Println("len:", len(ch))
	//fmt.Println( <- ch)
	//fmt.Println( <- ch)
	//fmt.Println( <- ch)
	//fmt.Println( <- ch)
}