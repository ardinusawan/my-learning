package main

import (
	"fmt"
)

// Stringer will override print methods of object
type Stringer interface {
	String() string
}

// Bitcoin is int type data
type Bitcoin int

// Wallet type data
type Wallet struct {
	balance Bitcoin
}

// Deposit will adds amount of money to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance will print total of money on wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
