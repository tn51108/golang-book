package main

import (
	"fmt"
	"time"
)

func main()  {
	volumn := 10
	start := time.Now()

	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
func order(volumn int) (container []string) {
	for i := 1; i <= volumn; i++ {
		//casier receive order
		casiers := make(chan string, 3)
		go casier(casiers)
		coffee := fmt.Sprintf("order: %d", i)

		//barista brew coffe
		baristas := make(chan string, 1)
		go barista(baristas)
		coffee = fmt.Sprintf("%s %s", coffee, "espresso")

		//waiter serve coffee
		waiters := make(chan string, 1)
		waiter(waiters)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}
	return
}

func casier(casiers chan<- string)  {
	time.Sleep(5 * time.Millisecond)
}
func barista(baristas chan<- string)  {
	time.Sleep(100 * time.Millisecond)
}
func waiter(waiters chan<- string)  {
	time.Sleep(5 * time.Millisecond)
}