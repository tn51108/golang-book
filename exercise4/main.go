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
		coffee := receiveOrder(i)
		c := make(chan int)
		coffee = brew(coffee)
		container = append(container, serve(coffee))
	}
	return
}

func receiveOrder(number int) string {
	time.Sleep(5 * time.Millisecond)
	return fmt.Sprintf("order: %d", number)
}
func brew(order chan<- string) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("%s %s", coffee, "espresso")
}
func serve(cofee string) string {
	time.Sleep(5 * time.Millisecond)
	fmt.Sprintf("%s %s", coffee, "ready :)")
}