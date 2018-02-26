
package main

import (
	"fmt"
)


type vendingMachine struct{
	insertedMoney int
	coins map[string]int
	items map[string]int
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

func (m *vendingMachine) SelectSD() string {
	price := 18
	change := m.insertedMoney - price
	return "SD" + m.change(change)
}
func (m *vendingMachine) SelectCC() string {
	price := 12
	change := m.insertedMoney - price
	//m.insertedMoney = 0
	return "CC" + m.change(change)
}
func  (m *vendingMachine) change(c int) string {
	var str string
	if c >= 5 {
		str += ", F"
		c -= 5
	}
	if c >= 2 {
		str += ", TW"
		c -= 2
	}
	if c >= 1 {
		str += ", O"
		c -= 1
	}
	return str
}
func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var items = map[string]int{"SD": 18, "CC": 12, "DK": 7}
	//vm := vendingMachine{}
	vm := vendingMachine{coins: coins, items: items}
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
	can := vm.SelectSD()
	fmt.Println(can)
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 10
	vm.InsertCoin("TW")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	// Inserted Money: 2
	can = vm.SelectCC()
	fmt.Println(can)
}