
package main

import (
	"fmt"
)


type vendingMachine struct{
	insertedMoney int
}

func (m vendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *vendingMachine) InsertCoin(coin string) {
	m.insertedMoney = 10	
}
func main() {
	vm := vendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money : 0
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
}