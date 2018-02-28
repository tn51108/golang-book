package main

import (
	"fmt"
	"github.com/tn51108/vendingMachine"
)
func main() {
	var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var items = map[string]int{"SD": 18, "CC": 12, "DK": 7}
	//vm := vendingMachine{}
	//vm := vendingMachine{coins: coins, items: items}
	vm := vendingMachine.NewVendingMachine(coins, items)
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
	coin := vm.CoinReturn()
	fmt.Println(coin)
}