package main

import (
	"fmt"
)

// Wallet type data
type Wallet struct {
	balance int
}

// Deposit will adds amount of money to wallet
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance will print total of money on wallet
func (w *Wallet) Balance() int {
	return w.balance
}
