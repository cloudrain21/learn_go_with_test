package main

import (
	"fmt"
	po "github.com/bob/pointer_error"
)

func main() {
	wallet := po.Wallet{}
	wallet.Deposit(100)
	fmt.Println(wallet.Balance())
	//fmt.Println(wallet.bal)
}
