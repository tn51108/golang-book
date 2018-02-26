
package main

import (
	"fmt"
)


type vendingMachine struct{
	insertedMoney int
	coins map[string]int
}

func (m vendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *vendingMachine) InsertCoin(coin string) {
	//m.insertedMoney = 10	
	/*
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
	}*/
	m.insertedMoney += m.coins[coin]
}
func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	//vm := vendingMachine{}
	vm := vendingMachine{coins: coins}
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money : 0
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 5
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 2
}