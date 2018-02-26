
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
	//m.insertedMoney = 10	
	if coin == "T" {
		m.insertedMoney += 10
	}
	if coin == "F" {
		m.insertedMoney += 5
	}
	if coin == "TW" {
		m.insertedMoney += 2
	}
	if coin == "O" {
		m.insertedMoney += 1
	}
}
func main() {
	vm := vendingMachine{}
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money : 0
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 5
}