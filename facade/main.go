// The objective is to implement an interface that uses a complex subsystem with alot of moving parts
// to make them simple to use
// Here we can see a wallet that uses the walletFacade to implement various methods from a complex subsystem

package main

import (
	"fmt"
	"log"
)

func main() {
	walletFacade := newWalletFacade("abc", 1234)

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}
