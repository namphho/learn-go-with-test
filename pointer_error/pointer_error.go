package pointer_error

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("deposit address of balance in test is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("Balance address of balance in test is %v \n", &w.balance)
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}