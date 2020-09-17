package main

import (
	"errors"
	"fmt"
)

// Bitcoin is int type data
type Bitcoin int

// Stringer will override print methods of object
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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

// Withdraw take amount to decrase wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}
