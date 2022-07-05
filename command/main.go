// Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request.
// The Command pattern suggests that GUI objects shouldn’t send these requests directly. Instead, you should extract all of the request details,
// such as the object being Doed, the name of the method and the list of arguments into a separate command class with a single method that triggers this request.

package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount,
		"\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount,
			"\b, balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Do()
	Undo()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

// the standalone object
type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Do() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Do()
	cmd.Undo()
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Do()
	fmt.Println(ba)
}
