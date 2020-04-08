package pointer_error

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin)String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	bal Bitcoin
}

func (w *Wallet)Deposit(amount Bitcoin) {
	(*w).bal += amount
}

var ErrInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet)Withdraw(amount Bitcoin) error {
	if amount > w.bal {
		return ErrInsufficientFunds
	}
	(*w).bal -= amount
	return nil
}

func (w Wallet)Balance() Bitcoin {
	return w.bal
}